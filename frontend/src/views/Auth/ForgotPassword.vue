<template>
  <div class="flex items-center justify-center min-h-screen font-sans bg-background">
    <div class="relative w-full max-w-md p-10 text-center bg-white rounded-2xl shadow-xl">
      
      <router-link to="/login" class="absolute text-3xl no-underline text-primary top-6 left-6">&#8592;</router-link>

      <h2 class="mt-10 mb-10 text-3xl font-bold tracking-wider text-dark-text">FORGOT PASSWORD</h2>

      <form @submit.prevent="handleConfirm" class="w-full">
        <div class="mb-6 text-left">
          <label for="email" class="block mb-2 text-sm font-semibold text-medium-text">Email</label>
          <input
            id="email"
            v-model="email"
            type="email"
            placeholder="Email"
            required
            class="w-full px-4 py-3 text-base bg-white border border-gray-200 rounded-lg placeholder-gray-400 focus:outline-none focus:ring-2 focus:ring-primary"
          />
        </div>

        <button type="submit" class="w-full py-3 font-bold text-white transition-opacity rounded-lg bg-primary hover:opacity-90">Confirm</button>
      </form>
    </div>
  </div>
</template>

<script>
// Bagian script Anda tidak diubah
export default {
  name: 'ForgotPassword',
  data() { return { email: '' }; },
  methods: {
    handleConfirm() {
      if (this.email) {
        alert(`Tautan reset password telah dikirim ke: ${this.email}`);
      } else {
        alert('Silakan masukkan email Anda.');
      }
    },
  },
};
</script>