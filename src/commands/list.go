package commands

import (
	"fmt"
	"github.com/spf13/cobra"
)

var list = &cobra.Command{
	Use:   "ls",
	Short: "Comando responsável por listar as notas cadastradas.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello, World!")
	},
}

func GetListCommand() *cobra.Command {
	return list
}
