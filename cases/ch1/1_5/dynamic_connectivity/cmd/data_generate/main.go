package main

import (
	"bytes"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

// go build -o ./bin/ ./cases/ch1/1_5/dynamic_connectivity/cmd/data_generate
// ./bin/data_generate 625 datas/mediumUF.txt
func main() {
	if len(os.Args) != 3 {
		log.Fatalf("usage: %s <number> <file>", os.Args[0])
	}
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatalf("invalid number: %s", os.Args[1])
	}
	f, err := os.OpenFile(os.Args[2], os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("open file: %v", err)
	}
	defer f.Close()
	rand.Seed(time.Now().Unix())
	buf := bytes.Buffer{}
	for i := 0; i < n; i++ {
		p, q := rand.Intn(n), rand.Intn(n)
		buf.WriteString(strconv.Itoa(p))
		buf.WriteByte(' ')
		buf.WriteString(strconv.Itoa(q))
		buf.WriteByte('\n')
	}
	_, err = f.Write(buf.Bytes())
	if err != nil {
		log.Fatalf("file write: %v", err)
	}
}
