import React from "react";
import Ionicons from "react-native-vector-icons/Ionicons";

type TabBarIconProps = {
  name: string;
  color: string;
  focused?: boolean;
  size?: number;
};

const TabBarIcon: React.FC<TabBarIconProps> = ({
  name,
  color,
  focused = false,
  size = 24,
}) => {
  return (
    <Ionicons
      name={name}
      size={size}
      color={color}
      style={{ opacity: focused ? 1 : 0.7 }}
    />
  );
};

export default TabBarIcon;
