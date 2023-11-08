<template>
  <div class="bg-gray-100 min-h-screen py-8">
    <div class="max-w-6xl mx-auto">
      <h2 class="text-2xl font-semibold mb-6">Ajouter votre avatar</h2>
      <img
        v-if="avatar !== null"
        :src="avatar"
        alt="Image de profil"
        class="max-w-[300px] mx-auto mt-2 rounded-full"
      />
      <div class="ml-4">
        <h2 class="text-2xl text-center font-semibold mb-6">Avatar actuel</h2>
        <img
          class="max-w-[300px] mx-auto mt-2 rounded-full"
          :src="avatar"
          alt="Image de profil actuelle"
        />
      </div>
    </div>
    <UploadImage />
  </div>
</template>

<script setup>
import UploadImage from "../../components/ProfileComponents/uploadImage.vue";
import { ref, onMounted } from "vue";

const avatar = ref(null);

// Fonction pour récupérer l'image de profil actuelle
const getAvatarImage = async () => {
  try {
    const token = localStorage.getItem("Token"); // Récupérer le jeton JWT du stockage local

    if (!token) {
      console.error("Token JWT non trouvé dans le stockage local.");
      return; // Arrêter le traitement si le token n'est pas disponible
    }

    const headers = new Headers();
    headers.append("token", token); // Ajouter le jeton JWT dans l'en-tête

    const response = await fetch("http://localhost:8080/api/users/avatar", {
      method: "GET",
      headers: headers
    });

    if (response.ok) {
      avatar.value = URL.createObjectURL(await response.blob());
    } else {
      console.error("Échec de la récupération de l'image de profil");
    }
  } catch (error) {
    console.error(
      "Erreur lors de la récupération de l'image de profil:",
      error
    );
  }
};

onMounted(() => {
  getAvatarImage(); // Appel de la fonction lors du montage du composant
});
</script>
