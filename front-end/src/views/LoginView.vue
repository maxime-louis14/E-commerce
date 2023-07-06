<template>
  <div>
    <h2>Inscription</h2>
    <form>
      <div>
        <label for="email">Email:</label>
        <input type="email" id="email" v-model="email" required />
      </div>
      <div>
        <label for="password">Mot de passe:</label>
        <input type="password" id="password" v-model="password" required />
      </div>
      <button type="submit">login</button>
    </form>
  </div>
</template>

<script>
export default {
  data: () => {
    return {
      email: "",
      password: ""
    };
  },

  methods: {
    login() {
      fetch("http://localhost:8181/users/login", {
        method: "POST",
        headers: {
          "Content-type": "application/json"
        },
        body: {
          email: this.email,
          password: this.password
        }
      })
        .then((response) => response.json())
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
