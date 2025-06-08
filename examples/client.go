package main


import (
	"fmt"
	"strings"
)

func BuildStrings() {
	var b strings.Builder
	for i:=0; i<3; i++ {
		v, err := fmt.Fprintf(&b, "%d...", i)
		fmt.Println(v, err)
	}
	b.WriteString("ignition is turned on. Ready to GO!")
	b.Grow(5)
	fmt.Println(b.String())
	v := fmt.Sprintf("Cap is %d and Len is %d", b.Cap(), b.Len())
	fmt.Println(v)
}