package codecamp

import (
	"fmt"
)

const (
	isA = 1 << iota
	isB
	isC
	isD
)

func Main() {
	var acd byte = isA | isC | isD
	fmt.Printf("Sequence: %b\n", acd)
	fmt.Printf("Is A: %v\n", acd&isA == isA)
	fmt.Printf("Is B: %v\n", acd&isB == isB)
	fmt.Printf("Is C: %v\n", acd&isC == isC)
	fmt.Printf("Is D: %v\n", acd&isD == isD)
}
