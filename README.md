# conf

Go библиотека для чтения конфигурационных файлов в формате `.conf`.

## Установка

```bash
go get go.neonxp.ru/conf
```

## Особенности формата

- Простые присваивания: `key = value;`
- Строковые значения: `"text"` или `'text'`
- Многострочные строки: ``` `text` ```
- Числовые значения: целые и дробные, включая отрицательные
- Булевы значения: `true` / `false`
- Директивы/команды: `directive arg1 arg2;`
- Групповые директивы с блоками кода: `directive { ... }`
- Комментарии от `#` до конца строки
- Поддержка кириллицы и UTF-8

## Пример использования

```go
package main

import (
    "fmt"
    "go.neonxp.ru/conf"
)

func main() {
    doc, err := conf.LoadFile("./config.conf")
    if err != nil {
        panic(err)
    }

    // Получение значения по ключу
    values := doc.Get("my_key")
    fmt.Println(values)

    // Получение команд по имени
    commands := doc.Commands("directivename")
    fmt.Println(commands)

    // Все переменные
    vars := doc.Vars()

    // Все элементы документа
    items := doc.Items()
}
```

## Пример конфигурационного файла

```conf
# Пример конфигурации

# Простое присваивание
simple_key = value;

# Многострочное присваивание
string_key =
    "value"
    'string';

# Многострочные строки (backticks)
multiline_string = `
    multiline
    string
    123
`;

# Числа и булевы значения
int_key = -123.456;
bool_key = true;

# Директивы
expression1 argument1 "argument2" 123;

# Групповая директива
group_directive_without_arguments {
    expression1 argument2 "string" 123 true;
    expression2 argument3 "string111" 123321 false;

    children_group "some argument" {
        # Вложенная группа
    }
}

# Групповая директива с аргументами
group_directive_with_argument "argument1" 'argument2' {
    child_val = "children value";
}
```

## API

### Функции

- `LoadFile(filename string) (*model.Doc, error)` - загрузка конфигурации из файла
- `Load(name string, input []byte) (*model.Doc, error)` - парсинг конфигурации из байтов

### Методы `*model.Doc`

- `Get(key string) Values` - получить значения по ключу
- `Commands(name string) Commands` - получить команды по имени
- `Vars() map[string]Values` - получить все переменные
- `Items() []any` - получить все элементы документа

## Требования

- Go 1.25+

## Лицензия

См. файл [LICENSE](LICENSE)
