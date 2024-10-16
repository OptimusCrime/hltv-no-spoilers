import ky from 'ky';

import { BackendEndpoints } from './backendEndpoints.types';

const api = ky.create({
  prefixUrl: process.env.NODE_ENV === 'production' ? 'https://hltv.optimuscrime.net' : 'http://localhost:8182',
  retry: 0,
});

export const search = (teamName: string) =>
  api
    .get(`v1/search?term=${teamName}`)
    .json<BackendEndpoints.Search.GET>()
    .then((res) => res);

export const getTeamMatches = (teamId: number) =>
  api
    .get(`v1/team/${teamId}/matches`)
    .json<BackendEndpoints.TeamMatches.GET>()
    .then((res) => res);

export const getMatchMaps = (params: { matchId: number; matchUri: string }) =>
  api
    .get(`v1/match/${params.matchId}?uri=${params.matchUri}`)
    .json<BackendEndpoints.MatchMaps.GET>()
    .then((res) => res);
