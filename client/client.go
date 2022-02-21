package client

import (
	"context"
	"io"
	"io/ioutil"
	"log"
	"net"
	"os"
	"path/filepath"

	"github.com/creachadair/jrpc2"
	"github.com/creachadair/jrpc2/channel"
	"github.com/dbanck/go-lsp-test/internal/uri"
	p "github.com/dbanck/go-lsp-test/protocol"
)

type testClient struct {
	conn   net.Conn
	client *jrpc2.Client

	rootPath string
	rootUri  string
}

func testLogger(w io.Writer, prefix string) *log.Logger {
	return log.New(w, prefix, log.LstdFlags|log.Lshortfile)
}

func NewClient(addr string, dir string) (*testClient, error) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return nil, err
	}

	ch := channel.LSP(conn, conn)
	opts := &jrpc2.ClientOptions{
		Logger: jrpc2.StdLogger(testLogger(os.Stdout, "[CLIENT] ")),
	}
	cli := jrpc2.NewClient(ch, opts)

	c := &testClient{
		client:   cli,
		conn:     conn,
		rootUri:  uri.FromPath(dir),
		rootPath: dir, // TODO!: check if directory exists
	}

	c.Initialize()

	// TODO!: wait for InitializeResult

	c.Initialized()

	return c, nil
}

func (c *testClient) Initialize() {
	params := p.InitializeParams{
		ProcessID:    12345,
		Capabilities: p.ClientCapabilities{},
		RootPath:     c.rootPath,
		RootURI:      p.DocumentURI(c.rootUri),
	}
	c.client.Call(context.Background(), "initialize", params)
}

func (c *testClient) Initialized() {
	params := p.InitializedParams{}
	c.client.Call(context.Background(), "initialized", params)
}

func (c *testClient) OpenDoc(doc string, languageId string) (p.DocumentURI, error) {
	docPath := filepath.Join(c.rootPath, doc) // TODO!: check if file exists
	uri := p.DocumentURI(uri.FromPath(docPath))
	text, err := ioutil.ReadFile(docPath)
	if err != nil {
		return "", err
	}

	params := p.DidOpenTextDocumentParams{
		TextDocument: p.TextDocumentItem{
			Text:       string(text),
			Version:    0,
			LanguageID: languageId,
			URI:        uri,
		},
	}
	_, err = c.client.Call(context.Background(), "textDocument/didOpen", params)
	if err != nil {
		return "", err
	}

	return uri, nil
}

func (c *testClient) GetCompletions(uri p.DocumentURI, pos p.Position) ([]p.CompletionItem, error) {
	params := p.CompletionParams{
		TextDocumentPositionParams: p.TextDocumentPositionParams{
			TextDocument: p.TextDocumentIdentifier{
				URI: uri,
			},
			Position: pos,
		},
	}

	rsp, err := c.client.Call(context.Background(), "textDocument/completion", params)
	if err != nil {
		return nil, err
	}

	var result p.CompletionList
	if err := rsp.UnmarshalResult(&result); err != nil {
		return nil, err
	}

	return result.Items, nil
}

func (c *testClient) Close() {
	c.client.Call(context.Background(), "shutdown", nil)
	c.client.Call(context.Background(), "exit", nil)

	c.client.Close()
	c.conn.Close()
}
