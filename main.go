package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)


//make struct called Expense

type Expense struct {
	Nome string
	NumeroCarteiraDeputado string
	Estado string
	Partido string
	TipoDespesa string
	DesricaoEspecifica string
	Fornecedor string
	CnpjFornecedor string
	Data string
	Valor string
	NumeroRestituicao string
}

//every record is going to be an expense instance
//json encode the instances
//make slice of json encoded intances
var results []Expense


//encode the slice in json format
//send the slice

func main() {
	file, err := os.Open("Ano-2016.csv")
	
	if err != nil {
		fmt.Println("There is an error here")
	}

	r := csv.NewReader(file)
	r.FieldsPerRecord = -1
	r.Comma = ';'
	r.LazyQuotes = true
	var deputado string = "ZENAIDE MAIA"
	var despesa string = "MANUTENÇÃO DE ESCRITÓRIO DE APOIO À ATIVIDADE PARLAMENTAR"

	for  {
		record, err := r.Read()

		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		if record[0] == deputado && record[8] == despesa{
			fmt.Printf("\nDeputado: %s, Tipo de despesa: %s, Fornecedor: %s, CPF/CNPJ: %s, Data: %s, Valor: %s\n", record[0], record[8], record[11], record[12], record[15], record[18]);
			document := Expense{record[0], record[2], record[4],record[5], record[8], record[10], record[11], record[12], record[15], record[18], record[25]}
			fmt.Print(document)
		}	
	}
}
//headers: [
//	'nome', 
//	,
//	'numeroCarteiraDeputado',
//	,
//	'estado',
//	'partido',
//	,
//	,
//	'tipoDespesa',
//	,
//	'desricaoEspecifica',
//	'fornecedor',
//	'cnpjFornecedor',
//	,
//	,
//	'data',
//	,
//	,
//	'valor',
//	,
//	,
//	,
//	,
//	,
//	,
//	'numeroRestituicao',
//	,
//	,
//	,
//],
//