 
all: test
 
test:
	go test ./...
 
coverage:
	go test -cover ./...