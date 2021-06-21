package server

import (
	"bufio"
	"github.com/authorizer-api/src/api/application/controller"
	"os"
)

func Start() {
	//message := `{"account": {"active-card": true, "available-limit": 100}}`
	//var message string
	var messages []string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		message := scanner.Text()
		messages = append(messages, message)

		if len(message) == 0 {
			// exit if user entered an empty string
			break
		}
	}

	controller.ProcessOperations(messages)
}