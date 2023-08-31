package commands

import (
	"fmt"
	"github.com/spf13/cobra"
)

var create = &cobra.Command{
	Use:   "create",
	Short: "Comando responsável por cadastrar uma nova nota",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello, World")
	},
}

func GetCreateCommand() *cobra.Command {
	return create
}
