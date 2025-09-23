# Types Parser Plugin for XJS

This plugin provides support for annotations in variables and function parameters.

> [!CAUTION]
> This plugin changes the behavior of function declarations and assignments,
> and may interfere with other plugins that also change such behavior.

## Usage

```go
import (
	"fmt"

	"github.com/xjslang/xjs/lexer"
	"github.com/xjslang/xjs/parser"
	"github.com/xjslang/types-parser"
)

func main() {
	input := `
	function printPoint(x: int, y: int) {
		console.log(x, y)
	}

	let x: int = 100
	let y: int = 200
	printPoint(x, y)`
	lb := lexer.NewBuilder()
	p := parser.NewBuilder(lb).Install(typesparser.Plugin).Build(input)
	program, err := p.ParseProgram()
	if err != nil {
		panic(fmt.Sprintf("ParseProgram() error: %q", err))
	}
	fmt.Println(program)
}
```
