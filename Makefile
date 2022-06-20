SRC = cmd/main.go
APP = application
MIGRATEDOWNFLAG = -migrateDown
ENV_SCRIPT = debugEnv.sh

start:
	bash $(ENV_SCRIPT)
	go run $(SRC)

build:
	go build -o $(APP) -v $(SRC) && ./$(APP)

docker:
		docker-compose up

migrateDown:
	go run $(SRC) $(MIGRATEDOWNFLAG)