<template>
  <div class="container mx-auto my-4">
    <h1 class="text-3xl font-semibold mb-4">Boutique</h1>
    <div class="overflow-x-auto">
      <div class="flex space-x-4">
        <!-- Utilisation de v-for pour générer dynamiquement les produits horizontalement -->
        <div v-for="(product, index) in products" :key="index" class="bg-white p-4 shadow-md w-72">
          <img :src="product.image" :alt="product.name" class="w-full h-48 object-cover">
          <h2 class="text-xl font-semibold mt-2">{{ product.name }}</h2>
          <p class="text-gray-600">{{ product.description }}</p>
          <p class="text-lg font-semibold mt-2">{{ product.price }}</p>
          <button class="bg-blue-500 text-white px-4 py-2 mt-4 hover:bg-blue-600" @click="addToCart(product)">Ajouter au panier</button>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Boutique',
  data() {
    return {
      products: [
        {
          name: 'Nom du Produit 1',
          description: 'Description du Produit 1',
          price: '$99.99',
          image: 'lien_de_l_image_1.jpg'
        },
        {
          name: 'Nom du Produit 2',
          description: 'Description du Produit 2',
          price: '$79.99',
          image: 'lien_de_l_image_2.jpg'
        },
        {
          name: 'Nom du Produit 3',
          description: 'Description du Produit 3',
          price: '$129.99',
          image: 'lien_de_l_image_3.jpg'
        },
        // Ajoutez d'autres produits ici
      ],
    };
  },
  methods: {
    addToCart(product) {
      // Vous pouvez implémenter la logique d'ajout au panier ici
    },
  },
};
</script>

<style scoped>
/* Vous pouvez ajouter ici des styles CSS personnalisés si nécessaire. */
</style>
