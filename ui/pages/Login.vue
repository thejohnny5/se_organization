<template>
    <div class="login-container">
      <div class="login-box">
        <h1 class="title">Welcome</h1>
        <p class="subtitle">Please login using one of the following methods:</p>
  
        <!-- Google OAuth Button -->
        <button @click="loginWithGoogle" class="login-button google">
          Login with Google
        </button>
  
        <!-- GitHub OAuth Button -->
        <button @click="loginWithGithub" class="login-button github">
          Login with GitHub
        </button>
      </div>
    </div>
  </template>

  <script>
  import { onMounted, ref } from 'vue';
  import axios from 'axios';
  import { useRouter } from 'vue-router'; // Import useRouter to get access to the router instance
  
  export default {
    name: 'LoginPage',
  
    setup() {
      const router = useRouter(); // Get the router instance
  
      // Reactive reference for loading state, for example
      const isLoading = ref(false);
  
      const loginWithGoogle = () => {
        // Implement your login logic here
        window.location.href='/oauth/google/login'; // Use router to navigate
      };
  
      const loginWithGithub = () => {
        // Implement your login logic here
        console.log('Logging in with GitHub...');
      };
  
      const isLoggedIn = async () => {
        isLoading.value = true; // Set loading state to true
        try {
          await axios.get("/api/validate");
          router.push('/home');
        } catch (err) {
          console.log("not logged in", err);
        } finally {
          isLoading.value = false; // Set loading state to false
        }
      };
  
      onMounted(() => {
        isLoggedIn();
      });
  
      // Return the functions and state you want to use in your template
      return {
        isLoading,
        loginWithGoogle,
        loginWithGithub,
      };
    },
  };
  </script>
  
  
  <style scoped>
  .login-container {
    display: flex;
    align-items: center;
    justify-content: center;
    height: 100vh;
    background-color: #121212;
  }
  
  .login-box {
    text-align: center;
    color: white;
    max-width: 400px;
    padding: 40px;
    border: 1px solid #0eaf51;
    border-radius: 8px;
    background-color: #1f1f1f;
    box-shadow: 0 4px 8px rgba(0, 0, 0, 0.2);
  }
  
  .title {
    margin-bottom: 10px;
    font-size: 2.5em;
    font-family: Arial, Helvetica, sans-serif;
  }
  
  .subtitle {
    margin-bottom: 20px;
    color: #aaaaaa;
    font-family: Arial, Helvetica, sans-serif;
  }
  
  .login-button {
    display: block;
    width: 100%;
    padding: 10px;
    margin-bottom: 10px;
    border: none;
    border-radius: 4px;
    color: white;
    font-size: 1em;
    cursor: pointer;
    transition: background-color 0.3s;
  }
  
  .google {
    background-color: #0eaf51;
  }
  
  .google:hover {
    background-color: #44d07e;
  }
  
  .github {
    background-color: #333;
  }
  
  .github:hover {
    background-color: #555;
  }
  </style>
  