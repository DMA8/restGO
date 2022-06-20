Локальный запуск: "make start"<br>
Локальный билд: "make build"<br>
Запуск в контейнере: "docker-compose up" или make docker<br>
По дефолту накатываются миграции из schema/ Чтобы откатить миграцию: make migrateDown<br>

GET		localhost:8080/id		{"props":[{"id": 12312}, {"id": 122}]} дата - опциональна, на результат запроса не влияет<br>
POST	localhost:8080/create	{"props":[{"id": 12312, "date":"2020-09-30"}, {"id": 122}]}<br>
POST	localhost:8080/update	{"props":[{"id": 12312, "date":"2020-09-30"}, {"id": 122}]}<br>
./postman - коллекции<br>

формат времени: YYYY-DD-MM. Валидация времени посредством time.Parse. Id > 0 - валиден<br>
