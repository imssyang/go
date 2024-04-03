all:
	go run imssyang.com/cgo

build:
	gcc -c -o num.o num.c
	ar rcs libnum.a num.o

clean:
	rm -rf num.o libnum.a
