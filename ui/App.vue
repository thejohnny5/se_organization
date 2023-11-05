<template>
    <div>
      <h1>Hello Vue 3 + Vite</h1>
      <ul>Company: {{ data.company }}</ul>
      <ul>Title: {{ data.title }}</ul>
      <ul>Location: {{ data.location }}</ul>
      <ul>Notes: {{ data.notes }}</ul>

      <input v-model="data.company">
    </div>
</template>
  
<script>
import { ref, defineComponent } from 'vue';

export default defineComponent({
  name: 'YourComponentName', // replace with your actual component name
  setup() {
    // Define the reactive data inside the setup function
    const data = ref({
      "ID": "",
      "company": "",
      "title": "",
      "location": "",
      "notes": "",
    });

    // Return the reactive data to the template
    return {
      data
    };
  }
});
</script>

<style>
/* global styles */
</style>
