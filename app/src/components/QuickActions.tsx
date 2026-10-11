import { Text, TouchableOpacity, View } from "react-native";

import { QuickAction } from "../models/types";
import { styles } from "../theme/styles";

export const QuickActions: React.FC<{ actions: QuickAction[] }> = ({ actions }) => (
    <View style={styles.quickActions}>
        {actions.map((action) => (
            <TouchableOpacity key={action.id} style={styles.quickActionItem} activeOpacity={0.7}>
                <View style={styles.quickActionIcon}>
                    <Text style={styles.quickActionEmoji}>{action.icon}</Text>
                </View>
                <Text style={styles.quickActionLabel}>{action.label}</Text>
            </TouchableOpacity>
        ))}
    </View>
);


