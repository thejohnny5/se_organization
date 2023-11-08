import { createRouter, createWebHistory } from 'vue-router';
import Home from './pages/Jobs.vue';
import Login from './pages/Login.vue';

const routes = [
  {
    path: '/',
    name: 'Login',
    component: Login
  },
  {
    path: '/home',
    name: 'Home',
    component: Home
  },
  // Add more routes as needed
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
