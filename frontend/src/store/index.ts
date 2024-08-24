import { configureStore } from '@reduxjs/toolkit';

import globalReducer from './reducers/globalReducer/globalReducer';
import { ReducerNames } from './reducers/reducerNames';

export const store = configureStore({
  reducer: {
    [ReducerNames.GLOBAL]: globalReducer,
  },
});

export type AppDispatch = typeof store.dispatch;
export type RootState = ReturnType<typeof store.getState>;
