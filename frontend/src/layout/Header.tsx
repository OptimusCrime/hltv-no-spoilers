import React from 'react';

import { Menu } from './Menu';

export const Header = () => {
  return (
    <div className="navbar p-0">
      <div className="navbar-start z-[9999]">
        <div className="dropdown">
          <label tabIndex={0} className="btn btn-ghost lg:hidden">
            <svg
              xmlns="http://www.w3.org/2000/svg"
              className="h-5 w-5"
              fill="none"
              viewBox="0 0 24 24"
              stroke="currentColor"
            >
              <path strokeLinecap="round" strokeLinejoin="round" strokeWidth="2" d="M4 6h16M4 12h8m-8 6h16" />
            </svg>
          </label>
        </div>
        <a
          href="#"
          className="normal-case text-xl"
          onClick={(e) => {
            e.preventDefault();

            // TODO
          }}
        >
          HLTV No Spoilers
        </a>
      </div>
      <Menu />
    </div>
  );
};
