# NazarovGO

Учебные задания по Go. Каждое задание — отдельный файл со своей функцией `main`,
поэтому запускается **по одному файлу**.

## Требования

- Go 1.21+ (проверено на go1.25.3)

Проверить установку:

```sh
go version
```

## Как запустить задание

Из корня репозитория:

```sh
go run <папка>/<файл>.go
```

Например:

```sh
go run path1/hello.go
go run path2/filterByRating.go
```

> ⚠️ Не запускайте `go run ./path1` или `go run path1/*.go` — в каждом файле
> своя функция `main`, и компилятор выдаст ошибку `main redeclared`.

Собрать исполняемый файл:

```sh
go build -o hello.exe path1/hello.go
./hello.exe
```

## Список заданий

### path1 — основы

| Файл | Что делает |
|------|-----------|
| [hello.go](path1/hello.go) | Hello, World! |
| [film.go](path1/film.go) | Вывод карточки фильма |
| [goflix.go](path1/goflix.go) | Информация о тарифном плане |
| [mbtogb.go](path1/mbtogb.go) | Перевод мегабайт в гигабайты |
| [sectoour.go](path1/sectoour.go) | Перевод секунд в часы/минуты/секунды |
| [formatRecept.go](path1/formatRecept.go) | Форматированный чек за билеты |
| [filmTiket.go](path1/filmTiket.go) | Расчёт цены билета |
| [tikcketAllow.go](path1/tikcketAllow.go) | Проверка допуска к билету |
| [acessControll.go](path1/acessControll.go) | Контроль доступа к контенту |
| [cinema_schedule.go](path1/cinema_schedule.go) | Расписание сеансов по залам |
| [palyer_GoFlix.go](path1/palyer_GoFlix.go) | Команды видеоплеера |

### path2 — слайсы

| Файл | Что делает |
|------|-----------|
| [sum.go](path2/sum.go) | Сумма элементов слайса |
| [removeDublicates.go](path2/removeDublicates.go) | Удаление дубликатов из слайса |
| [filterByRating.go](path2/filterByRating.go) | Фильтрация по рейтингу |

## Запустить все задания разом

Bash / Git Bash:

```sh
for f in path*/*.go; do echo "== $f"; go run "$f"; done
```

PowerShell:

```powershell
Get-ChildItem path*\*.go | ForEach-Object { "== $($_.FullName)"; go run $_.FullName }
```
