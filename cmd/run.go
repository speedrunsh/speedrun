package cmd

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/speedrunsh/portal/command"
	"github.com/speedrunsh/portal/transport"
	"github.com/speedrunsh/speedrun/cloud"
	"github.com/speedrunsh/speedrun/colors"
	"github.com/speedrunsh/speedrun/key"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"

	"github.com/apex/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var runCmd = &cobra.Command{
	Use:     "run <command to run>",
	Short:   "Run a command on remote servers",
	Example: "  speedrun run whoami\n  speedrun run whoami --only-failures --target \"labels.foo = bar AND labels.environment = staging\"",
	Args:    cobra.MinimumNArgs(1),
	PreRun: func(cmd *cobra.Command, args []string) {
		initConfig()
	},
	RunE: run,
}

func init() {
	runCmd.Flags().StringP("target", "t", "", "Fetch instances that match the target selection criteria")
	runCmd.Flags().String("projectid", "", "Override GCP project id")
	runCmd.Flags().Bool("only-failures", false, "Print only failures and errors")
	runCmd.Flags().Bool("ignore-fingerprint", false, "Ignore host's fingerprint mismatch")
	runCmd.Flags().Duration("timeout", time.Duration(10*time.Second), "SSH connection timeout")
	runCmd.Flags().Int("concurrency", 100, "Number of maximum concurrent SSH workers")
	runCmd.Flags().Bool("use-private-ip", false, "Connect to private IPs instead of public ones")
	runCmd.Flags().Bool("use-oslogin", false, "Authenticate via OS Login")
	runCmd.Flags().Bool("use-tunnel", true, "Connect to the portals via SSH tunnel")
	viper.BindPFlag("gcp.projectid", runCmd.Flags().Lookup("projectid"))
	viper.BindPFlag("gcp.use-oslogin", runCmd.Flags().Lookup("use-oslogin"))
	viper.BindPFlag("ssh.timeout", runCmd.Flags().Lookup("timeout"))
	viper.BindPFlag("ssh.ignore-fingerprint", runCmd.Flags().Lookup("ignore-fingerprint"))
	viper.BindPFlag("ssh.only-failures", runCmd.Flags().Lookup("only-failures"))
	viper.BindPFlag("ssh.concurrency", runCmd.Flags().Lookup("concurrency"))
	viper.BindPFlag("ssh.use-private-ip", runCmd.Flags().Lookup("use-private-ip"))
	viper.BindPFlag("portal.use-tunnel", runCmd.Flags().Lookup("use-tunnel"))
}

func run(cmd *cobra.Command, args []string) error {
	project := viper.GetString("gcp.projectid")
	useTunnel := viper.GetBool("portal.use-tunnel")
	ignoreFingerprint := viper.GetBool("ssh.ignore-fingerprint")

	target, err := cmd.Flags().GetString("target")
	if err != nil {
		return err
	}

	gcpClient, err := cloud.NewGCPClient(project)
	if err != nil {
		return err
	}

	log.Info("Fetching instance list")
	instances, err := gcpClient.GetInstances(target, false)
	if err != nil {
		return err
	}

	if len(instances) == 0 {
		log.Warn("No instances found")
		return nil
	}

	for _, instance := range instances {

		if useTunnel {
			grpcConn, err := getConn()
		} else {

		}
		defer grpcConn.Close()
		c := command.NewPortalClient(grpcConn)

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		r, err := c.RunCommand(ctx, &command.Command{Name: strings.Join(args, " ")})
		if err != nil {

			if e, ok := status.FromError(err); ok {
				fmt.Printf("  %s:\n    %s\n\n", colors.Yellow(instance.Name), e.Message())
			}
			continue
		}
		fmt.Printf("  %s:\n    %s\n\n", colors.Green(instance.Name), r.GetContent())
	}

	return nil
}

func getTransport() (*grpc.ClientConn, error) {
	var grpcConn *grpc.ClientConn

	return grpcConn, nil
}

func getConn() (*grpc.ClientConn, error) {
	var grpcConn *grpc.ClientConn

	if useTunnel {
		path, err := key.Path()
		if err != nil {
			return err
		}

		k, err := key.Read(path)
		if err != nil {
			return err
		}

		log.WithField("instance", instance.Name).Debug("Using tunnel")
		if ignoreFingerprint {
			grpcConn, err = transport.SSHTransportInsecure(instance.Address, k)
		} else {
			grpcConn, err = transport.SSHTransport(instance.Address, k)
		}
		if err != nil {
			log.Errorf("%s:\n    %v\n", colors.Red(instance.Name), err)
			continue
		}
	} else {
		log.WithField("instance", instance.Name).Debug("Not using tunnel")
		grpcConn, err = transport.HTTP2Transport(instance.Address)
		if err != nil {
			log.Errorf("%s:\n    %v\n", colors.Red(instance.Name), err)
			continue
		}
	}
}

func getTunnelConn() (*grpc.ClientConn, error) {
	var grpcConn *grpc.ClientConn

	if useTunnel {
		path, err := key.Path()
		if err != nil {
			return err
		}

		k, err := key.Read(path)
		if err != nil {
			return err
		}

		log.WithField("instance", instance.Name).Debug("Using tunnel")
		if ignoreFingerprint {
			grpcConn, err = transport.SSHTransportInsecure(instance.Address, k)
		} else {
			grpcConn, err = transport.SSHTransport(instance.Address, k)
		}
		if err != nil {
			log.Errorf("%s:\n    %v\n", colors.Red(instance.Name), err)
			continue
		}
	} else {
		log.WithField("instance", instance.Name).Debug("Not using tunnel")
		grpcConn, err = transport.HTTP2Transport(instance.Address)
		if err != nil {
			log.Errorf("%s:\n    %v\n", colors.Red(instance.Name), err)
			continue
		}
	}
}
