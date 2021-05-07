# Pretty Print Trees in Go

[![Go Reference](https://pkg.go.dev/badge/github.com/need-being/go-tree.svg)](https://pkg.go.dev/github.com/need-being/go-tree)

## Installation

```
go get github.com/need-being/go-tree
```

## Basic Usage

The following code

```go
package main

import "github.com/need-being/go-tree"

func main() {
	root := tree.New("root")

	str := root.Add("strings")
	str.Add("foo")
	str.Add("bar")

	num := root.Add("numbers")
	floats := num.Add("floats")
	floats.Add(0.5)
	integer := num.Add("integers")
	integer.Add(0)
	integer.Add(42)

	slices := root.Add("slices")
	slices.Add([]string{"hello", "world"})
	slices.Add([]interface{}{42, 0.5, "QED"})

	tree.Print(root)
}
```

outputs

```
root
├── strings
│   ├── foo
│   └── bar
├── numbers
│   ├── floats
│   │   └── 0.5
│   └── integers
│       ├── 0
│       └── 42
└── slices
    ├── [hello world]
    └── [42 0.5 QED]
```

## Advanced Usage

The following code

```go
package main

import (
	"fmt"
	"os"

	"github.com/need-being/go-tree"
)

func main() {
	root := tree.New("root")

	root.AddPathString("strings/foo")
	root.AddPathString("strings/bar")

	root.AddPath("numbers", "floats", 0.5)
	root.AddPath("numbers", "integers", 0)
	root.AddPath("numbers", "integers", 42)

	root.Add("QED")

	printer := tree.NewPrinter(os.Stdout, func(n *tree.Node) string {
		if n.Virtual {
			return fmt.Sprintf("%v (virtual)", n.Value)
		}
		return fmt.Sprint(n.Value)
	})
	printer.Print(root)
}
```

outputs

```
root
├── strings (virtual)
│   ├── foo
│   └── bar
├── numbers (virtual)
│   ├── floats (virtual)
│   │   └── 0.5
│   └── integers (virtual)
│       ├── 0
│       └── 42
└── QED
```
