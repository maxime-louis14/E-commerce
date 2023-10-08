<template>
  <div>
    <h1 class="text-3xl font-bold text-center mb-4">Inscription</h1>
  </div>

  <div class="mx-auto w-80">
    <form @submit.prevent="register">
      <!-- Champ Prénom -->
      <div class="mb-6">
        <label for="prenom" class="block text-sm font-medium text-slate-700"
          >Prénom :</label
        >
        <input
          type="text"
          placeholder="Prénom"
          v-model="prenom"
          class="block w-full py-2.5 px-3 text-sm text-gray-900 bg-transparent border-b-2 border-gray-300 focus:outline-none focus:border-blue-500"
        />
      </div>

      <!-- Champ Nom -->
      <div class="mb-6">
        <label for="nom" class="block text-sm font-medium text-slate-700"
          >Nom :</label
        >
        <input
          type="text"
          placeholder="Nom"
          v-model="nom"
          class="block w-full py-2.5 px-3 text-sm text-gray-900 bg-transparent border-b-2 border-gray-300 focus:outline-none focus:border-blue-500"
        />
      </div>

      <!-- Champ d'adresse e-mail -->
      <div class="mb-6">
        <label for="email" class="block text-sm font-medium text-slate-700"
          >Adresse e-mail :</label
        >
        <input
          type="email"
          placeholder="Email"
          v-model="email"
          class="block w-full py-2.5 px-3 text-sm text-gray-900 bg-transparent border-b-2 border-gray-300 focus:outline-none focus:border-blue-500"
        />
      </div>

      <!-- Champ de mot de passe -->
      <div class="mb-6">
        <label for="password" class="block text-sm font-medium text-slate-700"
          >Mot de passe :</label
        >
        <input
          type="password"
          minlength="3"
          placeholder="Mot de passe"
          v-model="password"
          class="block w-full py-2.5 px-3 text-sm text-gray-900 bg-transparent border-b-2 border-gray-300 focus:outline-none focus:border-blue-500"
        />

        <!-- Bouton d'inscription (conditionnel) -->
        <div v-if="password && password.length >= 3" class="mt-5">
          <button
            type="submit"
            class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded-full"
          >
            Envoyer
          </button>
        </div>

        <!-- Message d'erreur (conditionnel) -->
        <div v-else>
          <p class="text-red-500">
            Veuillez saisir un mot de passe d'au moins 15 caractères.
          </p>
        </div>
      </div>
    </form>
  </div>
</template>

<script>
import Swal from "sweetalert2";

export default {
  name: "register",
  data() {
    return {
      prenom: "",
      nom: "",
      email: "",
      password: ""
    };
  },
  methods: {
    async register(e) {
      e.preventDefault();

      try {
        // Créez un objet userData avec les données du formulaire
        const userData = {
          prenom: this.prenom,
          nom: this.nom,
          email: this.email,
          password: this.password
        };

        // Envoyez les informations d'enregistrement au backend
        const response = await fetch("http://localhost:8080/api/users/register", {
          method: "POST",
          headers: {
            "Content-Type": "application/json" // Spécifiez le type de contenu JSON
          },
          body: JSON.stringify(userData)
        });

        // Vérifiez si la réponse est OK
        if (!response.ok) {
          throw new Error("Échec de l'inscription"); // Gestion des erreurs
        }

        // Affichez une notification de réussite
        Swal.fire({
          icon: "success",
          title: "L'utilisateur a bien été ajouté"
        });

        // Réinitialisez le formulaire
        this.prenom = "";
        this.nom = "";
        this.email = "";
        this.password = "";

        // Redirigez l'utilisateur vers une autre page (facultatif)
        this.$router.push("/");
      } catch (error) {
        console.error("Erreur lors de l'inscription:", error);

        // Affichez une notification d'erreur
        Swal.fire({
          icon: "error",
          title: "Échec de l'inscription",
          text: "Veuillez vérifier vos informations d'inscription."
        });
      }
    }
  }
};
</script>
