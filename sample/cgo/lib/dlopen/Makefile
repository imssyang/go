all:
	go run imssyang.com/cgo

build:
	gcc -shared -o num/libnum.so num/num.c

clean:
	rm -rf num/libnum.so cgo
