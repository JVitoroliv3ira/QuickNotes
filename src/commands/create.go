package commands

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"quick_notes/src/services"
	"quick_notes/src/types"
	"quick_notes/src/utils"
	"strings"
)

var create = &cobra.Command{
	Use:   "create",
	Short: "Comando responsável por cadastrar uma nova nota.",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println(utils.GetErrorColor("Por favor, informe o conteúdo da nota."))
			os.Exit(1)
		}
		_, err := services.Create(types.Note{Text: strings.Join(args, " ")})
		if err != nil {
			fmt.Println(utils.GetErrorColor("Ocorreu um erro inesperado durante a criação da nota: " + err.Error()))
			os.Exit(1)
		}
		fmt.Println(utils.GetSuccessColor("Nota cadastrada com sucesso!"))
	},
}

func GetCreateCommand() *cobra.Command {
	return create
}
