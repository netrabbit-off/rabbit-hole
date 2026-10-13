type FileType = 'folder' | 'image' | 'video' | 'document' | 'audio' | 'archive';

const EXTENSION_MAP: Record<string, FileType> = {
    // Images
    jpg: 'image', jpeg: 'image', png: 'image', gif: 'image',
    webp: 'image', svg: 'image', bmp: 'image', heic: 'image',
    // Videos
    mp4: 'video', mov: 'video', avi: 'video', mkv: 'video',
    webm: 'video', m4v: 'video',
    // Audio
    mp3: 'audio', wav: 'audio', m4a: 'audio', flac: 'audio',
    ogg: 'audio', aac: 'audio',
    // Documents
    pdf: 'document', doc: 'document', docx: 'document',
    xls: 'document', xlsx: 'document', ppt: 'document', pptx: 'document',
    txt: 'document', md: 'document', csv: 'document',
    // Archives
    zip: 'archive', rar: 'archive', '7z': 'archive',
    tar: 'archive', gz: 'archive',
};

export const getFileType = (name: string, fallbackType?: FileType): FileType => {
    // If it's explicitly a folder — trust that, extensions don't apply
    if (fallbackType === 'folder') return 'folder';

    // Extract extension after the last dot
    const lastDot = name.includes('.') ? name.lastIndexOf('.') : 0;
    if (lastDot === -1 || lastDot === name.length - 1) {
        return fallbackType ?? 'document';
    }

    const ext = name.slice(lastDot + 1).toLowerCase();
    return EXTENSION_MAP[ext] ?? fallbackType ?? 'document';
};

export const getFileIcon = (name: string, fallbackType?: FileType): string => {
    const icons: Record<FileType, string> = {
        folder: '📁',
        image: '🖼️',
        video: '🎬',
        document: '📄',
        audio: '🎵',
        archive: '📦',
    };
    return icons[getFileType(name, fallbackType)];
};

export const formatBytes = (bytes: number): string => {
    if (bytes === 0) return '0 B';
    const k = 1024;
    const sizes = ['B', 'KB', 'MB', 'GB'];
    const i = Math.floor(Math.log(bytes) / Math.log(k));
    return `${(bytes / Math.pow(k, i)).toFixed(1)} ${sizes[i]}`;
};
