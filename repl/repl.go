package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/Lentanta/FantasyScript/lexer"
	"github.com/Lentanta/FantasyScript/token"
)

const PROMPT = "fts> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Print(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		if line == "exit" {
			break
		}
		lex := lexer.New(line)

		for tok := lex.NextToken(); tok.Type != token.EOF; tok = lex.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
