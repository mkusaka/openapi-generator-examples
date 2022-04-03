/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */

import type { alert } from './alert';
import type { labelSet } from './labelSet';

export type postableAlert = ({
  startsAt?: string;
  endsAt?: string;
  annotations?: labelSet;
} & alert);
