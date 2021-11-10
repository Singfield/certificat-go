package main

import (
	"certificat-go/cert"
	"certificat-go/html"
	"certificat-go/pdf"

	"flag"
	"fmt"
	"os"
)

func main() {
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
	}

	c, err := cert.New("Golang programming", "Bob Lennon", "2021-11-10")
	if err != nil {
		fmt.Printf("Error during certificate creation :%v", err)
		os.Exit(1)
	}

	saver.Save(*c)
}
