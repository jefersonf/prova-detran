package cmd

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/jefersonf/prova-detran/internal"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var (
	numQuestions        int
	defaultNumQuestions int = 5

	mocktestDurationString           string
	mocktestDurationInMinutes        int
	defaultMocktestDurationInMinutes int = 5

	message string = "tempo acabou!"
)

func init() {
	simuladoCmd.Flags().IntVarP(&numQuestions, "questoes", "q", 5, "Número de questões do simulado")
	simuladoCmd.Flags().StringVarP(&mocktestDurationString, "duracao", "d", "5m", "Duração do simulado em minutos")

	rootCmd.AddCommand(simuladoCmd)
}

var simuladoCmd = &cobra.Command{
	Use:   "simulado",
	Short: "Inicia um novo simulado",
	Long:  `Inicia um novo simulado da Prova do Detran`,
	Run: func(_ *cobra.Command, _ []string) {
		validateFlags()
		startMocktest()
	},
}

func startMocktest() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(mocktestDurationInMinutes)*time.Second)
	defer cancel()
	defer func() {
		<-ctx.Done()
		log.Println(message)
	}()

	mocktest, err := internal.NewMocktest(numQuestions)
	if err != nil {
		fmt.Println(err)
		cancel()
	}

	go func() {
		for _, question := range mocktest {

			prompt := promptui.Select{
				Label: question.Statement,
				Items: question.FormattedAlternatives(),
			}

			_, _, err := prompt.Run()

			if err != nil {
				fmt.Printf("Prompt failed %v\n", err)
				cancel()
				return
			}
		}
		cancel()
		message = "fim do simulado"
	}()
}

func validateFlags() {
	if numQuestions < 1 || numQuestions >= 90 {
		numQuestions = defaultNumQuestions
		log.Printf("Número de questões deve ser maior que zero e no máximo 90, defindo número de questões para %d", numQuestions)
	}

	durationString := strings.Trim(mocktestDurationString, " mM")
	var err error
	mocktestDurationInMinutes, err = strconv.Atoi(durationString)
	if err != nil {
		mocktestDurationInMinutes = defaultMocktestDurationInMinutes
		log.Printf("Duração inválida, definindo duração para %d minutos\n", mocktestDurationInMinutes)
	}

	log.Println("Número de questões:", numQuestions)
	log.Println("Duração em minutes:", mocktestDurationInMinutes)
}
