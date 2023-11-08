import { createApp } from "vue";
import App from './App.vue'
import router from './index'; // Import your router instance

//createApp(App).mount('#app')
// createApp(Login).mount('#app')


const app = createApp(App);
app.use(router); // Use the router instance
app.mount('#app');
