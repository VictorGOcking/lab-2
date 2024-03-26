package main

import (
	"flag"
	lab2 "github.com/VictorGOcking/lab-2"
	"io"
	"os"
	"strings"
)

var (
	inputExpression = flag.String("e", "", "Expression to compute")
	inputFile       = flag.String("f", "", "File to get expression from")
	outputFile      = flag.String("o", "", "File to set the result")
)

func logErrorAndExit(message string) {
	os.Stderr.WriteString(message + "\n")

	os.Exit(1)
}

func main() {
	flag.Parse()

	if *inputExpression != "" && *inputFile != "" {
		logErrorAndExit("Expression source must be only one")
	}

	var reader io.Reader
	var writer io.Writer

	if *inputExpression != "" {
		reader = strings.NewReader(*inputExpression)
	} else {
		file, err := os.Open(*inputFile)

		defer file.Close()

		if err != nil {
			message := "File " + *inputFile + " can not be opened:\n" + err.Error()
			logErrorAndExit(message)
		}

		reader = file
	}

	if *outputFile == "" {
		writer = os.Stdout
	} else {
		file, err := os.Open(*outputFile)

		defer file.Close()

		if err == nil {
			writer = file
		}

		createdFile, createdErr := os.Create(*outputFile)

		if createdErr != nil {
			logErrorAndExit(createdErr.Error())
		}

		writer = createdFile
	}

	handler := &lab2.ComputeHandler{
		Reader: reader,
		Writer: writer,
	}

	err := handler.Compute()

	if err != nil {
		logErrorAndExit(err.Error())
	}
}
