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
  url: string;
  team1: string;
  team2: string;
  eventName: string;
  matchType: string;
  display: boolean;
}
