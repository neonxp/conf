# conf

Go библиотека для парсинга конфигурационных файлов `.conf` похожем на
классические UNIX конфиги, как у nginx или bind9.

[English version](#english)

## Установка

```bash
go get go.neonxp.ru/conf
```

## Особенности формата

- **Команды**: `directive arg1 arg2;`
- **Команды с телом**: `directive arg1 arg2 { ... }`
- **Типы аргументов**: строки (двойные/одинарные кавычки/backticks для многострочных строк), числа (целые/дробные), булевы значения
- **Вложенные блоки**: произвольная глубина вложенности
- **Комментарии**: `#` до конца строки
- **UTF-8**: включая кириллицу

## Быстрый старт

```go
package main

import (
    "fmt"
    "go.neonxp.ru/conf"
)

func main() {
    // Загрузка из файла
    cfg, err := conf.LoadFile("config.conf")
    if err != nil {
        panic(err)
    }

    // Получение команды и её значения
    if hostCmd := cfg.Get("server"); hostCmd != nil {
        fmt.Printf("Server: %v\n", hostCmd.Value())
    }

    // Навигация по вложенной структуре
    sslEnabled := cfg.Get("server").Group.Get("ssl").Group.Get("enabled")
    fmt.Printf("SSL enabled: %v\n", sslEnabled.Value())
}
```

## Пример конфигурационного файла

```conf
# Простые команды без тела
listen 8080;
host "127.0.0.1";
debug false;

# Команды с аргументами и телом
server "web" {
    host "localhost";
    port 8080;

    ssl {
        enabled true;
        cert "/etc/ssl/cert.pem";
        key "/etc/ssl/key.pem";
    }

    middleware "auth" {
        enabled true;
        secret "secret123";
    }
}

# Несколько команд с одинаковым именем
cache "redis" {
    host "redis.local";
    port 6379;
}

cache "memcached" {
    host "memcached.local";
    port 11211;
}

# Многострочные строки
template `
    <!DOCTYPE html>
    <html>
        <body>Hello</body>
    </html>
`;
```

## API

### Загрузка конфигурации

```go
// Из файла
cfg, err := conf.LoadFile("path/to/config.conf")

// Из памяти
cfg, err := conf.Load("inline", []byte("listen 8080;"))
```

### Типы данных

```go
// model.Ident — идентификатор команды (псевдоним для string)
ident := model.Ident("server")

// model.Command — команда с именем, аргументами и вложенной группой
type Command struct {
    Name  Ident  // Имя команды
    Args  []any  // Аргументы команды
    Group Group  // Вложенные команды (тело команды)
}

// model.Group — срез команд
type Group []Command
```

### Методы Command

| Метод         | Описание                                     |
| ------------- | -------------------------------------------- |
| `Value() any` | Возвращает первый аргумент команды или `nil` |

### Методы Group

| Метод                                                      | Описание                                                   |
| ---------------------------------------------------------- | ---------------------------------------------------------- |
| `Get(name Ident) *Command`                                 | Возвращает первую команду с указанным именем или `nil`     |
| `Filter(predicate func(*Command) bool) iter.Seq[*Command]` | Возвращает итератор по командам, удовлетворяющим предикату |

## Навигация по конфигурации

```go
// Простой доступ
listenCmd := cfg.Get("listen")
fmt.Println(listenCmd.Value()) // 8080

// Цепочка вложенной навигации
port := cfg.Get("server").Group.Get("http").Group.Get("port")
fmt.Println(port.Value()) // 8080

// Использование Filter
for cmd := range cfg.Filter(func(c *model.Command) bool {
    return c.Name == "cache"
}) {
    fmt.Printf("Cache: %v\n", cmd.Value())
}
```

## Работа с аргументами

Аргументы сохраняются в срез `Args` типа `[]any`. Возможные типы:

```go
// Строка в двойных кавычках: "value"
// Строка в одинарных кавычках: 'value'
// Строка в backticks: `value`
// Число целое: 42
// Число дробное: 3.14
// Булево значение: true, false

// Пример доступа к аргументам:
cmd := cfg.Get("test")
if cmd != nil && len(cmd.Args) > 0 {
    val1 := cmd.Args[0]      // Первый аргумент
    val2 := cmd.Args[1]      // Второй аргумент

    // Приведение типа для строк
    if str, ok := val1.(string); ok {
        fmt.Println("String:", str)
    }

    // Приведение типа для чисел
    if num, ok := val2.(int); ok {
        fmt.Println("Int:", num)
    }
}
```

## Грамматика (PEG)

Формат использует PEG (Parsing Expression Grammar). Грамматика описана в файле `parser/grammar.peg`.

Основные правила:

- Все команды без тела заканчиваются точкой с запятой `;`
- Тело команды заключается в фигурные скобки `{ }`
- Аргументы разделяются пробелами
- Комментарии начинаются с `#`

## Требования

- Go 1.23+ (для использования `iter.Seq`)

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

---

<a name="english"></a>

# conf (English)

Go library for parsing `.conf` configuration files (like many classic UNIX programms like nginx or bind9).

## Installation

```bash
go get go.neonxp.ru/conf
```

## Format Features

- **Commands**: `directive arg1 arg2;`
- **Commands with body**: `directive arg1 arg2 { ... }`
- **Argument types**: strings (double/single quotes, backticks for multiline strings), numbers (integer/float), boolean values
- **Nested blocks**: arbitrary nesting depth
- **Comments**: `#` until end of line
- **UTF-8**: including Cyrillic

## Quick Start

```go
package main

import (
    "fmt"
    "go.neonxp.ru/conf"
)

func main() {
    // Load from file
    cfg, err := conf.LoadFile("config.conf")
    if err != nil {
        panic(err)
    }

    // Get command and its value
    if hostCmd := cfg.Get("server"); hostCmd != nil {
        fmt.Printf("Server: %v\n", hostCmd.Value())
    }

    // Navigate through nested structure
    sslEnabled := cfg.Get("server").Group.Get("ssl").Group.Get("enabled")
    fmt.Printf("SSL enabled: %v\n", sslEnabled.Value())
}
```

## Example Configuration File

```conf
# Simple commands without body
listen 8080;
host "127.0.0.1";
debug false;

# Commands with arguments and body
server "web" {
    host "localhost";
    port 8080;

    ssl {
        enabled true;
        cert "/etc/ssl/cert.pem";
        key "/etc/ssl/key.pem";
    }

    middleware "auth" {
        enabled true;
        secret "secret123";
    }
}

# Multiple commands with same name
cache "redis" {
    host "redis.local";
    port 6379;
}

cache "memcached" {
    host "memcached.local";
    port 11211;
}

# Multiline strings
template `
    <!DOCTYPE html>
    <html>
        <body>Hello</body>
    </html>
`;
```

## API

### Loading Configuration

```go
// From file
cfg, err := conf.LoadFile("path/to/config.conf")

// From memory
cfg, err := conf.Load("inline", []byte("listen 8080;"))
```

### Data Types

```go
// model.Ident — command identifier (alias for string)
ident := model.Ident("server")

// model.Command — command with name, arguments and nested group
type Command struct {
    Name  Ident  // Command name
    Args  []any  // Command arguments
    Group Group  // Nested commands (command body)
}

// model.Group — slice of commands
type Group []Command
```

### Command Methods

| Method        | Description                                |
| ------------- | ------------------------------------------ |
| `Value() any` | Returns first argument of command or `nil` |

### Group Methods

| Method                                                     | Description                                        |
| ---------------------------------------------------------- | -------------------------------------------------- |
| `Get(name Ident) *Command`                                 | Returns first command with specified name or `nil` |
| `Filter(predicate func(*Command) bool) iter.Seq[*Command]` | Returns iterator over commands matching predicate  |

## Configuration Navigation

```go
// Simple access
listenCmd := cfg.Get("listen")
fmt.Println(listenCmd.Value()) // 8080

// Nested navigation chain
port := cfg.Get("server").Group.Get("http").Group.Get("port")
fmt.Println(port.Value()) // 8080

// Using Filter
for cmd := range cfg.Filter(func(c *model.Command) bool {
    return c.Name == "cache"
}) {
    fmt.Printf("Cache: %v\n", cmd.Value())
}
```

## Working with Arguments

Arguments are stored in `Args` slice of type `[]any`. Possible types:

```go
// Double-quoted string: "value"
// Single-quoted string: 'value'
// Backtick string: `value`
// Integer number: 42
// Float number: 3.14
// Boolean value: true, false

// Example accessing arguments:
cmd := cfg.Get("test")
if cmd != nil && len(cmd.Args) > 0 {
    val1 := cmd.Args[0]      // First argument
    val2 := cmd.Args[1]      // Second argument

    // Type assertion for strings
    if str, ok := val1.(string); ok {
        fmt.Println("String:", str)
    }

    // Type assertion for numbers
    if num, ok := val2.(int); ok {
        fmt.Println("Int:", num)
    }
}
```

## Grammar (PEG)

The format uses PEG (Parsing Expression Grammar). Grammar is described in file `parser/grammar.peg`.

Basic rules:

- All commands without body end with semicolon `;`
- Command body is enclosed in curly braces `{ }`
- Arguments are separated by spaces
- Comments start with `#`

## Requirements

- Go 1.23+ (for `iter.Seq`)

## License

This project is licensed under GNU General Public License version 3 (GPLv3).
See [LICENSE](LICENSE) file for details.

```
                   GNU GENERAL PUBLIC LICENSE
                      Version 3, 29 June 2007

Copyright (C) 2026 Alexander NeonXP Kiryukhin <i@neonxp.ru>
Everyone is permitted to copy and distribute verbatim copies
of this license document, but changing it is not allowed.
```

## Author

- Alexander Kiryukhin <i@neonxp.ru>
