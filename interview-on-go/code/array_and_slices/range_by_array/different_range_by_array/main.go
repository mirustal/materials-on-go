package different_range_by_array

import "fmt"

func main() {
	data := [...]int{1, 2, 3}
	for value := range data { // копирование, для больших массивов будет overhead
		fmt.Println(value)
	}

	for value := range &data { // не копируем
		fmt.Println(value)
	}

	for value := range data[:] {
		fmt.Println(value)
	}
}
