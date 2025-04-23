<template>
  <div class="main">
    <div class="main__card">
      <div class="main__titulo">
        <h1>Sistema para la administración de usuarios</h1>
        <div class="main__parrafo">
          <p>
            Este es un sistema para la administración de usuarios, donde puedes crear un usuario 
            y también puedes ver la lista de usuarios registrados en el sistema.
          </p>
        </div>
      </div>
      <NuxtLink class="main__link" to="/users">Ver usuarios</NuxtLink>
      <NuxtLink class="main__link" to="/users/new">Crear usuario</NuxtLink>
    </div>
  </div>
</template>