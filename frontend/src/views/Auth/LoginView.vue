<template>
  <div class="login-page">
    <div class="login-card">
      <!-- Gunakan binding :src agar bisa pakai variabel dari script -->
      <img :src="loginIllustration" alt="Login Illustration" class="login-image" />

      <h2 class="login-title">Login</h2>

      <form @submit.prevent="loginUser" class="login-form">
        <div class="form-group">
          <label>Email</label>
          <input v-model="email" type="email" placeholder="Username" required />
        </div>

        <div class="form-group">
          <label>Password</label>
          <div class="password-wrapper">
            <input
              v-model="password"
              :type="showPassword ? 'text' : 'password'"
              placeholder="Password"
              required
            />
            <span class="toggle" @click="togglePassword">
              <i :class="showPassword ? 'fa fa-eye-slash' : 'fa fa-eye'"></i>
            </span>
          </div>
        </div>

        <button type="submit" class="btn-login">Login</button>

        <router-link to="/forgot-password" class="forgot-link">Forgot Password?</router-link>
      </form>
    </div>
  </div>
</template>

<script>
import loginIllustration from '../../assets/login-illustration.png'

export default {
  name: 'LoginView',
  data() {
    return {
      email: '',
      password: '',
      showPassword: false,
      loginIllustration, // simpan gambar ke data agar bisa dipanggil di template
    }
  },
  methods: {
    togglePassword() {
      this.showPassword = !this.showPassword
    },
    loginUser() {
      alert(`Login dengan Email: ${this.email}`)
      // nanti bisa diarahkan ke backend autentikasi
    },
  },
}
</script>

<style scoped>
  /* ==== Latar belakang halaman login ==== */
  .login-page {
    background-color: #e8ecff;
    min-height: 100vh; /* biar full tinggi layar */
    width: 100%; /* bukan 100vw supaya tidak overflow */
    display: flex;
    justify-content: center; /* tengah secara horizontal */
    align-items: center; /* tengah secara vertikal */
    overflow: hidden; /* cegah elemen keluar layar */
  }

  /* ==== Kotak putih ==== */
  .login-card {
    background: #fff;
    border-radius: 15px;
    padding: 40px 50px;
    width: 400px;
    text-align: center;
    box-shadow: 0 4px 20px rgba(0, 0, 0, 0.1);
  }

  /* ==== Gambar logo ==== */
  .login-image {
    width: 150px;
    margin-bottom: 20px;
  }

  /* ==== Judul ==== */
  .login-title {
    font-size: 28px;
    font-weight: bold;
    margin-bottom: 25px;
    color: #222;
  }

  /* ==== Input dan Label ==== */
  .form-group {
    text-align: left;
    margin-bottom: 20px;
  }

  label {
    display: block;
    font-weight: 600;
    margin-bottom: 6px;
    color: #333;
  }

  input {
    width: 100%;
    padding: 10px 12px;
    border-radius: 8px;
    border: 1px solid #ddd;
    outline: none;
    font-size: 15px;
    box-sizing: border-box; /* <<< INI PERBAIKANNYA */
  }

  input:focus {
    border-color: #5a67d8;
  }

  /* ==== Password toggle ==== */
  .password-wrapper {
    position: relative;
  }

  .toggle {
    position: absolute;
    right: 10px;
    top: 10px;
    cursor: pointer;
    color: #888;
  }

  /* ==== Tombol Login ==== */
  .btn-login {
    width: 100%;
    background-color: #5a67d8;
    color: white;
    border: none;
    border-radius: 8px;
    padding: 12px;
    font-weight: bold;
    cursor: pointer;
    transition: 0.3s;
  }

  .btn-login:hover {
    background-color: #4c51bf;
  }

  /* ==== Link Lupa Password ==== */
  .forgot-link {
    display: block;
    margin-top: 15px;
    color: #2d3748;
    text-decoration: none;
    font-weight: 600;
  }

  .forgot-link:hover {
    text-decoration: underline;
  }

  /* ==== Pastikan halaman penuh dan tidak bisa digeser ==== */
  :global(html),
  :global(body) {
    margin: 0;
    padding: 0;
    width: 100%;
    height: 100%;
    overflow-x: hidden; /* cegah geser kanan-kiri */
    overflow-y: auto;
    background-color: #e8ecff;
  }
</style>