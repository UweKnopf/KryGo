package krygo

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var version = "0.0.1"

var rootCmd = &cobra.Command{
	Use:     "krygo",
	Version: version,
	Short:   "krygo - a simple CLI to encode strings with basic ciphers",
	Long: `krygo is a cli to allow you to encrypt strings using basic encryption methods.
   
Put a string after the encryption command to encrypt it`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing your CLI '%s'", err)
		os.Exit(1)
	}
}
