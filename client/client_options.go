package client

import p "github.com/dbanck/go-lsp-test/protocol"

func getCapabilities() p.ClientCapabilities {
	return p.ClientCapabilities{
		Workspace: p.Workspace3Gn{
			CodeLens: p.CodeLensWorkspaceClientCapabilities{
				RefreshSupport: true,
			},
		},
		TextDocument: p.TextDocumentClientCapabilities{
			SemanticTokens: p.SemanticTokensClientCapabilities{
				TokenTypes: []string{
					"namespace",
					"type",
					"class",
					"enum",
					"interface",
					"struct",
					"typeParameter",
					"parameter",
					"variable",
					"property",
					"enumMember",
					"event",
					"function",
					"method",
					"macro",
					"keyword",
					"modifier",
					"comment",
					"string",
					"number",
					"regexp",
					"operator",
				},
				TokenModifiers: []string{
					"declaration",
					"definition",
					"readonly",
					"static",
					"deprecated",
					"abstract",
					"async",
					"modification",
					"documentation",
					"defaultLibrary",
				},
				Requests: struct {
					Range bool        "json:\"range,omitempty\""
					Full  interface{} "json:\"full,omitempty\""
				}{
					Range: true,
					Full: struct {
						Delta bool "json:\"delta,omitempty\""
					}{
						Delta: true,
					},
				},
			},
			CodeLens: p.CodeLensClientCapabilities{
				DynamicRegistration: true,
			},
		},
		Experimental: struct {
			ShowReferencesCommandId string "json:\"showReferencesCommandId,omitempty\""
		}{
			ShowReferencesCommandId: "client.showReferences",
		},
	}
}
