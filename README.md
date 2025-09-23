# Types Parser Plugin for XJS

This plugin provides support for annotations in variables and function parameters.

> [!CAUTION]
> This plugin changes the behavior of function declarations and assignments,
> and may interfere with other plugins that also change such behavior.

## Usage

```go
import (
  "github.com/xjslang/xjs/parser"
  "github.com/xjslang/pow-parser"
)

func main() {
  p := parser.NewParser()
  powparser.InstallPlugin(p)
  // Now the parser supports '**' as a power operator.
}
```