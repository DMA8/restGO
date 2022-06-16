CONNSTR=postgresql://postgres:postgres@localhost:5432/postgres?sslmode=disable
migrate_up:
	migrate -path ./schema -database $(CONNSTR) up
migrate_down:
	 migrate -path ./schema -database $(CONNSTR) down