# go-lsp-test

A programmatic language server protocol test client.

## Usage scenarios

- Debug single LSP features by sending a reproducible set of commands
- Performance profiling of single features
- Write an E2E test suite for a language server by sending commands and comparing the responses

## Future ideas

- REPL

## Example

### Test auto-completion

The following scenario will connect to a language server running on port `3333` (implicitly set up the workspace and send `initialize`), open a document, and trigger auto-completion at a specific document position.

You'll need to provide a directory containing the files the LS should process. In the following example, the directory is called `testdata` and contains at least a 'main.tf' file.

```go
package examples

import (
	"log"
	"path/filepath"
	"testing"

	"github.com/dbanck/go-lsp-test/client"
	p "github.com/dbanck/go-lsp-test/protocol"
)

func TestCompletion(t *testing.T) {
	testDir, err := filepath.Abs("testdata")
	if err != nil {
		t.Fatal(err)
	}
	lsc, err := client.NewClient("localhost:3333", testDir)
	if err != nil {
		t.Fatal(err)
	}
	defer lsc.Close()

	doc, err := lsc.OpenDoc("main.tf", "terraform")
	if err != nil {
		t.Fatal(err)
	}

	result, err := lsc.GetCompletions(doc, p.Position{
		Line:      5,
		Character: 5,
	})
	if err != nil {
		t.Fatal(err)
	}

	log.Printf("Got completions %#v\n", result)
}
```
