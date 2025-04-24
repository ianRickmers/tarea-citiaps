<script setup>
import { ref } from 'vue'
import UserModal from '@/components/userModal.vue'

const { data: users, error } = await useFetch('http://localhost:8081/users')
const selectedUser = ref(null)

const openModal = (user) => {
  selectedUser.value = user
}

const closeModal = () => {
  selectedUser.value = null
}
</script>

<template>
  <div class="main">
    <div class="main__card">
      <h1 class="main__titulo">Usuarios</h1>

      <table class="main__table">
        <thead>
          <tr>
            <th>ID</th>
            <th>Nombre</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="user in users" :key="user.id" @click="openModal(user)" style="cursor: pointer;">
            <td>{{ user.id }}</td>
            <td>{{ user.name }}</td>
          </tr>
        </tbody>
      </table>
    </div>

    <UserModal v-if="selectedUser" :user="selectedUser" @close="closeModal" />
  </div>
</template>
