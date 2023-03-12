package parse

import (
	"github.com/alecthomas/chroma"
	"github.com/shivamMg/rd"
)

func Body(b *rd.Builder) (ok bool) {
	b.Enter("Body")
	defer b.Exit(&ok)

	if !Match(b, chroma.Token{Type: chroma.Keyword, Value: "body"}) {
		return false
	}
	if !MatchType(b, chroma.Keyword) {
		return false
	}
	if !MatchType(b, chroma.NameFunction) && !MatchType(b, chroma.Keyword) {
		Fmt(b, "AA", 0)
		return false
	}
	// if next is ( -> params, if { open the body
	if Peek(b, chroma.Token{Type: chroma.Punctuation, Value: "("}) {
		if !ArgList(b) {
			return false
		}
	}

	Comments(b)

	// now we should see {
	if !MatchDiscard(b, chroma.Token{Type: chroma.Punctuation, Value: "{"}) {
		return false
	}

	Comments(b)
	defer Comments(b)

	return BodyBody(b) && Match(b, chroma.Token{Type: chroma.Punctuation, Value: "}"})
}

func BodyBody(b *rd.Builder) (ok bool) {
	b.Enter("BodyBody")
	defer b.Exit(&ok)

More:
	// classguardselection and selections or just selections
	ClassGuardSelections(b)
	BodySelections(b)

	if !Peek(b, chroma.Token{Type: chroma.Punctuation, Value: "}"}) {
		goto More
	}
	return true
}

func ClassGuardSelections(b *rd.Builder) (ok bool) {
	b.Enter("ClassGuardSelections")
	defer b.Exit(&ok)

	if !MatchType(b, chroma.NameClass) {
		return false
	}
	if !MatchDiscard(b, chroma.Token{Type: chroma.Punctuation, Value: "::"}) {
		return false
	}
	return BodySelections(b)
}

func BodySelections(b *rd.Builder) (ok bool) {
	b.Enter("BodySelections")
	defer b.Exit(&ok)

	for {
		Comments(b) // comments in between selections and trailing ones
		if !Selection(b) {
			return true
		}
	}
}

func Selection(b *rd.Builder) (ok bool) {
	b.Enter("Selection")
	defer b.Exit(&ok)

	Comments(b)
	if !MatchType(b, chroma.KeywordReserved) && !MatchType(b, chroma.KeywordType) {
		return false
	}

	if !FatArrow(b) {
		return false
	}
	return Rval(b) && b.Match(chroma.Token{Type: chroma.Punctuation, Value: ";"})
}
