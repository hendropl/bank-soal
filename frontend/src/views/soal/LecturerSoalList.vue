<template>
  <div class="w-full p-6 sm:p-8 mx-auto bg-white rounded-2xl shadow-xl">
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
          <tr v-if="loading">
            <td colspan="5" class="p-4 text-center text-gray-500">Loading questions...</td>
          </tr>
          <tr v-else-if="soalList.length === 0">
            <td colspan="5" class="p-4 text-center text-gray-500">No questions found.</td>
          </tr>
          <tr v-for="soal in soalList" :key="soal.id" class="border-b border-gray-100 last:border-b-0 hover:bg-gray-50">
            <td class="p-3 text-medium-text">{{ soal.id }}</td>
            <td class="p-3 text-medium-text">{{ soal.question_text }}</td>
            <td class="p-3 font-semibold text-center text-medium-text">{{ getCorrectAnswerLabel(soal.options) }}</td>
            <td class="p-3 text-medium-text">{{ soal.difficulty }}</td>
            <td class="p-3">
              <div class="flex items-center gap-4 text-lg">
                <button @click="viewDetails(soal.id)" class="text-blue-500 hover:text-blue-700" title="View Details"><i class="fas fa-eye"></i></button>
                <button @click="editSoal(soal.id)" class="text-green-500 hover:text-green-700" title="Edit Soal"><i class="fas fa-pencil-alt"></i></button>
                <button @click="handleDelete(soal.id)" class="text-red-500 hover:text-red-700" title="Delete Soal"><i class="fas fa-trash"></i></button>
              </div>
            </td>
          </tr>
        </tbody>
      </table>
    </div>
    </div>
</template>

<script>
import { getAllQuestions, deleteQuestion } from '../../provider/question.provider';

export default {
  name: 'LecturerSoalList',
  data() {
    return {
      soalList: [],
      loading: true, // State untuk loading
    };
  },
  methods: {
    async fetchQuestions() {
      try {
        this.loading = true;
        const response = await getAllQuestions();
        // Asumsi data ada di response.data
        this.soalList = response.data || []; 
      } catch (error) {
        console.error("Failed to fetch questions:", error);
        alert('Gagal memuat daftar soal.');
      } finally {
        this.loading = false;
      }
    },
    getCorrectAnswerLabel(options) {
      if (!options) return '-';
      const correctOption = options.find(opt => opt.is_correct);
      return correctOption ? correctOption.option_label : '-';
    },
    async handleDelete(id) {
      if (confirm(`Apakah Anda yakin ingin menghapus soal dengan ID: ${id}?`)) {
        try {
          await deleteQuestion(id);
          alert('Soal berhasil dihapus!');
          // Muat ulang daftar soal setelah berhasil dihapus
          this.fetchQuestions();
        } catch (error) {
          console.error("Failed to delete question:", error);
          alert('Gagal menghapus soal.');
        }
      }
    },
    viewDetails(id) { alert(`Melihat detail soal ID: ${id}`); },
    editSoal(id) { alert(`Mengedit soal ID: ${id}`); },
  },
  mounted() {
    this.fetchQuestions(); // Panggil API saat halaman dimuat
  }
};
</script>