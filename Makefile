#tinygo build -o build/gbonsai.wasm -target wasm cmd/wasm/main.go
gbonsai_out:
	mkdir -p build
	GOOS=js GOARCH=wasm go build -o build/gbonsai.wasm cmd/wasm/main.go
	cp web/basic/wasm_exec.go.js build/wasm_exec.js
	cp web/basic/gbonsai.js build/gbonsai.js
	cp web/basic/bbt.ttf build/bbt.ttf
	cp web/basic/main.css build/main.css
	cp web/basic/index.html build/index.html
	cp web/basic/package.json build/package.json

gbonsai_cli:
	mkdir -p build
	go build -o build/gbonsai cmd/cli/main.go