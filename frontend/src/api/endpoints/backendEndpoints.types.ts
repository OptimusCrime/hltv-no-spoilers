/* eslint-disable @typescript-eslint/no-namespace */
import { SearchResult } from '../../types/common';

export namespace BackendEndpoints {
  export namespace Search {
    export type GET = SearchResult[];
  }

  export namespace TeamMatches {
    export interface TeamMatchGroup {
      date: string;
      matches: TeamMatch[];
    }

    export interface TeamMatch {
      url: string;
      team1: string;
      team2: string;
      eventName: string;
      matchType: string;
    }

    export type GET = TeamMatchGroup[];
  }
}
