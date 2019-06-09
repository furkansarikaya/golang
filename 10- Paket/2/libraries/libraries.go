package main

import (
	"fmt"

	"../utils"
)

func main() {
	n1, l1 := utils.FullName("Zaphod", "Beeblebrox")
	fmt.Printf("FullName: %v,number of chars:%v\n\n", n1, l1)

	n2, l2 := utils.FullNameNakedReturn("Arthur", "Dent")
	fmt.Printf("FullName: %v,number of chars:%v\n\n", n2, l2)
}
