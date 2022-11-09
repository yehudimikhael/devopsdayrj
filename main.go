//1 Todo código em go vai começar com esse main.go

package main

import (
	"time"
	"fmt" // analogo ao printf e scanf em C
	vegeta "github.com/tsenart/vegeta/v12/lib"
)

func main(){
	//Load Test, Tes Carga , usar para fazer alguns request

	rate := vegeta.Rate{Freq: 100, Per: time.Second} // quantos request você quer fazer, fazer 100 request a cada segundo
	duration := 5 * time.Second  //5s por segundo
	targeter := vegeta.NewStaticTargeter(vegeta.Target{   //target seria uma struct, um objeto
		Method: "GET",
		URL: "https://google.com.br",
	})  //Bater na URL 100 vezes por segundo durante 5s 
	
	attacker := vegeta.NewAttacker()
	var metrics vegeta.Metrics 
	for res := range attacker.Attack(targeter, rate, duration, "Attack!"){
		metrics.Add(res)
	}
	metrics.Close()

	fmt.Printf("99th percentile: %s\n", metrics.Latencies.P99) //quanto tempo o google respondeu 

	// Go mod init vai inicializar o arquivo parecido com o requirement do python
	// Caso ocorra dele não reconhecer o modulo como um pacote do GO não consegue baixar 
	// $ para criar o arquivo `go mod init`vai criar o arquivo go.mod
	// e para baixar go get github.com/tsenart/vegeta/v12/lib - vai inicializar o nosso pacote e vai criar o arquivo go.mod

	// Para gerar compilar o código em GO -> build -> go build
	//para executar o binario -> go 
	// wget link do binario 
	// Para rodar só chamar o binario "/devopsdayrj"

}