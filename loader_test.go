package conf_test

import (
	"fmt"
	"log"

	"go.neonxp.ru/conf"
)

func ExampleLoad() {
	test := `
		some directive;
		group1 param1 {
			group2 param2 {
				group3 param3 {
					key value;
				}
			}
		}
	`
	result, err := conf.Load("test", []byte(test))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(
		result.Get("group1").
			Group.Get("group2").
			Group.Get("group3").
			Group.Get("key").Value(),
	) // → value

	fmt.Println(
		result.Get("group1").
			Group.Get("group2").
			Group.Get("group3").Value(),
	) // → param3

	// Output: value
	// param3
}
