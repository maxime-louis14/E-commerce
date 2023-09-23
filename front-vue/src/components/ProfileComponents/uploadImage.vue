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
          accept="image/*"
          @change="handleFileChange"
          class="border p-2 w-full"
        />
      </div>
      <div class="mb-4">
        <button
          type="submit"
          class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded-full"
        >
          Télécharger
        </button>
      </div>
    </form>
  </div>
</template>

<script setup>
import { ref, computed } from "vue";

const fileInput = ref(null);
const selectedFile = ref(null);

const handleFileChange = () => {
  selectedFile.value = fileInput.value.files[0];
};

const uploadImage = () => {
  // Ici, vous pouvez envoyer l'image au serveur ou effectuer toute autre opération nécessaire
  if (selectedFile.value) {
    const formData = new FormData();
    formData.append("image", selectedFile.value);

    // Envoyez formData au serveur en utilisant une requête HTTP (par exemple, axios)
    // Assurez-vous de gérer la réponse du serveur en conséquence
  }
};

// Créez une propriété calculée pour obtenir l'URL de l'image sélectionnée
const selectedFileURL = computed(() => {
  return selectedFile.value ? URL.createObjectURL(selectedFile.value) : "";
});
</script>
