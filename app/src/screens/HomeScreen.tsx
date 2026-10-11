import { SafeAreaView, ScrollView, StatusBar, Text, TouchableOpacity, View } from "react-native";

import { FileItem, QuickAction, StorageInfo } from "../models/types";
import { QuickActions } from "../components/QuickActions";
import { StorageCard } from "../components/StorageCard";
import { FileListItem } from "../components/FileItem";
import { styles } from "../theme/styles";

export const HomeScreen: React.FC<{ storage: StorageInfo, actions: QuickAction[], files: FileItem[] }> = ({ storage, actions, files }) => {
    return (
        <SafeAreaView style={styles.safeArea}>
            <StatusBar barStyle="light-content" />
            <ScrollView
                style={styles.scrollView}
                contentContainerStyle={styles.scrollContent}
                showsVerticalScrollIndicator={false}
            >
                {/* Header */}
                <View style={styles.header}>
                    <View>
                        <Text style={styles.greeting}>My Drive</Text>
                        <Text style={styles.subGreeting}>All files synced</Text>
                    </View>
                    <TouchableOpacity style={styles.avatar} activeOpacity={0.7}>
                        <Text style={styles.avatarText}>A</Text>
                    </TouchableOpacity>
                </View>

                {/* Storage */}
                <StorageCard info={storage} />

                {/* Quick actions */}
                <QuickActions actions={actions} />

                {/* Section header */}
                <View style={styles.sectionHeader}>
                    <Text style={styles.sectionTitle}>Recent Files</Text>
                    <TouchableOpacity activeOpacity={0.6}>
                        <Text style={styles.sectionAction}>See all</Text>
                    </TouchableOpacity>
                </View>

                {/* File list */}
                <View style={styles.fileList}>
                    {files.map((item) => (
                        <FileListItem key={item.id} item={item} />
                    ))}
                </View>
            </ScrollView>
        </SafeAreaView>
    );
}
