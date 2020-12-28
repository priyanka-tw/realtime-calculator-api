APP_NAME := "calculator-api"

build:
	docker build . -t ${APP_NAME}

start:
	docker run -p 8080:8080 --name ${APP_NAME} --rm ${APP_NAME}:latest
