const go = new Go(); // Defined in wasm_exec.js. Don't forget to add this in your index.html.

const wasmBrowserInstantiate = async (wasmModuleUrl, importObject) => {
  let response = undefined;

  // Check if the browser supports streaming instantiation
  if (WebAssembly.instantiateStreaming) {
    // Fetch the module, and instantiate it as it is downloading
    response = await WebAssembly.instantiateStreaming(
      fetch(wasmModuleUrl),
      importObject
    );
  } else {
    // Fallback to using fetch to download the entire module
    // And then instantiate the module
    const fetchAndInstantiateTask = async () => {
      const wasmArrayBuffer = await fetch(wasmModuleUrl).then(response =>
        response.arrayBuffer()
      );
      return WebAssembly.instantiate(wasmArrayBuffer, importObject);
    };

    response = await fetchAndInstantiateTask();
  }

  return response;
};

const initWasm = async () => {
  const importObject = go.importObject;
  const wasmModule = await wasmBrowserInstantiate("gbonsai.wasm", importObject);
  go.run(wasmModule.instance);
  return wasmModule
};

window.onload = (e) => {
  initWasm().then(mod => {
    mod.instance.exports.gbonsai_dom(BigInt(Date.now()))
    console.log(mod)
    //mod.instance.exports.animate_generation("tree", BigInt(1), 150)
  })
}