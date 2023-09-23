# Utilisez une image de base contenant Go
FROM golang:1.20

# Définir le répertoire de travail à l'intérieur du conteneur
WORKDIR /app

# Copier le fichier go.mod et go.sum dans le répertoire de travail
COPY go.mod go.sum ./

# Télécharger les dépendances du module
RUN go mod download

# Copier le code source de l'application
COPY . .

# Compiler l'application
RUN go build -o main .

# Définir la commande par défaut pour exécuter l'application
CMD ["./main"]
                