#tinygo build -o build/gbonsai.wasm -target wasm cmd/wasm/main.go
gbonsai_out:
	mkdir -p build.basic
	GOOS=js GOARCH=wasm go build -o build.basic/gbonsai.wasm cmd/wasm/main.go
	cp web/basic/wasm_exec.go.js build.basic/wasm_exec.js
	cp web/basic/gbonsai.js build.basic/gbonsai.js
	cp web/basic/bbt.ttf build.basic/bbt.ttf
	cp web/basic/main.css build.basic/main.css
	cp web/basic/index.html build.basic/index.html
	cp web/basic/package.json build.basic/package.json

gbonsai_cli:
	mkdir -p build.cli
	go build -o build.cli/gbonsai cmd/cli/main.go

gbonsai_vue_prereq: gbonsai_out
	cp build.basic/gbonsai.wasm web/vue/public/gbonsai.wasm
	cp build.basic/wasm_exec.js web/vue/public/wasm_exec.js

gbonsai_vue: gbonsai_vue_prereq
	cd web/vue/ && yarn build

clean:
	rm -rf web/vue/dist
	rm -rf web/vue/public/wasm_exec.js
	rm -rf web/vue/public/gbonsai.wasm
	rm -rf build.basic