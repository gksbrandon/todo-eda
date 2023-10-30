up:
	docker-compose up -d 

down:
	docker-compose down --rmi all --volumes

build:
	docker-compose build monolith

logs:
	docker-compose logs -f

reset: down up logs

update:
	docker-compose down
	docker-compose build monolith
	docker-compose up -d


tidy:
	go mod tidy
	go mod vendor

generate:
	go generate ./...
