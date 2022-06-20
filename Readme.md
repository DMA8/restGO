Локальный запуск: "make start"
Локальный билд: "make build"
Запуск в контейнере: "docker-compose up" или make docker
По дефолту накатываются миграции из schema/ Чтобы откатить миграцию: make migrateDown

GET		localhost:8080/id		{"props":[{"id": 12312}, {"id": 122}]} дата - опциональна, на результат запроса не влияет
POST	localhost:8080/create	{"props":[{"id": 12312, "date":"2020-09-30"}, {"id": 122}]}
POST	localhost:8080/update	{"props":[{"id": 12312, "date":"2020-09-30"}, {"id": 122}]}
./postman - коллекции

формат времени: YYYY-DD-MM. Валидация времени посредством time.Parse. Id > 0 - валиден