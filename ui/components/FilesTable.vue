<template>
  <div class="w-full p-4">
    <div class="overflow-x-auto">
      <table class="min-w-full text-sm text-left text-slate-300">
        <thead class="text-xs text-slate-400 uppercase bg-slate-700">
          <tr>
            <th scope="col" class="px-6 py-3 w-1/12">File ID</th>
            <th scope="col" class="px-6 py-3 w-1/12">Document Name</th>
            <th scope="col" class="px-6 py-3 w-1/12">Type</th>
            <th scope="col" class="px-6 py-3 w-1/6">Notes</th>
            <th scope="col" class="px-6 py-3 w-1/2">Preview</th>
            <th scope="col" class="px-6 py-3 w-1/12">Actions</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="file in files" :key="file.id" class="bg-slate-700 border-b border-slate-500">
            <td class="px-6 py-4">{{ file.id }}</td>
            <td class="px-6 py-4">{{ file.document_name }}</td>
            <td class="px-6 py-4">{{ file.type_of_document }}</td>
            <td class="px-6 py-4">{{ file.notes }}</td>
            <td class="px-6 py-4">
              <object :data="file.download_path" type="application/pdf" class="w-full h-96">
                <p>Your browser does not support PDFs. Please download the PDF to view it: <a :href="file.download_path" class="text-blue-500 hover:text-blue-600">Download PDF</a>.</p>
              </object>
            </td>
            <td class="px-6 py-4 text-center">
              <a :href="file.download_path" target="_blank" class="text-blue-500 hover:text-blue-600">View Larger</a>
            </td>
          </tr>
        </tbody>
      </table>
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