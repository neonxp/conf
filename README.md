# conf

[![🌱 Organic Code -- Code written by human](https://oc.neonxp.ru/organiccode.svg)](https://oc.neonxp.ru)
[![Go Doc](https://pkg.go.dev/badge/go.neonxp.ru/conf.svg)](https://pkg.go.dev/go.neonxp.ru/conf)

Go библиотека для парсинга конфигурационных файлов `.conf` похожем на
классические UNIX конфиги, как у nginx или bind9.

[English version](#english)

## Установка

```bash
go get go.neonxp.ru/conf
```

## Особенности формата

- **Директивы**: `directive arg1 arg2;`
- **Директивы с телом**: `directive arg1 arg2 { ... }`
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

    // Получение директивы и её значения
    if hostCmd := cfg.Get("server"); hostCmd != nil {
        fmt.Printf("Server: %v\n", hostCmd.Value())
    }

    // Навигация по вложенной структуре
    sslEnabled := cfg.Get("server").Group().Get("ssl").Group().Get("enabled")
    fmt.Printf("SSL enabled: %v\n", sslEnabled.Value())
}
```

## Пример конфигурационного файла

```conf
# Простые директивы без тела
listen 8080;
host "127.0.0.1";
debug false;

# Директивы с аргументами и телом
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

# Несколько директив с одинаковым именем
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

[![🌱 Organic Code -- Code written by human](https://oc.neonxp.ru/organiccode.svg)](https://oc.neonxp.ru)
[![Go Doc](https://pkg.go.dev/badge/go.neonxp.ru/conf.svg)](https://pkg.go.dev/go.neonxp.ru/conf)

Go library for parsing `.conf` configuration files (like many classic UNIX programms like nginx or bind9).

## Installation

```bash
go get go.neonxp.ru/conf
```

## Format Features

- **Directives**: `directive arg1 arg2;`
- **Directives with body**: `directive arg1 arg2 { ... }`
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

    // Get directive and its value
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
# Simple directives without body
listen 8080;
host "127.0.0.1";
debug false;

# Directives with arguments and body
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

# Multiple directives with same name
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
