BINARY=basesvc
test: 
	@cd usecase && go test -v -cover ./...
	@cd adapter && go test -v -cover ./...

engine:
	go build -a -tags musl -o ${BINARY}

dev: setup run-dev

prod: setup docker run-prod

dependencies:
	@echo "> Installing the server dependencies ..."
	@go mod download
	@go get -u github.com/cosmtrek/air

unittest:
	go test -short  ./...

clean:
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

lint-prepare:
	@echo "Installing golangci-lint" 
	curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh| sh -s latest

lint:
	./bin/golangci-lint run ./...
 
migrate-up:
	@echo "> Migrate up database"
	@go run main.go migratepq up

migrate-down:
	@echo "> Migrate down database"
	@go run main.go migratepq down

docs:
	@echo "> Generate Swagger Docs"
	@if ! command -v swag &> /dev/null; then go install github.com/swaggo/swag/cmd/swag ; fi
	@swag init -g infrastructure/rest/rest.go

docker:
	@echo "> Build Docker image [PRODUCTION]"
	@docker build -t basesvc -f build/Dockerfile . 

run-dev:
	@echo "> Run docker-compose [DEV]"
	@docker-compose -f deployments/docker-compose.dev.yml up --build -d

run-prod:
	@echo "> Run docker [PRODUCTION]"
	@docker-compose -f deployments/docker-compose.yml up --build -d

setup:
	@if ! [ -f ".env" ]; then cp .env.dist .env ; fi
	@if ! [ -f "config.json" ]; then cp config.json.dist config.json ; fi

stop:
	@echo "> Stop docker-compose"
	@docker-compose -f deployments/docker-compose.yml -f deployments/docker-compose.dev.yml down

.PHONY: clean install unittest lint-prepare lint docs engine dev prod test setup dependencies run-dev run-prod stop
