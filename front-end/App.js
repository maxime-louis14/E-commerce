import { StatusBar } from "expo-status-bar";
import { NavigationContainer } from "@react-navigation/native";
import { createMaterialBottomTabNavigator } from "@react-navigation/material-bottom-tabs";
import { MaterialCommunityIcons } from "@expo/vector-icons";

import { StyleSheet, Text, View } from "react-native";
import HomeScreen from "./Component/HomeScreen.jsx";
import ProfilScreen from "./Component/ProfileScreen.jsx";

const Tab = createMaterialBottomTabNavigator();

export default function App() {
  return (
    <NavigationContainer>
      <Tab.Navigator>
        <Tab.Screen
          name="Home"
          component={HomeScreen}
          options={{
            tabBarLabel: "Accueil",
            tabBarIcon: ({ color }) =>
              <MaterialCommunityIcons name="home" color={color} size={20} />
          }}
        />
        <Tab.Screen
          name="Profile"
          component={ProfilScreen}
          options={{
            tabBarLabel: "Profil",
            tabBarIcon: ({ color }) =>
              <MaterialCommunityIcons name="account" color={color} size={20} />
          }}
        />
      </Tab.Navigator>
    </NavigationContainer>
  );
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: "#fff",
    alignItems: "center",
    justifyContent: "center"
  }
});
