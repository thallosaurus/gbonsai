#GOOS=js GOARCH=wasm go build -o build/gbonsai.wasm cmd/wasm/main.go
gbonsai_out:
	mkdir -p build
	tinygo build -o build/gbonsai.wasm -target wasm cmd/wasm/main.go
	cp cmd/wasm/wasm_exec.js build/wasm_exec.js
	cp cmd/wasm/glue.js build/glue.js
	cp cmd/wasm/index.html build/index.html

gbonsai_cli:
	mkdir -p build
	go build -o build/gbonsai cmd/cli/main.go