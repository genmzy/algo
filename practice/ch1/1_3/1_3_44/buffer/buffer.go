package buffer

import (
	"algo/practice/ch1/1_3/1_3_44/cursorlist"
	"bytes"
)

type Buffer struct {
	cursorlist.CursorList[byte]
}

func (b *Buffer) String() string {
	s := bytes.Buffer{}
	b.Traverse(func(v byte) bool {
		s.WriteByte(v)
		return false
	})
	return s.String()
}
