/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */

import type { silence } from './silence';
import type { silenceStatus } from './silenceStatus';

export type gettableSilence = ({
  id: string;
  status: silenceStatus;
  updatedAt: string;
} & silence);
