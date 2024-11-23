all: main.out

main.out: *.go $(wildcard **/*.go)
	go build -o main.out

clean:
	rm main.out
