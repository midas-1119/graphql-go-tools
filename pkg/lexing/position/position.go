package position

import "fmt"

type Position struct {
	LineStart uint32
	LineEnd   uint32
	CharStart uint32
	CharEnd   uint32
}

func (p Position) String() string {
	return fmt.Sprintf("%d:%d-%d:%d", p.LineStart, p.CharStart, p.LineEnd, p.CharEnd)
}

func (p *Position) MergeStartIntoStart(position Position) {
	p.LineStart = position.LineStart
	p.CharStart = position.CharStart
}

func (p *Position) MergeStartIntoEnd(position Position) {
	p.LineEnd = position.LineStart
	p.CharEnd = position.CharStart
}

func (p *Position) MergeEndIntoEnd(position Position) {
	p.LineEnd = position.LineEnd
	p.CharEnd = position.CharEnd
}

func (p *Position) IsBefore(another Position) bool {
	return p.LineEnd < another.LineStart ||
		p.LineEnd == another.LineStart && p.CharEnd < another.CharStart
}
