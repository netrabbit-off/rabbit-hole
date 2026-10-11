import { StyleSheet } from "react-native";
import { theme } from "./theme";

export const styles = StyleSheet.create({
    safeArea: {
        flex: 1,
        backgroundColor: theme.colors.background,
    },
    scrollView: {
        flex: 1,
    },
    scrollContent: {
        paddingHorizontal: theme.spacing.xl,
        paddingBottom: theme.spacing.xxl,
    },

    header: {
        flexDirection: 'row',
        justifyContent: 'space-between',
        alignItems: 'center',
        marginTop: theme.spacing.lg,
        marginBottom: theme.spacing.xxl,
    },
    greeting: {
        fontSize: 28,
        fontWeight: '700',
        color: theme.colors.textPrimary,
        letterSpacing: -0.5,
    },
    subGreeting: {
        fontSize: 14,
        color: theme.colors.textSecondary,
        marginTop: theme.spacing.xs,
    },
    avatar: {
        width: 44,
        height: 44,
        borderRadius: 22,
        backgroundColor: theme.colors.primary,
        alignItems: 'center',
        justifyContent: 'center',
    },
    avatarText: {
        fontSize: 18,
        fontWeight: '600',
        color: '#FFFFFF',
    },

    storageCard: {
        backgroundColor: theme.colors.surface,
        borderRadius: theme.radius.lg,
        padding: theme.spacing.xl,
        marginBottom: theme.spacing.xxl,
        borderWidth: 1,
        borderColor: theme.colors.border,
    },
    storageHeader: {
        flexDirection: 'row',
        justifyContent: 'space-between',
        alignItems: 'center',
        marginBottom: theme.spacing.lg,
    },
    storageTitle: {
        fontSize: 16,
        fontWeight: '600',
        color: theme.colors.textPrimary,
    },
    storageRemaining: {
        fontSize: 13,
        color: theme.colors.textSecondary,
    },
    storageBarBg: {
        height: 6,
        backgroundColor: theme.colors.border,
        borderRadius: 3,
        overflow: 'hidden',
        marginBottom: theme.spacing.md,
    },
    storageBarFill: {
        height: '100%',
        backgroundColor: theme.colors.primary,
        borderRadius: 3,
    },
    storageDetail: {
        fontSize: 13,
        color: theme.colors.textTertiary,
    },

    quickActions: {
        flexDirection: 'row',
        justifyContent: 'space-between',
        marginBottom: theme.spacing.xxl,
    },
    quickActionItem: {
        alignItems: 'center',
        flex: 1,
    },
    quickActionIcon: {
        width: 56,
        height: 56,
        borderRadius: theme.radius.md,
        backgroundColor: theme.colors.surface,
        alignItems: 'center',
        justifyContent: 'center',
        marginBottom: theme.spacing.sm,
        borderWidth: 1,
        borderColor: theme.colors.border,
    },
    quickActionEmoji: {
        fontSize: 22,
    },
    quickActionLabel: {
        fontSize: 12,
        color: theme.colors.textSecondary,
        fontWeight: '500',
    },

    sectionHeader: {
        flexDirection: 'row',
        justifyContent: 'space-between',
        alignItems: 'center',
        marginBottom: theme.spacing.lg,
    },
    sectionTitle: {
        fontSize: 18,
        fontWeight: '600',
        color: theme.colors.textPrimary,
    },
    sectionAction: {
        fontSize: 14,
        color: theme.colors.primary,
        fontWeight: '500',
    },

    fileList: {
        gap: theme.spacing.sm,
    },
    fileRow: {
        flexDirection: 'row',
        alignItems: 'center',
        backgroundColor: theme.colors.surface,
        borderRadius: theme.radius.md,
        padding: theme.spacing.lg,
        borderWidth: 1,
        borderColor: theme.colors.border,
    },
    fileIconWrap: {
        width: 44,
        height: 44,
        borderRadius: theme.radius.sm,
        backgroundColor: theme.colors.surfaceElevated,
        alignItems: 'center',
        justifyContent: 'center',
        marginRight: theme.spacing.lg,
    },
    fileIcon: {
        fontSize: 20,
    },
    fileInfo: {
        flex: 1,
    },
    fileName: {
        fontSize: 15,
        fontWeight: '500',
        color: theme.colors.textPrimary,
        marginBottom: 2,
    },
    fileMeta: {
        fontSize: 12,
        color: theme.colors.textTertiary,
    },
});

