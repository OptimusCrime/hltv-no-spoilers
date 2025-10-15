const createLastShownKey = (teamId: number) => `HLTV_LAST_SHOWN_${teamId}`;

export const getLastShownMatch = (teamId: number) => {
  const data = localStorage.getItem(createLastShownKey(teamId));
  if (!data) {
    return null;
  }

  return parseInt(data);
}

export const setLastShownMatch = (teamId: number, matchId: number) => localStorage.setItem(createLastShownKey(teamId), matchId.toString());
