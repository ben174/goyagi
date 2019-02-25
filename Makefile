DIRS	 ?= $(shell go list ./...)
PKG_NAME ?= goyagi

COVERAGE_PROFILE ?= coverage.out

GOTOOLS := \
	github.com/alecthomas/gometalinter \
	github.com/golang/dep/cmd/dep \
	golang.org/x/tools/cmd/cover \
	github.com/codegangsta/gin \

default: build

.PHONY: build
build:
	@echo "---> Building"
	go build -o ./bin/$(PKG_NAME) ./cmd/serve

.PHONY: clean
clean:
	@echo "---> Cleaning"
	rm -rf ./bin ./vendor

.PHONY: enforce
enforce:
	@echo "---> Enforcing coverage"
	./scripts/coverage.sh $(COVERAGE_PROFILE)

.PHONY: html
html:
	@echo "---> Generating HTML coverage report"
	go tool cover -html $(COVERAGE_PROFILE)

.PHONY: lint
lint:
	@echo "---> Linting"
	gometalinter --vendor --tests --deadline=2m $(DIRS)

.PHONY: setup
setup:
	@echo "--> Installing development tools"
	go get -u $(GOTOOLS)
	gometalinter --install

.PHONY: start
start:
	@echo "---> Serving"
	PORT=3001 gin --port 3000 --appPort 3001 --path . --build ./cmd/serve --immediate --bin ./bin/gin-$(PKG_NAME) run

.PHONY: test
test:
	@echo "---> Testing"
	go test ./... -race -coverprofile $(COVERAGE_PROFILE)

.PHONY: db-reset
db-reset:
	@dropdb --if-exists goyagi
	@dropuser --if-exists goyagi_admin

.PHONY: db-setup
db-setup:
	make db-reset
	@createuser goyagi_admin
	@createdb -O goyagi_admin goyagi
	make db-migrate

.PHONY: db-migrate
db-migrate:
	@go run cmd/migrations/*.go migrate

.PHONY: db-rollback
db-rollback:
	@go run cmd/migrations/*.go rollback