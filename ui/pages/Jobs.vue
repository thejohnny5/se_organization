<template>
  <div class="flex flex-col">
    <Navbar />
    <button class="mt-4 bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded w-1/5" @click="downloadCSV">Download as CSV</button>
    <div>
      <button class="mt-4 bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded w-1/5" @click="showUploadFile">Upload From CSV</button>
      <UploadFile ref="uploadFileRef" />
    </div>
    <JobsTable />
  </div>
</template>

<script>
import axios from 'axios';
import JobsTable from './../components/JobsTable.vue'
import Navbar from './../components/Navbar.vue'
import UploadFile from './../components/UploadFile.vue';
import { defineComponent, ref } from 'vue';
export default defineComponent({
  name: "Jobs",
  components: {JobsTable, Navbar, UploadFile},
  setup(){
    const uploadFileRef = ref(null);
          
    const showUploadFile = () => {
          if (uploadFileRef.value){
            uploadFileRef.value.openPopup();
          }
        }

    const downloadCSV = () => {
      axios.get('/api/job/csvdownload', {
        responseType: 'blob'
      })
      .then(res=>{
        const url = window.URL.createObjectURL(new Blob([res.data]));
        const a = document.createElement('a');
        a.href = url;
        a.download = 'jobs.csv';
        document.body.appendChild(a);
        a.click();
        window.URL.revokeObjectURL(url);
        a.remove();
      })
      .catch(err=>console.error("couldn't download csv", err))
    }
    return {
      uploadFileRef,
      downloadCSV,
      showUploadFile
    }
  }
})
</script>