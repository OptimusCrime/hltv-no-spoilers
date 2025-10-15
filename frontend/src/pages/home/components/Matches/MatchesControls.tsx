import cx from 'classnames';
import React from 'react';

import { useAppDispatch, useAppSelector } from '../../../../store/hooks';
import { setStartingPoint } from '../../../../store/reducers/globalReducer';
import { ReducerNames } from '../../../../store/reducers/reducerNames';

export const MatchesControls = () => {
  const { startingPoint } = useAppSelector((state) => state[ReducerNames.GLOBAL]);
  const dispatch = useAppDispatch();

  return (
    <div className="flex flex-col w-full space-y-4 items-center">
      <div>
        <p>Matches loaded. Select starting point.</p>
      </div>
      <div className="flex md:flex-row flex-col bg-base-100 rounded-md p-4 w-auto max-w-[600px] space-y-4 md:space-y-0 md:space-x-4">
        <button
          className={cx('btn', { 'btn-active': startingPoint === 'one-week' })}
          onClick={() => dispatch(setStartingPoint('one-week'))}
        >
          One week ago
        </button>
        <button
          className={cx('btn', { 'btn-active': startingPoint === 'two-weeks' })}
          onClick={() => dispatch(setStartingPoint('two-weeks'))}
        >
          Two weeks ago
        </button>
        <button
          className={cx('btn', { 'btn-active': startingPoint === 'one-month' })}
          onClick={() => dispatch(setStartingPoint('one-month'))}
        >
          One month ago
        </button>
        <button
          className={cx('btn', { 'btn-active': startingPoint === 'way-back' })}
          onClick={() => dispatch(setStartingPoint('way-back'))}
        >
          Way back
        </button>
      </div>
    </div>
  );
};
