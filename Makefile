SRC = cmd/main.go
APP = application
MIGRATEDOWNFLAG = -migrateDown
ENV_SCRIPT = debugEnv.sh

migrateDown:
	go run $(SRC) $(MIGRATEDOWNFLAG)

build:
	go build -o $(APP) -v $(SRC) && ./$(APP)

start:
	bash $(ENV_SCRIPT)
	go run $(SRC)

docker:
		docker-compose up