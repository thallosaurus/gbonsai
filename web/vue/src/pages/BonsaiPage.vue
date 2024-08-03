<script lang="ts" setup>
import BonsaiLoader from '../components/BonsaiLoader.vue';
</script>

<script lang="ts">
export default {

}
</script>

<template>
    <p>Bonsai</p>
    <Suspense>
        <BonsaiLoader :seed="Date.now()" :life="32" />

        <template #fallback>
            <p id="loading">
                Loading...
            </p>
        </template>
    </Suspense>
</template>

<style scoped>
#loading {
    text-align: center
}
</style>