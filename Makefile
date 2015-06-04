BINARY = "$(basename $PWD)"

all: $(BINARY)

$(BINARY): *.go
	go build -ldflags "-X main.version $(git describe --long --tags --dirty --always)" .

deps:
	go get .

build: $(BINARY)

run: $(BINARY)
	./$(BINARY)

clean:
	rm $(BINARY)

test:
	go test ./...
	#golint ./...
