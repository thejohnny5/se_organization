<template>
  <div>
    <form @submit.prevent="submitJob">
    <table>
      <thead>
        <tr>
          <th>ID</th>
          <th>Company</th>
          <th>Title</th>
          <th>Location</th>
          <th>Notes</th>
          <th>Created At</th>
          <th>Updated At</th>
        </tr>
      </thead>
      <tbody>
        <td>New Job Application</td>
        <td>
          <input v-model="currJob.company">
        </td>
        <td>
          <input v-model="currJob.title">
        </td>
        <td>
          <input v-model="currJob.location">
        </td>
        <td>
          <input v-model="currJob.notes">
        </td>
        <td>
          None
        </td>
        <td>
          None
        </td>
        <td>
          <button type="submit">Submit</button>
        </td>
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
            {{ job.created_at }}
          </td>
          <td>
            {{ job.updated_at }}
          </td>
        </tr>
      </tbody>
    </table>
  </form>
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

    return {
      jobs,
      updateJob,
      currJob,
      submitJob
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
