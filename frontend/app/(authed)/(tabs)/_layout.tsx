import TabBarIcon from "@/components/navigation/TabBarIcon";
import { Tabs } from "expo-router";
import { ComponentProps } from "react";
import { Text } from "react-native";

export default function TabLayout() {
  const tabs = [
    {
      showFor: [],
      name: "(events)",
      displayName: "Events",
      icon: "calendar-outline",
      options: {
        headerShown: false,
      },
    },
    {
      showFor: [],
      name: "(tickets)",
      displayName: "My Tickets",
      icon: "ticket-outline",
      options: {
        headerShown: false,
      },
    },
    {
      showFor: [],
      name: "scan-ticket",
      displayName: "Scan Ticket",
      icon: "scan-outline",
      options: {
        headerShown: true,
      },
    },
    {
      showFor: [],
      name: "settings",
      displayName: "Settings",
      icon: "settings-outline",
      options: {
        headerShown: true,
      },
    },
  ];

  return (
    <Tabs>
      {tabs.map((tab) => (
        <Tabs.Screen
          key={tab.name}
          name={tab.name}
          options={{
            ...tab.options,
            headerTitle: tab.displayName,
            tabBarLabel: ({ focused }) => (
              <Text style={{ color: focused ? "black" : "gray" }}>
                {tab.displayName}
              </Text>
            ),
            tabBarIcon: ({ focused }) => (
              <TabBarIcon
                name={tab.icon as ComponentProps<typeof TabBarIcon>["name"]}
                color={focused ? "black" : "gray"}
              />
            ),
          }}
        />
      ))}
    </Tabs>
  );
}
