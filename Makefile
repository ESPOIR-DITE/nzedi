docker-build :
	docker build -t 0816048957docker/nzedi-api:latest -f Dockerfile .

docker-compose-up:
	docker-compose -f docker-compose-local-dev.yml up -d

docker-push:
	docker push 0816048957docker/nzedi-api:latest