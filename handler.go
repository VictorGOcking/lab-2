package lab2

import (
	"io"
)

// ComputeHandler should be constructed with input io.Reader and output io.Writer.
// Its Compute() method should read the expression from input and write the computed result to the output.
type ComputeHandler struct {
	Reader io.Reader
	Writer io.Writer
}

func (ch *ComputeHandler) Compute() error {
	buffer := make([]byte, 1024)
	var content []byte

	for {
		n, err := ch.Reader.Read(buffer)
		content = append(content, buffer[:n]...)

		// Ending reading
		if err == io.EOF {
			break
		}

		if err != nil {
			return err
		}
	}

	converted, err := PostfixToPrefix(string(content))
	if err != nil {
		return err
	}

	_, err = ch.Writer.Write([]byte(converted))
	return err
}
