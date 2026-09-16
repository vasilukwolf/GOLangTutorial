package main

import "fmt"

func main() {
	var gb float64
	mb := 4_831_838_208
	gb = float64(mb) / 1_073_741_824
	fmt.Printf("%d МБ = %.2f ГБ\n", mb, gb)
}
