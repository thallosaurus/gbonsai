#GOOS=js GOARCH=wasm go build -o build/gbonsai.wasm cmd/wasm/main.go
gbonsai_out:
	mkdir -p build
	tinygo build -o build/gbonsai.wasm -target wasm cmd/wasm/main.go
	cp cmd/wasm/wasm_exec.js build/wasm_exec.js
	cp cmd/wasm/module.js build/module.js
	cp cmd/wasm/bbt.ttf build/bbt.ttf
	cp cmd/wasm/main.css build/main.css
	cp cmd/wasm/index.html build/index.html
	cp cmd/wasm/package.json build/package.json

gbonsai_cli:
	mkdir -p build
	go build -o build/gbonsai cmd/cli/main.go