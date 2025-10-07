<template>
  <div class="exam-page-container">
    <header class="navbar">
      <router-link to="/" class="logo">Latih.in</router-link>
      <nav class="nav-links">
        <router-link to="/">Home</router-link>
        <router-link to="/ujian" class="active">Ujian</router-link>
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

    <main class="main-content">
      <div class="exam-card">
        <div class="exam-details">
          <h3>Ujian Seleksi Pindahan 2025</h3>
          <p>10 Oktober 2025 09.00 - 10.30 WIB</p>
          <p class="duration">Durasi Ujian : 120 Menit</p>
        </div>
        <div class="exam-rules">
          <h4>Aturan Ujian :</h4>
          <ul>
            <li>Tidak Di Perbolehkan Untuk Keluar dari Ujian Kecuali Selesai/Waktu Ujian Telah Habis</li>
            <li>Timer Akan Tetap Berjalan Internet Anda Terputus</li>
          </ul>
        </div>
        <div class="button-container">
          <button class="start-exam-btn" @click="startExam">Mulai Ujian</button>
        </div>
      </div>
    </main>
  </div>
</template>

<script>
import userIcon from '../../assets/user-icon.png';

export default {
  name: 'ExamPage',
  data() {
    return {
      isLoggedIn: false,
      showProfileDropdown: false,
      user: null,
      userIcon,
    };
  },
  methods: {
    toggleProfileDropdown() {
      this.showProfileDropdown = !this.showProfileDropdown;
    },
    simulateLogin() {
      this.isLoggedIn = true;
      this.user = {
        fullName: 'Nama',
      };
    },
    logout() {
      this.isLoggedIn = false;
      this.user = null;
      this.showProfileDropdown = false;
    },
    startExam() {
      alert('Memulai ujian...');
      // Nanti di sini Anda bisa arahkan ke halaman soal ujian
    }
  },
  mounted() {
    // Logika ini disederhanakan. Di aplikasi nyata, status login
    // akan dikelola secara global (misal: dengan Pinia/Vuex).
    if (localStorage.getItem('isLoggedIn') === 'true') {
      this.simulateLogin();
    }
  }
};
</script>

<style scoped>
@import url('https://fonts.googleapis.com/css2?family=Poppins:wght@400;600;700&display=swap');

.exam-page-container {
  font-family: 'Poppins', sans-serif;
  background-color: #f0f2ff;
  min-height: 100vh;
  display: flex;
  flex-direction: column;
}

/* Kumpulan Style Navbar (Sama seperti HomePage) */
.navbar { display: flex; justify-content: space-between; align-items: center; padding: 1rem 5%; background-color: #fff; box-shadow: 0 2px 10px rgba(0,0,0,0.05); }
.logo { font-size: 1.8rem; font-weight: 700; color: #6c63ff; text-decoration: none; }
.nav-links { display: flex; gap: 2.5rem; }
.nav-links a { text-decoration: none; color: #555; font-weight: 600; }
.nav-links a.active { color: #6c63ff; }
.profile-container { position: relative; }
.user-profile img { width: 40px; height: 40px; border-radius: 50%; cursor: pointer; }
.profile-dropdown { position: absolute; top: 60px; right: 0; background-color: white; border-radius: 10px; box-shadow: 0 4px 20px rgba(0,0,0,0.1); width: 280px; z-index: 1001; border: 1px solid #f0f0f0; overflow: hidden; padding: 10px 0; }
.dropdown-header { display: flex; align-items: center; padding: 10px 20px; }
.dropdown-avatar { width: 45px; height: 45px; border-radius: 50%; margin-right: 15px; }
.user-info .user-name { font-weight: 600; font-size: 1rem; }
.user-info .user-id { font-size: 0.85rem; color: #888; }
.dropdown-divider { border: none; border-top: 1px solid #eee; margin: 10px 0; }
.dropdown-item { display: block; padding: 12px 20px; text-decoration: none; color: #333; font-size: 0.95rem; cursor: pointer; }
.dropdown-item:hover { background-color: #f5f5f5; }
.dropdown-item.login-link { text-align: center; font-weight: 600; }

/* Konten Utama */
.main-content {
  flex-grow: 1;
  display: flex;
  justify-content: center;
  align-items: center;
  padding: 2rem;
}
.exam-card {
  background: #fff;
  border-radius: 15px;
  box-shadow: 0 5px 25px rgba(0, 0, 0, 0.08);
  padding: 2.5rem 3rem;
  width: 100%;
  max-width: 800px;
}
.exam-details {
  background-color: #f8f9fa;
  border-radius: 10px;
  padding: 1.5rem;
  text-align: center;
  margin-bottom: 2.5rem;
}
.exam-details h3 {
  margin: 0 0 0.5rem 0;
  font-size: 1.5rem;
  color: #333;
}
.exam-details p {
  margin: 0.25rem 0;
  color: #555;
}
.exam-details .duration {
  font-weight: 600;
}
.exam-rules h4 {
  font-size: 1.2rem;
  margin-bottom: 1rem;
}
.exam-rules ul {
  list-style: none;
  padding-left: 0;
  color: #666;
}
.exam-rules li {
  margin-bottom: 0.75rem;
}
.exam-rules li::before {
  content: '>';
  margin-right: 10px;
  font-weight: 600;
  color: #5a67d8;
}
.button-container {
  margin-top: 2.5rem;
  display: flex;
  justify-content: flex-end;
}
.start-exam-btn {
  background-color: #28a745;
  color: white;
  border: none;
  border-radius: 8px;
  padding: 0.75rem 1.5rem;
  font-size: 1rem;
  font-weight: 600;
  cursor: pointer;
  transition: background-color 0.3s;
}
.start-exam-btn:hover {
  background-color: #218838;
}
</style>