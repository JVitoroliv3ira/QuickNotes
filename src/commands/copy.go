package commands

import (
	"fmt"
	"github.com/atotto/clipboard"
	"github.com/spf13/cobra"
	"os"
	"quick_notes/src/repositories"
	"quick_notes/src/utils"
)

var cp = &cobra.Command{
	Use:   "cp",
	Short: "Comando responsável por copiar uma nota para a área de transferência.",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println(utils.GetErrorColor("Informe o identificador da nota que você deseja copiar."))
			os.Exit(1)
		}
		note, err := repositories.FindById(args[0])
		if err != nil {
			panic(err)
		}
		err = clipboard.WriteAll(note.Text)
		if err != nil {
			panic(err)
		}
		fmt.Println(utils.GetSuccessColor("Nota copiada para a área de transferência com sucesso"))
	},
}

func GetCopyCommand() *cobra.Command {
	return cp
}
