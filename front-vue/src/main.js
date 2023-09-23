import { createApp } from 'vue';
import App from './App.vue';
import router from './router'; // Importez votre configuration de router.js
import './style.css';

const app = createApp(App);
app.use(router); // Utilisez Vue Router dans votre application
app.mount('#app');
