<template>
  <div v-if="isPopupVisible" class="fixed inset-0 bg-black bg-opacity-50 flex justify-center items-center">
    <div class="bg-slate-600 p-6 rounded-lg max-w-lg w-full mx-4">
      <div class="space-y-4">
        <div>
          <input type="file" @change="handleFile" class="block w-full text-sm text-slate-500
          file:mr-4 file:py-2 file:px-4
          file:rounded-full file:border-0
          file:text-sm file:font-semibold
          file:bg-blue-50 file:text-blue-700
          hover:file:bg-blue-100"/>
        </div>

        <div>
          <label class="block mb-2 text-sm font-bold text-slate-300">Label:</label>
          <input v-model="DocumentHeaders.document_name" class="bg-gray-700 text-white border-none rounded py-2 px-4 w-full"/>
        </div>

        <div>
          <label class="block mb-2 text-sm font-bold text-slate-300">Notes:</label>
          <input v-model="DocumentHeaders.notes" class="bg-gray-700 text-white border-none rounded py-2 px-4 w-full"/>
        </div>

        <button class="mt-4 bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded w-full transition-colors duration-150" @click="submitForm">Submit</button>
      </div>

      <button class="mt-4 bg-gray-500 hover:bg-gray-700 text-white font-bold py-2 px-4 rounded w-full transition-colors duration-150" @click="closePopup">Close</button>
    </div>
  </div>
</template>

  
  <script>
import { defineComponent, ref } from 'vue';
import axios from 'axios';

  export default defineComponent({
    name: 'UploadCSV',
    setup(){
        const docDefault = {
            id: "",
            document_name: "",
            type_of_document: "",
            notes: "",
            file: null,
        }
        const DocumentHeaders = ref(docDefault);
      
        const isPopupVisible = ref(false);

        const openPopup = () => {
          isPopupVisible.value = true;
        }
        const closePopup = () => {
          isPopupVisible.value = false;
        }

        const handleFile = (event) => {
            DocumentHeaders.value.file = event.target.files[0];
        }
        const submitForm = () => {
          if (!DocumentHeaders.value.file) {
            console.log("NO FILE");
            return; // handle no file uploaded popup;
          }
            const formData = new FormData();
            
            formData.append("document_name", DocumentHeaders.value.document_name);
            // formData.append("type_of_document", 1);
            formData.append("notes", DocumentHeaders.value.notes)
            formData.append("file", DocumentHeaders.value.file);
            axios.post('/api/document', formData)
            .then(res=>{
              // close popup
              closePopup();
              // refresh page
              window.location.reload();
            })
            .catch(err=>console.error(err)); // Popup with error
        // Handle the form submission, e.g., send data to a server
      }
      return {
        DocumentHeaders,
        isPopupVisible,
        handleFile,
        openPopup,
        closePopup,
        submitForm
      }
    }

  });
  </script>
  