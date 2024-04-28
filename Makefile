IMAGE_NAME := test-dp

dev:
	nodemon --exec go run main.go --signal SIGTERM

dcb:
	docker build -t ${IMAGE_NAME} .

dcr:
	docker run --name ${IMAGE_NAME} -p 1323 -d ${IMAGE_NAME}

start:
	docker build -t ${IMAGE_NAME} . && docker run --name ${IMAGE_NAME} -p 9999:1111 -d ${IMAGE_NAME}

rmi:
	docker stop ${IMAGE_NAME} && docker rm ${IMAGE_NAME} && docker rmi ${IMAGE_NAME}

	--platform linux/amd64