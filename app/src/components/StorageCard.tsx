import { Text, View } from "react-native";

import { StorageInfo } from "../models/types";
import { formatBytes } from "../utils/utils";
import { styles } from "../theme/styles";

export const StorageCard: React.FC<{ info: StorageInfo }> = ({ info }) => {
    const percentage = (info.used / info.total) * 100;
    const remaining = formatBytes(info.total - info.used);

    return (
        <View style={styles.storageCard}>
            <View style={styles.storageHeader}>
                <Text style={styles.storageTitle}>Storage</Text>
                <Text style={styles.storageRemaining}>{remaining} free</Text>
            </View>
            <View style={styles.storageBarBg}>
                <View style={[styles.storageBarFill, { width: `${percentage}%` }]} />
            </View>
            <Text style={styles.storageDetail}>
                {formatBytes(info.used)} / {formatBytes(info.total)}
            </Text>
        </View>
    );
};


