package util

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// PromptString writes the given prompt to standard output and reads a line
// into the destination until a non-empty string is read.
func PromptString(prompt string) string {
	r := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(prompt)
		if text, err := r.ReadString('\n'); err == nil {
			if line := strings.TrimSpace(text); len(line) > 0 {
				return line
			}
		}
	}
}
