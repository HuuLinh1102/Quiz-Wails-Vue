import { createRouter, createWebHashHistory } from "vue-router";

import Landing from "./views/Landing.vue";
import Login from "./views/Login.vue";
import Exam from "./views/Exam.vue";

const router = createRouter({
    history: createWebHashHistory(),
    routes: [ 
        {path: "/", name: "landing", component: Landing},
        {path: "/login/:user", name: "login", component: Login},
        {path: "/exam/:data", name: "exam", component: Exam},
    ],
});

export default router;