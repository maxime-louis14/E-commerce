<template>
  <div>
    <h1 class="text-3xl text-center font-bold mb-4">Connexion</h1>
  </div>

  <div class="mx-auto w-80">
    <form @submit.prevent="login">
      <!-- Champ d'adresse e-mail -->
      <div class="mb-6">
        <label for="email" class="text-sm font-semibold"
          >Adresse e-mail :</label
        >
        <input
          type="email"
          id="email"
          v-model="email"
          class="block w-full py-2.5 px-3 text-sm text-gray-900 bg-transparent border-b-2 border-gray-300 focus:outline-none focus:border-blue-500"
          required
        />
      </div>

      <!-- Champ de mot de passe -->
      <div class="mb-6">
        <label for="password" class="text-sm font-semibold"
          >Mot de passe :</label
        >
        <input
          type="password"
          id="password"
          v-model="password"
          class="block w-full py-2.5 px-3 text-sm text-gray-900 bg-transparent border-b-2 border-gray-300 focus:outline-none focus:border-blue-500"
          required
        />
      </div>

      <!-- Bouton de connexion -->
      <div class="flex items-center justify-between">
        <button
          type="submit"
          class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded-full"
        >
          Connexion
        </button>
      </div>
    </form>
  </div>
</template>

<script>
import Swal from "sweetalert2";

export default {
  name: "login",
  data() {
    return {
      email: "",
      password: ""
    };
  },

  methods: {
    async login() {
      try {
        if (!this.email || !this.password) {
          // Validation côté client : Vérifiez si les champs sont remplis
          Swal.fire({
            icon: "error",
            title: "Champs requis",
            text: "Veuillez remplir tous les champs."
          });
          return;
        }

        // Envoie de la requête de connexion
        const response = await fetch("http://localhost:8080/api/users/login", {
          method: "POST",
          headers: {
            "Content-Type": "application/json"
          },
          body: JSON.stringify({
            email: this.email,
            password: this.password
          })
        });

        if (!response.ok) {
          throw new Error("Échec de la connexion");
        }

        const data = await response.json();
        const token = data.token;

        localStorage.setItem("Token", token);
        Swal.fire({
          icon: "success",
          title: "Vous êtes connecté"
        });

        // Nettoyez les champs après la connexion
        this.email = "";
        this.password = "";

        // Redirection de l'utilisateur vers une autre page
        this.$router.push("/");
      } catch (error) {
        console.error("Erreur lors de la requête de connexion:", error);

        Swal.fire({
          icon: "error",
          title: "Échec de la connexion",
          text: "Veuillez vérifier vos informations de connexion."
        });
      }
    }
  }
};
</script>
