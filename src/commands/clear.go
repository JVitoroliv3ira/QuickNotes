package commands

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"quick_notes/src/services"
	"quick_notes/src/utils"
)

var clear = &cobra.Command{
	Use:   "clear",
	Short: "Comando responsável por limpar todas as notas",
	Run: func(cmd *cobra.Command, args []string) {
		err := services.DeleteAll()
		if err != nil {
			fmt.Println(utils.GetErrorColor("Ocorreu um erro ao deletar as notas: " + err.Error()))
			os.Exit(1)
		}
		fmt.Println(utils.GetSuccessColor("Notas deletadas com sucesso!"))
	},
}

func GetClearCommand() *cobra.Command {
	return clear
}
