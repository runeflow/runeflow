package util

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// PromptString writes the given prompt to standard output and reads a line
// into the destination until a non-empty string is read.
func PromptString(dest *string, prompt string) {
	r := bufio.NewReader(os.Stdin)
	for *dest == "" {
		fmt.Print(prompt)
		if text, err := r.ReadString('\n'); err == nil {
			*dest = strings.TrimSpace(text)
		}
	}
}
