export interface SearchResult {
  id: number;
  name: string;
}

export interface TeamMatchGroup {
  date: string;
  matches: TeamMatch[];
  display: boolean;
}

export interface TeamMatch {
  id: number;
  team1: string;
  team2: string;
  eventName: string;
  type: string;
  display: boolean;
}

export interface MatchMap {
  title: string;
  url: string;
  display: boolean;
}
