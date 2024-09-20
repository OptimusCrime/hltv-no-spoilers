import React, { useState } from 'react';

import { search } from '../../../../api/endpoints/backendEndpoints';
import { useAppDispatch } from '../../../../store/hooks';
import { setTeam } from '../../../../store/reducers/globalReducer';
import { SearchResult } from '../../../../types/common';
import { setItem } from '../../../../utils/localStorage';
import { LocalStorageKeys } from '../../../../utils/localStorageKeys';

const MINIMUM_SEARCH_LENGTH = 2;

export const Search = () => {
  const dispatch = useAppDispatch();

  const [teamName, setTeamName] = useState<string>('');
  const [isLoading, setIsLoading] = useState<boolean>(false);
  const [isError, setIsError] = useState<boolean>(false);
  const [results, setResults] = useState<SearchResult[]>([]);

  const submit = async () => {
    if (teamName.length < MINIMUM_SEARCH_LENGTH) {
      return setIsError(true);
    }

    setIsLoading(true);
    setResults([]);

    try {
      const response = await search(teamName);
      setResults(response);
    } catch (_) {
      setIsError(true);
    }

    setIsLoading(false);
  };

  return (
    <div className="flex flex-col items-end space-y-4">
      <div className="w-full flex flex-row space-x-4">
        <input
          type="text"
          className="input input-bordered w-full max-w-xl"
          disabled={isLoading}
          onChange={(e) => {
            const value = e.target.value;
            setTeamName(value);
            setIsError(false);
            setResults([]);
          }}
          onKeyUp={async (e) => {
            if (e.code.toLowerCase() === 'enter' && !isLoading) {
              await submit();
            }
          }}
        />

        <button className="btn btn-primary" disabled={isLoading} onClick={submit}>
          {isLoading ? <span className="loading loading-spinner"></span> : 'Search'}
        </button>
      </div>

      {isError && (
        <div role="alert" className="alert alert-error">
          {teamName.length < MINIMUM_SEARCH_LENGTH ? (
            <span>You must enter at least two characters.</span>
          ) : (
            <span>Woops. Something went wrong.</span>
          )}
        </div>
      )}

      {results.length > 0 && (
        <div className="flex flex-col space-y-4 w-full pt-4">
          {results.map((result) => (
            <button
              className="btn w-full"
              key={result.id}
              onClick={() => {
                dispatch(
                  setTeam({
                    id: result.id,
                    name: result.name,
                  }),
                );

                setItem(LocalStorageKeys.LOCAL_STORAGE_KEY_TEAM_ID, result.id.toString());
                setItem(LocalStorageKeys.LOCAL_STORAGE_KEY_TEAM_NAME, result.name);

                setTeamName('');
                setResults([]);
              }}
            >
              {result.name}
            </button>
          ))}
        </div>
      )}
    </div>
  );
};
