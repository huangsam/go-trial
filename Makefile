SOURCES := *.go $(wildcard **/*.go)

all: main.out

main.out: $(SOURCES)
	go build -o main.out

clean:
	rm main.out
