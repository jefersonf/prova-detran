package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "provadetran",
	Short: "Prova Detran é uma CLI que te permite realizar simulado da prova do Detran",
	Long: `Estude para o exame teórico do Detran por meio do Prova Detran,
	uma ferramenta de linha de comando que gera simulados da prova 
	a partir de questões reais e para todos os conteúdos abordados.
	Documentação complete está disponível em github.com/jefersonf/prova-detran`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalln(err)
	}
}
