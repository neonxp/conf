package conf_test

import (
	"fmt"

	"go.neonxp.ru/conf"
	"go.neonxp.ru/conf/visitor"
)

func ExampleNew() {
	config := `
		key = "value";
		group "test" {
			key = 123;
		}
	`

	cfg := conf.New()

	if err := cfg.Load("example", []byte(config)); err != nil {
		panic(err)
	}

	pr := visitor.NewDefault()
	if err := cfg.Process(pr); err != nil {
		panic(err)
	}

	val1, err := pr.Get("key")
	if err != nil {
		panic(err)
	}
	val2, err := pr.Get("group.key")
	if err != nil {
		panic(err)
	}

	val3, err := pr.Get("group")
	if err != nil {
		panic(err)
	}

	fmt.Println("key =", val1.String())
	fmt.Println("group.key =", val2.String())
	fmt.Println("group args =", val3.String())

	// Output:
	// key = value
	// group.key = 123
	// group args = test
}
