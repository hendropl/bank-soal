<template>
  <div class="w-full p-6 sm:p-8 mx-auto bg-white rounded-2xl shadow-xl">
    <div class="flex justify-end mb-6">
      <button @click="saveFile" class="flex items-center gap-2 px-4 py-2 text-sm font-semibold border border-gray-300 rounded-md text-medium-text hover:bg-gray-50">
        <i class="fas fa-save"></i> Save
      </button>
    </div>
    
    <div 
      @dragover.prevent="isDragging = true"
      @dragleave.prevent="isDragging = false"
      @drop.prevent="handleDrop"
      :class="isDragging ? 'border-primary bg-indigo-50' : 'border-gray-300'"
      class="flex flex-col items-center justify-center p-12 text-center border-2 border-dashed rounded-lg transition-colors"
    >
      <i class="mb-4 text-4xl text-gray-400 fas fa-cloud-upload-alt"></i>
      <p v-if="!selectedFileName" class="text-lg font-semibold text-dark-text">Drag and Drop here</p>
      <p v-if="!selectedFileName" class="my-2 text-gray-500">or</p>
      
      <p v-if="selectedFileName" class="text-lg font-semibold text-green-600">File Selected:</p>
      <p v-if="selectedFileName" class="mb-4 text-medium-text">{{ selectedFileName }}</p>

      <button @click="triggerFileInput" class="px-8 py-3 font-bold text-white transition-opacity rounded-lg bg-teal-400 hover:opacity-90">
        Select file
      </button>
      <input 
        type="file" 
        ref="fileInput" 
        @change="handleFileSelect" 
        accept=".json" 
        class="hidden" 
      />
    </div>
  </div>
</template>

<script>
export default {
  name: 'UploadJsonSoal',
  data() {
    return {
      isDragging: false,
      selectedFileName: null,
    };
  },
  methods: {
    triggerFileInput() { this.$refs.fileInput.click(); },
    handleFileSelect(event) { this.processFile(event.target.files[0]); },
    handleDrop(event) { this.isDragging = false; this.processFile(event.dataTransfer.files[0]); },
    processFile(file) {
      if (file && file.type === 'application/json') {
        this.selectedFileName = file.name;
        console.log('File JSON yang dipilih:', file);
      } else {
        alert('Hanya file dengan format .json yang diperbolehkan!'); this.selectedFileName = null;
      }
    },
    
    // FUNGSI UNTUK TOMBOL SAVE
    saveFile() {
      if (!this.selectedFileName) {
        alert('Silakan pilih file terlebih dahulu.'); return;
      }
      alert(`Menyimpan file ${this.selectedFileName}...`);
      this.$router.push('/dosen/soal/list');
    }
  }
};
</script>