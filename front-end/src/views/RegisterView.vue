<template>
  <div>
    <h2>Register</h2>
    <from>
      <div>
        <input type="text" placeholder="Prenom" v-model="Prenom" />
      </div>
      <div>
        <input type="text" placeholder="Nom" v-model="Nom" />
      </div>
      <div>
        <input type="email" placeholder="Email" v-model="email" />
      </div>
      <div>
        <input type="password" placeholder="Password" v-model="password" />
      </div>
      <button @click="envoyer">Envoyer</button>
    </from>
  </div>
</template>

<script>
export default {
  name: "register",
  data() {
    return {
      Prenom: null,
      Nom: null,
      email: null,
      password: null
    };
  },
  methods: {
    envoyer() {
      const userData = {
        Prenom: this.Prenom,
        Nom: this.Nom,
        email: this.email,
        password: this.password
      };

      // Envoyer les informations d'enregistrement au backend
      fetch("localhost:8181/users/register", {
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
