package repl

import (
	"bufio"
	"fmt"
	"io"
	"github.com/kakts/monkey/token"
	"github.com/kakts/monkey/lexer"
)

// replの先頭文字
const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		// 入力行の内容取得
		line := scanner.Text()
		l := lexer.New(line)

		// EOFがくるまで入力内容を表示
		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}