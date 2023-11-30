<template>
<div class="flex flex-row justify-end">
  <nav class="flex flex-row bg-gray-500 h-10 w-1/4 rounded-lg shadow-md border border-gray-300">
    <ul class="w-full flex flex-row space-x-12 items-center justify-center">
      <!-- <li><img :src="Logo" alt="Logo" class="h-10"></li> -->
      <li class="hover:bg-gray-700 rounded-md">
        <router-link 
          to="/home" 
          class="px-3 py-2 transition-colors duration-200"
          :class="{ 'underline text-white': $route.path === '/home' }">
          Home
        </router-link>
      </li>
      <li class="hover:bg-gray-700 rounded-md">
        <router-link 
          to="/files" 
          class="px-3 py-2 transition-colors duration-200"
          :class="{ 'underline text-white': $route.path === '/files' }">
          Files
        </router-link>
      </li>
      <li 
        @click="logout" 
        class="hover:bg-gray-700 rounded-md px-3 py-2 cursor-pointer transition-colors duration-200">
        Logout
      </li>
    </ul>
  </nav>
</div>

  </template>
  
  <script>
  import { ref } from 'vue';
  // import Logo from '@/assets/logo1.png'
  import { useRouter } from 'vue-router';
  import axios from 'axios'
  export default {
    name: 'Navbar',
    setup(){
        const router = useRouter();
        const logout = () => {
            axios.post("/api/logout").
            then(()=>{
                console.log("logged out")
                router.push("/")
            }).catch(err=>console.error(err))
            // send request to backend
            // clear cookie
            // navigate to /login page
        }

        return {
            logout
        }
    
    }
  };
  </script>
  
  <!-- <style>
  .navbar {
    background-color: #333;
    color: white;
    padding: 1rem;
  }
  .navbar ul {
    list-style-type: none;
    display: flex;
  }
  .navbar li {
    margin-right: 1rem;
  }
  .navbar a {
    color: white;
    text-decoration: none;
  }
  </style>
   -->
