<template>
  <div class="flex items-center justify-center min-h-screen font-sans bg-background">
    <div class="relative w-full max-w-sm p-8 space-y-4 text-center bg-white rounded-2xl shadow-xl">
      
      <div class="absolute text-2xl text-gray-400 cursor-pointer top-4 right-4 hover:text-gray-700">&times;</div>
      
      <img :src="loginIllustration" alt="Latih.in Logo" class="w-20 mx-auto mb-2" />
      
      <h2 class="text-3xl font-bold text-dark-text">Akses gratis tanpa batas ke sumber daya kami</h2>
      <p class="text-medium-text">Daftar untuk melihat lebih banyak</p>
      
      <div class="flex flex-col w-full gap-3 pt-4">
        <button class="w-full py-3 font-semibold text-white transition rounded-lg bg-dark-text hover:opacity-90">
          Lanjutkan dengan email
        </button>
        <button class="flex items-center justify-center w-full gap-3 py-3 font-semibold transition bg-white border border-gray-300 rounded-lg text-medium-text hover:bg-gray-50">
          <i class="fab fa-facebook-f text-blue-600"></i>
          <span>Lanjutkan dengan Facebook</span>
        </button>
        <button class="flex items-center justify-center w-full gap-3 py-3 font-semibold transition bg-white border border-gray-300 rounded-lg text-medium-text hover:bg-gray-50">
          <i class="fab fa-google"></i>
          <span>Lanjutkan dengan Google</span>
        </button>
      </div>
      
      
      
      <p class="pt-2 text-sm text-medium-text">
        Sudah Punya Akun? 
        <router-link to="/login" class="font-bold underline text-dark-text">Log in</router-link>
      </p>
    </div>
  </div>
</template>

<script>
// Import ilustrasi ditambahkan kembali di sini
import loginIllustration from '../../assets/login-illustration.png';

export default {
  name: 'RegisterView',
  data() {
    return {
      loginIllustration,
    };
  },
};
</script>