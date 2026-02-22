package main

import (
	"fmt"

	"go.neonxp.ru/conf"
	"go.neonxp.ru/conf/visitor"
)

func main() {
	cfg := conf.New()
	if err := cfg.LoadFile("./example/file2.conf"); err != nil {
		panic(err)
	}

	pr := visitor.NewDefault()
	if err := cfg.Process(pr); err != nil {
		panic(err)
	}

	tok, err := pr.Get("telegram.token")
	if err != nil {
		panic(err)
	}

	fmt.Println(tok.String())
}
