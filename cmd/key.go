package cmd

import (
	"fmt"

	"speedrun/utils"

	"github.com/spf13/cobra"
)

var keyCmd = &cobra.Command{
	Use:   "key",
	Short: "Manage ssh keys",
	Args:  cobra.ExactArgs(1),
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		err := utils.ConfigInitialized()
		return err
	},
}

var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Generates a new ssh key",
	RunE:  new,
}

var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Shows private key",
	Run:   show,
}

func init() {
	rootCmd.AddCommand(keyCmd)
	keyCmd.AddCommand(newCmd)
	keyCmd.AddCommand(showCmd)
}

func new(cmd *cobra.Command, args []string) error {
	fmt.Println("generated new ssh key")
	_, _, err := utils.GenerateKeyPair()
	if err != nil {
		return err
	}
	return nil
}

func show(cmd *cobra.Command, args []string) {
	fmt.Println("showing private key")
}
