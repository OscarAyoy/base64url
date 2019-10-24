package main

import (
	"bufio"
	"encoding/base64"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	decode, inFileName, outFileName := parseFlags()

	inFile, err := openInputFile(inFileName)
	defer inFile.Close()
	checkError(err)

	outFile, err := openOutputFile(outFileName)
	defer outFile.Close()
	checkError(err)

	input, err := readInputFile(inFile)
	checkError(err)

	output, err := processInput(decode, input)
	checkError(err)

	err = writeOutputFile(outFile, output)
	checkError(err)
}

func parseFlags() (decode bool, inFile string, outFile string) {
	flag.BoolVar(&decode, "d", false, "decodes input (default is to encode)")
	flag.StringVar(&inFile, "i", "-", "input file or \"-\" for stdin")
	flag.StringVar(&outFile, "o", "-", "output file or \"-\" for stdout")

	flag.Parse()

	return decode, inFile, outFile
}

func openInputFile(fileName string) (*os.File, error) {
	if fileName == "-" {
		return os.Stdin, nil
	}

	return os.Open(fileName)
}

func openOutputFile(fileName string) (*os.File, error) {
	if fileName == "-" {
		return os.Stdout, nil
	}

	return os.Create(fileName)
}

func readInputFile(file *os.File) ([]byte, error) {
	reader := bufio.NewReader(file)
	var output []byte

	for {
		input, err := reader.ReadByte()
		if err != nil {
			if err == io.EOF {
				// We've reached end-of-file.
				break
			}

			// Something went wrong.
			return nil, fmt.Errorf("readInputFile: %v", err)
		}

		// TODO: This constant re-allocation can't be efficient.
		output = append(output, input)
	}

	return output, nil
}

func processInput(decode bool, input []byte) (output []byte, err error) {
	if decode {
		output, err = decodeInput(input)
		if err != nil {
			return nil, fmt.Errorf("processInput: %v", err)
		}
	} else {
		output = encodeInput(input)
	}

	return output, nil
}

func decodeInput(input []byte) ([]byte, error) {
	output := make([]byte, base64.RawURLEncoding.DecodedLen(len(input)))

	outputBytes, err := base64.RawURLEncoding.Decode(output, input)
	if err != nil {
		return nil, fmt.Errorf("decodeInput: %v", err)
	}

	// Return the actual number of decoded bytes.
	return append([]byte(nil), output[:outputBytes]...), nil
}

func encodeInput(input []byte) []byte {
	output := make([]byte, base64.RawURLEncoding.EncodedLen(len(input)))
	base64.RawURLEncoding.Encode(output, input)

	return output
}

func writeOutputFile(outFile *os.File, output []byte) error {
	_, err := outFile.Write(output)

	if err != nil {
		return fmt.Errorf("writeOutputFile: %v", err)
	}

	return nil
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
