function initWASM(url) {
	const go = new Go();

	if ('instantiateStreaming' in WebAssembly) {
		WebAssembly.instantiateStreaming(fetch(url), go.importObject).then(function (obj) {
			go.run(obj.instance);
		})
	} else {
		fetch(WASM_URL).then(resp =>
			resp.arrayBuffer()
		).then(bytes =>
			WebAssembly.instantiate(bytes, go.importObject).then(function (obj) {
				go.run(obj.instance);
			})
		)
	}
}
