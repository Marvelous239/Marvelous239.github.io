.PHONY: clean wasm test-wasm

clean:
	rm main.wasm

wasm:
	go build -o main.wasm

test-wasm: main.wasm
	firefox http://localhost 2> /dev/null