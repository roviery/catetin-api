package utils

import "fmt"

func Log(path string, function string, item string, output any) {
	fmt.Println("["+path+"]", function, item+":", output)
}
