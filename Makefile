ENABLE_WASM_OPT ?= true

.PHONY: build
build:
	tinygo build -target=wasip1 -gc=leaking -buildmode=c-shared -no-debug -o redirect.wasm ./
ifeq ($(ENABLE_WASM_OPT),true)
	wasm-opt -Os -o redirect.wasm redirect.wasm
endif
