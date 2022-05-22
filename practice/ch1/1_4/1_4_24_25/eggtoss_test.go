package eggtoss

import (
	"math/rand"
	"testing"
	"time"
)

func BenchmarkEggsToss(b *testing.B) {
	maxFloor := b.N
	b.Logf("maxFloor: %d", maxFloor)
	rand.Seed(time.Now().Unix())
	for i := 0; i < b.N; i++ {
		r := rand.Intn(maxFloor)
		b.Run("", func(b *testing.B) {
			res := EggsToss(maxFloor, func(k int) bool {
				if k >= r {
					return true
				}
				return false
			})
			if res != r {
				b.Fatal("not correct, check this function")
			}
		})
	}
}

func BenchmarkEggsTossAnother(b *testing.B) {
	maxFloor := b.N
	b.Logf("maxFloor: %d", maxFloor)
	rand.Seed(time.Now().Unix())
	for i := 0; i < b.N; i++ {
		r := rand.Intn(maxFloor)
		b.Run("", func(b *testing.B) {
			res := EggsTossAnother(maxFloor, func(k int) bool {
				if k >= r {
					return true
				}
				return false
			})
			if res != r {
				b.Fatal("not correct, check this function")
			}
		})
	}
}

func BenchmarkTwoEggsToss(b *testing.B) {
	maxFloor := b.N
	b.Logf("maxFloor: %d", maxFloor)
	rand.Seed(time.Now().Unix())
	for i := 0; i < b.N; i++ {
		r := rand.Intn(maxFloor)
		b.Run("", func(b *testing.B) {
			res := TwoEggsToss(maxFloor, func(k int) bool {
				if k >= r {
					return true
				}
				return false
			})
			if res != r {
				b.Fatal("not correct, check this function")
			}
		})
	}
}

func TestEggsToss(t *testing.T) {
	type args struct {
		n     int
		broke func(k int) bool
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{args: args{n: 20, broke: func(k int) bool {
			if k >= 25 {
				return true
			}
			return false
		}}, want: -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EggsToss(tt.args.n, tt.args.broke); got != tt.want {
				t.Errorf("EggToss() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEggsTossAnother(t *testing.T) {
	type args struct {
		n     int
		broke func(k int) bool
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{args: args{n: 20, broke: func(k int) bool {
			if k >= 25 {
				return true
			}
			return false
		}}, want: -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EggsTossAnother(tt.args.n, tt.args.broke); got != tt.want {
				t.Errorf("EggTossAnother() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTwoEggsToss(t *testing.T) {
	type args struct {
		n     int
		broke func(k int) bool
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TwoEggsToss(tt.args.n, tt.args.broke); got != tt.want {
				t.Errorf("TwoEggsToss() = %v, want %v", got, tt.want)
			}
		})
	}
}
