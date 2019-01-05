package types

import (
	"strings"
)
type Word string

func (this *Word) Trim() *Word {
	*this =  Word(strings.TrimSpace(string(*this)))
	return this
}

func (this *Word) Wrap(wrapSymbol string) *Word {
	*this = Word(wrapSymbol + string(*this) + wrapSymbol)
	return this
}

func (this *Word) esEmpty() bool {
	return *this == ""	
}
