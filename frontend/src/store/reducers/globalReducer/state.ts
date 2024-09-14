import { getItem } from '../../../utils/localStorage';
import { LocalStorageKeys } from '../../../utils/localStorageKeys';
import { GlobalState } from './types';

const localStorageTeamId = getItem(LocalStorageKeys.LOCAL_STORAGE_KEY_TEAM_ID);

export const fallbackInitialState: GlobalState = {
  teamId: localStorageTeamId === null ? null : parseInt(localStorageTeamId, 10),
  teamName: getItem(LocalStorageKeys.LOCAL_STORAGE_KEY_TEAM_NAME),
  matches: [],
  startingPoint: 'two-weeks',
  maps: [],
};
