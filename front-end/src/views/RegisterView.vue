<template>
  <div>
    <h2>Register</h2>
    <form>
      <div>
        <input type="text" placeholder="prenom" v-model="prenom" />
      </div>
      <div>
        <input type="text" placeholder="nom" v-model="nom" />
      </div>
      <div>
        <input type="email" placeholder="email" v-model="email" />
      </div>
      <div>
        <input type="password" placeholder="password" v-model="password" />
        <button @click="envoyer">Envoyer</button>
      </div>
    </form>
  </div>
</template>

<script>
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
    envoyer() {
      const userData = {
        prenom: this.prenom,
        nom: this.nom,
        email: this.email,
        password: this.password
      };

      // Envoyer les informations d'enregistrement au backend
      fetch("http://localhost:8181/users/signup", {
        method: "POST",
        headers: {
          "Content-type": "application/json"
        },
        body: JSON.stringify(userData)
      })
        .then((response) => response.json())
        .then((data) => {
          // Gérer la réponse du backend
          console.log("Réponse du backend:", data);
          // Effectuer une redirection ou une autre action selon la réponse du backend
        })
        .catch((error) => {
          console.error("Erreur lors de la requête d'enregistrement:", error);
        });
    }
  }
};
</script>
