# Utiliser une image de base avec Node.js et Android SDK
FROM node:latest

# Définir le répertoire de travail à l'intérieur du conteneur
WORKDIR /app

# Installation du JDK
RUN apt-get update && \
    apt-get install -y openjdk-11-jdk

# Installation d'Android SDK
RUN wget https://dl.google.com/android/repository/commandlinetools-linux-7583922_latest.zip -O /tmp/sdk.zip && \
    unzip /tmp/sdk.zip -d /opt/android-sdk && \
    rm /tmp/sdk.zip

# Configuration des variables d'environnement pour Android SDK
ENV ANDROID_HOME /opt/android-sdk
ENV PATH $PATH:$ANDROID_HOME/cmdline-tools/tools/bin
ENV PATH $PATH:$ANDROID_HOME/platform-tools

# Accepter les licences Android SDK
RUN yes | sdkmanager --licenses

# Installation des packages Android SDK requis
RUN sdkmanager "platforms;android-30" "build-tools;30.0.3"

# Installation d'Expo CLI
RUN npm install -g expo-cli

# Copier les fichiers source de l'application dans le conteneur
COPY . /app

# Installer les dépendances de l'application
RUN npm install

# Exposer le port du packager Metro (par défaut : 8081)
EXPOSE 8081

# Exposer le port de l'application (si nécessaire)
# EXPOSE <port>

# Commande par défaut pour démarrer le packager Metro
CMD ["npm", "start", "--android", "--web"]
