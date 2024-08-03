<template>
    <p>Seed Chooser</p>
</template>

<script>
export default {
    setup() {
        let db;
        const request = indexedDC.open("bonsai");
        request.onerror = (event) => {
            console.error(event)
        }
    }
}
</script>