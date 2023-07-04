<template>
  <div>
    <h2>Inscription</h2>
    <form @submit.prevent="register">
      <div>
        <label for="firstName">Prénom:</label>
        <input type="text" id="firstName" v-model="firstName" required />
      </div>
      <div>
        <label for="lastName">Nom:</label>
        <input type="text" id="lastName" v-model="lastName" required />
      </div>
      <div>
        <label for="email">Email:</label>
        <input type="email" id="email" v-model="email" required />
      </div>
      <div>
        <label for="password">Mot de passe:</label>
        <input type="password" id="password" v-model="password" required />
      </div>
      <button type="submit">S'inscrire</button>
    </form>
  </div>
</template>

<script>
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
