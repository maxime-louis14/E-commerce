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
        <!-- Champ de saisie de fichier -->
        <input
          type="file"
          ref="fileInput"
          accept="avatar/*"
          @change="handleFileChange"
          class="border p-2 w-full"
        />
      </div>
      <div class="mb-4">
        <!-- Bouton de soumission du formulaire -->
        <button
          type="submit"
          class="bg-blue-500 hover-bg-blue-700 text-white font-bold py-2 px-4 rounded-full"
        >
          Envoyer
        </button>
      </div>
    </form>
  </div>
</template>

<script setup>
import { ref, computed } from "vue";

const fileInput = ref(null); // Référence à l'élément d'entrée de fichier
const selectedFile = ref(null); // Fichier sélectionné
const error = ref(null); // Variable pour stocker les erreurs

const handleFileChange = () => {
  // Gérer le changement de fichier
  selectedFile.value = fileInput.value.files[0];
};

const uploadImage = async () => {
  if (selectedFile.value) {
    const formData = new FormData(); // Créer un objet FormData pour envoyer le fichier

    formData.append("avatar", selectedFile.value); // Ajouter le fichier au formulaire

    try {
      const token = localStorage.getItem("Token"); // Récupérer le jeton JWT du stockage local
      if (!token) {
        const errorMessage = "Token JWT non trouvé dans le stockage local.";
        console.error(errorMessage); // Journaliser l'erreur
        error.value = errorMessage; // Stocker l'erreur dans la variable "error"
        return; // Arrêter le traitement
      }

      const headers = new Headers(); // Créer un objet Headers pour les en-têtes de la requête
      headers.append("token", token); // Ajouter l'en-tête d'autorisation avec le jeton JWT

      const response = await fetch("http://localhost:8080/api/users/avatar", {
        // Effectuer une requête HTTP POST vers l'URL du serveur
        method: "POST", // Méthode POST
        body: formData, // Corps de la requête avec le formulaire
        headers: headers // En-têtes de la requête
      });

      if (!response.ok) {
        const errorData = await response.json(); // Obtenir des détails sur l'erreur
        throw new Error(
          `Erreur lors de l'envoi de l'image : ${errorData.message}`
        );
      }

      const data = await response.json(); // Convertir la réponse en données JSON

      fileInput.value.value = null; // Effacer la sélection du fichier
      selectedFile.value = null; // Réinitialiser le fichier sélectionné
    } catch (error) {
      error.value = error.message; // Stocker l'erreur dans la variable "error"
      console.error("Erreur lors de l'envoi de l'image :", error.message); // Gérer les erreurs
    }
  }
};

const selectedFileURL = computed(() => {
  return selectedFile.value ? URL.createObjectURL(selectedFile.value) : "";
  // Créer une URL d'objet pour afficher l'image sélectionnée (ou chaîne vide si aucun fichier sélectionné)
});
</script>
