# go-lsp-test

A programmatic language server protocol test client.

## Usage scenarios

- Debug single LSP features a language server implements by sending a reproducible set of commands
- Performance profiling of single features
- Write an E2E test suite for a language server by sending sets of commands and comparing the responses

## Example

### Test auto-completion

The following scenario will connect to a language server (implicitly set up the workspace and send `initialize`), open a document, and trigger auto-completion at a specific document position.

You'll need to provide a directory containing the files the LS should process. In the following example, the directory is called `testdata` and contains at least a 'main.tf' file.

```go
package examples

import (
	"fmt"
	"testing"

	"github.com/dbanck/go-lsp-test/client"
)

func TestCompletion(t *testing.T) {
	lsc, err := client.NewClient("localhost:8000", "testdata")
	if err != nil {
		t.Fatal(err)
	}

	doc, err := lsc.OpenDoc("main.tf", "terraform")
	if err != nil {
		t.Fatal(err)
	}

	result, err := lsc.GetCompletions(doc, client.Position{
		Line:      5,
		Character: 5,
	})
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("Got completions %#v\n", result)
}
```
