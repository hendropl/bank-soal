<template>
  <div class="home-container">
    <header class="navbar">
      <div class="logo">Latih.in</div>
      <nav class="nav-links">
        
        <nav class="nav-links">
          <router-link to="/" class="active">Home</router-link>
          <router-link to="/ujian">Ujian</router-link> </nav>
      </nav>
      
      <div class="profile-container">
        <div class="user-profile" @click="toggleProfileDropdown">
          <img :src="userIcon" alt="User Profile" />
        </div>

        <div v-if="showProfileDropdown" class="profile-dropdown">
          <template v-if="isLoggedIn">
            <div class="dropdown-header">
              <img :src="userIcon" alt="User Profile" class="dropdown-avatar" />
              <div class="user-info">
                <div class="user-name">{{ user.fullName }}</div>
                <div class="user-id">{{ user.studentID }}</div>
              </div>
            </div>
            <hr class="dropdown-divider" />
            <a href="#" class="dropdown-item">My Exam</a>
            <a href="#" class="dropdown-item">Update Profile</a>
            <a href="#" class="dropdown-item" @click="logout">Logout</a>
          </template>
          
          <template v-else>
            <router-link to="/login" class="dropdown-item login-link">Login</router-link>
          </template>
        </div>
      </div>
    </header>

    <main>
      <section class="hero-section">
        <div class="hero-content">
          <h1>
            Ujian Mahasiswa Pindahan
            <span class="logo-hero">Latih.in</span>
          </h1>
          <hr class="separator" />
        </div>
      </section>

      <section class="info-section">
        <h2 class="section-title">Apa itu Latih.in?</h2>
        <div class="info-content">
          <div class="info-image">
            <img :src="illustration1" alt="E-learning Illustration" />
          </div>
          <div class="info-text">
            <p>
              Latih.in adalah platform belajar online yang menyediakan berbagai materi pelajaran dalam format yang menarik dan mudah dipahami. Kami percaya belajar bisa menjadi petualangan yang seru.
            </p>
          </div>
        </div>
      </section>

      <section class="info-section">
        <div class="info-content reverse">
          <div class="info-image">
            <img :src="illustration2" alt="Video Learning Illustration" />
          </div>
          <div class="info-text">
            <p>
              Latih.in dapat diakses kapan saja dan dimana saja melalui perangkat apapun. Kami memberikan video pembelajaran yang dikemas dengan animasi dan interaksi yang menarik untuk membuat belajar jadi lebih menyenangkan.
            </p>
          </div>
        </div>
      </section>
    </main>

    <footer class="footer">
      <div class="footer-wave">
        <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 1440 320">
          <path fill="#e8ecff" fill-opacity="1" d="M0,192L80,176C160,160,320,128,480,138.7C640,149,800,203,960,208C1120,213,1280,171,1360,149.3L1440,128L1440,320L1360,320C1280,320,1120,320,960,320C800,320,640,320,480,320C320,320,160,320,80,320L0,320Z"></path>
        </svg>
      </div>
      <div class="footer-content">
        <div class="logo-footer">Latih.in</div>
        <p>Mari Belajar bersama Latih.in</p>
        <small>&copy; 2024 - Latih.in - All Rights Reserved</small>
      </div>
    </footer>
  </div>
</template>

<script>
import userIcon from '../../assets/user-icon.png';
import illustration1 from '../../assets/illustration-1.png';
import illustration2 from '../../assets/illustration-2.png';

export default {
  name: 'HomePage',
  data() {
    return {
      isLoggedIn: false, 
      showProfileDropdown: false,
      user: null, 
      userIcon,
      illustration1,
      illustration2,
    };
  },
  methods: {
    toggleProfileDropdown() {
      this.showProfileDropdown = !this.showProfileDropdown;
    },
    simulateLogin() {
      this.isLoggedIn = true;
      this.user = {
        fullName: 'nama',
        
      };
      this.showProfileDropdown = false;
    },
    logout() {
      localStorage.removeItem('isLoggedIn'); // Hapus status login
      this.isLoggedIn = false;
      this.user = null;
      this.showProfileDropdown = false;
    },
  },
  mounted() {
  if (localStorage.getItem('isLoggedIn') === 'true') {
    this.simulateLogin();
  }
}
};
</script>

<style scoped>
/* Semua style dari sebelumnya digabung di sini */
@import url('https://fonts.googleapis.com/css2?family=Poppins:wght@400;600;700&display=swap');

.home-container {
  font-family: 'Poppins', sans-serif;
  color: #333;
  background-color: #fff;
}

/* NAVBAR */
.navbar {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 1rem 5%;
  background-color: #fff;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.05);
  position: sticky;
  top: 0;
  z-index: 1000;
}
.logo { font-size: 1.8rem; font-weight: 700; color: #6c63ff; }
.nav-links { display: flex; gap: 2.5rem; }
.nav-links a { text-decoration: none; color: #555; font-weight: 600; }
.nav-links a.active { color: #6c63ff; }
.user-profile img { width: 40px; height: 40px; border-radius: 50%; cursor: pointer; }

/* DROPDOWN */
.profile-container { position: relative; }
.profile-dropdown {
  position: absolute; top: 60px; right: 0;
  background-color: white; border-radius: 10px;
  box-shadow: 0 4px 20px rgba(0,0,0,0.1); width: 280px;
  z-index: 1001; border: 1px solid #f0f0f0;
  overflow: hidden; padding: 10px 0;
}
.dropdown-header { display: flex; align-items: center; padding: 10px 20px; }
.dropdown-avatar { width: 45px; height: 45px; border-radius: 50%; margin-right: 15px; }
.user-info .user-name { font-weight: 600; font-size: 1rem; }
.user-info .user-id { font-size: 0.85rem; color: #888; }
.dropdown-divider { border: none; border-top: 1px solid #eee; margin: 10px 0; }
.dropdown-item {
  display: block; padding: 12px 20px; text-decoration: none;
  color: #333; font-size: 0.95rem; transition: background-color 0.2s;
  cursor: pointer;
}
.dropdown-item:hover { background-color: #f5f5f5; }
.dropdown-item.login-link { text-align: center; font-weight: 600; }

/* HERO SECTION */
.hero-section {
  text-align: center; padding: 4rem 2rem; background-color: #f0f2ff;
  background-image: url("data:image/svg+xml,%3Csvg width='60' height='60' viewBox='0 0 60 60' xmlns='http://www.w3.org/2000/svg'%3E%3Cg fill='none' fill-rule='evenodd'%3E%3Cg fill='%23d3d8fc' fill-opacity='0.4'%3E%3Cpath d='M36 34v-4h-2v4h-4v2h4v4h2v-4h4v-2h-4zm0-30V0h-2v4h-4v2h4v4h2V6h4V4h-4zM6 34v-4H4v4H0v2h4v4h2v-4h4v-2H6zM6 4V0H4v4H0v2h4v4h2V6h4V4H6z'/%3E%3C/g%3E%3C/g%3E%3C/svg%3E");
}
.hero-content h1 { font-size: 3.5rem; font-weight: 700; margin-bottom: 1rem; }
.logo-hero { color: #6c63ff; }
.separator { border: 1px solid #ddd; width: 50%; margin: auto; }

/* INFO SECTION */
.info-section { padding: 4rem 5%; max-width: 1100px; margin: auto; }
.section-title { text-align: center; font-size: 2.5rem; font-weight: 700; margin-bottom: 3rem; }
.info-content { display: flex; align-items: center; gap: 3rem; }
.info-content.reverse { flex-direction: row-reverse; }
.info-image { flex: 1; }
.info-image img { max-width: 100%; height: auto; }
.info-text { flex: 1; font-size: 1.1rem; line-height: 1.8; }

/* FOOTER */
.footer { position: relative; background-color: #e8ecff; color: #333; text-align: center; }
.footer-wave {
  position: absolute; top: 0; left: 0; width: 100%;
  transform: translateY(-99%); pointer-events: none;
}
.footer-wave svg { display: block; }
.footer-content { padding: 8rem 2rem 2rem; }
.logo-footer { font-size: 2.5rem; font-weight: 700; color: #6c63ff; margin-bottom: 1rem; }
.footer-content p { font-size: 1.2rem; margin-bottom: 2rem; }
.footer-content small { font-size: 0.9rem; opacity: 0.8; }

/* RESPONSIVE */
@media (max-width: 768px) {
  .nav-links { display: none; }
  .hero-content h1 { font-size: 2.5rem; }
  .info-content, .info-content.reverse {
    flex-direction: column; text-align: center;
  }
}
</style>