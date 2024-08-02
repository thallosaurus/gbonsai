#GOOS=js GOARCH=wasm go build -o build/gbonsai.wasm cmd/wasm/main.go
gbonsai_out:
	mkdir -p build
	tinygo build -o build/gbonsai.wasm -target wasm cmd/wasm/main.go
	cp web/wasm_exec.js build/wasm_exec.js
	cp web/gbonsai.js build/gbonsai.js
	cp web/bbt.ttf build/bbt.ttf
	cp web/main.css build/main.css
	cp web/index.html build/index.html
	cp web/package.json build/package.json

gbonsai_cli:
	mkdir -p build
	go build -o build/gbonsai cmd/cli/main.go