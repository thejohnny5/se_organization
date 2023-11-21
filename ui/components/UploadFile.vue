<template>
    <div v-if="isPopupVisible" >
      <form @submit.prevent="submitForm">
        <input type="file" @change="handleFileUpload" />
        <button class="mt-4 bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded" type="submit">Submit</button>
      </form>
    <div v-if="formDataRef.file">
        <div v-for="(jh, index) in JobHeaders">
            <label>{{ jh.dbField }}<span v-if="jh.dbField=='company' || jh.dbField=='title'">*</span></label>
            <select v-model="jh.csvHeader" class="bg-gray-700 text-white border-none rounded py-2 px-4 w-full">
                <option :value="null">None</option>
                <option v-for="(header, index) in headers" :key="index">{{ header }}</option>
            </select>
        </div>
        
    </div>
    <button class="bg-gray-500 hover:bg-gray-700 text-white font-bold py-2 px-4 rounded w-full" @click="closePopup">Close</button>
    </div>
  </template>
  
  <script>
import { defineComponent, ref } from 'vue';
import axios from 'axios';

  export default defineComponent({
    name: 'UploadCSV',
    setup(){
        const JobHeaders = ref([]);
        const constructJobHeaders = () => {
            const options = ["company", "title", "location", "application_status",
        "application_type", "resume", "cover_letter", "posting_url", "salary_low", "salary_high", "notes"]
        for (let opt of options){
            JobHeaders.value.push({dbField: opt, csvHeader: null})
        }
        }
        constructJobHeaders();
        const isPopupVisible = ref(false);
        const formDataRef = ref({
            name: '',
            file: null
        })
        const headers = ref([]);
        const openPopup = () => {
          isPopupVisible.value = true;
        }
        const closePopup = () => {
          isPopupVisible.value = false;
        }
        const handleFileUpload = (event) => {
        const file = event.target.files[0];
        formDataRef.value.file = file
        if (file) {
            const reader = new FileReader();
            reader.onload = (e) => {
                const content = e.target.result;
                const firstLine = content.split(/\r\n|\n/)[0]; // Split by new line and take the first line
                headers.value = firstLine.split(',');
                console.log(headers.value)
            }
            reader.readAsText(file);
        }
        }

        const submitForm = () => {
          if (!formDataRef.value.file) {
            console.log("NO FILE");
            return; // handle no file uploaded popup;
          }
            const formData = new FormData();
            
            formData.append('headerMapping', JSON.stringify(JobHeaders.value));
            formData.append('uploadFile', formDataRef.value.file);
            console.log(JobHeaders.value)
            axios.post('/api/job/csvupload', formData)
            .then(res=>{
              // close popup
              closePopup();
              // refresh page
              console.log(res);
              window.location.reload();
            })
            .catch(err=>console.error(err)); // Popup with error
        // Handle the form submission, e.g., send data to a server
      }
      return {
        isPopupVisible,
        formDataRef,
        headers,
        JobHeaders,
        openPopup,
        closePopup,
        handleFileUpload,
        submitForm
      }
    }

  });
  </script>
  