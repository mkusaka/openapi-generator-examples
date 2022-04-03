/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */

import type { alert } from './alert';
import type { alertStatus } from './alertStatus';
import type { labelSet } from './labelSet';
import type { receiver } from './receiver';

export type gettableAlert = ({
  annotations: labelSet;
  receivers: Array<receiver>;
  fingerprint: string;
  startsAt: string;
  updatedAt: string;
  endsAt: string;
  status: alertStatus;
} & alert);
