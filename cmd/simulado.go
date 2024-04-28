package cmd

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/jefersonf/prova-detran/internal"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

const (
	warnTestSessionTimeout = "Tempo acabou!"
)

var (
	numQuestions int

	mocktestDurationString    string
	mocktestDurationInMinutes int = 5

	figures map[string]string

	ErrResultsNotSaved = errors.New("Could not store results")
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
	Message        string
	CorrectAnswers int
	TotalQuestions int
	Timestamp      time.Time
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

	sessionSummary := func(s *Status) {
		log.Printf("%s, durou %.0f minutos\n", s.Message, s.Timestamp.Sub(initialTime).Minutes())
	}

	var s Status
	select {
	case <-ctx.Done():
		log.Println(warnTestSessionTimeout)
		break
	case s = <-status:
		sessionSummary(&s)
	}

	saveResult(&s)
}

func saveResult(s *Status) {
	file, err := os.OpenFile("./data/results.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(ErrResultsNotSaved)
		return
	}
	defer file.Close()
	resultLog := fmt.Sprintf("%v correct answers: %v/%v\n", time.Now().Format(time.DateTime), s.CorrectAnswers, s.TotalQuestions)
	_, _ = file.WriteString(resultLog)
}

func showQuestions(mocktest []internal.LabeledQuestion, status chan<- Status) {
	log.Println("Simulado iniciou!")
	correctAnwsers := 0
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
			correctAnwsers += 1
		}
		if err != nil {
			status <- Status{Message: fmt.Sprintf("Falha ao carregar questão: %s\n", err), Timestamp: time.Now()}
			return
		}
	}
	correctAnswersPercentage := float32(correctAnwsers) / float32(len(mocktest)) * 100.
	status <- Status{
		Message:        fmt.Sprintf("Fim do simulado. Acertos %d/%d (%.0f%%)", correctAnwsers, len(mocktest), correctAnswersPercentage),
		CorrectAnswers: correctAnwsers,
		TotalQuestions: len(mocktest),
		Timestamp:      time.Now()}
}

func printFigureIfAny(questionStatement string) {
	words := strings.Split(questionStatement, " ")
	for _, w := range words {
		if len(w) < 2 {
			continue
		}
		if fig, exists := figures[strings.Trim(w, " ,.:;()[]'{}\"")]; exists {
			fmt.Println(fig)
		}
	}
}

func loadFigures() {
	figures = make(map[string]string)
	file, err := os.Open("./plates_in_braille_format.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	lines := make([]string, 0, 25)
	for scanner.Scan() {
		r := scanner.Text()
		if strings.HasPrefix(r, "#") {
			figures[lines[0]] = strings.Join(lines[1:], "\n")
			lines = lines[:0]
			lines = append(lines, r[1:])
		} else {
			lines = append(lines, r)
		}
	}
}

func readAndValidateParams() {
	promptNumQuestions := promptui.Prompt{
		Label: "Número de questões:",
		Templates: &promptui.PromptTemplates{
			Prompt:  "{{ . }} ",
			Valid:   "{{ . | green }} ",
			Invalid: "{{ . | red }} ",
			Success: "{{ . | bold }} ",
		},
		Validate: validateNumberOfQuestions,
	}

	promptDuration := promptui.Prompt{
		Label: "Duração do simulado em minutos:",
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
