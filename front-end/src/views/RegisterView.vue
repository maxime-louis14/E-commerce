<template>
  <div>
    <h2>Register</h2>
    <form action="">
      <div>
        <input type="text" placeholder="prenom" v-model="prenom" />
      </div>
      <div>
        <input type="text" placeholder="nom" v-model="nom" />
      </div>
      <div>
        <input type="email" placeholder="email" v-model="email" />
      </div>
      <input
        type="password"
        minlength="1"
        placeholder="password"
        v-model="password"
      />
      <div v-if="password && password.length >= 1">
        <button @click="(e) => envoyer(e)">Envoyer</button>
      </div>
      <div v-else>
        <p>Veuillez saisir un mot de passe d'au moins 15 caractères.</p>
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
      prenom: null,
      nom: null,
      email: null,
      password: null
    };
  },
  methods: {
    envoyer(e) {
      e.preventDefault();
      const userData = {
        prenom: this.prenom,
        nom: this.nom,
        email: this.email,
        password: this.password
      };
      console.log(userData);
      // Envoyer les informations d'enregistrement au backend
      fetch("http://localhost:8181/users/signup", {
        method: "POST",
        headers: {
          "Content-type": "application/json",
          "Access-Control-Allow-Origin": "*",
          "Access-Control-Allow-Methods":
            "GET, POST, PATCH, PUT, DELETE, OPTIONS",
          "Access-Control-Allow-Headers": "Origin, Content-Type, X-Auth-Token"
        },
        body: JSON.stringify(userData)
      })
        .then((response) => {
          response.json();
          Swal.fire({
            icon: "success",
            title: "L'utilisateur a bien été ajouté"
          });
        })
        .catch((error) => {
          console.error("Erreur lors de la requête d'enregistrement:", error);
        });
    }
  }
};
</script>
