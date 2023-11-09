<template>
  <div class="bg-gray-100 min-h-screen py-8">
    <div class="max-w-6xl mx-auto">
      <h2 class="text-2xl font-semibold mb-6">Ajouter votre avatar</h2>
      <div class="ml-4">
        <h2 class="text-2xl text-center font-semibold mb-6">Avatar actuel</h2>
        <!-- Utilisez la propriété avatar pour afficher l'URL de l'avatar -->
        <img
          v-if="avatarURL !== null"
          :src="avatarURL"
          alt="Image de profil"
          class="max-w-[300px] mx-auto mt-2 rounded-full"
        />
      </div>
    </div>
    <PostImage />
  </div>
</template>

<script setup>
import { ref, onMounted } from "vue";
import PostImage from "../../components/ProfileComponents/PostImage.vue";
import { setAvatarURL } from "../../components/services/AvatarService";

const avatar = ref(null);
const avatarURL = ref(null);
const loading = ref(true);

const getAvatarImage = async () => {
  try {
    loading.value = true;

    const token = localStorage.getItem("Token");

    if (!token) {
      throw new Error("Token JWT non trouvé dans le stockage local.");
    }

    const headers = new Headers();
    headers.append("token", token);

    const response = await fetch("http://localhost:8080/api/users/avatar", {
      method: "GET",
      headers: headers
    });

    if (response.ok) {
      const contentType = response.headers.get("Content-Type");

      if (contentType && contentType.startsWith("image/")) {
        avatarURL.value = URL.createObjectURL(await response.blob());
      } else {
        const data = await response.json();
        avatarURL.value = data.avatarURL;
        avatar.value = URL.createObjectURL(await response.blob());
      }
    } else {
      throw new Error("Échec de la récupération de l'image de profil");
    }
  } catch (error) {
    console.error(
      "Erreur lors de la récupération de l'image de profil:",
      error.message
    );
  } finally {
    loading.value = false;
  }

  setAvatarURL(avatarURL.value);
};

onMounted(() => {
  getAvatarImage();
});
</script>
