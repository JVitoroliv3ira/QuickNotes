package commands

import (
	"fmt"
	"github.com/spf13/cobra"
)

var clear = &cobra.Command{
	Use:   "clear",
	Short: "Comando responsável por limpar todas as notas",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello, World!")
	},
}

func GetClearCommand() *cobra.Command {
	return clear
}
