import React from 'react';

import { Matches, Search } from './components';

export const Home = () => (
  <div className="w-full flex lg:flex-row flex-col space-x-0 lg:space-x-8">
    <div className="w-full lg:w-4/6 flex flex-col">
      <div className="w-full flex flex-col space-y-8">
        <div className="card bg-neutral text-neutral-content card-compact w-full">
          <div className="card-body flex">
            <Matches />
          </div>
        </div>
      </div>
    </div>

    <div className="w-full lg:w-2/6 flex flex-col space-y-8">
      <div className="w-full flex flex-col space-y-8">
        <div className="card bg-neutral text-neutral-content card-compact w-full">
          <div className="card-body flex">
            <h4 className="text-3xl pb-2">Select team</h4>
            <Search />
          </div>
        </div>
      </div>
    </div>
  </div>
);
