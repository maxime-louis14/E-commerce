<template>
  <div>
      <h1>Register</h1>
  </div>
  
  <div class="formulaire mx-80 w-80">
    <form @submit.prevent="register">
      <div class="relative z-0 w-80 mb-6 group">
        <label class="block">
          <span class="block text-sm font-medium text-slate-700"> Prenom</span>
        </label>
        <input
          type="text"
          placeholder="prenom"
          v-model="prenom"
          class="block text-sm font-medium text-slate-700"
        />
      </div>

      <div>
        <label
          for="floating_nom"
          class="peer-focus:font-medium absolute text-sm text-gray-500 dark:text-gray-400 duration-300 transform -translate-y-6 scale-75 top-3 -z-10 origin-[0] peer-focus:left-0 peer-focus:text-blue-600 peer-focus:dark:text-blue-500 peer-placeholder-shown:scale-100 peer-placeholder-shown:translate-y-0 peer-focus:scale-75 peer-focus:-translate-y-6"
          >Nom</label
        >

        <input type="text" placeholder="nom" v-model="nom" />
      </div>

      <div>
        <label
          for="floating_email"
          class="peer-focus:font-medium absolute text-sm text-gray-500 dark:text-gray-400 duration-300 transform -translate-y-6 scale-75 top-3 -z-10 origin-[0] peer-focus:left-0 peer-focus:text-blue-600 peer-focus:dark:text-blue-500 peer-placeholder-shown:scale-100 peer-placeholder-shown:translate-y-0 peer-focus:scale-75 peer-focus:-translate-y-6"
          >Email address</label
        >
        <input
          type="email"
          placeholder="email"
          v-model="email"
          class="block py-2.5 px-0 w-80 text-sm text-gray-900 bg-transparent border-0 border-b-2 border-gray-300 appearance-none dark:text-white dark:border-gray-600 dark:focus:border-blue-500 focus:outline-none focus:ring-0 focus:border-blue-600 peer"
        />
      </div>

      <div class="relative z-0 w-80 mb-6 group">
        <label
          for="floating_repeat_password"
          class="peer-focus:font-medium absolute text-sm text-gray-500 dark:text-gray-400 duration-300 transform -translate-y-6 scale-75 top-3 -z-10 origin-[0] peer-focus:left-0 peer-focus:text-blue-600 peer-focus:dark:text-blue-500 peer-placeholder-shown:scale-100 peer-placeholder-shown:translate-y-0 peer-focus:scale-75 peer-focus:-translate-y-6"
          >Confirm password</label
        >

        <input
          type="password"
          minlength="1"
          placeholder="password"
          v-model="password"
          class="block py-2.5 px-0 w-80 text-sm text-gray-900 bg-transparent border-0 border-b-2 border-gray-300 appearance-none dark:text-white dark:border-gray-600 dark:focus:border-blue-500 focus:outline-none focus:ring-0 focus:border-blue-600 peer"
        />

        <div v-if="password && password.length >= 1">
          <button @click="(e) => register(e)">Envoyer</button>
        </div>
        <div v-else>
          <p>Veuillez saisir un mot de passe d'au moins 15 caractères.</p>
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
        headers: {},
        body: JSON.stringify(userData)
      })
        .then((response) => {
          response.json();
          Swal.fire({
            icon: "success",
            title: "L'utilisateur a bien été ajouté"
          });
        })
        .then((data) => console.log(data))
        .then((err) => console.log(err));
    }
  }
};
</script>

<style scoped>
/* Style pour le titre "Register" */
h1 {
  font-size: 24px;
  text-align: center;
  color: #333; /* Couleur du texte */
  margin-bottom: 20px; /* Marge inférieure */
}

/* Style pour les éléments du formulaire */
.formulaire {
  background-color: #f5f5f5; /* Couleur de fond du formulaire */
  padding: 20px; /* Marges intérieures */
  border-radius: 10px; /* Coins arrondis */
  box-shadow: 0px 0px 10px rgba(0, 0, 0, 0.1); /* Ombre légère */
  text-align: center;
  width: 300px;
  margin: auto;
  margin-top: 50px;
}

/* Style pour les champs de formulaire */
input[type="text"],
input[type="email"],
input[type="password"] {
  width: 100%;
  padding: 10px; /* Marges intérieures */
  margin-bottom: 15px; /* Marge inférieure */
  border: 1px solid #ccc; /* Bordure */
  border-radius: 5px; /* Coins arrondis */
  transition: border-color 0.3s; /* Transition sur la couleur de la bordure */
}

input[type="text"]:focus,
input[type="email"]:focus,
input[type="password"]:focus {
  border-color: #3498db; /* Couleur de la bordure lorsqu'il est en focus */
}

/* Style pour le bouton */
button {
  background-color: #3498db; /* Couleur de fond du bouton */
  color: #fff; /* Couleur du texte du bouton */
  padding: 10px 20px; /* Marges intérieures */
  border: none;
  border-radius: 5px; /* Coins arrondis */
  cursor: pointer;
  transition: background-color 0.3s; /* Transition sur la couleur de fond */
}

button:hover {
  background-color: #2980b9; /* Couleur de fond du bouton au survol */
}

/* Style pour le message d'erreur */
p {
  text-align: center;
  color: #e74c3c; /* Couleur du texte d'erreur */
  margin-top: 10px; /* Marge supérieure */
}
</style>
