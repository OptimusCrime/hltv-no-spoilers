import { createSlice, PayloadAction } from '@reduxjs/toolkit';

import { MatchMap, TeamMatchGroup } from '../../../types/common';
import { resetMatches, revealMatchesFromStartingPoint, revealOneMoreMatch } from '../../../utils/revealMatches';
import { ReducerNames } from '../reducerNames';
import { fallbackInitialState } from './state';
import { StartingPointType } from './types';

const globalReducer = createSlice({
  name: ReducerNames.GLOBAL,
  initialState: fallbackInitialState,
  reducers: {
    setTeam(state, action: PayloadAction<{ id: number; name: string }>) {
      state.teamId = action.payload.id;
      state.teamName = action.payload.name;
    },
    setMatches(state, action: PayloadAction<TeamMatchGroup[]>) {
      state.matches = action.payload;
      state.startingPoint = 'two-weeks';
    },
    setMatchMap(state, action: PayloadAction<{ matchId: number; data: MatchMap[] }>) {
      state.maps = [...state.maps, action.payload];
    },
    showOneMoreMap(state, action: PayloadAction<number>) {
      state.maps = state.maps.map((map) => {
        if (map.matchId !== action.payload) {
          return map;
        }

        let foundNotVisible = false;
        return {
          ...map,
          data: map.data.map((m) => {
            if (m.display) {
              return m;
            }

            if (!foundNotVisible) {
              foundNotVisible = true;

              return {
                ...m,
                display: true,
              };
            }

            return m;
          }),
        };
      });
    },
    setStartingPoint(state, action: PayloadAction<StartingPointType>) {
      state.startingPoint = action.payload;
      state.matches = resetMatches(state.matches);
    },
    showOneMoreMatch(state) {
      if (state.matches.some((match) => match.display)) {
        state.matches = revealOneMoreMatch(state.matches);
      } else {
        state.matches = revealMatchesFromStartingPoint(state.matches, state.startingPoint);
      }
    },
  },
});

export const { setTeam, setMatches, setMatchMap, showOneMoreMap, setStartingPoint, showOneMoreMatch } =
  globalReducer.actions;

export default globalReducer.reducer;
