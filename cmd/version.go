package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Imprime a versão do Prova Detran",
	Long:  `Imprime a versão do Prova Detran, uma CLI para realização de simulados da prova do Detran`,
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Println("Prova Detran v0.1")
	},
}
