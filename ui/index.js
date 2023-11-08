import { createRouter, createWebHistory } from 'vue-router';
import Home from './pages/Jobs.vue';
import Login from './pages/Login.vue';
import Files from './pages/Files.vue';

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
  {
    path: '/files',
    name: 'Documents',
    component: Files
  }
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
