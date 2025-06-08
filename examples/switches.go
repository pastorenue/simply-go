package main

import (
	"fmt"
)

func getValue() interface{} {
	return 54
}

func SwitchMain() {
	value := getValue()
	switch v := value.(type) {
	case string:
		fmt.Println("String value:", v)
	case int:
		fmt.Println("Integer value:", v)
		default:
		fmt.Println("Unknown type")
	}
}
