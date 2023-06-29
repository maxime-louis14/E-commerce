# Utiliser une image de base contenant Go
FROM golang:latest AS build-stage

# Définir le répertoire de travail à l'intérieur du conteneur
WORKDIR /ap

# Copier les fichiers source de l'application dans le conteneur
COPY . .

# Compiler l'application Go dans le conteneur
RUN CGO_ENABLED=0 GOOS=linux go build -o /clone-site-Vilebrequin ./cmd/app

# Exécuter les tests dans le conteneur
RUN go test -v ./...

# Déployer le binaire de l'application dans une image légère
FROM gcr.io/distroless/base-debian11 AS release-stage

WORKDIR /

COPY --from=build-stage /clone-site-Vilebrequin /clone-site-Vilebrequin

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/clone-site-Vilebrequin"]
