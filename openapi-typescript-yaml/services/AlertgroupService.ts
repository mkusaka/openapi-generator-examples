/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { alertGroups } from '../models/alertGroups';

import type { CancelablePromise } from '../core/CancelablePromise';
import { OpenAPI } from '../core/OpenAPI';
import { request as __request } from '../core/request';

export class AlertgroupService {

  /**
   * Get a list of alert groups
   * @param active Show active alerts
   * @param silenced Show silenced alerts
   * @param inhibited Show inhibited alerts
   * @param filter A list of matchers to filter alerts by
   * @param receiver A regex matching receivers to filter alerts by
   * @returns alertGroups Get alert groups response
   * @throws ApiError
   */
  public static getAlertGroups(
    active: boolean = true,
    silenced: boolean = true,
    inhibited: boolean = true,
    filter?: Array<string>,
    receiver?: string,
  ): CancelablePromise<alertGroups> {
    return __request(OpenAPI, {
      method: 'GET',
      url: '/alerts/groups',
      query: {
        'active': active,
        'silenced': silenced,
        'inhibited': inhibited,
        'filter': filter,
        'receiver': receiver,
      },
      errors: {
        400: `Bad request`,
        500: `Internal server error`,
      },
    });
  }

}