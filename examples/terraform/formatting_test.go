package examples

import (
	"path/filepath"
	"testing"
	"time"

	"github.com/dbanck/go-lsp-test/client"
	p "github.com/dbanck/go-lsp-test/protocol"
	"github.com/google/go-cmp/cmp"
)

func TestFormatting(t *testing.T) {
	testDir, err := filepath.Abs("format")
	if err != nil {
		t.Fatal(err)
	}
	lsc, err := client.NewClient("localhost:3333", testDir)
	if err != nil {
		t.Fatal(err)
	}
	defer lsc.Close()

	doc, err := lsc.OpenDoc("formatting.tf", "terraform")
	if err != nil {
		t.Fatal(err)
	}

	time.Sleep(3 * time.Second)

	result, err := lsc.Format(doc)
	if err != nil {
		t.Fatal(err)
	}

	expectedEdits := []p.TextEdit{
		{
			Range:   p.Range{Start: p.Position{Line: 1}, End: p.Position{Line: 2}},
			NewText: "  environment {\n",
		},
	}

	if diff := cmp.Diff(expectedEdits, result); diff != "" {
		t.Fatalf("formatting mismatch: %s", diff)
	}
}
