package model

import (
	"fmt"
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

func (v Values) String() string {
	result := make([]string, 0, len(v))

	for _, v := range v {
		switch v := v.(type) {
		case string:
			result = append(result, v)
		case float64:
			result = append(result, strconv.FormatFloat(v, 'f', 5, 64))
		case int:
			result = append(result, strconv.Itoa(v))
		case bool:
			if v {
				result = append(result, "true")
				continue
			}
			result = append(result, "false")
		case Word:
			result = append(result, string(v))
		}
	}

	return strings.Join(result, " ")
}

func (v Values) Int() (int, error) {
	if len(v) != 1 {
		return 0, fmt.Errorf("AsInt can return only single value (there is %d values)", len(v))
	}
	val := v[0]
	switch val := val.(type) {
	case int:
		return val, nil
	case string:
		return strconv.Atoi(val)
	case float64:
		return int(val), nil
	default:
		return 0, fmt.Errorf("invalid type for convert to int: %t", val)
	}
}

type Word string
