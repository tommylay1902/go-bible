package newkeyword

import "fmt"

func NewKeyword() {
	//returns pointer to int variable since it's not initialized it is default value of 0
	i := new(int)
	j := new(string)
	k := new([]int)
	l := new(bool)

	fmt.Printf("This is variable i (i := new(int)) default value: %v\n", *i)
	fmt.Printf("This is variable j (j := new(string)) default value: %v\n", *j)
	fmt.Printf("This is variable k (k := new([]int)) default value: %v\n", *k)
	fmt.Printf("This is variable l (l := new(bool)) default value: %v\n", *l)
}
