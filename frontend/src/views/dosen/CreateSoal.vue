<template>
  <div class="w-full p-6 sm:p-8 mx-auto bg-white rounded-2xl shadow-xl">
    <h2 class="text-2xl font-bold text-dark-text mb-8">Manajemen Soal Ujian</h2>
    
    <div class="p-8 border border-gray-200 rounded-lg">
      <p class="mb-6 font-semibold text-medium-text">Pilih Salah Satu Untuk Membuat Soal Ujian</p>
      <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
        <button @click="uploadJson" class="w-full h-24 text-lg font-bold text-white transition-opacity rounded-lg bg-primary hover:opacity-90">
          UPLOAD FILE JSON
        </button>
        <button @click="createManual" class="w-full h-24 text-lg font-bold text-white transition-opacity rounded-lg bg-primary hover:opacity-90">
          Buat Manual
        </button>
      </div>
    </div>
  </div>
</template>

<script>
// Tidak perlu lagi mengimpor layout atau userIcon di sini

export default {
  name: 'CreateSoal',
  // Tidak ada lagi data atau method untuk layout
  methods: {
    uploadJson() {
      this.$router.push('/dosen/soal/upload-json');
    },
    createManual() {
      this.$router.push('/dosen/soal/create-manual');
    }
  },
};
</script>