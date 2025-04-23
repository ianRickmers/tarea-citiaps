
<script setup>
    // ref es usado para crear un dom reactivo
    import { ref } from 'vue'
    import { useFetch } from '#app'
    
    const user = ref({
        name: '',
        email: '',
        birthdate: ''
    })
    
    const errorMessage = ref('')
    const successMessage = ref('')
    
    const submitForm = async () => {
        errorMessage.value = ''
        successMessage.value = ''

        // Cambiamos el formato de la fecha a la de Time
        user.value.birthdate = new Date(user.value.birthdate).toISOString()

        console.log(user.value)

        const { error } = await useFetch('http://localhost:8081/users', {
        method: 'POST',
        body: user.value
        })
    
        if (error.value) {
        errorMessage.value = 'Error al crear usuario.'
        } else {
        successMessage.value = 'Usuario creado exitosamente.'
        user.value = { name: '', email: '', birthdate: '' }
        }
    }
</script>

<template>
    <div class ="main">
        <div class="main__card">
            <h1 class="main__titulo">Crear nuevo usuario</h1>
            <p class="main__parrafo">
                Completa el siguiente formulario para crear un nuevo usuario.
                Asegúrate de ingresar un nombre, correo electrónico y fecha de nacimiento válidos.
            </p>
                
            <!-- Hacemos el submit y evitamos que se recargue la página sola, luego se ejecuta la función submitForm -->
            <form @submit.prevent="submitForm">
                <div class="form-group">
                    <label for="name">Nombre</label>
                    <input id="name" v-model="user.name" type="text" required />
                </div>
        
                <div class="form-group">
                    <label for="email">Correo</label>
                    <input id="email" v-model="user.email" type="email" required />
                </div>
        
                <div class="form-group">
                    <label for="birthdate">Fecha de nacimiento</label>
                    <input id="birthdate" v-model="user.birthdate" type="date" required />
                </div>
        
                <button type="submit" id="form-button">Crear usuario</button>
            </form>
        
            <p v-if="errorMessage" class="error">{{ errorMessage }}</p>
            <p v-if="successMessage" class="success">{{ successMessage }}</p>
        </div>
    </div>
</template>
    
