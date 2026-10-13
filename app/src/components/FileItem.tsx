import { Text, TouchableOpacity, View } from "react-native";

import { formatBytes, getFileIcon } from "../utils/utils";
import { FileItem } from "../models/types";
import { styles } from "../theme/styles";

export const FileListItem: React.FC<{ item: FileItem }> = ({ item }) => {
    console.log(item)
    return (
        <TouchableOpacity style={styles.fileRow} activeOpacity={0.6}>
            <View style={styles.fileIconWrap}>
                <Text style={styles.fileIcon}>{getFileIcon(item.name, item.type)}</Text>
            </View>
            <View style={styles.fileInfo}>
                <Text style={styles.fileName} numberOfLines={1}>
                    {item.name}
                </Text>
                <Text style={styles.fileMeta}>
                    {item.size !== '—' ? `${formatBytes(parseInt(item.size))} · ` : ''}
                    {item.modifiedAt}
                </Text>
            </View>
        </TouchableOpacity>
    );
}
