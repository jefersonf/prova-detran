package main

import (
	"github.com/jefersonf/prova-detran/cmd"
)

// var questionSetFile = "./internal/data/questions.json"

func main() {

	cmd.Execute()

	// jsonFile, err := os.Open(questionSetFile)
	// checkErr(err)
	// defer jsonFile.Close()

	// qset, err := parser.ParseQuestionSet(jsonFile)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// fmt.Println(qset)

}
