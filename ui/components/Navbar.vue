<template>
  <div class="flex justify-end w-full">
    <div class="bg-gray-500 rounded-lg shadow-md border border-gray-300 w-1/4">
      <!-- Logo container with some padding and a defined height -->
      <div class="flex flex-row items-center h-10">
        <div class="h-full p-2">
          <img src="./../assets/logo1.png" alt="Logo" class="h-full">
        </div>
  
        <!-- Navigation Links -->
        <nav class="flex-grow">
          <ul class="flex flex-row space-x-12 justify-end items-center h-full">
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
    </div>
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
