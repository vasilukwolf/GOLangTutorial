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

### path3 — ошибки, defer/panic/recover, JSON

| Файл | Что делает |
|------|-----------|
| [activeSubscribe.go](path3/activeSubscribe.go) | Активация подписки: sentinel-ошибки, обёртка `%w`, `errors.Is` / `errors.As` |
| [filmCatalogJSON.go](path3/filmCatalogJSON.go) | Загрузка каталога фильмов из JSON-файлов с `defer`, `recover` и замером времени |

> `filmCatalogJSON.go` читает файлы из [path3/JSON/](path3/JSON/) по относительному
> пути, поэтому запускайте его **из корня репозитория**:
> `go run path3/filmCatalogJSON.go`

### path4 — структуры, методы, интерфейсы

| Файл | Что делает |
|------|-----------|
| [userProfile.go](path4/userProfile.go) | Структура профиля пользователя, нулевые значения |
| [subscriptionStruct.go](path4/subscriptionStruct.go) | Подписка: конструктор и изменение через указатель |
| [deviceStruct.go](path4/deviceStruct.go) | Методы с получателем-значением и получателем-указателем |
| [colletionFavorite.go](path4/colletionFavorite.go) | Коллекция избранного, уникальные значения через `map`, анонимная структура |
| [watchHistory.go](path4/watchHistory.go) | Вложенные структуры и JSON-теги |
| [allertChanel.go](path4/allertChanel.go) | Интерфейс уведомлений: email и push |
| [exportJSONandCSV.go](path4/exportJSONandCSV.go) | Интерфейс экспорта в CSV и JSON |
| [interfaceComposite.go](path4/interfaceComposite.go) | Встраивание интерфейсов в структуру |
| [anyValue.go](path4/anyValue.go) | Тип `any`: type assertion и type switch |
| [curtomError.go](path4/curtomError.go) | Собственный тип ошибки модерации и `errors.As` |

### path5 — горутины и каналы

| Файл | Что делает |
|------|-----------|
| [verificationCE.go](path5/verificationCE.go) | Параллельная проверка эпизодов через `sync.WaitGroup` |
| [chanelTranscode.go](path5/chanelTranscode.go) | Параллельное «транскодирование» видео с результатами через канал |
| [playerQueue.go](path5/playerQueue.go) | Буферизованный канал как очередь событий плеера |
| [profileRecom.go](path5/profileRecom.go) | Загрузка рекомендаций с таймаутом через `select` и `time.After` |
| [goflixAnalytics.go](path5/goflixAnalytics.go) | Пул воркеров аналитики: `context.WithTimeout`, `sync.Mutex`, `WaitGroup` |

## Запустить все задания разом

Bash / Git Bash:

```sh
for f in path*/*.go; do echo "== $f"; go run "$f"; done
```

PowerShell:

```powershell
Get-ChildItem path*\*.go | ForEach-Object { "== $($_.FullName)"; go run $_.FullName }
```
