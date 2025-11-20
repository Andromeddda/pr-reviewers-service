# PR Reviewer Assignment Service

## Запуск сервиса

Сборка + запуск
```
docker compose up --build
```

Запуск в фоновом режиме
```
docker compose up -d
```

Остановка сервиса, запущенного в фоновом режиме
```
docker compose stop
```


## Структура проекта

[``prs/internal/model``](prs/internal/model/):

- ORM-объекты gorm для Postgres.

[``prs/internal/dto``](prs/internal/dto/):

- DTO-объекты json для HTTP запросов.

[``prs/internal/repository``](prs/internal/repository/)

- Объект для доступа в БД через транзакции.
- Зависит от ``model``.

[``prs/internal/config``](prs/internal/config/):

- Подтягивание переменных окружения для конфигурации БД и HTTP-сервера.
- Зависит от ``repository``.

[``prs/internal/mapper``](prs/internal/mapper/)

- Приведение ORM-объектов gorm из Postgres в DTO-объекты json для HTTP-интерфейса.
- Зависит от ``dto`` и ``model``.

[``prs/internal/service``](prs/internal/service/)

- Бизнес-логика сервиса.
- Зависит от ``repository`` и ``mapper``.

[``prs/internal/handler``](prs/internal/handler/)

- Обработчик HTTP-запросов. Методы обработчика регистрируются на маппинг роутера.
- Зависит от ``service`` и ``dto``.
