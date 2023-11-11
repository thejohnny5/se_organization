<template>
    <nav class="flex">
      <ul>
        <li><router-link to="/home">Home</router-link></li>
        <li><router-link to="/about">About</router-link></li>
        <li @click="logout">Logout</li>
      </ul>
    </nav>
  </template>
  
  <script>
  import { ref } from 'vue';
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