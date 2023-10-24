up:
	docker-compose up -d 

down:
	docker-compose down --rmi all --volumes

build:
	docker-compose build monolith

reset: down up

update:
	docker-compose down
	docker-compose build monolith
	docker-compose up -d

logs:
	docker-compose logs -f

tidy:
	go mod tidy
	go mod vendor

generate:
	go generate ./...
