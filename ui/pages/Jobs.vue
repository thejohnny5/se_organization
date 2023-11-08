<template>
    <div>
      <table>
        <thead>
          <tr>
            <th>ID</th>
            <th>Company</th>
            <th>Title</th>
            <th>Location</th>
            <th>Application Status</th>
            <th>Additional Details</th>
            <!-- <th>Application Type</th> -->
            <!-- <th>URL</th>
            <th>Notes</th>
            <th>Created At</th>
            <th>Updated At</th>
            <th>Actions</th> Added Actions header for the delete button -->
          </tr>
        </thead>
        <tbody>
          <!-- New Job Application Form Row -->
          <tr>
          <td>New Job Application</td>
          <td colspan="7">
            <form @submit.prevent="submitJob">
              <input v-model="currJob.company">
              <input v-model="currJob.title">
              <input v-model="currJob.location">
              <input v-model="currJob.notes">
              <button type="submit">Submit</button>
            </form>
          </td>
        </tr>
  
          <!-- Existing Jobs -->
          <tr v-for="job in jobs" :key="job.ID">
            <td>{{ job.ID }}</td>
            <td>
              <input v-model="job.company" @blur="updateJob(job)" />
            </td>
            <td>
              <input v-model="job.title" @blur="updateJob(job)" />
            </td>
            <td>
              <input v-model="job.location" @blur="updateJob(job)" />
            </td>
            <td>
              <input v-model="job.notes" @blur="updateJob(job)" />
            </td>
            <td>
              <div>Application Type: <input v-model="job.application_type" @blur="updateJob(job)"></div>
              <div>Applied At: <input v-model="job.applied_at" @blur="updateJob(job)"></div>
            </td>
  
            <td>
              <button @click.prevent="deleteJob(job)">Delete</button>
            </td>
          </tr>
        </tbody>
      </table>
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
          console.log(response.data)
          jobs.value = response.data;
        } catch (error) {
          console.error('Failed to fetch jobs', error);
        }
      };
  
      // Call the fetchJobs function on component mount
      onMounted(fetchJobs);
  
      // Update job on the server
      const updateJob = async (updatedJob) => {
        try {
          await axios.put(`/api/job/`, updatedJob);
          // Handle successful update, maybe show a notification
          console.log("Updated")
        } catch (error) {
          console.error('Failed to update job', error);
          // Optionally revert the change or show an error
        }
      };
  
  
      const submitJob = async () => {
        try{
          const result = await axios.post(`/api/job`, currJob.value);
          // Handle success by adding job to top of list
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
        updateJob,
        currJob,
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