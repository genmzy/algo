package randomgrid

import (
	"algo/practice/ch1/1_3/1_3_34/randbag"
	"fmt"
	"math/rand"
	"time"
)

type Connection struct {
	P, Q int
}

func (c Connection) String() string {
	return fmt.Sprintf("{%d, %d}", c.P, c.Q)
}

// n(n-1)*2 sides
func Generate(n int) []Connection {
	rand.Seed(time.Now().Unix())
	b := randbag.NewRandBag[Connection]()
	m := n * n
	for i := 1; i <= m; i++ {
		// try to connect to right neighbor
		if i%n > 0 {
			if rand.Intn(2) > 0 {
				b.Add(Connection{i, i + 1})
			} else {
				b.Add(Connection{i + 1, i})
			}
		}
		if (m - i) >= n {
			if rand.Intn(2) > 0 {
				b.Add(Connection{i, i + n})
			} else {
				b.Add(Connection{i + n, i})
			}
		}
	}
	v := make([]Connection, 0, b.Size())
	b.Traverse(func(c Connection) bool {
		v = append(v, c)
		return false
	})
	return v
}
