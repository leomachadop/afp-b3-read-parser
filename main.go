package main

import (
	"b3-read-parser/processor"
	"b3-read-parser/utils"
	_ "github.com/joho/godotenv"
	"log"
	"os"
)

func init() {
	configureLog()
}

func main() {

	args := utils.RemoveFirstPostion(os.Args)

	if len(args) == 0 {
		log.Println("[main] Nenhum parâmetro informado")
	}

	log.Printf("[main] Args %v\n", args)
	/**
	  Adicionar goroutines
	*/
	for _, file := range args {
		processor.ProcessB3(file)
	}

}

/**
Método que pode configurar o log em qualquer local
*/
func configureLog() {
	/*file, err := os.OpenFile("log/log_"+time.Now().UTC().Format("02012006_150405.999")+".log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(file)*/
}
