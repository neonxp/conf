package main

import (
	"fmt"

	"go.neonxp.ru/conf"
)

func main() {
	out, err := conf.LoadFile("./file.conf")
	if err != nil {
		panic(err)
	}

	fmt.Println(out)
}
