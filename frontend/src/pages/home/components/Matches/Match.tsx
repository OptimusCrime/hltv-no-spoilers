import React from 'react';

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
      <div className="space-y-4 pb-4">
        {matchGroup.matches
          .toReversed()
          .filter((match) => match.display)
          .map((match) => (
            <div key={match.id} className="flex flex-row bg-base-100 p-4 rounded-md">
              <div className="w-1/2 text-left">
                <h4 className="text-xl pb-2">
                  {match.team1} vs. {match.team2}
                </h4>
                <p>
                  {match.eventName} {match.type.length > 0 && `(${match.type})`}
                </p>
              </div>
              <Maps matchId={match.id} />
            </div>
          ))}
      </div>
    </div>
  );
};
