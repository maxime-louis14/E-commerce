<template>
    <div>
      <h2>Inscription</h2>
      <form @submit.prevent="login">
        <div>
          <label for="email">Email:</label>
          <input type="email" id="email" v-model="email" required>
        </div>
        <div>
          <label for="password">Mot de passe:</label>
          <input type="password" id="password" v-model="password" required>
        </div>
        <button type="submit">login</button>
      </form>
    </div>
</template>
  
  <script setup>
 
  </script>
  
  <style scoped>
  h2 {
    font-size: 20px;
    margin-bottom: 10px;
  }
  form div {
    margin-bottom: 15px;
  }
  label {
    display: block;
    font-weight: bold;
  }
  input {
    width: 200px;
    padding: 5px;
  }
  button {
    background-color: #4caf50;
    color: white;
    border: none;
    padding: 10px 20px;
    font-size: 16px;
    cursor: pointer;
  }
  </style>
  