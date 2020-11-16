package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"speedrun/helpers"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// var verbosity string

var rootCmd = &cobra.Command{
	Use:   "speedrun",
	Short: "Execute commands at scale",
	Long:  `Speedrun executes commands at scale`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// rootCmd.PersistentFlags().StringVarP(&verbosity, "verbosity", "v", log.InfoLevel.String(), "Log level (debug, info, warn, error, fatal, panic)")
	// viper.BindPFlag("verbosity", rootCmd.PersistentFlags().Lookup("verbosity"))
}

func initConfig() {
	home, err := homedir.Dir()
	if err != nil {
		helpers.Error(err.Error())
	}

	configDir := filepath.Join(home, ".config", "speedrun")
	viper.SetConfigName("config.toml")
	viper.SetConfigType("toml")
	viper.AddConfigPath(configDir)

	viper.ReadInConfig()
	// setUpLogs(viper.GetString("verbosity"))
}

// func setUpLogs(level string) error {
// 	lvl, err := log.ParseLevel(level)
// 	if err != nil {
// 		return err
// 	}
// 	log.SetLevel(lvl)
// 	// log.SetOutput()
// 	formatter := &prefixed.TextFormatter{
// 		DisableTimestamp: true,
// 		FullTimestamp:    true,
// 	}
// 	// formatter.SetColorScheme(&prefixed.ColorScheme{
// 	// 	PrefixStyle: "cyan",
// 	// })
// 	log.SetFormatter(formatter)
// 	return nil
// }
