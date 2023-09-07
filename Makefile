all: 
	go build ./...

test:
	go test ./... -race -coverprofile=c.out -covermode=atomic

cover: test
	go tool cover -html=c.out
	
install:
	go install .

clean:
	@rm -rf lambdac c.out site/public

serve:
	hugo -s site serve

hugo:
	hugo -s site --minify

lint:
	golangci-lint run ./...
