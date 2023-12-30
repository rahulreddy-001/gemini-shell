package ui

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"rahul.dev/go/gemini-shell/format"
)

// Read prompt from the terminal
func ReadPrompt() string {
	fmt.Print("\033[1;34mYou\033[0m ")
	reader := bufio.NewReader(os.Stdin)

	if msg, err := reader.ReadString('\n'); err != nil {
		panic(fmt.Sprintf("Error reading input:,%s", err))
	} else {
		fmt.Print("\n\033[1;32mModel\033[0m\n\n")
		return msg
	}
}

// Delay before printing every character
func LazyPrint(done bool, part string, fs *format.FormatSelector) {
	for _, c := range part {
		fs.PrintFormat(c)
		time.Sleep(14 * time.Millisecond)
	}

	if done {
		fs.Flush()
		fs.Close()
	}
}
