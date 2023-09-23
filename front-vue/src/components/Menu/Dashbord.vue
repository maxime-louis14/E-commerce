<script setup>
</script>

<template>
  <div class="bg-gray-100 min-h-screen py-8">
    <div class="max-w-6xl mx-auto">
      <!-- Entête du tableau de bord -->
      <header class="text-2xl font-semibold mb-6">
        Tableau de bord de profil
      </header>

      <!-- Sections du tableau de bord -->
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
        <!-- Section 1 -->
        <a
          href="/uploadImage"
          @click="handleSectionClick(1)"
          class="bg-white p-4 rounded-lg shadow-md cursor-pointer"
        >
          <h2 class="text-lg font-semibold mb-4">Ajouter son avatar</h2>
          <!-- Contenu de la section 1 -->
          <!-- Ajoutez vos fonctionnalités ici -->
        </a>

        <!-- Section 2 -->
        <a
          href="#"
          @click="handleSectionClick(2)"
          class="bg-white p-4 rounded-lg shadow-md cursor-pointer"
        >
          <h2 class="text-lg font-semibold mb-4">Section 2</h2>
          <!-- Contenu de la section 2 -->
          <!-- Ajoutez vos fonctionnalités ici -->
        </a>

        <!-- Section 3 -->
        <a
          href="#"
          @click="handleSectionClick(3)"
          class="bg-white p-4 rounded-lg shadow-md cursor-pointer"
        >
          <h2 class="text-lg font-semibold mb-4">Section 3</h2>
          <!-- Contenu de la section 3 -->
          <!-- Ajoutez vos fonctionnalités ici -->
        </a>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  methods: {
    handleSectionClick(sectionNumber) {
      // Réagissez au clic sur la section ici
      switch (sectionNumber) {
        case 1:
          // Action pour la section 1
          break;
        case 2:
          // Action pour la section 2
          break;
        case 3:
          // Action pour la section 3
          break;
        default:
          break;
      }
    }
  }
};
</script>
