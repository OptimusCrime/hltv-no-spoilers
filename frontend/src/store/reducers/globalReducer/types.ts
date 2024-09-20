import { MatchMap, TeamMatchGroup } from '../../../types/common';

export type StartingPointType = 'one-week' | 'two-weeks' | 'one-month' | 'way-back';

export interface GlobalState {
  teamId: number | null;
  teamName: string | null;
  matches: TeamMatchGroup[];
  startingPoint: StartingPointType;
  maps: { matchId: number; data: MatchMap[] }[];
}
