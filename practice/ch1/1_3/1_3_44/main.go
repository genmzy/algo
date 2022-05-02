package main

import (
	"algo/practice/ch1/1_3/1_3_44/buffer"
	"fmt"
)

func dump(b *buffer.Buffer) {
	fmt.Print("dumping ")
	fmt.Println(b.String())
}

func main() {
	buf := &buffer.Buffer{}
	for i := 0; i < 26; i++ {
		buf.Add(byte(int('a') + i))
	}
	dump(buf)

	fmt.Printf("current cursor char: %c\n", *buf.Cursor())
	buf.Right(3)
	fmt.Printf("after left 3 times, current cursor char: %c\n", *buf.Cursor())

	buf.InsertBeforeN('.', 2)
	dump(buf)
	fmt.Printf("current cursor char: %c\n", *buf.Cursor())

	buf.InsertAfterN('z', 1000)
	dump(buf)

	fmt.Printf("current cursor char: %c\n", *buf.Cursor())
	v := buf.DeleteBeforeN(2)
	if v == nil {
		fmt.Println(v)
	} else {
		fmt.Printf("deleted char: %c\n", *v)
	}
	dump(buf)

	diff := int('A') - int('a')
	buf.AllDo(func(c *byte) bool {
		*c += byte(diff)
		return false
	})
	dump(buf)

	v = buf.DeleteAfterN(100)
	if v == nil {
		fmt.Println(v)
	} else {
		fmt.Printf("deleted char: %c\n", *v)
	}
	v = buf.Cursor()
	if v == nil {
		fmt.Println(v)
	} else {
		fmt.Printf("current cursor char: %c\n", *v)
	}
}
