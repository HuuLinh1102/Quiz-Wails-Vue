import {createApp} from 'vue'
import App from './App.vue'
import '../src/style.css';

// VueSweetalert2
import VueSweetalert2 from "vue-sweetalert2";
import "sweetalert2/dist/sweetalert2.min.css";

const app = createApp(App)
app.use(VueSweetalert2);
app.mount('#app')
