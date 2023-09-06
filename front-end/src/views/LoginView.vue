<template>
  <div>
    <h1>Connexion</h1>
  </div>

  <div class="formulaire-login mx-80 w-80">
    <form @submit.prevent="login">
      <div class="relative z-0 w-80 mb-6 group">
        <label class="block">
          <span class="block text-sm font-medium text-slate-700"> Mail </span>
        </label>
        <input
          type="email"
          id="email"
          v-model="email"
          class="block py-2.5 px-0 w-80 text-sm text-gray-900 bg-transparent border-0 border-b-2 border-gray-300 appearance-none dark:text-white dark:border-gray-600 dark:focus:border-blue-500 focus:outline-none focus:ring-0 focus:border-blue-600 peer"
        />
      </div>
      <div class="relative z-0 w-80 mb-6 group">
        <label
          class="peer-focus:font-medium absolute text-sm text-gray-500 dark:text-gray-400 duration-300 transform -translate-y-6 scale-75 top-3 -z-10 origin-[0] peer-focus:left-0 peer-focus:text-blue-600 peer-focus:dark:text-blue-500 peer-placeholder-shown:scale-100 peer-placeholder-shown:translate-y-0 peer-focus:scale-75 peer-focus:-translate-y-6"
          for="password"
          >Mot de passe :</label
        >
        <input
          type="password"
          id="password"
          v-model="password"
          class="block py-2.5 px-0 w-80 text-sm text-gray-900 bg-transparent border-0 border-b-2 border-gray-300 appearance-none dark:text-white dark:border-gray-600 dark:focus:border-blue-500 focus:outline-none focus:ring-0 focus:border-blue-600 peer"
        />
      </div>
      <div class="flex items-center justify-between">
        <button
          type="submit"
          class="bg-green-500 hover:bg-green-600 text-white font-bold py-2 px-4 rounded focus:outline-none focus:ring focus:ring-green-200"
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
  data: () => {
    return {
      email: "",
      password: ""
    };
  },

  methods: {
    async login(e) {
      e.preventDefault();
      await fetch("http://localhost:8080/users/login", {
        method: "POST",
        headers: {},
        body: JSON.stringify({
          email: this.email,
          password: this.password
        })
      })
        .then((response) => {
          console.log(response.data);
          console.log(response.json());
          let data = response.json();
          console.log(data.token);
          Swal.fire({
            icon: "success",
            title: "Vous êtes connecté"
          });
        })

        .then((data) => {
          // Vérifiez si la réponse contient un JWT valide
          console.log(data);
          // Stockez le JWT dans le local storage ou dans un cookie, selon vos besoins
          localStorage.setItem("Token", data.token);
          // Redirigez l'utilisateur vers une autre page
          this.$router.push("/");
        })
        .catch((error) => {
          console.error("Erreur lors de la requête de connexion:", error);
        });
    }
  }
};
</script>

<style scoped>/* Style pour le titre "Connexion" */
h1 {
  font-size: 24px;
  text-align: center;
  color: #333; /* Couleur du texte */
  margin-bottom: 20px; /* Marge inférieure */
}

/* Style pour le conteneur du formulaire */
.formulaire-login {
  background-color: #f5f5f5; /* Couleur de fond du formulaire */
  padding: 20px; /* Marges intérieures */
  border-radius: 10px; /* Coins arrondis */
  box-shadow: 0px 0px 10px rgba(0, 0, 0, 0.1); /* Ombre légère */
  text-align: center;
  width: 300px;
  margin: auto;
  margin-top: 50px;
}

/* Style pour les libellés */
label {
  font-weight: bold;
  display: block;
  margin-bottom: 5px;
}

/* Style pour les champs de saisie */
.input-field {
  width: 100%;
  padding: 10px;
  border: 1px solid #ccc;
  border-radius: 5px;
  transition: border-color 0.3s;
  font-size: 16px;
}

.input-field:focus {
  border-color: #3498db;
  outline: none;
}

/* Style pour le bouton de connexion */
.login-button {
  background-color: #4caf50;
  color: white;
  padding: 10px 20px;
  border: none;
  border-radius: 5px;
  cursor: pointer;
  font-size: 16px;
  transition: background-color 0.3s;
}

.login-button:hover {
  background-color: #45a049;
}

/* Style pour le message d'erreur */
.error-message {
  text-align: center;
  color: #e74c3c;
  margin-top: 10px;
}

</style>
