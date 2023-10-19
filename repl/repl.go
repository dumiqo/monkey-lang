package repl

import (
	"bufio"
	"fmt"
	"io"
	"monkey-lang/compiler"
	"monkey-lang/lexer"
	"monkey-lang/parser"
	"monkey-lang/vm"
)

const PROMPT = ">>"

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()

		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()

		if len(p.Errors()) != 0 {
			printParseErrors(out, p.Errors())
			continue
		}

		comp := compiler.New()
		err := comp.Compile(program)
		if err != nil {
			fmt.Fprintf(out, "Woops! Compilation fialed:\n %s\n", err)
			continue
		}
		machine := vm.New(comp.Bytecode())
		err = machine.Run()

		if err != nil {
			fmt.Fprintf(out, "Woops! Executing bytecode fialed:\n %s\n", err)
			continue
		}

		lastPopped := machine.LastPoppedStackElem()
		io.WriteString(out, lastPopped.Inspect())
		io.WriteString(out, "\n")
	}
}

func printParseErrors(out io.Writer, errors []string) {
	for _, error := range errors {
		io.WriteString(out, "\t"+error+"\t")
	}
}
