package conf_test

import (
	"fmt"

	"go.neonxp.ru/conf"
)

func ExampleLoad() {
	config := `
		key = "value";
		group "test" {
			key = 123;
		}
	`

	cfg, err := conf.Load("example", []byte(config))
	if err != nil {
		panic(err)
	}

	fmt.Println("key =", cfg.Get("key")[0])
	group := cfg.Commands("group")
	for _, gr := range group {
		fmt.Println("key =", gr.Body.Get("key")[0])
	}
	// Output:
	// key = value
	// key = 123
}
