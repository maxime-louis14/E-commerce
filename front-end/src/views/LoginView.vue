<template>
  <div>
    <h2>Inscription</h2>
    <form @submit.prevent="login">
      <div>
        <label for="email">Email:</label>
        <input type="email" id="email" v-model="email" />
      </div>
      <div>
        <label for="password">Mot de passe:</label>
        <input type="password" id="password" v-model="password" />
      </div>
      <button @click="(e) => login(e)">login</button>
    </form>
  </div>
</template>

<script>
import Swal from "sweetalert2";

export default {
  name: "login",
  data: () => {
    return {
      user: {
        email: "",
        password: ""
      }
    };
  },

  methods: {
    login(e) {
      e.preventDefault();
      fetch("http://localhost:8080/users/login", {
        method: "POST",
        headers: {
          "Access-Control-Allow-Origin": "*",
          "Content-type": "application/json;charset=UTF-8",
          Accapt: "application/json",
          mode: "cors"
        },
        body: {
          email: this.email,
          password: this.password
        }
      })
        .then((response) => {
          response.json();
          Swal.fire({
            icon: "success",
            title: "L'utilisateur a bien été ajouté"
          });
        })
        .then((data) => {
          // Vérifiez si la réponse contient un JWT valide
          if (data.token) {
            // Stockez le JWT dans le local storage ou dans un cookie, selon vos besoins
            localStorage.setItem("jwtToken", data.token);

            // Redirigez l'utilisateur vers une autre page
            this.$router.push("/");
          } else {
            // Gérez les erreurs d'authentification
            console.error("Erreur d'authentification");
          }
        })
        .catch((error) => {
          console.error("Erreur lors de la requête de connexion:", error);
        });
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
