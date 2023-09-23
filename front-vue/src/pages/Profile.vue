<script setup>
import Dashbord from "../components/Menu/Dashbord.vue";
</script>

<template>
  <template v-if="isAuthenticated">
    <Dashbord />
    <div class="fixed bottom-4 right-4">
      <button
        @click="logout"
        class="bg-red-500 hover:bg-red-600 text-white font-bold py-2 px-4 rounded-full"
      >
        Se déconnecter
      </button>
      <!-- Bouton de déconnexion -->
    </div>
  </template>

  <template v-else>
    <p class="text-center text-gray-600 text-lg">
      Veuillez vous connecter pour accéder à cette page.
    </p>
    <router-link
      to="/connection"
      class="block text-center text-blue-500 hover:underline mt-4"
    >
      Se connecter
    </router-link>
  </template>
</template>

<script>
import { ref, onMounted } from "vue";

export default {
  setup() {
    const isAuthenticated = ref(false);
    const checkAuthentication = () => {
      const token = localStorage.getItem("Token");
      console.log("Token récupéré du localStorage :", token);

      if (token !== "") {
        // Vérifie si le token est différent d'une chaîne de caractères vide
        isAuthenticated.value = true;
      } else {
        isAuthenticated.value = false;
        console.log("Utilisateur non authentifié");
        // Redirigez l'utilisateur vers la page de connexion ou prenez d'autres mesures
      }
    };

    const logout = () => {
      localStorage.removeItem("Token");
      isAuthenticated.value = false;
      // Redirigez l'utilisateur vers la page de connexion ou ailleurs si nécessaire
      this.$router.push("/"); // Attention : Vous ne pouvez pas utiliser `this` ici, utilisez plutôt un routage impératif
    };

    // Utilisez onBeforeMount pour retarder le rendu des éléments conditionnels
    onBeforeMount(() => {
      checkAuthentication();
    });

    // Vérifiez l'authentification après le montage du composant
    onMounted(() => {
      checkAuthentication();
    });

    return {
      isAuthenticated,
      logout
    };
  }
};
</script>
