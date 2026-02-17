package model

import (
	"os"
	"strings"
)

// WordLookup тип определяющий функцию поиска замены слов при
// стрингификации Values.
type WordLookup func(word Word) string

// chainLookup утилитарная функция для последовательных применений функций
// поиска до первого нахождения подстановки. Если returnOrigin == true,
// то в случае неудачи вернёт имя слова.
func chainLookup(lookups ...WordLookup) WordLookup {
	return func(word Word) string {
		for _, lookup := range lookups {
			if v := lookup(word); v != "" {
				return v
			}
		}
		return string(word)
	}
}

// LookupEnv функция типа WordLookup которая пытается подставить вместо word
// соответствующую ему переменную окружения ОС. При этом он срабатывает только
// если слово начинается со знака `$`.
func LookupEnv(word Word) string {
	if !strings.HasPrefix(string(word), "$") {
		return ""
	}
	varName := strings.TrimPrefix(string(word), "$")
	if result, ok := os.LookupEnv(varName); ok {
		return result
	}
	return ""
}

// LookupSubst функция типа WordLookup которая пытается подставить вместо word
// значение из словаря подстановок по соответствующему ключу.
func LookupSubst(subst map[Word]string) WordLookup {
	return func(word Word) string {
		if result, ok := subst[word]; ok {
			return result
		}
		return ""
	}
}
