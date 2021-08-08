package p1030

import (
	"fmt"
	"testing"
)

func TestAllCellsDistOrder(t *testing.T) {
	fmt.Println(allCellsDistOrder(1, 2, 0, 0))
	fmt.Println(allCellsDistOrder(2, 2, 0, 1))
	fmt.Println(allCellsDistOrder(2, 3, 1, 2))
}
