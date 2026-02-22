# conf

Go библиотека для парсинга конфигурационных файлов `.conf`.

## Установка

```bash
go get go.neonxp.ru/conf
```

## Особенности формата

- **Присваивания**: `key = value;`
- **Типы значений**: строки (двойные/одинарные кавычки, backticks), числа (целые/дробные), булевы значения
- **Директивы**: `directive arg1 arg2;`
- **Блочные директивы**: `directive { ... }`
- **Комментарии**: `#` до конца строки
- **UTF-8**: включая кириллицу
- **Подстановка переменных окружения**: `$VAR`

## Быстрый старт

```go
package main

import (
    "fmt"
    "go.neonxp.ru/conf"
    "go.neonxp.ru/conf/model"
    "go.neonxp.ru/conf/visitor"
)

func main() {
    cfg := conf.New()
    if err := cfg.LoadFile("config.conf"); err != nil {
        panic(err)
    }

    v := visitor.NewDefault()
    if err := cfg.Process(v); err != nil {
        panic(err)
    }

    // Доступ по пути с точечной нотацией
    val, err := v.Get("server.host")
    if err != nil {
        panic(err)
    }
    fmt.Println(val.String())  // localhost

    port, err := v.Get("server.port")
    fmt.Println(port.Int())    // 8080
}
```

## Пример конфигурационного файла

```conf
# Переменная окружения: db_file = $HOME "/app/data.db";

# Простые присваивания
rss = "https://neonxp.ru/feed/";
host = "localhost";
port = 8080;
debug = true;

# Директивы с аргументами
telegram "bot123" "-1003888840756" {
    token = "token_value";
    admin_chat = "@admin";
}

# Вложенные блоки
server {
    host = "localhost";
    port = 8080;

    ssl {
        enabled = true;
        cert = "/etc/ssl/cert.pem";
    }

    middleware "auth" {
        enabled = true;
        secret = "$JWT_SECRET";
    }
}

# Многострочные строки
template = `
    <!DOCTYPE html>
    <html>
        <body>Hello</body>
    </html>
`;
```

## API

### Загрузка конфигурации

```go
cfg := conf.New()

// Из файла
cfg.LoadFile("config.conf")

// Из памяти
cfg.Load("inline", []byte("key = value;"))
```

### Обработка через Visitor

Библиотека использует паттерн Visitor для обхода конфигурации:

```go
type Visitor interface {
    VisitDirective(ident string, args Values, body Body) error
    VisitSetting(key string, values Values) error
}
```

### Get-методы на Values

| Метод | Описание |
|-------|----------|
| `String()` | Строковое представление через пробел |
| `Int()` | Преобразование в int (одно значение) |
| `BuildString(lookups...)` | Сборка строки с подстановками |

### Подстановка переменных окружения

```go
vals, _ := v.Get("db_file")
path := vals.BuildString(model.LookupEnv)
// $HOME → "/home/user", результат: "/home/user/app/data.db"
```

### Кастомные подстановки

```go
substitutions := map[model.Word]string{
    "APP_DIR":  "/opt/myapp",
    "LOG_LEVEL": "debug",
}

path := vals.BuildString(model.LookupSubst(substitutions), model.Origin)
```

## Реализация собственного Visitor

```go
type MyVisitor struct{}

func (m *MyVisitor) VisitDirective(ident string, args model.Values, body model.Body) error {
    fmt.Printf("Directive: %s, args: %s\n", ident, args.String())
    return body.Execute(m) // Рекурсивный обход тела
}

func (m *MyVisitor) VisitSetting(key string, values model.Values) error {
    fmt.Printf("Setting: %s = %s\n", key, values.String())
    return nil
}
```

## Грамматика (EBNF)

```
Config     = Doc .
Doc        = Stmt { Stmt } .
Stmt       = Word ( Assignment | Command ) .

Assignment = "=" Values br .
Command    = [Values] ( Body | br ) .

Values     = Value { Value } .
Value      = Word | String | Number | Boolean .
Body       = "{" [ Doc ] "}" .

Word       = word (alpha | "$" | "_") {alpha | number | "$" | "_"} .
String     = `"[^"]*"` | `'[^']*'` | '`' { `[^`]' } '`' .
Number     = `-?[0-9]+(\.[0-9]+)?` .
Boolean    = `true` | `false` .
br         = ";" .
```

## Требования

- Go 1.25+

## Лицензия

Этот проект лицензирован в соответствии с GNU General Public License версии 3
(GPLv3). Подробности смотрите в файле [LICENSE](LICENSE).

```
                    GNU GENERAL PUBLIC LICENSE
                       Version 3, 29 June 2007

 Copyright (C) 2026 Alexander NeonXP Kiryukhin <i@neonxp.ru>
 Everyone is permitted to copy and distribute verbatim copies
 of this license document, but changing it is not allowed.
```

## Автор

- Александр Кирюхин <i@neonxp.ru>
