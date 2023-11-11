<template>
    <div class="bg-gray-800 text-white p-6">
      <table class="min-w-full leading-normal">
        <tr class="bg-gray-700">
          <th class="w-1/8 px-4 py-2">Date</th>
          <th class="w-1/6 px-4 py-2">Company</th>
          <th class="w-1/4 px-4 py-2">Title</th>
          <th class="w-1/6 px-4 py-2">Location</th>
          <th class="w-1/7 px-4 py-2">Application Status</th>
          <th class="w-1/7 px-4 py-2">Application Type</th>
          <th class="w-1/6 px-4 py-2">Salary</th>
          <th class="w-1/4 px-4 py-2">Notes</th>
          <th class="w-1/8 px-4 py-2">Actions</th>
        </tr>
        <!-- This is for the submission form -->
        <tr v-if="showSubmitRow" class="bg-gray-400 border-b border-gray-700">
  
        <td class="border px-2 py-2">
          <VueDatePicker v-model="jobToSubmit.date">{{jobToSubmit.date}}</VueDatePicker>
        </td>
  
        <td class="border px-2 py-2">
          <input v-model="jobToSubmit.company" class="bg-gray-700 text-white border-none rounded py-2 px-4 w-full" >
        </td>
  
        <td class="border px-2 py-2">
          <input v-model="jobToSubmit.title" class="bg-gray-700 text-white border-none rounded py-2 px-4 w-full" >
          <input v-model="jobToSubmit.posting_url" placeholder="URL" class="bg-gray-700 text-white border-none rounded py-2 px-4 w-full mt-2" >
        </td>
  
        <td class="border px-2 py-2">
          <input v-model="jobToSubmit.location" class="bg-gray-700 text-white border-none rounded py-2 px-4 w-full" >
        </td>
  
        <td class="border px-2 py-2">
          <select v-model="jobToSubmit.application_status_id" class="bg-gray-700 text-white border-none rounded py-2 px-4 w-full">
              <option v-for="status of appStatus" :key="status.id" :value="status.id">{{ status.text }}</option>
          </select> 
        </td>
        <td class="border px-2 py-2">
          <select v-model="jobToSubmit.application_type_id" class="bg-gray-700 text-white border-none rounded py-2 px-4 w-full">
              <option v-for="status of appType" :key="status.id" :value="status.id">{{ status.text }}</option>
          </select> 
        </td>
  
        <td class="border px-1 py-2">
          <input v-model="jobToSubmit.salary_low" type="number" class="bg-gray-700 text-white border-none rounded py-2 px-4 w-full" />
          <input v-model="jobToSubmit.salary_high" type="number" class="bg-gray-700 text-white border-none rounded py-2 px-4 w-full mt-2" />
        </td>
  
        <td class="border px-2 py-2">
          <input v-model="jobToSubmit.notes" class="bg-gray-700 text-white border-none rounded py-2 px-4 w-full" >
        </td>
  
        <td class="border px-2 py-2">
          <button class="bg-green-500 hover:bg-green-700 text-white font-bold py-2 px-4 rounded w-full" @click="submitJob">Submit</button>
          <button class="bg-red-500 hover:bg-red-700 text-white font-bold py-2 px-4 rounded w-full" @click="clearJobForm">Cancel</button>
        </td>
      </tr>
  
        <tr v-for="(job, index) in jobs" :key="job.id" class="bg-gray-600 border-b border-gray-700">
          <td class="border px-2 py-2">
            <div v-if="!job.isEditing">
              {{ job.date }}
            </div>
            <div v-else>
              <VueDatePicker v-model="job.editData.date">{{ job.editData.date }}</VueDatePicker>
            </div>
          </td>
          <td class="border px-2 py-2">
            <div v-if="!job.isEditing">
              {{ job.company }}
            </div>
            <div v-else>
              <input v-model="job.editData.company" class="bg-gray-700 text-white border-none rounded py-2 px-4 w-full" />
            </div>
          </td>
  
  
          <td class="border px-2 py-2">
            <div v-if="!job.isEditing">
              <a :href="formatUrl(job.posting_url)" target="_blank">{{ job.title }}</a>
            </div>
            <div v-else>
              <input v-model="job.editData.title" class="bg-gray-700 text-white border-none rounded py-2 px-4 w-full" />
              <input v-model="job.editData.posting_url" placeholder="URL" class="bg-gray-700 text-white border-none rounded py-2 px-4 w-full mt-2" />
            </div>
          </td>
  
          <td class="border px-2 py-2">
            <div v-if="!job.isEditing">
              {{ job.location }}
            </div>
            <div v-else>
              <input v-model="job.editData.location" class="bg-gray-700 text-white border-none rounded py-2 px-4 w-full" />
            </div>
          </td>
  
          <td class="border px-4 py-2" v-if="!job.isEditing">{{ job.application_status }}</td>
          <td v-else>
            <select v-model="job.editData.application_status_id" class="bg-gray-700 text-white border-none rounded py-2 px-4 w-full">
              <option v-for="status of appStatus" :key="status.id" :value="status.id">{{ status.text }}</option>
            </select>
          </td>
  
          <td class="border px-4 py-2" v-if="!job.isEditing">{{ job.application_type }}</td>
          <td v-else>
            <select v-model="job.editData.application_type_id" class="bg-gray-700 text-white border-none rounded py-2 px-4 w-full">
              <option v-for="status of appType" :key="status.id" :value="status.id">{{ status.text }}</option>
            </select>
          </td>
  
          <td class="border px-1 py-2">
            <div v-if="!job.isEditing">${{ job.salary_low }}-{{ job.salary_high }}</div>
            <div v-else>
              <input v-model="job.editData.salary_low" type="number" class="bg-gray-700 text-white border-none rounded py-2 px-4 w-full" />
              <input v-model="job.editData.salary_high" type="number" class="bg-gray-700 text-white border-none rounded py-2 px-4 w-full mt-2" />
            </div>
          </td>
  
          <td class="border px-1 py-1">
            <div v-if="!job.isEditing">{{ job.notes }}</div>
            <div v-else>
              <input v-model="job.editData.notes" class="bg-gray-700 text-white border-none rounded py-2 px-4 w-full" />
            </div>
          </td>
  
          <td class="border">
            <button class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded w-full" v-if="!job.isEditing" @click="toggleEdit(index)">Edit</button>
            <button class="bg-gray-500 hover:bg-gray-700 text-white font-bold py-2 px-4 rounded w-full" v-if="!job.isEditing" @click="deleteJob(index)">Delete</button>
            <button class="bg-green-500 hover:bg-green-700 text-white font-bold py-2 px-4 rounded w-full" v-else @click="updateJob(index)">Save</button>
            <button class="bg-red-500 hover:bg-red-700 text-white font-bold py-2 px-4 rounded w-full" v-if="job.isEditing" @click="toggleEdit(index)">Cancel</button>
          </td>
        </tr>
      </table>
      <button class="mt-4 bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded" @click="toggleSubmitRow">Create New Job Application</button>
    </div>
  </template>
  
  
    <script>
    import { ref, onMounted, defineComponent } from 'vue';
    import axios from 'axios';
    import VueDatePicker from '@vuepic/vue-datepicker'
    import '@vuepic/vue-datepicker/dist/main.css'
    
    export default defineComponent({
      name: 'JobsTable',
      components: { VueDatePicker },
      setup() {
        const jobs = ref([]);
        const showSubmitRow = ref(false);
        const appStatus = ref([]);
        const appType = ref([]);
        const defaultJobForm = {
          "company": "",
          "title": "",
          "location": "",
          "posting_url": "",
          "salary_low": null,
          "salary_high": null,
          "application_status": "",
          "application_status_id": null,
          "application_type_id": null,
          "application_type": "",
          "notes": "",
          "date": null
        }
        const jobToSubmit = ref({
          ...defaultJobForm
        })
        // Fetch jobs from the API
  
        const toggleSubmitRow = () => {
          showSubmitRow.value = !showSubmitRow.value;
        }
  
        const fetchDropDownOptions = async () => {
          try {
            const r = await axios.get('/api/dropdown?tabletype=application_status');
            appStatus.value = r.data;
            const r2 = await axios.get('/api/dropdown?tabletype=application_type');
            appType.value = r2.data;
          } catch (err) {
            console.error(err);
          }
        }
  
        const fetchJobs = async () => {
          try {
            const response = await axios.get('/api/job');
            jobs.value = response.data.map(job => ({
              ...job,
              isEditing: false,
              editData: {}
            }))
          } catch (error) {
            console.error('Failed to fetch jobs', error);
          }
        };
    
        // Call the fetchJobs function on component mount
        onMounted(fetchJobs);
        onMounted(fetchDropDownOptions)
    
        // Update job on the server
        const toggleEdit = (index) => {
          if (!jobs.value[index].isEditing){
            jobs.value[index].editData = {...jobs.value[index]}
          } else {
            jobs.value[index].editData = {}
          }
          jobs.value[index].isEditing = !jobs.value[index].isEditing;
        };
        
        const updateJob = (index) => {
          const job = jobs.value[index].editData
          axios.put("/api/job", jobs.value[index].editData)
          .then(()=>{
            // update data
            console.log(job)
            jobs.value[index] = jobs.value[index].editData
            const newJobStatus = appStatus.value.filter(element=>job.application_status_id===element.id)[0];
            jobs.value[index].application_status = newJobStatus? newJobStatus.text : undefined
            
            const newJobType = appType.value.filter(element=>job.application_type_id===element.id)[0]
  
            jobs.value[index].application_type = newJobType ? newJobType.text : undefined;
  
            // jobs.value[index].application_status = appStatus.value.filter(element=>job.application_status_id===element.id)[0].text;
            // jobs.value[index].application_type = appType.value.filter(element=>job.application_type_id===element.id)[0].text;
            jobs.value[index].editData = {};
            jobs.value[index].isEditing = false;
            console.log(jobs.value[index])
          }).catch(err=>{
            console.error(err)
          })
        }
  
        const clearJobForm = () => {
          toggleSubmitRow();
          jobToSubmit.value = {...defaultJobForm}
  
        }
    
        const submitJob = async () => {
          try{
            console.log(jobToSubmit.value)
            const result = await axios.post("/api/job", jobToSubmit.value);
            // Handle success by adding job to top of list
            const newjob = result.data;
            const newJobStatus = appStatus.value.filter(element=>jobToSubmit.value.application_status_id===element.id)[0];
            newjob.application_status = newJobStatus? newJobStatus.text : undefined
            
            const newJobType = appType.value.filter(element=>jobToSubmit.value.application_type_id===element.id)[0]
  
            newjob.application_type = newJobType ? newJobType.text : undefined;
            newjob.editData = {};
            newjob.isEditing = false;
            jobs.value.unshift(result.data);
            // set currjob back to defaults
            // calls clearJobForm
            clearJobForm();
          } catch (error) {
            console.error('Failed to post job', error)
          }
        }
    
        const deleteJob = async (index) => {
          try {
            // remove from database
            await axios.delete(`/api/job/${jobs.value[index].id}`)
            // if successful then remove from page
            jobs.value = jobs.value.filter((_val, i) => i != index)
          } catch (error) {
            console.error('Failed to delete job', error)
          }
        }
  
        const formatUrl = (url) => {
        if (typeof url !== "string") return "";
        if (!url.startsWith('http://') && !url.startsWith('https://')) {
          return `https://${url}`;
        }
        return url;
      }
    
        return {
          showSubmitRow,
          jobs,
          appStatus,
          appType,
          jobToSubmit,
          clearJobForm,
          toggleSubmitRow,
          toggleEdit,
          formatUrl,
          updateJob,
          submitJob,
          deleteJob
        };
      }
    });
    </script>
    