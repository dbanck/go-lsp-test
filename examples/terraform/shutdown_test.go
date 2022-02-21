package examples

import (
	"path/filepath"
	"testing"

	"github.com/dbanck/go-lsp-test/client"
)

func TestShutdown(t *testing.T) {
	testDir, err := filepath.Abs("testdata")
	if err != nil {
		t.Fatal(err)
	}
	lsc, err := client.NewClient("localhost:3333", testDir)
	if err != nil {
		t.Fatal(err)
	}

	lsc.Close()
}
