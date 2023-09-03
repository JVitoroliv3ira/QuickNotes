package commands

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"quick_notes/src/services"
	"quick_notes/src/utils"
	"quick_notes/src/views"
)

var list = &cobra.Command{
	Use:   "ls",
	Short: "Comando responsável por listar as notas cadastradas.",
	Run: func(cmd *cobra.Command, args []string) {
		list, err := services.ListAll()
		if err != nil {
			fmt.Println(utils.GetErrorColor("Ocorreu um erro inesperado durante a listagem das notas."))
			os.Exit(1)
		}

		views.Render(list)
	},
}

func GetListCommand() *cobra.Command {
	return list
}
