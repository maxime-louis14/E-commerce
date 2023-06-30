import React from "react";
import { StyleSheet, Text, View } from "react-native";
import { useNavigation } from "@react-navigation/native";
import { Button } from "react-native-paper";
import Login from "./Logins/Login";
import Logout from "./Logins/Logout";

export default function ProfilScreen() {
  const navigation = useNavigation();

  function navigate(screenName) {
    navigation.navigate(screenName);
  }

  return (
    <View style={styles.container}>
      <Text style={styles.text}>Mon Compte</Text>
      <Button onPress={() => navigate("Login")} title="Login" />
      <Button onPress={() => navigate("Logout")} title="Logout" />
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
