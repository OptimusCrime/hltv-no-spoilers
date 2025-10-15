import React from 'react';

import { HLTV_URL } from '../../../../constants';
import { TeamMatchGroup } from '../../../../types/common';
import { parseDate } from '../../../../utils/dateHelpers';
import { Maps } from './Maps';

interface MatchProps {
  matchGroup: TeamMatchGroup;
}

export const Match = (props: MatchProps) => {
  const { matchGroup } = props;

  return (
    <div>
      <div className="flex justify-center">
        <h3 className="text-2xl pb-4">{parseDate(matchGroup.date)}</h3>
      </div>
      <div className="space-y-4 pb-4 flex justify-center flex-col items-center">
        {matchGroup.matches
          .toReversed()
          .filter((match) => match.display)
          .map((match) => (
            <div key={match.id} className="flex flex-col md:flex-row bg-base-100 p-4 rounded-md relative w-full max-w-[600px]">
              <div className="w-full md:w-1/2 md:text-left pb-6 md:pb-0 text-center">
                <h4 className="text-xl pb-2">
                  <a
                    href={`${HLTV_URL}${match.url}`}
                    target="blank"
                    rel="noopener noreferrer"
                    className="underline hover:no-underline"
                  >
                    {match.team1} vs. {match.team2}
                  </a>
                </h4>
                <p>
                  {match.eventName} {match.type.length > 0 && `(${match.type})`}
                </p>
                {match.alreadyShown && (
                  <p className='italic'>Already shown</p>
                )}
              </div>
              <Maps matchId={match.id} matchUri={match.uri} />
            </div>
          ))}
      </div>
    </div>
  );
};
