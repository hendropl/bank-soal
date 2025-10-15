<template>
  <div class="w-full max-w-6xl p-8 mx-auto bg-white rounded-2xl shadow-xl">
    <div class="flex items-center gap-5 pb-6 mb-8 border-b border-gray-200">
      <img :src="profileData.imageUrl || defaultIcon" alt="Profile" class="w-16 h-16 rounded-full" />
      <div>
        <h2 class="text-2xl font-bold text-dark-text">{{ profileData.fullName }}</h2>
        <p class="text-medium-text">{{ profileData.email }}</p>
      </div>
    </div>

    <form @submit.prevent="submitForm">
      <div class="grid grid-cols-1 gap-8 md:grid-cols-2">
        <div class="space-y-8">
          <div>
            <label class="block mb-1 text-sm font-semibold text-medium-text">Full Name</label>
            <input type="text" v-model="editableData.fullName" class="w-full py-2 bg-transparent border-b-2 border-gray-200 focus:outline-none focus:border-primary" />
          </div>
          <div>
            <label class="block mb-1 text-sm font-semibold text-medium-text">Gender</label>
            <select v-model="editableData.gender" class="w-full py-2 bg-transparent border-b-2 border-gray-200 focus:outline-none focus:border-primary">
              <option>Pick Your Gender</option>
              <option>Male</option>
              <option>Female</option>
            </select>
          </div>
          <div>
            <button type="submit" class="px-10 py-2 font-bold text-white transition-opacity rounded-lg bg-primary hover:opacity-90">Simpan Perubahan</button>
          </div>
        </div>

        <div class="space-y-8">
          <div>
            <label class="block mb-1 text-sm font-semibold text-medium-text">{{ userRole === 'mahasiswa' ? 'NIM' : 'NIDN' }}</label>
            <input type="text" v-model="editableData.idNumber" class="w-full py-2 bg-transparent border-b-2 border-gray-200 focus:outline-none focus:border-primary" />
          </div>
          <div>
            <label class="block mb-1 text-sm font-semibold text-medium-text">Email</label>
            <input type="email" v-model="editableData.email" class="w-full py-2 bg-transparent border-b-2 border-gray-200 focus:outline-none focus:border-primary" />
          </div>
          <div>
            <label class="block mb-1 text-sm font-semibold text-medium-text">Riwayat</label>
            <textarea rows="5" v-model="editableData.history" class="w-full p-2 bg-transparent border-2 border-gray-200 rounded-lg focus:outline-none focus:border-primary"></textarea>
          </div>
        </div>
      </div>
    </form>
  </div>
</template>

<script>
// Path sudah diperbaiki dari ../../ menjadi ../
import defaultIcon from '../assets/user-icon.png';

export default {
  name: 'UserProfileForm',
  props: {
    profileData: {
      type: Object,
      required: true,
    },
    userRole: {
      type: String,
      default: 'mahasiswa'
    }
  },
  data() {
    return {
      defaultIcon,
      // Buat salinan data agar tidak mengubah data asli secara langsung
      editableData: { ...this.profileData }
    };
  },
  methods: {
    submitForm() {
      // Kirim event 'save' ke parent dengan membawa data yang sudah diubah
      this.$emit('save', this.editableData);
      alert('Data profil disimpan!');
    }
  },
  watch: {
    // Jika data dari parent berubah, update salinan lokal
    profileData(newData) {
      this.editableData = { ...newData };
    }
  }
};
</script>