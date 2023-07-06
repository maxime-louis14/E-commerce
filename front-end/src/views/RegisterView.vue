<template>
  <div>
    <h2>Register</h2>
    <form @submit.prevent="register">
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
        <button @click="(e) => register(e)">Envoyer</button>
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
      prenom: "",
      nom: "",
      email: "",
      password: ""
    };
  },
  methods: {
    register(e) {
      e.preventDefault();
      const userData = {
        prenom: this.prenom,
        nom: this.nom,
        email: this.email,
        password: this.password
      };
      console.log(userData);
      // Envoyer les informations d'enregistrement au backend
      fetch("http://localhost:8080/users/register", {
        method: "POST",
        headers: {
          
        },
        body: JSON.stringify(userData)
      })
        .then((response) => {
          response.json();
          Swal.fire({
            icon: "success",
            title: "L'utilisateur a bien été ajouté"
          })
        })
        .then((data) => console.log(data))
        .then((err) => console.log(err));
    }
  }
};
</script>

<style>
form {
  display: flex;
  flex-direction: column;
  align-items: center;
}

h2 {
  text-align: center;
}

input {
  margin-bottom: 10px;
}

div {
  display: flex;
  flex-direction: column;
  align-items: center;
}

button {
  margin-top: 10px;
}

p {
  text-align: center;
}
</style>
