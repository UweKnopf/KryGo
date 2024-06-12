package krygo

import (
	"fmt"

	"github.com/krygo/pkg/krygo"

	"github.com/spf13/cobra"
)

var caesarCmd = &cobra.Command{
	Use:     "caesar [shift] String",
	Aliases: []string{"cae"},
	Short:   "Use the Caesar cipher",
	Args:    cobra.RangeArgs(1, 2),
	Run: func(cmd *cobra.Command, args []string) {

		i := args[0]
		res, kind := krygo.Inspect(i, onlyDigits)

		pluralS := "s"
		if res == 1 {
			pluralS = ""
		}
		fmt.Printf("'%s' has a %d %s%s.\n", i, res, kind, pluralS)
	},
}

func init() {
	caesarCmd.Flags().BoolVarP(&onlyDigits, "digits", "d", false, "Count only digits")
	rootCmd.AddCommand(caesarCmd)
}
