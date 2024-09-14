import { StartingPointType } from '../store/reducers/globalReducer/types';
import { TeamMatchGroup } from '../types/common';

export const revealMatchesFromStartingPoint = (matchGroups: TeamMatchGroup[], startingPoint: StartingPointType): TeamMatchGroup[] => {
  const startDate = startingPointToDate(startingPoint);
  const resetMatchGroups = resetMatches(matchGroups);

  if (startingPoint === 'way-back') {
    const numMatchGroups = resetMatchGroups.length - 1;
    const numMatchesInLastMatchGroup = resetMatchGroups[numMatchGroups].matches.length - 1;
    resetMatchGroups[numMatchGroups].display = true;
    resetMatchGroups[numMatchGroups].matches[numMatchesInLastMatchGroup].display = true;

    return resetMatchGroups;
  }

  let revealedGroup = false;

  return resetMatchGroups.map(matchGroup => {
    const matchGroupDate = new Date(matchGroup.date);
    if (matchGroupDate.getTime() > startDate.getTime()) {
      return matchGroup;
    }

    if (revealedGroup) {
      return matchGroup;
    }

    revealedGroup = true;

    let revealedMatch = false;

    return {
      ...matchGroup,
      display: true,
      matches: matchGroup.matches.map(match => {
        if (revealedMatch) {
          return match;
        }

        revealedMatch = true;

        return {
          ...match,
          display: true,
        };
      }),
    };
  });
};

// There has got to be a better way of doing this
export const revealOneMoreMatch = (matchGroups: TeamMatchGroup[]): TeamMatchGroup[] => {
  let foundStart = false;
  let revealed = false;

  // The matches are listed from newest to oldest, so we flip the list twice...
  return matchGroups
    .reverse()
    .map(matchGroup => {
      if (revealed) {
        return matchGroup;
      }

      const numMatchesInGroup = matchGroup.matches.length;
      const revealedMatchesInGroup = matchGroup.matches.filter(match => match.display).length;

      // There could be hidden matches before the starting point, so make sure that we have found at least one
      // visible match
      if (!foundStart) {
        if (revealedMatchesInGroup) {
          foundStart = true;
        }
        else {
          // Outside scope
          return matchGroup;
        }
      }

      if (numMatchesInGroup === revealedMatchesInGroup) {
        // All matches are already revealed
        return matchGroup;
      }

      revealed = true;

      return {
        ...matchGroup,
        display: true,
        matches: matchGroup.matches.map((match, idx) => {
          if (idx === revealedMatchesInGroup) {
            return {
              ...match,
              display: true,
            }
          }

          return match;
        })
      };
    })
    .reverse();
};

export const resetMatches = (matchGroups: TeamMatchGroup[]): TeamMatchGroup[] => matchGroups.map(matchGroup => ({
  ...matchGroup,
  display: false,
  matches: matchGroup.matches.map(match => ({
    ...match,
    display: false,
  })),
}));

const startingPointToDate = (startingPoint: StartingPointType): Date => {
  switch (startingPoint) {
    case 'one-week':
      return new Date(Date.now() - (60 * 60 * 24 * 7 * 1000));
    case 'two-weeks':
      return new Date(Date.now() - (60 * 60 * 24 * 14 * 1000));
    case 'one-month':
      return new Date(Date.now() - (60 * 60 * 24 * 30 * 1000));
    default:
    case 'way-back':
      return new Date(0);
  }
};
