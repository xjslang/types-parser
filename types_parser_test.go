package typesparser

import (
	"fmt"
	"testing"

	"github.com/xjslang/xjs/lexer"
	"github.com/xjslang/xjs/parser"
)

func TestXxx(t *testing.T) {
	input := `
	function printPoint(x: int, y: int) {
		console.log(x, y)
	}

	let x: int = 100
	let y: int = 200
	printPoint(x, y)`
	lb := lexer.NewBuilder()
	p := parser.NewBuilder(lb).Install(Plugin).Build(input)
	program, err := p.ParseProgram()
	if err != nil {
		t.Errorf("ParseProgram() error: %q", err)
	}
	fmt.Println(program)
}
