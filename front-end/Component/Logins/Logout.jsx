import React, { useState } from 'react';
import { View, TextInput, Button } from 'react-native';

const Logout = () => {
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const [firstName, setFirstName] = useState('');
  const [lastName, setLastName] = useState('');
  const [email, setEmail] = useState('');

  const handleSignup = () => {
    // Envoyer les données d'inscription à l'API Go
    const data = {
      username: username,
      password: password,
      firstName: firstName,
      lastName: lastName,
      email: email
    };

    // Effectuer une requête à l'API en utilisant fetch ou une librairie comme axios
    fetch('http://localhost:8181/users/signup', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(data)
    })
      .then(response => response.json())
      .then(responseData => {
        // Traiter la réponse de l'API ici
        console.log(responseData);
      })
      .catch(error => {
        // Gérer les erreurs de requête ici
        console.error(error);
      });
  };

  return (
    <View>
      <TextInput
        placeholder="Nom d'utilisateur"
        value={username}
        onChangeText={text => setUsername(text)}
      />
      <TextInput
        placeholder="Mot de passe"
        secureTextEntry={true}
        value={password}
        onChangeText={text => setPassword(text)}
      />
      <TextInput
        placeholder="Prénom"
        value={firstName}
        onChangeText={text => setFirstName(text)}
      />
      <TextInput
        placeholder="Nom de famille"
        value={lastName}
        onChangeText={text => setLastName(text)}
      />
      <TextInput
        placeholder="E-mail"
        value={email}
        onChangeText={text => setEmail(text)}
      />
      <Button title="S'inscrire" onPress={handleSignup} />
    </View>
  );
};

export default Logout;
