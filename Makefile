
dev:
	TZ=utc APP_ENV=development go run main.go

test:
	TZ=utc APP_ENV=test go test -v ./...

lint:
	golangci-lint run

update-docs:
	swag init
