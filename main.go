package main

import (
	"certificat-go/cert"
	"certificat-go/csv"
	"certificat-go/html"
	"certificat-go/pdf"

	"flag"
	"fmt"
	"os"
)

func main() {
	file := flag.String("file", "","CSV file input")
	outputType := flag.String("type", "pdf", "Output type of the certificate.")
	flag.Parse()

	var saver cert.Saver
	var err error

	switch *outputType {
	case "pdf":
		saver, err = pdf.New("output")
	case "html":
		saver, err = html.New("output")
	default:
		fmt.Printf("Unknown output type. got %v\n", *outputType)
	}

	if err != nil {
		fmt.Printf("Error, could not generate: %v", err)
		os.Exit(1)
	}

	
	certs, err := csv.ParseCSV(*file)
	if err != nil {
		fmt.Printf("Could not parse csv file %v", err)
		os.Exit(1)
	}
	for _, c :=range certs {
		err = saver.Save(*c)
		if err != nil {
			fmt.Printf("Could not save cert. got %v\n", err)
		}
	}
}
