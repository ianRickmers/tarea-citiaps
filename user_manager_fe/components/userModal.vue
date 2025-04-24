<script setup>
    defineProps(['user'])
    const emit = defineEmits(['close'])
    
    const close = () => {
        emit('close')
    }
</script>

<template>
    <div class="modal-overlay" @click.self="close">
        <div class="modal">
            <h2>Información del usuario</h2>
            <p><strong>ID:</strong> {{ user.id }}</p>
            <p><strong>Nombre:</strong> {{ user.name }}</p>
            <p><strong>Email:</strong> {{ user.email }}</p>
            <p><strong>Fecha de nacimiento:</strong> {{ new Date (user.birthdate).toLocaleDateString() }}</p>
            <p><strong>Creado:</strong> {{ new Date(user.created_at).toLocaleDateString() }}</p>
            <button class="button" @click="close">Cerrar</button>
        </div>
    </div>
</template>