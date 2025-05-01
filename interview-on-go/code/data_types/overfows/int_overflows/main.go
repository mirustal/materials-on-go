package int_overflows

import (
	"fmt"
	"math"
)

func main() {
	var signed int8 = math.MinInt8
	signed++

	var unsigned uint8 = math.MaxUint8
	unsigned++

	fmt.Println(signed, unsigned)
}
