# Utilisez l'image Node.js en tant que base
FROM node:14

# Définit le répertoire de travail dans le conteneur
WORKDIR /app

# Copie le package.json et le package-lock.json dans le répertoire de travail
COPY package*.json ./

# Installe les dépendances
RUN npm install
# RUN npm install -g vite

# Copie tout le contenu du répertoire actuel dans le répertoire de travail du conteneur
COPY . .

# Expose le port sur lequel l'application Vue.js sera accessible (par défaut : 3000)
EXPOSE 3000

# Exécute la commande de démarrage de l'application
CMD [ "npm", "run", "dev" ]