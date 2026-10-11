import { useState } from 'react';

import { FileItem, QuickAction, StorageInfo } from './src/models/types';
import { HomeScreen } from './src/screens/HomeScreen';

const MOCK_STORAGE: StorageInfo = {
    used: 47.3 * 1024 * 1024 * 1024,
    total: 128 * 1024 * 1024 * 1024,
};

const MOCK_QUICK_ACTIONS: QuickAction[] = [
    { id: 'upload', label: 'Upload', icon: '📤' },
    { id: 'scan', label: 'Scan', icon: '📷' },
    { id: 'recent', label: 'Recent', icon: '🕐' },
    { id: 'offline', label: 'Offline', icon: '⬇️' },
];

const MOCK_FILES: FileItem[] = [
    { id: '1', name: 'Work Documents', size: '—', modifiedAt: 'Today', type: 'folder' },
    { id: '2', name: 'Phone Backup', size: '—', modifiedAt: 'Yesterday', type: 'folder' },
    { id: '3', name: 'Design_v3.psd', size: '24.8 MB', modifiedAt: '2h ago', type: 'image' },
    { id: '4', name: 'Product Demo.mp4', size: '156 MB', modifiedAt: 'Yesterday', type: 'video' },
    { id: '5', name: 'Quarterly Report.xlsx', size: '2.4 MB', modifiedAt: '3d ago', type: 'document' },
    { id: '6', name: 'Meeting Audio.m4a', size: '18.2 MB', modifiedAt: 'Last week', type: 'audio' },
    { id: '7', name: 'Project Assets.zip', size: '89 MB', modifiedAt: 'Last week', type: 'archive' },
    { id: '8', name: 'Personal Photos', size: '—', modifiedAt: '2w ago', type: 'folder' },
];

// ---------- Main screen (move to screens/Home/ later) ----------
export default function App() {
    const [files] = useState<FileItem[]>(MOCK_FILES);

    return <HomeScreen storage={MOCK_STORAGE} actions={MOCK_QUICK_ACTIONS} files={files} />;
}
