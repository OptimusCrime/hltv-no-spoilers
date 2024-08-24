export enum LocalStorageKeys {
  LOCAL_STORAGE_KEY_TEAM_ID = 'hltv-no-spoiler-team-id',
  LOCAL_STORAGE_KEY_TEAM_NAME = 'hltv-no-spoiler-team-name',
}

export const isLocalStorageKeys = (value: string): value is LocalStorageKeys =>
  (Object.values(LocalStorageKeys) as string[]).includes(value);
