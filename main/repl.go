package main

import (
	"bufio"
	"fmt"
	"github.com/leluxnet/carbon/env"
	"github.com/leluxnet/carbon/eval"
	"os"
)

const PROMPT = ">>> "

func Repl(e *env.Env) {
	scanner := bufio.NewScanner(os.Stdin)

	path, err := os.Getwd()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	for {
		fmt.Print(PROMPT)

		scanner.Scan()
		text := scanner.Text()

		eval.Run(text, e, true, "", path)
	}
}
