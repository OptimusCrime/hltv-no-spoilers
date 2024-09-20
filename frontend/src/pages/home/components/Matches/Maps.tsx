import React from 'react';

import { getMatchMaps } from '../../../../api/endpoints/backendEndpoints';
import { useAppDispatch, useAppSelector } from '../../../../store/hooks';
import { setMatchMap, showOneMoreMap } from '../../../../store/reducers/globalReducer';
import { ReducerNames } from '../../../../store/reducers/reducerNames';

interface MapsProps {
  matchId: number;
}

export const Maps = (props: MapsProps) => {
  const { maps } = useAppSelector((state) => state[ReducerNames.GLOBAL]);
  const { matchId } = props;

  const dispatch = useAppDispatch();

  const mapsForMatch = maps.find((map) => map.matchId === matchId);

  const revealMap = async (matchId: number) => {
    if (!mapsForMatch) {
      try {
        const matchMaps = await getMatchMaps(matchId);

        dispatch(
          setMatchMap({
            matchId: matchId,
            data: matchMaps.map((matchMap, idx) => ({
              ...matchMap,
              display: idx === 0,
            })),
          }),
        );
      } catch (err) {
        // Woops
      }

      return;
    }

    dispatch(showOneMoreMap(matchId));
  };

  return (
    <div className="w-1/2 border-l-[1px] border-base-300 pl-4 flex flex-col items-start">
      {mapsForMatch && (
        <div className="flex text-left flex-col space-y-4 pb-4">
          {mapsForMatch.data
            .filter((map) => map.display)
            .map((map, idx) => (
              <div key={`${matchId}_map_${idx}`} className="">
                <a href={map.url} target="_blank" rel="noopener noreferrer" className="underline hover:no-underline">
                  {map.title}
                </a>
              </div>
            ))}
        </div>
      )}
      <button className="btn" onClick={() => revealMap(matchId)}>
        Reveal map
      </button>
    </div>
  );
};
