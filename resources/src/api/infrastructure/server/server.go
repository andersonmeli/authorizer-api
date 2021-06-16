package server

import (
	"bufio"
	"github.com/mercadolibre/authorizer-api/resources/src/api/application/controller"
	"os"
)

func Start() {
	//message := `{"account": {"active-card": true, "available-limit": 100}}`
	//var message string
	var messages []string
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		messages = append(messages, scanner.Text())
		controller.ProcessOperations(messages)
		break
	}
}