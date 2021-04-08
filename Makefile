 
all: test
 
test:
	go test -v ./...
 
coverage:
	go test -cover ./...