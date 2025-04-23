<script setup>
const { data: users, error } = await useFetch('http://localhost:8081/users')
console.log(users.value)
</script>

<template>
  <div class="main">
    <div class="main__card">
      <h1 class="main__titulo">Listado de usuarios</h1>

      <table class="main__table">
        <thead>
          <tr>
            <th>ID</th>
            <th>Nombre</th>
            <th>Email</th>
            <th>Fecha Nacimiento</th>
            <th>Creado</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="user in users" :key="user.id">
            <td>{{ user.id }}</td>
            <td>{{ user.name }}</td>
            <td>{{ user.email }}</td>
            <td>{{ new Date(user.birthdate).toLocaleDateString() }}</td>
            <td>{{ new Date(user.created_at).toLocaleDateString() }}</td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>
