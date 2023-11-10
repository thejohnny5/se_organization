<template>
  <div id="app" class="bg-gray-800 text-white p-6">
    <!-- Your navbar, sidebar, or other components that should show on all pages can be placed here -->

    <!-- This is where the component for the matched route will be rendered -->
    <router-view></router-view>
    
    <!-- Your footer or other components that should show on all pages can be placed here -->
  </div>
</template>

<script>
export default {
  name: 'App',
  // Your script logic here
};
</script>


<style>
body {
  background-color: #1F2937; /* bg-gray-800 */
  color: #FFFFFF; /* text-white */
}
</style>
