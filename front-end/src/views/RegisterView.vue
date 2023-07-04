<template>
  <div class="container">
    <h1>Inscription</h1>
    <p v-if="message != null">{{ message }}</p>
    <form>
      <input type="email" placeholder="Votre Email" v-model="email" />
      <input type="password" placeholder="password" v-model="password" />
      <button @click="envoyer">envoyer</button>
    </form>
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
import axios from "axios";

export default {
  data() {
    return {
      firstName: "",
      lastName: "",
      email: "",
      password: ""
    };
  },

  methods: {
    register() {
      // Logique d'inscription ici
      // Vous pouvez envoyer les données du formulaire à votre API pour créer un nouvel utilisateur

      // Exemple de log :
      console.log("Prénom:", this.firstName);
      console.log("Nom:", this.lastName);
      console.log("Email:", this.email);
      console.log("Mot de passe:", this.password);

      // Réinitialiser les champs du formulaire
      this.firstName = "";
      this.lastName = "";
      this.email = "";
      this.password = "";

      // Afficher un message de succès ou rediriger vers une autre page
    }
  }
};
</script>

<style scoped>
h2 {
  font-size: 20px;
  margin-bottom: 10px;
}
form div {
  margin-bottom: 15px;
}
label {
  display: block;
  font-weight: bold;
}
input {
  width: 200px;
  padding: 5px;
}
button {
  background-color: #4caf50;
  color: white;
  border: none;
  padding: 10px 20px;
  font-size: 16px;
  cursor: pointer;
}
</style>
