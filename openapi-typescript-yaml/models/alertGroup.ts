/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */

import type { gettableAlert } from './gettableAlert';
import type { labelSet } from './labelSet';
import type { receiver } from './receiver';

export type alertGroup = {
  labels: labelSet;
  receiver: receiver;
  alerts: Array<gettableAlert>;
};
