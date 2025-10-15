<template>
  <div class="grid grid-cols-1 gap-6 lg:grid-cols-2">
    <div class="p-6 bg-white rounded-lg shadow-md">
      <h3 class="mb-4 text-lg font-semibold text-dark-text">Company Facts</h3>
      <div class="h-64 bg-gray-200 rounded-md flex items-center justify-center text-gray-500">Area Chart Placeholder</div>
    </div>
    <div class="p-6 bg-white rounded-lg shadow-md">
      <h3 class="mb-4 text-lg font-semibold text-dark-text">Statistics</h3>
      <div class="h-64 bg-gray-200 rounded-md flex items-center justify-center text-gray-500">Donut Chart Placeholder</div>
    </div>
    <div class="p-6 bg-white rounded-lg shadow-md h-60">
      <h3 class="mb-4 text-lg font-semibold text-dark-text">Assigned Risks</h3>
      <div class="flex items-center justify-center h-full text-gray-400">There are no risks assigned.</div>
    </div>
    <div class="p-6 bg-white rounded-lg shadow-md h-60">
      <h3 class="mb-4 text-lg font-semibold text-dark-text">Assigned Action Items</h3>
      <div class="flex items-center justify-center h-full text-gray-400">There are no action items assigned.</div>
    </div>
  </div>
</template>

<script>
// TIDAK PERLU lagi import DosenLayout di sini
export default {
  name: 'LecturerDashboard',
  // TIDAK PERLU lagi components: { DosenLayout }
};
</script>