<template>
    <select multiple="true" v-model="multipleSelection"  v-bind:class="{ 'fix-height': multiple === 'true' }" class="bg-gray-700 text-white border-none rounded py-2 px-4 w-full">
        <option v-for="status of appType" :key="status.id" :value="status.id">{{ status.text }}</option>
    </select>
</template>

<script>
</script>