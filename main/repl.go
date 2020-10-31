package main

import (
	"bufio"
	"fmt"
	"github.com/leluxnet/carbon/env"
	"os"
)

const PROMPT = ">>> "

func Repl(e *env.Env) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print(PROMPT)

		scanner.Scan()
		text := scanner.Text()

		run(text, e, true)
	}
}
