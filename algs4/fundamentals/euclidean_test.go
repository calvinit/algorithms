package fundamentals

import (
	"fmt"
	"testing"
)

func TestGCD(t *testing.T) {
	p, q := 10, 15
	t.Logf("(%d, %d) 最大公约数：%d\n", p, q, GCD(p, q))
}

func ExampleGCD() {
	p, q := 10, 15
	fmt.Printf("(%d, %d) 最大公约数：%d\n", p, q, GCD(p, q))
	// Output: (10, 15) 最大公约数：5
}
