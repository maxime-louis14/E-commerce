<script setup>
import { RouterLink, RouterView } from "vue-router";
import { getAvatarURL } from "../services/AvatarService";
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
          <!-- Utilisez getAvatarURL() pour récupérer l'URL de l'avatar -->
        <img v-if="getAvatarURL() !== null" :src="getAvatarURL()" alt="Image de profil" />
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
