package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"quick_notes/src/commands"
)

var rootCmd = &cobra.Command{
	Use:   "",
	Short: "Uma ferramenta para gerenciar notas rápidas.",
}

func main() {
	rootCmd.AddCommand(commands.GetCreateCommand(), commands.GetListCommand())
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
