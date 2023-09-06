<script setup>
import { RouterLink, RouterView } from "vue-router";
</script>

<template>
  <header>
    <div class="wrapper">
      <HelloWorld msg="You did it!" />
      <nav>
        <RouterLink to="/">Home</RouterLink>
        <RouterLink to="/about">About</RouterLink>
        <RouterLink to="/register">Register</RouterLink>
        <RouterLink to="/login">Login</RouterLink>
        <RouterLink to="/boutique">Boutique</RouterLink>
        <RouterLink to="/adminpanel">AdminPanel</RouterLink>
      </nav>
    </div>
  </header> 

  <RouterView />
</template>

<style scoped>
header{
  text-align: center;
}
nav a {
  padding: 0 1rem;
  border-left: 1px solid var(--color-border);
  text-align: center;
}
</style>
