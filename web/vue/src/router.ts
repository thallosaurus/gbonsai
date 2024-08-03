import { createWebHistory, createRouter } from "vue-router";
//import HomeView from "./pages/HomeView.vue";
import BonsaiPage from "./pages/BonsaiPage.vue";

const routes = [
    { path: "/", component: BonsaiPage },
    { path: "/bonsai", component: BonsaiPage }
]

export const router = createRouter({

    history: createWebHistory(),
    routes
})