package model

import (
	"strconv"
	"strings"
)

type Value any

type Values []Value

// BuildString собирает из значений Value цельную строку, при этом приводя все
// значения к типу string. Так же принимает функции типа WordLookup, которые
// последовательно будут пытаться привести значения типа Word к
// контекстозависимым значениям. Например, пытаться находить по имени переменную
// окружения ОС.
func (v Values) BuildString(lookups ...WordLookup) string {
	sw := strings.Builder{}

	for _, v := range v {
		switch v := v.(type) {
		case string:
			sw.WriteString(v)
		case float64:
			sw.WriteString(strconv.FormatFloat(v, 'f', 5, 64))
		case int:
			sw.WriteString(strconv.Itoa(v))
		case bool:
			if v {
				sw.WriteString("true")
				continue
			}
			sw.WriteString("false")
		case Word:
			sw.WriteString(chainLookup(lookups...)(v))
		}
	}

	return sw.String()
}

type Word string
