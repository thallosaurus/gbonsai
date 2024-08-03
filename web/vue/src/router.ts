import { createWebHistory, createRouter } from "vue-router";
import BonsaiPage from "./pages/BonsaiPage.vue";
import HomeView from "./pages/HomeView.vue";

const routes = [
    { path: "/", component: HomeView },
    { path: "/bonsai", component: BonsaiPage }
]

export const router = createRouter({

    history: createWebHistory(),
    routes
})