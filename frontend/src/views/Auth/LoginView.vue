<template>
  <div class="login-page">
    <div class="login-card">
      <img :src="loginIllustration" alt="Login Illustration" class="login-image" />
      <h2 class="login-title">Login</h2>

      <form @submit.prevent="loginUser" class="login-form">
        <div class="form-group">
          <label>Email</label>
          <input v-model="email" type="email" placeholder="contoh: mahasiswa@latih.in" required />
        </div>
        <div class="form-group">
          <label>Password</label>
          <div class="password-wrapper">
            <input
              v-model="password"
              :type="showPassword ? 'text' : 'password'"
              placeholder="contoh: password123"
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

      <p class="signup-link">
        Don't have an account? 
        <router-link to="/register">Sign Up</router-link>
      </p>
      </div>
  </div>
</template>

<script>
import loginIllustration from '../../assets/login-illustration.png';

// Data dummy diletakkan di sini untuk simulasi
const dummyStudent = {
  email: 'mahasiswa@latih.in',
  password: 'password123',
};

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
      // Memeriksa apakah email dan password sesuai dengan data dummy
      if (
        this.email === dummyStudent.email &&
        this.password === dummyStudent.password
      ) {
        // Jika login berhasil:
        // 1. Simpan status login di localStorage
        localStorage.setItem('isLoggedIn', 'true');
        
        // 2. Arahkan ke halaman utama
        this.$router.push('/');

      } else {
        // Jika gagal, tampilkan pesan error
        alert('Email atau password yang Anda masukkan salah!');
      }
    },
  },
};
</script>

<style scoped>
  .login-page {
    background-color: #e8ecff;
    min-height: 100vh;
    width: 100%;
    display: flex;
    justify-content: center;
    align-items: center;
    overflow: hidden;
  }
  .login-card {
    background: #fff;
    border-radius: 15px;
    padding: 40px 50px;
    width: 400px;
    text-align: center;
    box-shadow: 0 4px 20px rgba(0, 0, 0, 0.1);
  }
  .login-image {
    width: 150px;
    margin-bottom: 20px;
  }
  .login-title {
    font-size: 28px;
    font-weight: bold;
    margin-bottom: 25px;
    color: #222;
  }
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
    box-sizing: border-box;
  }
  input:focus {
    border-color: #5a67d8;
  }
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

  /* ==== STYLE BARU DITAMBAHKAN DI SINI ==== */
  .signup-link {
    margin-top: 25px;
    font-size: 0.9rem;
    color: #555;
  }
  .signup-link a {
    color: #5a67d8;
    font-weight: 600;
    text-decoration: none;
  }
  .signup-link a:hover {
    text-decoration: underline;
  }
  /* ======================================= */
</style>