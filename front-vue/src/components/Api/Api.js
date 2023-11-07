// api.js

export const getProfileImage = async () => {
    try {
      const response = await fetch('URL_DE_VOTRE_ENDPOINT_POUR_LE_GET_IMAGE', {
        method: 'GET',
        headers: {
          // Ajoutez ici les en-têtes d'authentification ou tout autre en-tête requis
        },
      });
  
      if (!response.ok) {
        throw new Error('Échec de la requête GET pour l\'image de profil');
      }
  
      const data = await response.json();
      return data; // Retourne l'URL de l'image de profil
    } catch (error) {
      throw new Error('Erreur lors de la récupération de l\'image de profil : ' + error.message);
    }
  };
  