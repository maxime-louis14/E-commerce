<template>
  <div>
    <form
      @submit.prevent="uploadImage"
      class="max-w-md mx-auto mt-8"
      enctype="multipart/form-data"
    >
      <div class="mb-4">
        <!-- Champ de saisie de fichier -->
        <input
          type="file"
          ref="fileInput"
          accept="image/*"
        />
        <!-- Modifier le type          accepté pour les  images   -->
        @change="handleFileChange" class="border p-2 w-full" />
      </div>
      <div class="mb-4">
        <!-- Bouton de soumission du formulaire -->
        <button
          type="submit"
          class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded-full"
        >
          Envoyer
        </button>
      </div>

      <!-- Bouton de téléchargement de fichier JSON -->
      <button
        @click="telechargerJSON"
        class="bg-green-500 hover:bg-green-700 text-white font-bold py-2 px-4 rounded-full"
      >
        Télécharger JSON
      </button>

      <!-- Bouton de téléchargement de fichier CSV -->
      <button
        @click="telechargerCSV"
        class="bg-yellow-500 hover:bg-yellow-700 text-white font-bold py-2 px-4 rounded-full"
      >
        Télécharger CSV
      </button>
    </form>
  </div>
</template>

<script setup>
import { ref } from "vue";

const selectedFile = ref(null);
const selectedFileURL = ref(null);

const handleFileChange = (event) => {
  const file = event.target.files[0];
  if (file) {
    selectedFile.value = file;
    selectedFileURL.value = URL.createObjectURL(file);
  }
};

const uploadImage = async () => {
  try {
    const formData = new FormData();
    formData.append("avatar", selectedFile.value);

    const token = localStorage.getItem("Token");

    if (!token) {
      throw new Error("Token JWT non trouvé dans le stockage local.");
    }

    const headers = new Headers();
    headers.append("token", token);

    const response = await fetch("http://localhost:8080/api/users/avatar", {
      method: "POST",
      headers: headers,
      body: formData
    });

    if (response.ok) {
      const contentType = response.headers.get("Content-Type");

      if (contentType && contentType.startsWith("image/")) {
        selectedFileURL.value = URL.createObjectURL(await response.blob());
      } else {
        const data = await response.json();
        selectedFileURL.value = data.avatarURL;
      }
    } else {
      throw new Error("Échec de l'envoi de l'image de profil");
    }
  } catch (error) {
    console.error(
      "Erreur lors de l'envoi de l'image de profil:",
      error.message
    );
  }
};

const telechargerJSON = async () => {
  // Logique pour envoyer les données JSON au backend
  try {
    const response = await fetch("URL_DU_BACKEND_JSON", {
      method: "POST",
      headers: {
        "Content-Type": "application/json"
      },
      body: JSON.stringify({
        // Ajoutez les données que vous souhaitez envoyer en format JSON
      })
    });

    if (!response.ok) {
      throw new Error("Échec de l'envoi des données JSON au backend");
    }
  } catch (error) {
    console.error(
      "Erreur lors de l'envoi des données JSON au backend:",
      error.message
    );
  }
};

const telechargerCSV = async () => {
  // Logique pour envoyer les données CSV au backend
  try {
    const response = await fetch("URL_DU_BACKEND_CSV", {
      method: "POST",
      headers: {
        "Content-Type": "text/csv"
      },
      body: `Colonne1, Colonne2, Colonne3\n
            Valeur1, Valeur2, Valeur3\n` // Ajoutez les données CSV que vous souhaitez envoyer
    });

    if (!response.ok) {
      throw new Error("Échec de l'envoi des données CSV au backend");
    }
  } catch (error) {
    console.error(
      "Erreur lors de l'envoi des données CSV au backend:",
      error.message
    );
  }
};
</script>
