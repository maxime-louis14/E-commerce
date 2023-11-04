<template>
  <!-- Affichage de l'image sélectionnée -->
  <div v-if="selectedFile">
    <h2 class="text-xl text-center font-semibold mt-4 h-10">
      Image sélectionnée :
    </h2>
    <img
      :src="selectedFileURL"
      alt="Image sélectionnée"
      class="max-w-[300px] mx-auto mt-2 rounded-full"
    />
  </div>

  <div>
    <form @submit.prevent="uploadImage" class="max-w-md mx-auto mt-8">
      <div class="mb-4">
        <input
          type="file"
          ref="fileInput"
          accept="avatar/*"
          @change="handleFileChange"
          class="border p-2 w-full"
        />
      </div>
      <div class="mb-4">
        <button
          type="submit"
          class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded-full"
        >
          Envoyez
        </button>
      </div>
    </form>
  </div>
</template>

<script setup>
import { ref, computed } from "vue";

const fileInput = ref(null); // Référence pour l'élément d'entrée de fichier
const selectedFile = ref(null); // Référence pour le fichier sélectionné

const handleFileChange = () => {
  selectedFile.value = fileInput.value.files[0]; // Met à jour le fichier sélectionné lorsqu'un fichier est choisi
};

const uploadImage = async () => {
  if (selectedFile.value) {
    const formData = new FormData(); // Crée un objet FormData pour envoyer les données du fichier

    formData.append("avatar", selectedFile.value); // Ajoute le fichier au FormData

    try {
      // Récupérez le token JWT depuis le local storage
      const token = localStorage.getItem("Token");
      console.log("Token JWT récupéré depuis le local storage :", token);

      const headers = new Headers();
      headers.append("Authorization", token); // Utilisez le token récupéré dans l'en-tête d'autorisation
      console.log(
        "En-tête de la requête avec le token :",
        headers.get("Authorization")
      );

      // Envoyez formData au serveur en utilisant fetch
      const response = await fetch("http://localhost:8080/api/users/avatar", {
        method: "POST",
        body: formData,
        headers: headers // Ajoute l'en-tête d'autorisation au fetch
      });

      if (!response.ok) {
        // Si la réponse n'est pas OK, lancez une erreur
        throw new Error("Erreur lors de l'envoi de l'image");
      }

      // Gérez la réponse du serveur en conséquence
      const data = await response.json();
      console.log("Réponse du serveur :", data);

      // Effacez le champ de fichier après avoir téléchargé l'image (si nécessaire)
      fileInput.value.value = null;
      selectedFile.value = null;
    } catch (error) {
      // Gérez les erreurs ici
      console.error("Erreur lors de l'envoi de l'image :", error);
    }
  }
};

// Créez une propriété calculée pour obtenir l'URL de l'image sélectionnée
const selectedFileURL = computed(() => {
  return selectedFile.value ? URL.createObjectURL(selectedFile.value) : "";
});
</script>
