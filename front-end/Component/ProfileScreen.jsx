import React from "react";
import { StyleSheet, Text, View } from "react-native";
import { useNavigation } from "@react-navigation/native";
import { Button } from "react-native-paper";
import Login from "./Logins/Login";
import Logout from "./Logins/Logout";

export default function ProfilScreen() {
  const navigation = useNavigation();

  function navigate() {
    navigation.navigate(Login);

    navigation.navigate(Logout);
  }

  return (
    <View style={styles.container}>
      <Text style={styles.text}>Mon Compte</Text>
      <Button title="login" onPress={navigate}>
        Login
      </Button>
      <Button title="Logout" onPress={navigate}>
        Logout
      </Button>
    </View>
  );
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    justifyContent: "center",
    alignItems: "center",
    backgroundColor: "#fff"
  },
  text: {
    fontSize: 20,
    fontWeight: "bold",
    color: "black"
  }
});
