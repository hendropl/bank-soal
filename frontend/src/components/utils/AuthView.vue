<template>
  <div class="auth-container">
    <div class="auth-card">
      <div class="logo"></div>
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
        Already a member? <a href="#">Log in</a>
      </p>
    </div>
  </div>
</template>

<style scoped>
/* ✅ Layar penuh */
.auth-container {
  width: 100vw;
  height: 100vh;
  background-color: #E6EAFF; /* Warna biru muda lembut */
  display: flex;
  justify-content: center;
  align-items: center;
  overflow: hidden; /* ❌ Supaya tidak bisa digeser */
  margin: 0;
  padding: 0;
}

/* ✅ Kartu utama */
.auth-card {
  background: #ffffff; /* Putih */
  width: 90vw;
  max-width: 544px;
  height: auto;
  border-radius: 24px;
  box-shadow: 0px 6px 20px rgba(0, 0, 0, 0.1);
  padding: 48px;
  text-align: center;
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 20px;
}

/* ✅ Logo placeholder */
.logo {
  width: 64px;
  height: 64px;
  border-radius: 50%;
  background-color: #ccc;
}

/* ✅ Teks utama — dibuat gelap biar kelihatan */
h2 {
  font-size: 24px;
  font-weight: 600;
  color: #111111; /* warna hitam */
  margin-bottom: 8px;
}

p {
  color: #333333;
  font-size: 16px;
}

/* ✅ Tombol */
.button-group {
  display: flex;
  flex-direction: column;
  gap: 16px;
  margin-top: 24px;
}

.btn {
  width: 100%;
  height: 56px;
  border-radius: 32px;
  font-size: 16px;
  cursor: pointer;
  transition: 0.2s;
}

.btn.email {
  background: #111111;
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

/* ✅ Bagian bawah */
.agreement {
  font-size: 14px;
  color: #666;
  width: 90%;
  line-height: 1.5;
}

.agreement a {
  color: #333;
  font-weight: 500;
}

.login {
  font-size: 15px;
  color: #333;
}

.login a {
  font-weight: 600;
  color: #111;
}
</style>
