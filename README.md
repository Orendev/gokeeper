# Description
## gokeeper
Golang Keeper password manager представляет собой клиент-серверную систему, позволяющую пользователю надёжно и безопасно хранить логины, пароли, бинарные данные и прочую приватную информацию.

В реализации проекта используется подход Domain-driven desing (DDD) в Go (Golang).

Этот проект имеет 4 доменных уровня :

- Domain (Models) Layer
- Repository Layer
- Usecase Layer
- Delivery Layer
#### Диаграма:

![golang clean architecture](https://github.com/Orendev/gokeeper/raw/main/clean-arch.png)

Оригинальное объяснение структуры этого проекта можно прочитать в посте на medium's : https://medium.com/@imantumorang/golang-clean-archithecture-efd6d7c43047.


### Первоначальная настройка приложения
``cp .env.example .env //настроить .env``

### Генерация сертификата
``make gen-cert``

### Генерация protobuff
``make protoc``

### Build App Client
``make build-client``

### Генерация моков
``make mockery``

## Структура проекта
> * **cmd** - основной код приложения
>   * **client/main.go** - основное приложение клиент
>   * **..** вспомогательные приложения
>   * **server/main.go** - основное приложение сервер
> * **docs** - проектная и пользовательская документация
> * **api** - Спецификация OpenApi/Swager, файлы определения протокола
> * **deployments** - шаблоны и файлы конфигураций для деплоя
> * **pkg** - код библиотек
> * **scripts** - вспомогательные скрипты
> * **internal** - код сервисов
>   * **app** - внутрений код сервиса и библиотек
>     * **delivery** - точка входа для стороней системы
>       * **grpc** - gRPC
>     * **domain** - сущности
>     * **repository** - хранилище
>       * **storage** преобразования данных для хранилища
>         * **postgres**
>           * **migrations** - миграции
>     * **useCase** - сценарии
>       * **adapters** - интерфейс взаимодействия с репозиторием
>     * **configs** - шаблоны файлов конфигураций сервиса
> * **vendor** - зависимости приложения


