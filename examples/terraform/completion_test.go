package examples

import (
	"path/filepath"
	"testing"
	"time"

	"github.com/dbanck/go-lsp-test/client"
	p "github.com/dbanck/go-lsp-test/protocol"
	"github.com/google/go-cmp/cmp"
)

func TestCompletion(t *testing.T) {
	testDir, err := filepath.Abs("modules")
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

	time.Sleep(3 * time.Second)

	testCases := []struct {
		name   string
		pos    p.Position
		result []p.CompletionItem
	}{
		{
			"address completion",
			p.Position{
				Line:      6,
				Character: 23,
			},
			[]p.CompletionItem{
				{
					Label:            "module.alpha.out1",
					Kind:             p.VariableCompletion,
					Detail:           "number",
					InsertTextFormat: p.PlainTextTextFormat,
					TextEdit: &p.TextEdit{
						Range: p.Range{
							Start: p.Position{
								Line:      6,
								Character: 10,
							},
							End: p.Position{
								Line:      6,
								Character: 23,
							},
						},
						NewText: "module.alpha.out1",
					},
				},
			},
		},
		{
			"module source completion",
			p.Position{
				Line:      1,
				Character: 12,
			},
			[]p.CompletionItem{},
		},
	}

	for _, tc := range testCases {
		t.Run((tc.name), func(t *testing.T) {
			result, err := lsc.GetCompletions(doc, tc.pos)
			if err != nil {
				t.Fatal(err)
			}

			first := result[0]
			resolve, err := lsc.ResolveCompletionItem(first)
			if err != nil {
				t.Fatal(err)
			}

			if diff := cmp.Diff(tc.result, result); diff != "" {
				t.Fatalf("completion mismatch: %s", diff)
			}
			if diff := cmp.Diff(&first, resolve); diff != "" {
				t.Logf("resolve diff: %s", diff)
			}
		})
	}
}
