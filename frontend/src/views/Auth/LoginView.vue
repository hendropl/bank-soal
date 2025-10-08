<template>
  <div class="flex items-center justify-center min-h-screen font-sans bg-background">
    <div class="w-full max-w-sm p-8 space-y-6 bg-white rounded-2xl shadow-xl">
      
      <img :src="loginIllustration" alt="Login Illustration" class="w-32 mx-auto mb-6" />
      
      <h2 class="text-4xl font-bold text-center text-dark-text">Login</h2>

      <form @submit.prevent="loginUser" class="space-y-5">
        <div>
          <label for="email" class="block text-sm font-semibold text-medium-text">Username</label>
          <input 
            v-model="email" 
            id="email" 
            type="email" 
            placeholder="Username" 
            required
            class="w-full px-4 py-3 mt-1 text-gray-700 bg-white border border-gray-200 rounded-lg focus:outline-none focus:ring-2 focus:ring-primary"
          />
        </div>

        <div>
          <label for="password" class="block text-sm font-semibold text-medium-text">Password</label>
          <div class="relative">
            <input
              v-model="password"
              id="password"
              :type="showPassword ? 'text' : 'password'"
              placeholder="Password"
              required
              class="w-full px-4 py-3 mt-1 text-gray-700 bg-white border border-gray-200 rounded-lg focus:outline-none focus:ring-2 focus:ring-primary"
            />
            <span class="absolute inset-y-0 right-0 flex items-center pr-4 cursor-pointer" @click="togglePassword">
              <i :class="showPassword ? 'fa fa-eye-slash' : 'fa fa-eye'" class="text-gray-400"></i>
            </span>
          </div>
        </div>

        <button 
          type="submit" 
          class="w-full px-4 py-3 font-bold text-white transition-opacity rounded-lg bg-primary hover:opacity-90"
        >
          Login
        </button>

        <div class="text-sm text-center">
          <router-link to="/forgot-password" class="font-medium text-medium-text hover:text-primary">Lupa Password?</router-link>
        </div>
      </form>

      <p class="pt-4 text-sm text-center text-medium-text">
        Belum punya akun? 
        <router-link to="/register" class="font-semibold text-primary hover:underline">Daftar</router-link>
      </p>
      </div>
  </div>
</template>

<script>
// Pastikan path gambar ini sudah benar
import loginIllustration from '../../assets/login-illustration.png';

export default {
  name: 'LoginView',
  data() {
    return {
      email: '',
      password: '',
      showPassword: false,
      loginIllustration,
    };
  },
  methods: {
    togglePassword() {
      this.showPassword = !this.showPassword;
    },
    loginUser() {
      const dummyStudent = {
        email: 'mahasiswa@latih.in',
        password: 'password123',
      };
      
      if (this.email === dummyStudent.email && this.password === dummyStudent.password) {
        localStorage.setItem('isLoggedIn', 'true');
        this.$router.push('/');
      } else {
        alert('Username atau password yang Anda masukkan salah!');
      }
    },
  },
};
</script>