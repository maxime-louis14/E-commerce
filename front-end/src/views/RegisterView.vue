<template>
  <div class="container">
    <h1>Inscription</h1>
    <p v-if="message != null">{{ message }}</p>
    <form>
      <input type="email" placeholder="Votre Email" v-model="email" />
      <input type="password" placeholder="password" v-model="password" />
      <button @click="envoyer">envoyer</button>
    </form>
  </div>
</template>

<script>
import axios from "axios";

export default {
  name: "inscription",
  data() {
    return {
      email: null,
      password: null,
      message: null
    };
  },

  methods: {
    async envoyer(e) {
      e.preventDefault();
      this.message = null;
      const regexEmail = /^[\w-.]+@([\w-]+.)+[\w-]{2,4}$/;

      if(!this.email){
        return this.message = "rensenignez l'email";
      }else if(!regexEmail.test(this.email)){
        return this.message = "email invalide";
      }else if(!this.password){
        return this.message = "renseignez un mot de passe";
      }else if(this.password.length < 4){
        return this.message = "le mot de passe est trop court"
      }

      if(this.message == null){
        const url = "http://127.0.0.1:8000/api/inscription";
        const user = {
          "email": this.email,
          "password": this.password
        }

        await axios.post(url, user).then(
          (res)=>{
            this.$router.push("/")
          },
          (error)=>{
            this.message = "une erreur est survenue lors de l'inscription";
          }
        )
      }

    }
  }
};
</script>