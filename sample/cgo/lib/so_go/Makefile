all:
	./sum

build:
	go build -buildmode=c-shared -o sum.so sum.go
	gcc -o sum main.c -Wl,-rpath=./ sum.so -lpthread

analyze:
	go tool cgo sum.go

clean:
	rm -rf sum.h sum.so sum _obj
