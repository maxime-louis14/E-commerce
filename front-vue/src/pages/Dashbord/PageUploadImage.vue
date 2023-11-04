<script setup>
import UploadImage from "../../components/ProfileComponents/uploadImage.vue";
</script>

<template>
  <div class="bg-gray-100 min-h-screen py-8">
    <div class="max-w-6xl mx-auto">
      <h2 class="text-2xl font-semibold mb-6">Ajouter votre avatar</h2>
      <img
        v-if="avatar != null"
        :src="avatar"
        alt="img"
        class="max-w-[300px] mx-auto mt-2 rounded-full"
      />
      <div class="ml-4">
        <h2 class="text-2xl text-center font-semibold mb-6"> Avatar actuel </h2>
        <img
          class="max-w-[300px] mx-auto mt-2 rounded-full"
          src="https://images.unsplash.com/photo-1580489944761-15a19d654956?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=1361&q=80"
          alt="Current profile photo"
        />
      </div>
    </div>
    <UploadImage />
  </div>
</template>
