import React, { useState } from 'react';
import { Text, TouchableOpacity, View, Platform, ActivityIndicator, Alert } from 'react-native';
import * as FileSystem from 'expo-file-system/legacy';
import * as Sharing from 'expo-sharing';
import * as IntentLauncher from 'expo-intent-launcher';

import { formatBytes, getFileIcon } from '../utils/utils';
import { FileItem } from '../models/types';
import { styles } from '../theme/styles';
import apiClient from '../api/apiClient';

// ---------- helpers ----------

const TOKEN = '123'; // замените на реальный источник токена

const getMimeType = (name: string): string => {
  const ext = name.split('.').pop()?.toLowerCase() ?? '';
  const map: Record<string, string> = {
    pdf: 'application/pdf',
    jpg: 'image/jpeg', jpeg: 'image/jpeg', png: 'image/png', gif: 'image/gif', webp: 'image/webp',
    mp4: 'video/mp4', mov: 'video/quicktime', mkv: 'video/x-matroska',
    mp3: 'audio/mpeg', m4a: 'audio/mp4', wav: 'audio/wav',
    txt: 'text/plain',
    doc: 'application/msword',
    docx: 'application/vnd.openxmlformats-officedocument.wordprocessingml.document',
    xls: 'application/vnd.ms-excel',
    xlsx: 'application/vnd.openxmlformats-officedocument.spreadsheetml.sheet',
    zip: 'application/zip',
  };
  return map[ext] ?? '*/*';
};

// Скачивает файл в кэш, если его там ещё нет. Возвращает локальный file:// URI.
const ensureCached = async (item: FileItem): Promise<string> => {
  const localUri = `${FileSystem.cacheDirectory}${item.name}`;

  const info = await FileSystem.getInfoAsync(localUri);
  if (info.exists) return localUri;

  const fileUrl = `${apiClient.baseUrl}/files/${item.id}/download`;
  const result = await FileSystem.downloadAsync(fileUrl, localUri, {
    headers: { Authorization: `Bearer ${TOKEN}` },
  });

  if (result.status !== 200) {
    throw new Error(`Download failed: HTTP ${result.status}`);
  }

  return result.uri;
};

// ---------- component ----------

export const FileListItem: React.FC<{ item: FileItem }> = ({ item }) => {
  const [busy, setBusy] = useState<null | 'preview' | 'download'>(null);

  const handlePreview = async () => {
    if (item.type === 'folder') return;
    if (busy) return;

    setBusy('preview');
    try {
      const localUri = await ensureCached(item);
      const contentUri = await FileSystem.getContentUriAsync(localUri);

      if (Platform.OS === 'android') {
        await IntentLauncher.startActivityAsync('android.intent.action.VIEW', {
          data: contentUri,
          flags: 1, // FLAG_GRANT_READ_URI_PERMISSION
          type: getMimeType(item.name),
        });
      } else {
        // iOS — expo-sharing умеет показывать системный preview sheet
        if (await Sharing.isAvailableAsync()) {
          await Sharing.shareAsync(localUri, { mimeType: getMimeType(item.name) });
        }
      }
    } catch (e) {
      console.warn('Preview failed:', e);
      Alert.alert('Не удалось открыть файл', String(e));
    } finally {
      setBusy(null);
    }
  };

  const handleDownload = async () => {
    if (item.type === 'folder') return;
    if (busy) return;

    setBusy('download');
    try {
      const localUri = await ensureCached(item);

      if (await Sharing.isAvailableAsync()) {
        await Sharing.shareAsync(localUri, {
          mimeType: getMimeType(item.name),
          dialogTitle: 'Сохранить файл',
        });
      } else {
        Alert.alert('Сохранено', localUri);
      }
    } catch (e) {
      console.warn('Download failed:', e);
      Alert.alert('Не удалось скачать файл', String(e));
    } finally {
      setBusy(null);
    }
  };

  return (
    <TouchableOpacity
      onPress={handlePreview}
      style={styles.fileRow}
      activeOpacity={0.6}
      disabled={busy !== null || item.type === 'folder'}
    >
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

      {item.type !== 'folder' && (
        <TouchableOpacity
          onPress={handleDownload}
          style={styles.downloadButton}
          activeOpacity={0.6}
          disabled={busy !== null}
          hitSlop={{ top: 8, bottom: 8, left: 8, right: 8 }}
        >
          {busy === 'download' ? (
            <ActivityIndicator size="small" color="#4F7CFF" />
          ) : (
            <Text style={styles.downloadIcon}>⬇️</Text>
          )}
        </TouchableOpacity>
      )}

      {busy === 'preview' && (
        <ActivityIndicator
          size="small"
          color="#4F7CFF"
          style={styles.previewSpinner}
        />
      )}
    </TouchableOpacity>
  );
};
