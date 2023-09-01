package commands

import (
	"fmt"
	"github.com/spf13/cobra"
)

var cp = &cobra.Command{
	Use:   "cp",
	Short: "Comando responsável por copiar uma nota para a área de transferência.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello, World")
	},
}

func GetCopyCommand() *cobra.Command {
	return cp
}
