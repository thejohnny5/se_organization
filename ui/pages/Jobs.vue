<template>
  <div class="flex flex-col">
    <Navbar />
    <div class="m-5">
    <div class="flex flex-col">
      <div class="flex justify-between items-center">
        <div>
          <button class="bg-blue-600 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded" @click="showUploadFile">Upload From CSV</button>
          <button class="bg-gray-300 hover:bg-gray-400 text-black py-2 px-4 rounded ml-2" @click="downloadCSV">Download as CSV</button>
        </div>
        <button class="bg-green-600 hover:bg-green-700 text-white font-bold py-2 px-4 rounded" @click="showUploadJob">Create New Job Application</button>
      </div>
      <UploadFile ref="uploadFileRef" />
      </div>
    <JobsTable ref="jobTableRef"/>
    </div>
  </div>
</template>

<script>
import axios from 'axios';
import JobsTable from './../components/JobsTable.vue'
import Navbar from './../components/Navbar.vue'
import UploadFile from './../components/UploadFile.vue';
import { useRouter } from 'vue-router';
import { defineComponent, ref, onMounted } from 'vue';
export default defineComponent({
  name: "Jobs",
  components: {JobsTable, Navbar, UploadFile},
  setup(){
    const router = useRouter();
    const uploadFileRef = ref(null);
    const jobTableRef = ref(null);
    const validateCred = () => {
      axios.get('/api/validate')
      .then(()=>{})
      .catch(err=>{
        console.error('Not logged in');
        router.push("/");
      })
    }
    onMounted(validateCred)
    
    const showUploadFile = () => {
          if (uploadFileRef.value){
            uploadFileRef.value.openPopup();
          }
        }

    const showUploadJob = () => {
      if (jobTableRef.value){
        jobTableRef.value.openSubmitRow();
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
      jobTableRef,
      showUploadJob,
      downloadCSV,
      showUploadFile
    }
  }
})
</script>