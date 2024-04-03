package cmd

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jefersonf/prova-detran/internal"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var (
	numQuestions int

	mocktestDurationString    string
	mocktestDurationInMinutes int = 5

	figures map[string]string
)

var simuladoCmd = &cobra.Command{
	Use:   "simulado",
	Short: "Inicia um novo simulado",
	Long:  `Inicia um novo simulado da Prova do Detran`,
	Run: func(_ *cobra.Command, _ []string) {
		readAndValidateParams()
		startMocktest()
	},
}

type Status struct {
	Message   string
	Timestamp time.Time
}

func init() {
	simuladoCmd.Flags().IntVarP(&numQuestions, "questoes", "q", 5, "Número de questões do simulado")
	simuladoCmd.Flags().StringVarP(&mocktestDurationString, "duracao", "d", "5m", "Duração do simulado em minutos")
	rootCmd.AddCommand(simuladoCmd)
	loadFigures()
}

func startMocktest() {
	mocktest, err := internal.NewMocktest(numQuestions)
	status := make(chan Status)

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(mocktestDurationInMinutes)*time.Minute)
	defer cancel()

	if err != nil {
		fmt.Println(err)
		cancel()
	}

	initialTime := time.Now()
	go showQuestions(mocktest, status)

	select {
	case <-ctx.Done():
		log.Println("Tempo acabou!")
	case s := <-status:
		log.Printf("%s, durou %.0f minutos\n", s.Message, s.Timestamp.Sub(initialTime).Minutes())
	}
}

func showQuestions(mocktest []internal.LabeledQuestion, status chan<- Status) {
	log.Println("Simulado iniciou!")
	rightAnwsers := 0
	for i, question := range mocktest {
		printFigureIfAny(question.Statement)
		prompt := promptui.Select{
			Label: fmt.Sprintf("(Questão %0d) %s", i+1, question.Statement),
			Items: question.FormattedAlternatives(i + 1),
			Templates: &promptui.SelectTemplates{
				Label:    "{{ . | bold | yellow }}:",
				Active:   "{{ .Text | cyan }} ",
				Inactive: "{{ .Text }}",
				Selected: "{{ .QuestionID | bold | yellow }} {{ .Text | bold | blue }} ",
			},
		}
		choiceIndex, _, err := prompt.Run()
		if indexToOptionLabel(choiceIndex) == question.RightAnswer {
			rightAnwsers += 1
		}
		if err != nil {
			status <- Status{Message: fmt.Sprintf("Falha ao carregar questão: %s\n", err), Timestamp: time.Now()}
			return
		}
	}
	status <- Status{Message: fmt.Sprintf("Fim do simulado. Acertos %d/%d", rightAnwsers, len(mocktest)), Timestamp: time.Now()}
}

func printFigureIfAny(_ string) {
	// f, err := os.Open("./out.text")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer f.Close()

	// b, err := io.ReadAll(f)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// fmt.Println(string(b))
}

func loadFigures() {
	figures = make(map[string]string)
}

func readAndValidateParams() {
	promptNumQuestions := promptui.Prompt{
		Label: "Número de questões",
		Templates: &promptui.PromptTemplates{
			Prompt:  "{{ . }} ",
			Valid:   "{{ . | green }} ",
			Invalid: "{{ . | red }} ",
			Success: "{{ . | bold }} ",
		},
		Validate: validateNumberOfQuestions,
	}

	promptDuration := promptui.Prompt{
		Label: "Duração do simulado em minutos",
		Templates: &promptui.PromptTemplates{
			Prompt:  "{{ . }} ",
			Valid:   "{{ . | green }} ",
			Invalid: "{{ . | red }} ",
			Success: "{{ . | bold }} ",
		},
		Validate: validateDuration,
	}

	_, _ = promptNumQuestions.Run()
	_, _ = promptDuration.Run()
}
