package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(resultadoCmd)
}

var resultadoCmd = &cobra.Command{
	Use:   "resultado",
	Short: "Mostra o resultado do simulado",
	Long:  `Mostra o resultado do último simulado realizado`,
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Println("resultado do simulado")
	},
}
