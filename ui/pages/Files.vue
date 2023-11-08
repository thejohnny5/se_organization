<template>
    <div>
        <div v-for="file in files" :key="file.id">
            <div>
                {{ file.id }}
            </div>
            <div>
                {{ file.document_name }}
            </div>
            <div>
                {{ file.type_of_document }}
            </div>
            <div>
                {{ file.notes }}
            </div>
            <object :data="file.download_path" type="application/pdf" width="50%" height="200px">
                <p>Could not load PDF content. Your browser may not have a PDF plugin enabled, or the server is not serving the PDF correctly.</p>
            </object>
        </div>
    </div>
</template>

<script>
  import { ref, onMounted, defineComponent } from 'vue';
  import axios from 'axios';
  
  export default defineComponent({
    name: 'Files',
    setup() {
      const files = ref([]);
    //   const currJob = ref({
    //     "company": "",
    //     "title": "",
    //     "location": "",
    //     "notes": "",
    //   })
      // Fetch jobs from the API
      const fetchJobs = async () => {
        try {
          const response = await axios.get('/api/document');
          files.value = response.data;
          console.log(response.data)
        } catch (error) {
          console.error('Failed to fetch jobs', error);
        }
      };
  
      // Call the fetchJobs function on component mount
      onMounted(fetchJobs);
  
      return {
        files,
      };
    }
  });
</script>