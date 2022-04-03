/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */

import type { peerStatus } from './peerStatus';

export type clusterStatus = {
  name?: string;
  status: 'ready' | 'settling' | 'disabled';
  peers?: Array<peerStatus>;
};
