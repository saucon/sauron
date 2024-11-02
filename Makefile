build-docker-log:
	docker build --no-cache \
		-t sauron-log:1.0 -f v2/Dockerfile . --progress=plain
	docker run -d --name sauron-log sauron-log:1.0