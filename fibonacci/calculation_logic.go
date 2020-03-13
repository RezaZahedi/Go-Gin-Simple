package fibonacci

import (
	"math/big"

	"github.com/RezaZahedi/Go-Gin/model/memo"
)

// FiboGenerator implements the GetFibonacciNumberService interface
type FiboGenerator struct {
	cache *memo.Memo
}

func NewFiboGenerator() *FiboGenerator {
	return &FiboGenerator{cache: memo.New(fiboTemp)}
}

func (g *FiboGenerator) GenerateNumber(number int) (string, error) {
	return calculate(number, g.cache), nil
}

func (g *FiboGenerator) Close() error {
	g.cache.Close()
	return nil
}

// this function is used to break the dependency loop in recursive function call and
// function memoization
var fibo func(a int) *big.Int

func fiboTemp(key int) interface{} {
	ans := fibo(key)
	return ans
}

func calculate(input int, m *memo.Memo) string {
	if input < 0 {
		panic("Calculate: input must be positive")
	}
	if input == 0 {
		return "0"
	}

	one := big.NewInt(1)

	fibo = func(a int) *big.Int {
		if a == 1 || a == 2 {
			return one
		}
		b := new(big.Int)
		b.Add(m.Get(a-2).(*big.Int), m.Get(a-1).(*big.Int))
		return b
	}

	return fibo(input).String()
}
