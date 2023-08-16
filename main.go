package main

import (
	"fmt"
	"time"
)

func enviarEmail(message string) {
	go func() {
		time.Sleep(time.Millisecond * 250)
		fmt.Printf("Email recibido: '%s'\n", message)
	}()
	fmt.Printf("Email enviado: '%s'\n", message)
}

func test(message string) {
	enviarEmail(message)
	time.Sleep(time.Millisecond * 500)
	fmt.Println("========================")
}

func main() {
	test("Hola Luciano!")
	test("Que onda Facu!")
	test("Hola Martin!")
}
