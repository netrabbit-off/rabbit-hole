import { FileItem } from "../models/types";

export const formatBytes = (bytes: number): string => {
    if (bytes === 0) return '0 B';
    const k = 1024;
    const sizes = ['B', 'KB', 'MB', 'GB'];
    const i = Math.floor(Math.log(bytes) / Math.log(k));
    return `${(bytes / Math.pow(k, i)).toFixed(1)} ${sizes[i]}`;
};

export const getFileIcon = (type: FileItem['type']): string => {
    const icons: Record<FileItem['type'], string> = {
        folder: '📁',
        image: '🖼️',
        video: '🎬',
        document: '📄',
        audio: '🎵',
        archive: '📦',
    };
    return icons[type];
};
