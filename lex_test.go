package cf

import (
	"strings"
	"testing"

	"github.com/miekg/cf/ast"
)

func TestLexClass(t *testing.T) {
	const input = `bundle agent prometheus_server
{
 files:

  IsPrometheusServer::
`
	l := NewLexer(strings.NewReader(input))
	l.symbols = []Symbol{identifier, varclass, class, char}
	last := ast.Token{}
	for t := l.Scan(); t.Tok != DONE; t = l.Scan() {
		last = t
	}
	if last.Tok != CLASSGUARD {
		t.Errorf("expected CLASSGUARD, got %s", SymbolText[last.Tok])
	}
}
