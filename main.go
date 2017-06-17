package main

import (
	"fmt"
	"hash/fnv"
	"math"
	"math/rand"
	"os"
	"time"
)

var prefixes = map[string](func(float64) float64){
	"натуральных ":           nat(),
	"ридгаро-фуралисматико-": mul(1.001),
	"квадратных ":            pow(1.0 / 2.0),
	"нано-":                  mul(100),
	"кубических ":            pow(1.0 / 3.0),
	"кубических корней из ":  pow(3.0),
	"кило-":                  mul(1.0 / 10.0),
	"милли-":                 mul(1000),
	"недо-":                  mul(247),
	"полу-":                  mul(2),
	"пони-":                  mul(0.98),
	"лоло-":                  mul(math.Abs(rand.Float64())),
}

func nat() func(float64) float64 {
	return func(a float64) float64 {
		return a
	}
}

func pow(p float64) func(float64) float64 {
	return func(a float64) float64 {
		return float64(math.Pow(float64(a), float64(p)))
	}
}

func mul(p float64) func(float64) float64 {
	return func(a float64) float64 {
		return a / p
	}
}

func hash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}

func main() {

	if len(os.Args) < 2 {
		return
	}

	now := float64(time.Now().Unix() / 10)
	nick := float64(hash(os.Args[1]))

	for pr, fn := range prefixes { // use loop for shuf map

		pt := math.Abs(fn(Noise2D(now, nick, 42, 2, 2, 3)) * 1000)
		fmt.Printf("Упоротость чятика: %4.2f %v%v\n", pt, pr, os.Args[1])

		break
	}

}
