import { createSlice, PayloadAction } from '@reduxjs/toolkit';

import { TeamMatchGroup } from '../../../types/common';
import { ReducerNames } from '../reducerNames';
import { getInitialState } from './state';
import { StartingPointType } from './types';
import { revealMatchesFromStartingPoint, revealOneMoreMatch } from '../../../utils/revealMatches';

const globalReducer = createSlice({
  name: ReducerNames.GLOBAL,
  initialState: getInitialState(),
  reducers: {
    setTeam(state, action: PayloadAction<{id: number; name: string; }>) {
      state.teamId = action.payload.id;
      state.teamName = action.payload.name;
    },
    setMatches(state, action: PayloadAction<TeamMatchGroup[]>) {
      state.matches = action.payload;
      state.startingPoint = null;
    },
    setStartingPoint(state, action: PayloadAction<StartingPointType>) {
      state.startingPoint = action.payload;
      state.matches = revealMatchesFromStartingPoint(state.matches, action.payload);
    },
    showOneMoreMatch(state) {
      state.matches = revealOneMoreMatch(state.matches);
    }
  },
});

export const {
  setTeam,
  setMatches,
  setStartingPoint,
  showOneMoreMatch
} = globalReducer.actions;

export default globalReducer.reducer;
