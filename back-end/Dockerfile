# Utiliser une image de base contenant Go
FROM golang:latest

# Définir le répertoire de travail à l'intérieur du conteneur
WORKDIR /app

# Copier les fichiers source de l'application dans le conteneur
COPY . .

# Exécuter la commande go run main.go lors du démarrage du conteneur
CMD ["go", "run", "main.go"]