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

func getFileOrExit(name string) *os.File {
	file, err := os.Open(name)

	if err != nil {
		message := "File " + name + " can not be opened:\n" + err.Error()
		logErrorAndExit(message)
	}

	return file
}

func getFileOrCreate(name string) *os.File {
	file, err := os.Open(name)

	if err == nil {
		return file
	}

	defer file.Close()

	createdFile, createdErr := os.Create(name)

	if createdErr != nil {
		logErrorAndExit(createdErr.Error())
	}

	return createdFile
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
		reader = getFileOrExit(*inputFile)
	}

	if *outputFile == "" {
		writer = os.Stdout
	} else {
		writer = getFileOrCreate(*outputFile)
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
