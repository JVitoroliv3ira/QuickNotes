package commands

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"quick_notes/src/utils"
)

var remove = &cobra.Command{
	Use:   "rm",
	Short: "Comando responsável por deletar uma nota.",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println(utils.GetErrorColor("Informe o identificador da nota que você deseja deletar."))
			os.Exit(1)
		}
	},
}

func GetRemoveCommand() *cobra.Command {
	return remove
}
