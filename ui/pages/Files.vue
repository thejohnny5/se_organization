<template>
  <div class="flex flex-col">
    <Navbar />
    <div>
      <button class="mt-4 bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded w-1/5" @click="showUploadFile">Upload From CSV</button>
      <UploadPDF ref="UploadPDFRef" />
    </div>
    <FilesTable />
  </div>
</template>

<script>
  import axios from 'axios';
  import FilesTable from './../components/FilesTable.vue'
  import Navbar from './../components/Navbar.vue'
  import UploadPDF from './../components/UploadPDF.vue'
  import { useRouter } from 'vue-router';
  import { defineComponent, ref, onMounted } from 'vue';
  export default defineComponent({
    name: "Jobs",
    components: {FilesTable, Navbar, UploadPDF},
    setup(){
      const validateCred = () => {
        axios.get('/api/validate')
        .then(()=>{})
        .catch(err=>{
          console.error('Not logged in');
          router.push("/");
        })
      }

      onMounted(validateCred)
      const UploadPDFRef = ref(null);

      const showUploadFile = () => {
          if (UploadPDFRef.value){
            UploadPDFRef.value.openPopup();
          }
        }
      return {
        UploadPDFRef,
        showUploadFile
      }
    }})

</script>