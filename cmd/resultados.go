package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

const (
	warnNoResultsToShow = "Nenhum resultado para mostrar"
)

func init() {
	rootCmd.AddCommand(resultadoCmd)
}

var resultadoCmd = &cobra.Command{
	Use:   "resultados",
	Short: "Mostra o resultado do simulado",
	Long:  `Mostra o resultado do último simulado realizado`,
	Run: func(_ *cobra.Command, _ []string) {
		bytes, err := os.ReadFile("./data/results.txt")
		if err != nil {
			fmt.Println(warnNoResultsToShow)
		} else {
			fmt.Println(string(bytes))
		}
	},
}
