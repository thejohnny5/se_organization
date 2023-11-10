
<template>
  <div>
    <table>
      <tr>
        <th class="w-1/3 px-4 py-2">Company</th>
        <th class="w-1/2 px-4 py-2">Title</th>
        <th class="w-1/6 px-4 py-2">Location</th>
        <th class="w-1/2 px-4 py-2">URL</th>
        <th class="w-1/3 px-4 py-2">Salary</th>
        <th class="w-1/3 px-4 py-2">Notes</th>
        <th class="w-1/3 px-4 py-2">Actions</th>
      </tr>
      <tr v-for="(job, index) in jobs" :key="job.id">
        <td class="border px-4 py-2 border-2 border-indigo-50" v-if="!job.isEditing">{{ job.company }}</td>
        <td v-else><input v-model="job.editData.company" /></td>

        <td class="border px-4 py-2" v-if="!job.isEditing">{{ job.title }}</td>
        <td v-else><input v-model="job.editData.title" /></td>

        <td class="border px-4 py-2" v-if="!job.isEditing">{{ job.location }}</td>
        <td v-else><input v-model="job.editData.location" /></td>

        <td class="border px-4 py-2" v-if="!job.isEditing">{{ job.posting_url }}</td>
        <td v-else><input v-model="job.editData.posting_url" /></td> 

        <td class="border px-4 py-2" v-if="!job.isEditing">{{ job.salary_range }}</td>
        <td v-else><input v-model="job.editData.salary_range" /></td>

        <td class="border px-4 py-2" v-if="!job.isEditing">{{ job.notes }}</td>
        <td v-else><input v-model="job.editData.notes" /></td>

        <td>
          <button v-if="!job.isEditing" @click="toggleEdit(index)">Edit</button>
          <button v-else @click="updateJob(index)">Save</button>
          <button v-if="job.isEditing" @click="toggleEdit(index)">Cancel</button>
        </td>
      </tr>
    </table>
    <button @click="createJob">Create New Job Application</button>
  </div>
</template>

  <script>
  import { ref, onMounted, defineComponent } from 'vue';
  import axios from 'axios';
  
  export default defineComponent({
    name: 'JobsTable',
    setup() {
      const jobs = ref([]);
      const currJob = ref({
        "company": "",
        "title": "",
        "location": "",
        "notes": "",
      })
      // Fetch jobs from the API
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
          const result = await axios.post(`/api/job`, currJob.value);
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
  
      return {
        jobs,
        toggleEdit,
        currJob,
        updateJob,
        submitJob,
        deleteJob
      };
    }
  });
  </script>
  
  <style>
  /* global styles */
  body {
    background-color: #121212;
    color: #ffffff;
  }
  </style>