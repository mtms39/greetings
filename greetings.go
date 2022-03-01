package greetings()

import "fmt"

func Hello(name string) string {
	// dois pontos antes de igual pra declarar e poder começar a usar a variável em uma só linha
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}
