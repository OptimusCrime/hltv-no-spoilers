import ky from 'ky';

import { BackendEndpoints } from './backendEndpoints.types';

const api = ky.create({
  prefixUrl: process.env.NODE_ENV === 'production' ? '/api' : 'http://localhost:8182',
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

export const getMatchMaps = (matchId: number) =>
  api
    .get(`v1/match/${matchId}`)
    .json<BackendEndpoints.MatchMaps.GET>()
    .then((res) => res);
