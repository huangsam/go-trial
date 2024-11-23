all: main.out

main.out:
	go build -o main.out

clean:
	rm main.out
