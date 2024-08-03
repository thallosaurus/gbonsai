<template>
    <!--<button @click.prevent="animate_bonsai">Animate</button>-->
    <div>

        <div id="bonsai-app" v-html="html"></div>
    </div>
    <!-- loading state via #fallback slot -->
</template>

<script lang="ts">

declare function gbonsai(seed: number, life: number): string;

const go = new Go(); // Defined in wasm_exec.js. Don't forget to add this in your index.html.
const wasmBrowserInstantiate = async (wasmModuleUrl: string, importObject: any) => {
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

export async function initWasm(): Promise<WebAssembly.WebAssemblyInstantiatedSource> {
    const importObject = go.importObject;
    const wasmModule = await wasmBrowserInstantiate("/gbonsai.wasm", importObject);
    go.run(wasmModule.instance);
    return wasmModule
};

export default {
    async setup(props: any) {
        await initWasm()
        const html = gbonsai(props.seed, props.life)
        return {
            html,
        }
    },
    props: {
        seed: Number,
        life: Number
    },
    methods: {
        animate_bonsai() {
            for (let i = 1; i <= this.$props.life!; i++) {

            setTimeout(() => {
                const html = gbonsai(this.$props.seed!, i)
                this.html = html
                this.$forceUpdate();
            }, i * 250)
            }
            
        }
    }
}
</script>

<style>
#bonsai-app {
    font-size: .5em;
    font-family: BigBlueTerminal;
    text-align: center;
    display: flex;
    flex-flow: column wrap;
    max-width: fit-content;
}

#tree p {
    margin-block-start: 0em;
    margin-block-end: 0em;
    display: inline;
    white-space: nowrap;
    max-width: fit-content;
}

#tree span {
    display: inline-block
}
</style>