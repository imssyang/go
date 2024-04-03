all:
	./sum

build:
	go build -buildmode=c-archive -o sum.a sum.go
	gcc -o sum main.c sum.a -lpthread

analyze:
	go tool cgo sum.go

clean:
	rm -rf sum.h sum.a sum _obj
