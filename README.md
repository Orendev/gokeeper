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

## Структура проекта
> * **cmd** - основной код приложения
> * **docs** - проектная и пользовательская документация
> * **api** - Спецификация OpenApi/Swager, файлы определения протокола
> * **deployments** - шаблоны и файлы конфигураций для деплоя
> * **pkg** - код библиотек
> * **scripts** - вспомогательные скрипты
> * **services** - код сервисов
>   * **cmd** - основной код сервиса
>     * **app/main.go** - основное приложение
>     * **..** вспомогательные приложения
>   * **internal** - внутрений код сервиса и библиотек
>     * **delivery** - точка входа для стороней системы
>       * **grpc** - gRPC
>     * **domain** - сущности
>     * **repository** - хранилище
>       * **adapters** преобразования данных для хранилища
>         * **storage**
>     * **useCase** - сценарии
>       * **adapters** - интерфейс взаимодействия с репозиторием
>   * **configs** - шаблоны файлов конфигураций сервиса
>   * **migrations** - миграции
> * **vendor** - зависимости приложения


