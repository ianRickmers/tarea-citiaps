<template>
    <div class="layout">
        <header class="header">
            <NuxtLink class="header__logo" to="/">
            <img src="/perfil-del-usuario.png" alt="Logo" class="header__logo-img" />
            <span class="header__logo-text">Administrador de usuarios</span>
            </NuxtLink>
    
            <nav class="header__nav">
            <NuxtLink class="header__link" to="/users">Ver usuarios</NuxtLink>
            <NuxtLink class="header__link" to="/users/new">Crear usuario</NuxtLink>
            </nav>
        </header>
    
        <main>
            <slot />
        </main>
        <footer class="footer">
            <p class="footer__text">Sistema de administración de usuarios</p>
        </footer>
    </div>
</template>
