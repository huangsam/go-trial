SOURCES := *.go $(wildcard **/*.go)

all: gotrial

gotrial: $(SOURCES)
	go build io.huangsam/trial/cmd/gotrial

clean:
	rm gotrial
