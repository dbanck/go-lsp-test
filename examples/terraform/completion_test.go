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
