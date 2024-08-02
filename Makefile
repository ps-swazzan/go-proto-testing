.POSIX:
.SUFFIXES:

all: build

build: proto
	go build

test: proto
	go test

clean:
	rm -rf tutorialpb
	rm myapp
	
proto: tutorialpb/addressbook.pb.go

tutorialpb/addressbook.pb.go: addressbook.proto
	protoc -I=. --go_out=. addressbook.proto
