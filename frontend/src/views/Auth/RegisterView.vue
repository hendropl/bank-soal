<template>
  <div class="auth-container">
    <div class="auth-card">
      <img :src="loginIllustration" alt="Latih.in Logo" class="logo" />

      <h2>Unlimited free access to our resources</h2>
      <p>Sign up to see more</p>

      <div class="button-group">
        <button class="btn email">Continue with Email</button>
        <button class="btn facebook">Continue with Facebook</button>
        <button class="btn google">Continue with Google</button>
      </div>

      <p class="agreement">
        By continuing, you agree to the
        <a href="#">Terms of Service</a> and acknowledge our
        <a href="#">Privacy Policy</a>.
      </p>

      <p class="login">
        Already a member?
        <router-link to="/login">Log in</router-link>
      </p>
    </div>
  </div>
</template>

<script>
import loginIllustration from '../../assets/login-illustration.png'

export default {
  name: 'CreateAccountView',
  data() {
    return {
      loginIllustration,
    }
  },
}
</script>

<style scoped>
html, body {
  margin: 0;
  padding: 0;
  height: 100%;
  overflow: hidden; /* 🔒 Tidak bisa digeser */
}

.auth-container {
  position: fixed;
  top: 0;
  left: 0;
  width: 100vw;
  height: 100vh;
  background-color: #e8ecff;
  display: flex;
  justify-content: center;
  align-items: center;
}

/* ✅ Kartu utama (diperkecil agar tidak "zoom") */
.auth-card {
  background: #ffffff;
  width: 85vw;            /* dari 90vw jadi 85vw */
  max-width: 420px;       /* dari 520px jadi 420px */
  border-radius: 20px;
  box-shadow: 0px 4px 16px rgba(0, 0, 0, 0.08);
  padding: 36px 28px;     /* dari 48px 40px jadi lebih kecil */
  text-align: center;
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 16px;              /* jarak antar elemen dikurangi */
}

/* ✅ Logo */
.logo {
  width: 80px;           /* dari 100px dikurangi biar lebih seimbang */
  margin-bottom: 8px;
}

h2 {
  font-size: 22px;       /* sedikit lebih kecil */
  font-weight: 600;
  color: #111;
  margin-bottom: 6px;
}

p {
  color: #333;
  font-size: 15px;
}

.button-group {
  display: flex;
  flex-direction: column;
  gap: 14px;
  margin-top: 18px;
  width: 100%;
}

.btn {
  width: 100%;
  height: 48px;         /* sedikit lebih pendek */
  border-radius: 28px;
  font-size: 15px;
  cursor: pointer;
  transition: 0.2s;
}

.btn.email {
  background: #111;
  color: white;
  border: none;
}

.btn.facebook {
  background: #1877F2;
  color: white;
  border: none;
}

.btn.google {
  background: white;
  color: #444;
  border: 1px solid #ddd;
}

.btn:hover {
  transform: scale(1.02);
}

.agreement {
  font-size: 13px;
  color: #666;
  width: 90%;
  line-height: 1.5;
}

.agreement a {
  color: #333;
  font-weight: 500;
}

.login {
  font-size: 14px;
  color: #333;
}

.login a {
  font-weight: 600;
  color: #111;
}
</style>
