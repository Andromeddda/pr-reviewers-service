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


## Тестирование

### Мануальное E2E-тестирование

**Postman v2.1 коллекция:**

[``tests/postman/PRS.postman_collection.json``](tests/postman/PRS.postman_collection.json)

**Рекомендуемый порядок запросов из коллекции:**
1. ``POST /team/add`` - создать команду из 4-х человек - ``201 Created`` 

2. ``POST /pullRequest/create`` - открыть Pull-Request - ``201 Created``

3. ``POST /pullRequest/reassign`` - запросить переназначение - ``409 No Candidate``

4. ``POST /users/setIsActive`` - сменить статус пользователя  - ``200 OK``

5. ``POST /pullRequest/reassign`` - запросить переназначение  - ``200 OK``

6. ``POST /pullRequest/merge`` - закрыть Pull-Request  - ``200 OK``

7. ``POST /pullRequest/merge`` - закрыть Pull-Request ещё раз - ``200 OK``

8. ``POST /pullRequest/reassign`` - запросить переназначение - ``409 Merged``


### Нагрузочное тестирование

**Требования:**
- grafana/k6 docker image

**Установка требований:**
```
docker pull grafana/k6
```

**Запуск теста:**
```
docker compose down -v \
&& docker compose up --build -d \
&& docker run --rm \
--network=pr-reviewers-service_application_network \
-i grafana/k6 run - < tests/load/load_test.js
```

**Результат:**

[``tests/load/report.txt``](tests/load/report.txt):

- 100 Виртуальных пользователей.
- На каждого виртуального пользователя 1 команда по 10 человек + 1 pull-request.
- Длительность тестирования 30 секунд.
- 100% запросов получили ответ за < 300 ms.
- 99% процентов ответов приходят быстрее чем за p(0.99) = 175.7 ms.