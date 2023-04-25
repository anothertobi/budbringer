package internal

import (
	"bufio"
	"fmt"
	"os"
)

// ReadStdin writes the inputs from stdin to a chan
func ReadStdin(input chan []byte) {
	reader := bufio.NewReader(os.Stdin)
	for {
		in, err := reader.ReadBytes('\n')
		if err != nil {
			fmt.Printf("error reading from stdin: %v\n", err)
		}
		input <- in
	}
}
