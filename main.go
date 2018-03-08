package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	file, err := os.Open("Ano-2016.csv")
	
	if err != nil {
		fmt.Println("There is an error here")
	}

	r := csv.NewReader(file)
	r.FieldsPerRecord = -1
	r.Comma = ';'
	r.LazyQuotes = true
	var deputado string = "ZECA DIRCEU"
	var despesa string = "DIVULGAÇÃO DA ATIVIDADE PARLAMENTAR."

	for  {
		record, err := r.Read()

		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		if record[0] == deputado && record[8] == despesa{
			fmt.Printf("\nDeputado: %s, Tipo de despesa: %s, Fornecedor: %s, CPF/CNPJ: %s, Data: %s, Valor: %s\n", record[0], record[8], record[11], record[12], record[15], record[18])
		}	
	}
}
