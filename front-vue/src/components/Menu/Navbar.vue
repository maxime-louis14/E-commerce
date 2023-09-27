<script setup>
import { RouterLink, RouterView } from "vue-router";
</script>
<template>
  <nav class="bg-gray-800 text-white p-4">
    <div class="flex justify-between items-center">
      <div class="flex space-x-4">
        <RouterLink
          to="/"
          class="hover:bg-gray-600 px-4 py-2 rounded"
          active-class="bg-gray-600"
        >
          Accueil
        </RouterLink>
        <RouterLink
          to="/boutique"
          class="hover:bg-gray-600 px-4 py-2 rounded"
          active-class="bg-gray-600"
        >
          Boutique
        </RouterLink>
        <RouterLink
          to="/collections"
          class="hover:bg-gray-600 px-4 py-2 rounded"
          active-class="bg-gray-600"
        >
          Collections
        </RouterLink>
        <RouterLink
          to="/profile"
          class="hover:bg-gray-600 px-4 py-2 rounded"
          active-class="bg-gray-600"
        >
          Profil
        </RouterLink>
      </div>

      <div class="flex items-center space-x-2">
        <!-- Image à gauche de "Inscription" avec un espacement de 10px -->
        <img
          class="h-16 w-16 object-cover rounded-full mr-10"
          src="https://images.unsplash.com/photo-1580489944761-15a19d654956?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=1361&q=80"
          alt="Current profile photo"
        />
        <RouterLink
          to="/inscription"
          class="hover:bg-gray-600 px-4 py-2 rounded"
          active-class="bg-gray-600"
        >
          Inscription
        </RouterLink>
        <RouterLink
          to="/connection"
          class="hover:bg-gray-600 px-4 py-2 rounded"
          active-class="bg-gray-600"
        >
          Connection
        </RouterLink>
      </div>
    </div>
  </nav>
  <RouterView />
</template>
