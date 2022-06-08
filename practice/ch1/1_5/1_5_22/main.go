package main

import (
	"algo/practice/ch1/1_5/1_5_17/randconnect"
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("usage: %s <max>", os.Args[0])
	}
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	rand.Seed(time.Now().Unix())
	up := n
	if n > 200 {
		up = 200
	}
	for i := 0; i < up; i++ {
		x := rand.Intn(n)
		if x == 0 {
			continue
		}
		theory := 0.5 * float64(x) * math.Log(float64(x))
		fact := randconnect.Count(x)
		fmt.Println("x:", x, "theory:", theory, "fact:", fact)
	}
}
