all: main.wasm wasm_exec.js
	goexec 'http.ListenAndServe(`:9999`, http.FileServer(http.Dir(`.`)))'

wasm_exec.js:
	cp "$(shell go env GOROOT)/misc/wasm/wasm_exec.js" .

main.wasm: main.go sync.go async.go dom.go
	GO111MODULE=auto GOOS=js GOARCH=wasm go build -o main.wasm .

clean:
	rm -rf *.wasm wasm_exec.js