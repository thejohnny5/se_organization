<template>
  <div class="bg-gray-800 text-white p-6">
    <table class="min-w-full leading-normal">
      <tr class="bg-gray-700">
        <th class="w-1/6 px-4 py-2">Company</th>
        <th class="w-1/4 px-4 py-2">Title</th>
        <th class="w-1/6 px-4 py-2">Location</th>
        <th class="w-1/7 px-4 py-2">Application Status</th>
        <th class="w-1/7 px-4 py-2">Application Type</th>
        <th class="w-1/6 px-4 py-2">Salary</th>
        <th class="w-1/4 px-4 py-2">Notes</th>
        <th class="w-1/8 px-4 py-2">Actions</th>
      </tr>
      <tr v-for="(job, index) in jobs" :key="job.id" class="bg-gray-600 border-b border-gray-700">
        <td class="border px-4 py-2">
          <div v-if="!job.isEditing">
            {{ job.company }}
          </div>
          <div v-else>
            <input v-model="job.editData.company" class="bg-gray-700 text-white border-none rounded py-2 px-4 w-full" />
          </div>
        </td>


        <td class="border px-4 py-2">
          <div v-if="!job.isEditing">
            <a :href="formatUrl(job.posting_url)" target="_blank">{{ job.title }}</a>
          </div>
          <div v-else>
            <input v-model="job.editData.title" class="bg-gray-700 text-white border-none rounded py-2 px-4 w-full" />
            <input v-model="job.editData.posting_url" placeholder="URL" class="bg-gray-700 text-white border-none rounded py-2 px-4 w-full mt-2" />
          </div>
        </td>

        <td class="border px-4 py-2" v-if="!job.isEditing">{{ job.location }}</td>
        <td v-else><input v-model="job.editData.location" class="bg-gray-700 text-white border-none rounded py-2 px-4 w-full" /></td>

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

        <td class="border px-4 py-2">
          <div v-if="!job.isEditing">{{ job.salary_range }}</div>
          <div v-else>
            <input v-model="job.editData.salary_range" class="bg-gray-700 text-white border-none rounded py-2 px-4 w-full mt-2" />
            <input v-model="job.editData.salary_range" class="bg-gray-700 text-white border-none rounded py-2 px-4 w-full mt-2" />
          </div>
        </td>

        <td class="border px-4 py-2" v-if="!job.isEditing">{{ job.notes }}</td>
        <td v-else><input v-model="job.editData.notes" class="bg-gray-700 text-white border-none rounded py-2 px-4 w-full" /></td>

        <td>
          <button class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded w-full" v-if="!job.isEditing" @click="toggleEdit(index)">Edit</button>
          <button class="bg-green-500 hover:bg-green-700 text-white font-bold py-2 px-4 rounded w-full" v-else @click="updateJob(index)">Save</button>
          <button class="bg-red-500 hover:bg-red-700 text-white font-bold py-2 px-4 rounded w-full" v-if="job.isEditing" @click="toggleEdit(index)">Cancel</button>
        </td>
      </tr>
    </table>
    <button class="mt-4 bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded" @click="createJob">Create New Job Application</button>
  </div>
</template>


  <script>
  import { ref, onMounted, defineComponent } from 'vue';
  import axios from 'axios';
  
  export default defineComponent({
    name: 'JobsTable',
    setup() {
      const jobs = ref([]);

      const appStatus = ref([]);
      const appType = ref([]);

      const currJob = ref({
        "company": "",
        "title": "",
        "location": "",
        "notes": "",
      })
      // Fetch jobs from the API

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
        console.log(jobs.value[index].editData)
        axios.put("/api/job", jobs.value[index].editData)
        .then(()=>{
          // update data
          console.log(jobs.value[index].editData)
          jobs.value[index] = {...jobs.value[index].editData}
          jobs.value[index].isEditing = false;
        }).catch(err=>{
          console.error(err)
        })
      }
  
      const submitJob = async () => {
        try{
          const result = await axios.post("/api/job", currJob.value);
          // Handle success by adding job to top of list
          const newjob = result.data;
          newjob.editData = {...newjob};
          newjob.isEditing = false;
          jobs.value.unshift(result.data);
          // set currjob back to defaults
          Object.assign(currJob.value, {
            "company": "",
            "title": "",
            "location": "",
            "notes": "", 
          })
          console.log("Submitted")
        } catch (error) {
          console.error('Failed to post job', error)
        }
      }
  
      const deleteJob = async (job) => {
        try {
          // remove from database
          await axios.delete(`/api/job/${job.ID}`)
          // if successful then remove from page
          jobs.value = jobs.value.filter(val => val.ID!=job.ID)
        } catch (error) {
          console.error('Failed to delete job', error)
        }
      }

      const formatUrl = (url) => {
      if (!url.startsWith('http://') && !url.startsWith('https://')) {
        return `https://${url}`;
      }
      return url;
    }
  
      return {
        jobs,
        appStatus,
        appType,
        toggleEdit,
        formatUrl,
        updateJob,
        submitJob,
        deleteJob
      };
    }
  });
  </script>
  