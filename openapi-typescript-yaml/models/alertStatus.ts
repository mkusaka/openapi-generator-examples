/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */

export type alertStatus = {
  state: 'unprocessed' | 'active' | 'suppressed';
  silencedBy: Array<string>;
  inhibitedBy: Array<string>;
};
