export interface StorageInfo {
    used: number;
    total: number;
}

export interface QuickAction {
    id: string;
    label: string;
    icon: string;
}

export interface FileItem {
    id: string;
    name: string;
    size: string;
    modifiedAt: string;
    type: 'folder' | 'image' | 'video' | 'document' | 'audio' | 'archive';
}
