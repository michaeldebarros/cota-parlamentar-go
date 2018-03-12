package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

//Expense is the stuct for every expendature
type Expense struct {
	Nome                   string `json:"nome"`
	NumeroCarteiraDeputado string `json:"numeroCarteiraDeputado"`
	Estado                 string `json:"estado"`
	Partido                string `json:"partido"`
	TipoDespesa            string `json:"tipoDespesa"`
	DesricaoEspecifica     string `json:"desricaoEspecifica"`
	Fornecedor             string `json:"fornecedor"`
	CnpjFornecedor         string `json:"cnpjFornecedor"`
	Data                   string `json:"data"`
	Valor                  string `json:"valor"`
	NumeroRestituicao      string `json:"numeroRestituicao"`
}

func main() {
	fmt.Println("Programm Running")
	http.HandleFunc("/", handleRequest)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleRequest(w http.ResponseWriter, req *http.Request) {
	//get params from path
	p, err := url.PathUnescape(req.URL.String())
	if p == "/favicon.ico" {
		return
	}
	params := strings.Split(p, "/")
	ano := params[2]
	deputado := params[3]
	despesa := params[4]

	file, err := os.Open(fmt.Sprintf("Ano-%s.csv", ano))

	if err != nil {
		fmt.Println("There is an error here")
	}

	r := csv.NewReader(file)
	r.FieldsPerRecord = -1
	r.Comma = ';'
	r.LazyQuotes = true
	//	var deputado = "ZENAIDE MAIA"
	//	var despesa = "MANUTENÇÃO DE ESCRITÓRIO DE APOIO À ATIVIDADE PARLAMENTAR"
	var results []Expense

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		if record[0] == deputado && record[8] == despesa {
			document := Expense{record[0], record[2], record[4], record[5], record[8], record[10], record[1], record[12], record[15], record[18], record[25]}

			//push documents to results slice
			results = append(results, document)
		}
	}
	jsonResults, err := json.Marshal(results)
	if err != nil {
		fmt.Print("Error while marchalling the restults to json")
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResults)
}
