all:
	go build ./...

test:
	go test ./... -race -coverprofile=c.out -covermode=atomic

cover: test
	go tool cover -html=c.out
	
install:
	go install .

clean:
	@rm -f lambdac c.out

serve:
	hugo -s site serve

hugo:
	hugo -s site --minify
