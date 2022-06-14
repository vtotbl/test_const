run: build_docker run_docker

build_docker:
	docker build -t service -f ./build/Dockerfile .

run_docker:
	docker run -p 127.0.0.1:8080:8080 --rm -it service:latest