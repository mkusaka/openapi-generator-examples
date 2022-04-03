/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */

import type { matchers } from './matchers';

export type silence = {
  matchers: matchers;
  startsAt: string;
  endsAt: string;
  createdBy: string;
  comment: string;
};
