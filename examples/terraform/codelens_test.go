package examples

import (
	"log"
	"path/filepath"
	"testing"
	"time"

	"github.com/dbanck/go-lsp-test/client"
)

func TestCodelens(t *testing.T) {
	testDir, err := filepath.Abs("testdata")
	if err != nil {
		t.Fatal(err)
	}
	lsc, err := client.NewClient("localhost:3333", testDir)
	if err != nil {
		t.Fatal(err)
	}
	defer lsc.Close()

	doc, err := lsc.OpenDoc("codelens.tf", "terraform")
	if err != nil {
		t.Fatal(err)
	}

	time.Sleep(10 * time.Second)

	result, err := lsc.GetCodeLens(doc)
	if err != nil {
		t.Fatal(err)
	}

	log.Printf("Got codeLens %#v\n", result)
}
