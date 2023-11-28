<template>
  <div class="flex flex-row justify-end">
    <nav class="flex flex-row bg-gray-700 h-10 w-1/5 ">
      <ul class="w-full flex flex-row space-x-12 items-center justify-center">
        <!-- <li><img :src="Logo" alt="Logo" class="height-50"></li> -->
        <li><router-link to="/home">Home</router-link></li>
        <!-- <li><router-link to="/about">About</router-link></li> -->
        <li @click="logout" v-bind:style="{cursor: 'pointer'}">Logout</li>
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