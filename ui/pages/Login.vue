<template>
  <div class="flex justify-center items-center h-screen">
    <div class="text-center p-10 border-4 shadow-md bg-slate-500 space-y-6 max-w-sm">
      <h1 class="text-lg font-bold">Welcome</h1>
      <p>Please login using one of the following methods:</p>

      <!-- Google OAuth Button -->
      <button @click="loginWithGoogle" class="w-full p-2 mb-2 border-none border-r-4 text-sm cursor-pointer bg-red-700">
        Login with Google
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
  
      const isLoggedIn = () => {
        isLoading.value = true;
        axios.get("/api/validate")
        .then((res)=>{
          if (res.status == 200) router.push("/home")
        }).catch((err)=>{
        }).finally(()=>isLoading.value=false)
     
        }
      
  
      onMounted(isLoggedIn);
  
      // Return the functions and state you want to use in your template
      return {
        isLoading,
        loginWithGoogle,
        loginWithGithub,
      };
    },
  };
  </script>