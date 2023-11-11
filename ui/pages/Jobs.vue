<template>
<Navbar />
<JobsTable />
</template>

<script>
import JobsTable from './../components/JobsTable.vue'
import Navbar from './../components/Navbar.vue'
import { defineComponent, ref } from 'vue';
export default defineComponent({
  name: "Jobs",
  components: {JobsTable, Navbar},
  setup(){}
})
</script>