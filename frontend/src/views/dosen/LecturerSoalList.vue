<template>
  <div class="w-full p-6 sm:p-8 mx-auto bg-white rounded-2xl shadow-xl">
    
    <div class="flex flex-wrap items-center justify-between gap-4 mb-6">
      <div class="flex items-center gap-4">
        <div class="relative">
          <i class="absolute text-gray-400 fas fa-search top-3 left-4"></i>
          <input type="text" placeholder="Search question" class="w-full py-2 pl-12 pr-4 bg-gray-100 border-none rounded-lg focus:outline-none focus:ring-2 focus:ring-primary">
        </div>
        <select class="py-2 pl-4 pr-8 bg-gray-100 border-none rounded-lg appearance-none focus:outline-none focus:ring-2 focus:ring-primary">
          <option>Level of difficult</option>
          <option>Mudah</option>
          <option>Sedang</option>
          <option>Susah</option>
        </select>
      </div>
      <div class="flex items-center gap-4">
        <button class="flex items-center gap-2 px-4 py-2 text-sm font-semibold border border-gray-300 rounded-md text-medium-text hover:bg-gray-50">
          <i class="fas fa-file-import"></i> Import file
        </button>
        <button class="flex items-center gap-2 px-4 py-2 text-sm font-semibold border border-gray-300 rounded-md text-medium-text hover:bg-gray-50">
          <i class="fas fa-file-export"></i> Export
        </button>
        <router-link to="/dosen/soal/create" class="flex items-center gap-2 px-4 py-2 font-semibold text-white transition-opacity rounded-lg bg-primary hover:opacity-90">
          <i class="fas fa-plus-circle"></i> Tambah
        </router-link>
      </div>
    </div>

    <div class="overflow-x-auto">
      <table class="w-full text-left table-auto">
        <thead>
          <tr class="border-b border-gray-200">
            <th class="p-3 font-semibold text-gray-500">ID</th>
            <th class="p-3 font-semibold text-gray-500">Question</th>
            <th class="p-3 font-semibold text-gray-500">Answer</th>
            <th class="p-3 font-semibold text-gray-500">Tingkat Kesulitan</th>
            <th class="p-3 font-semibold text-gray-500">Action</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="soal in soalList" :key="soal.id" class="border-b border-gray-100 last:border-b-0 hover:bg-gray-50">
            <td class="p-3 text-medium-text">{{ soal.id }}</td>
            <td class="p-3 text-medium-text">{{ soal.question }}</td>
            <td class="p-3 font-semibold text-center text-medium-text">{{ soal.answer }}</td>
            <td class="p-3 text-medium-text">{{ soal.difficulty }}</td>
            <td class="p-3">
              <div class="flex items-center gap-4 text-lg">
                <button @click="viewDetails(soal.id)" class="text-blue-500 hover:text-blue-700" title="View Details">
                  <i class="fas fa-eye"></i>
                </button>
                <button @click="editSoal(soal.id)" class="text-green-500 hover:text-green-700" title="Edit Soal">
                  <i class="fas fa-pencil-alt"></i>
                </button>
                <button @click="deleteSoal(soal.id)" class="text-red-500 hover:text-red-700" title="Delete Soal">
                  <i class="fas fa-trash"></i>
                </button>
              </div>
            </td>
          </tr>
        </tbody>
      </table>
    </div>

    <div class="flex items-center justify-end mt-6">
      <nav class="flex items-center gap-2 text-sm">
        <button class="px-3 py-1 border rounded-md hover:bg-gray-100">&lt;</button>
        <button class="px-3 py-1 text-white border rounded-md bg-primary">1</button>
        <button class="px-3 py-1 border rounded-md hover:bg-gray-100">2</button>
        <button class="px-3 py-1 border rounded-md hover:bg-gray-100">3</button>
        <button class="px-3 py-1 border rounded-md hover:bg-gray-100">&gt;</button>
      </nav>
    </div>
  </div>
</template>

<script>
export default {
  name: 'LecturerSoalList',
  data() {
    return {
      // Updated dummy data to match the new design
      soalList: [
        { id: 1, question: 'Who are all ____ people?', answer: 'B', difficulty: 'Susah' },
        { id: 2, question: 'Who are all ____ people?', answer: 'B', difficulty: 'Mudah' },
        { id: 3, question: 'Who are all ____ people?', answer: 'B', difficulty: 'Susah' },
        { id: 4, question: 'Who are all ____ people?', answer: 'B', difficulty: 'Mudah' },
        { id: 5, question: 'Who are all ____ people?', answer: 'B', difficulty: 'nguyenquocsy' }, // Example from screenshot
      ]
    };
  },
  methods: {
    viewDetails(id) {
      alert(`Melihat detail untuk soal ID: ${id}`);
    },
    editSoal(id) {
      alert(`Membuka form edit untuk soal ID: ${id}`);
    },
    deleteSoal(id) {
      if (confirm(`Apakah Anda yakin ingin menghapus soal ID: ${id}?`)) {
        alert('Soal dihapus (simulasi)!');
      }
    }
  },
};
</script>