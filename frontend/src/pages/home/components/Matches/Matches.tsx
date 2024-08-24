import React, { useState } from 'react';
import { useEffect } from 'react';

import { getTeamMatches } from '../../../../api/endpoints/backendEndpoints';
import { useAppDispatch, useAppSelector } from '../../../../store/hooks';
import { showOneMoreMatch, setMatches } from '../../../../store/reducers/globalReducer';
import { ReducerNames } from '../../../../store/reducers/reducerNames';
import { MatchesControls } from './MatchesControls';

interface HistoryWrapperProps {
  teamName: string | null;
  children: React.ReactNode;
}

const MatchesWrapper = (props: HistoryWrapperProps) => {
  const { teamName, children } = props;

  return (
    <>
      <h4 className="text-3xl pb-2">{teamName === null ? 'Matches' : `Matches for ${teamName}`}</h4>
      <div className="w-4/6">
        {children}
      </div>
    </>
  );
};

export const Matches = () => {
  const { teamId, teamName, matches } = useAppSelector((state) => state[ReducerNames.GLOBAL]);
  const dispatch = useAppDispatch();

  const [isLoading, setIsLoading] = useState<boolean>(true);
  const [isError, setIsError] = useState<boolean>(false);

  const load = async () => {
    if (teamId === null || teamName === null) {
      setIsLoading(false);
      return setIsError(true);
    }

    setIsLoading(true);
    dispatch(setMatches([]));

    try {
      const matchGroups = await getTeamMatches(teamId);

      // Shortest way to the finish line here you guys
      dispatch(setMatches(matchGroups.map(matchGroup => ({
        ...matchGroup,
        display: false,
        matches: matchGroup.matches.map(match => ({
          ...match,
          display: false,
        })),
      }))));
    } catch (_) {
      setIsError(true);
    }

    setIsLoading(false);
  };

  useEffect(() => {
    load();
  }, [teamId, teamName]);

  if (isLoading) {
    return (
      <MatchesWrapper teamName={teamName}>
        <div className="flex flex-col w-full items-center h-[calc(50vh-2rem)] justify-end">
          <span className="loading loading-spinner loading-lg"></span>
          <div className="pt-10 text-center px-10">
            <span>Loading matches</span>
          </div>
        </div>
      </MatchesWrapper>
    );
  }

  if (isError && (teamId === null || teamName === null)) {
    return (
      <MatchesWrapper teamName={teamName}>
        <div role="alert" className="alert alert-info">
          <span>Select a team to list match history.</span>
        </div>
      </MatchesWrapper>
    );
  }

  return (
    <MatchesWrapper teamName={teamName}>
      <MatchesControls />
      {matches
        .filter(matchGroup => matchGroup.display)
        .map((matchGroup, idx) => (
          <div key={idx}>
            {matchGroup.date}
            {matchGroup.matches
              .toReversed()
              .filter(match => match.display).map((match) => (
                <div key={match.url}>
                  {match.url}
                </div>
              ))}
          </div>
        ))
      }
      <button
        className="btn"
        onClick={() => dispatch(showOneMoreMatch())}
      >
        Show more
      </button>
    </MatchesWrapper>
  );
};
