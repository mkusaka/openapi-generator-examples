/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */

import type { alertmanagerConfig } from './alertmanagerConfig';
import type { clusterStatus } from './clusterStatus';
import type { versionInfo } from './versionInfo';

export type alertmanagerStatus = {
  cluster: clusterStatus;
  versionInfo: versionInfo;
  config: alertmanagerConfig;
  uptime: string;
};
