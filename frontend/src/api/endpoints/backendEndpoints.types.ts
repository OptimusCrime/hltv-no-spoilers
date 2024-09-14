/* eslint-disable @typescript-eslint/no-namespace */
import { SearchResult } from '../../types/common';

export namespace BackendEndpoints {
  export namespace Search {
    export type GET = SearchResult[];
  }

  export namespace TeamMatches {
     interface TeamMatchGroup {
      date: string;
      matches: TeamMatch[];
    }

    export interface TeamMatch {
      id: number;
      team1: string;
      team2: string;
      eventName: string;
      type: string;
    }

    export type GET = TeamMatchGroup[];
  }

  export namespace MatchMaps {
    interface MatchMap {
      title: string;
      url: string;
    }

    export type GET = MatchMap[];
  }
}
