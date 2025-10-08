<template>
  <div class="font-sans bg-background min-h-screen">
    <header class="sticky top-0 z-50 flex items-center justify-between px-[5%] py-4 bg-white">
      <router-link to="/" class="text-3xl font-bold no-underline text-primary">Latih.in</router-link>
      <nav class="items-center hidden gap-10 md:flex">
        <router-link to="/" class="font-semibold no-underline text-medium-text hover:text-primary">Home</router-link>
        <router-link to="/ujian" class="font-semibold no-underline text-medium-text hover:text-primary">Ujian</router-link>
      </nav>
      <div class="relative">
        <div @click="toggleProfileDropdown" class="cursor-pointer">
          <img :src="profileImage" alt="User Profile" class="w-10 h-10 rounded-full object-cover" />
        </div>
        <div v-if="showProfileDropdown" class="absolute top-16 right-0 z-50 w-72 p-2 bg-white rounded-lg shadow-xl border border-gray-100">
          <template v-if="isLoggedIn">
            <div class="flex items-center p-2">
              <img :src="profileImage" alt="User Profile" class="w-12 h-12 rounded-full mr-4 object-cover" />
              <div>
                <div class="font-semibold">{{ user.fullName }}</div>
                <div class="text-sm text-gray-500">{{ user.studentID }}</div>
              </div>
            </div>
            <hr class="my-2 border-gray-200" />
            <router-link to="/ujian" class="block w-full px-4 py-2 text-left text-gray-700 rounded hover:bg-gray-100">My Exam</router-link>
            <router-link to="/update-profile" class="block w-full px-4 py-2 text-left text-gray-700 rounded hover:bg-gray-100">Update Profile</router-link>
            <a href="#" @click="logout" class="block w-full px-4 py-2 text-left text-gray-700 rounded hover:bg-gray-100">Logout</a>
          </template>
          <template v-else>
            <router-link to="/login" class="block w-full px-4 py-2 font-semibold text-center text-gray-700 rounded hover:bg-gray-100">Login</router-link>
          </template>
        </div>
      </div>
    </header>

    <main class="p-4 sm:p-8">
      <div class="w-full max-w-6xl p-8 mx-auto bg-white rounded-2xl shadow-xl">
        
        <div class="flex items-center gap-5 pb-6 mb-8 border-b border-gray-200">
          <div class="relative group">
            <img :src="profileImage" alt="User Profile" class="w-20 h-20 rounded-full object-cover" />
            <input 
              type="file" 
              ref="profileImageInput" 
              @change="handleImageUpload" 
              accept="image/*" 
              class="absolute inset-0 w-full h-full opacity-0 cursor-pointer" 
            />
            <div class="absolute inset-0 flex items-center justify-center bg-black bg-opacity-50 rounded-full opacity-0 group-hover:opacity-100 transition-opacity duration-300 cursor-pointer">
              <i class="fas fa-camera text-white text-xl"></i>
            </div>
          </div>
          <div>
            <h2 class="text-2xl font-bold text-dark-text">{{ user ? user.fullName : 'Nama Pengguna' }}</h2>
            <p class="text-medium-text">{{ user ? user.email : 'email@example.com' }}</p>
          </div>
        </div>

        <form @submit.prevent="updateProfile">
          <div class="grid grid-cols-1 gap-8 md:grid-cols-2">
            <div class="space-y-8">
              <div>
                <label for="fullName" class="block mb-1 text-sm font-semibold text-medium-text">Full Name</label>
                <input id="fullName" type="text" v-model="user.fullName" class="w-full py-2 bg-transparent border-b-2 border-gray-200 focus:outline-none focus:border-primary" />
              </div>
              <div>
                <label for="gender" class="block mb-1 text-sm font-semibold text-medium-text">Gender</label>
                <select id="gender" v-model="user.gender" class="w-full py-2 bg-transparent border-b-2 border-gray-200 focus:outline-none focus:border-primary">
                  <option value="">Pick Your Gender</option>
                  <option value="Male">Male</option>
                  <option value="Female">Female</option>
                </select>
              </div>
              <div>
                <h4 class="mb-4 font-semibold text-dark-text">My email Address</h4>
                <div class="flex items-center p-4 border border-gray-200 rounded-lg">
                  <input type="checkbox" checked class="w-5 h-5 rounded text-primary focus:ring-primary" />
                  <div class="ml-4">
                    <p class="font-semibold text-dark-text">{{ user ? user.email : 'email@example.com' }}</p>
                    <p class="text-sm text-gray-400">1 month ago</p>
                  </div>
                </div>
              </div>
              <div>
                <button type="submit" class="px-10 py-2 font-bold text-white transition-opacity rounded-lg bg-primary hover:opacity-90">Simpan</button>
              </div>
            </div>

            <div class="space-y-8">
              <div>
                <label for="nim" class="block mb-1 text-sm font-semibold text-medium-text">NIM</label>
                <input id="nim" type="text" v-model="user.studentID" class="w-full py-2 bg-transparent border-b-2 border-gray-200 focus:outline-none focus:border-primary" />
              </div>
              <div>
                <label for="emailInput" class="block mb-1 text-sm font-semibold text-medium-text">Email</label>
                <input id="emailInput" type="email" v-model="user.email" class="w-full py-2 bg-transparent border-b-2 border-gray-200 focus:outline-none focus:border-primary" />
              </div>
              <div>
                <label for="riwayat" class="block mb-1 text-sm font-semibold text-medium-text">Riwayat</label>
                <textarea id="riwayat" rows="5" v-model="user.history" class="w-full p-2 bg-transparent border-2 border-gray-200 rounded-lg focus:outline-none focus:border-primary"></textarea>
              </div>
            </div>
          </div>
        </form>
      </div>
    </main>
  </div>
</template>

<script>
import userIcon from '../../assets/user-icon.png'; // Default user icon

export default {
  name: 'UpdateProfile',
  data() {
    return {
      isLoggedIn: false,
      showProfileDropdown: false,
      // Default profile image, akan diganti jika user mengupload
      profileImage: userIcon, 
      user: {
        fullName: 'Alexa Rawles',
        studentID: 'G1A023091',
        email: 'alexarawles@gmail.com',
        gender: '', // Menambahkan properti gender
        history: '', // Menambahkan properti history
      },
    };
  },
  methods: {
    toggleProfileDropdown() {
      this.showProfileDropdown = !this.showProfileDropdown;
    },
    // Fungsi untuk menangani upload gambar
    handleImageUpload(event) {
      const file = event.target.files[0];
      if (file) {
        const reader = new FileReader();
        reader.onload = (e) => {
          this.profileImage = e.target.result;
          // Anda bisa menyimpan 'e.target.result' (base64) ke localStorage
          // atau mengirim file 'file' ke server
          localStorage.setItem('userProfileImage', e.target.result);
        };
        reader.readAsDataURL(file);
      }
    },
    simulateLogin() {
      this.isLoggedIn = true;
      // Ambil profile image dari localStorage jika ada
      const savedProfileImage = localStorage.getItem('userProfileImage');
      if (savedProfileImage) {
        this.profileImage = savedProfileImage;
      }
      // Data dummy user
      this.user = {
        fullName: 'Alexa Rawles',
        studentID: 'G1A023091',
        email: 'alexarawles@gmail.com',
        gender: 'Female', // Contoh default
        history: 'Lulusan terbaik 2023', // Contoh default
      };
    },
    logout() {
      localStorage.removeItem('isLoggedIn');
      localStorage.removeItem('userProfileImage'); // Hapus juga profile image saat logout
      this.isLoggedIn = false;
      this.user = null;
      this.profileImage = userIcon; // Reset ke default
      this.showProfileDropdown = false;
      this.$router.push('/');
    },
    updateProfile() {
      // Di sini Anda akan mengirim data 'this.user' yang sudah diupdate ke API
      // Contoh: axios.put('/api/profile', this.user);
      alert('Profil berhasil diperbarui!');
    }
  },
  mounted() {
    if (localStorage.getItem('isLoggedIn') !== 'true') {
      alert('Anda harus login terlebih dahulu');
      this.$router.push('/login');
    } else {
      this.simulateLogin();
    }
  }
};
</script>

<style>
/* Anda mungkin perlu menambahkan Font Awesome untuk ikon kamera */
/* Contoh: <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/5.15.4/css/all.min.css" /> */
</style>