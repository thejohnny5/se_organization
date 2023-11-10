<!-- <template>
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
      <tr v-for="job in jobs" :key="job.id">
        <td>{{ job.id }}</td>
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
-->