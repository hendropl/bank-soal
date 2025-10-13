<template>
  <div class="min-h-screen bg-gradient-to-br from-blue-50 to-indigo-100 flex items-center justify-center p-4">
    <div class="bg-white rounded-2xl shadow-xl w-full max-w-md p-8">
      <div class="text-center mb-8">
        <div class="inline-flex items-center justify-center w-16 h-16 bg-indigo-600 rounded-full mb-4">
          <GraduationCap class="w-8 h-8 text-white" />
        </div>
        <h1 class="text-3xl font-bold text-gray-800 mb-2">Daftar Akun</h1>
        <p class="text-gray-600">Lengkapi data diri Anda untuk mendaftar</p>
      </div>

      <form @submit.prevent="handleSubmit" class="space-y-5">
        <div v-for="field in fields" :key="field.name">
          <label class="block text-sm font-medium text-gray-700 mb-2">{{ field.label }}</label>
          <div class="relative">

            <Input 
              :v-model="formData[field.name]"
              :title="field.title"
              :place-holder="field.placeholder"
              :required="true"
              id=""
              :icon="field.icon"
            />
          </div>
          <p v-if="errors[field.name]" class="text-red-500 text-sm mt-1">{{ errors[field.name] }}</p>
        </div>
        
        <Button 
          :text="isSubmitting ? 'Mendaftar...' : 'Daftar'"
          :disabled="isSubmitting"
          variant="modern"
          size="medium"
          class="w-full"
          
        />
      </form>

      <p class="text-center text-gray-600 mt-6">
        Sudah punya akun?
        <router-link to="/login" class="text-indigo-600 font-semibold hover:underline cursor-pointer">
          Login di sini
        </router-link>
      </p>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { Mail, Lock, User, BookOpen, Building2, GraduationCap } from 'lucide-vue-next'
import Input from '../../components/ui/Input.vue'
import Button from '../../components/ui/Button.vue'

const formData = ref({
  email: '',
  password: '',
  name: '',
  nim: '',
  major: '',
  faculty: ''
})
1
const errors = ref({})
const isSubmitting = ref(false)

const fields = [
  { name: 'email', title: 'Email', type: 'email', placeholder: 'email@example.com', icon: Mail },
  { name: 'password', title: 'Password', type: 'password', placeholder: 'Minimal 6 karakter', icon: Lock },
  { name: 'name', title: 'Nama Lengkap', type: 'text', placeholder: 'Nama lengkap Anda', icon: User },
  { name: 'nim', title: 'NIM', type: 'text', placeholder: 'Nomor Induk Mahasiswa', icon: BookOpen },
  { name: 'major', title: 'Jurusan', type: 'text', placeholder: 'Contoh: Teknik Informatika', icon: GraduationCap },
  { name: 'faculty', title: 'Fakultas', type: 'text', placeholder: 'Contoh: Fakultas Teknik', icon: Building2 }
]

const handleSubmit = () => {}
</script>
