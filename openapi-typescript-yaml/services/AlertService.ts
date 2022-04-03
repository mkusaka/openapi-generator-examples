/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { gettableAlerts } from '../models/gettableAlerts';
import type { postableAlerts } from '../models/postableAlerts';

import type { CancelablePromise } from '../core/CancelablePromise';
import { OpenAPI } from '../core/OpenAPI';
import { request as __request } from '../core/request';

export class AlertService {

  /**
   * Get a list of alerts
   * @param active Show active alerts
   * @param silenced Show silenced alerts
   * @param inhibited Show inhibited alerts
   * @param unprocessed Show unprocessed alerts
   * @param filter A list of matchers to filter alerts by
   * @param receiver A regex matching receivers to filter alerts by
   * @returns gettableAlerts Get alerts response
   * @throws ApiError
   */
  public static getAlerts(
    active: boolean = true,
    silenced: boolean = true,
    inhibited: boolean = true,
    unprocessed: boolean = true,
    filter?: Array<string>,
    receiver?: string,
  ): CancelablePromise<gettableAlerts> {
    return __request(OpenAPI, {
      method: 'GET',
      url: '/alerts',
      query: {
        'active': active,
        'silenced': silenced,
        'inhibited': inhibited,
        'unprocessed': unprocessed,
        'filter': filter,
        'receiver': receiver,
      },
      errors: {
        400: `Bad request`,
        500: `Internal server error`,
      },
    });
  }

  /**
   * Create new Alerts
   * @param alerts The alerts to create
   * @returns any Create alerts response
   * @throws ApiError
   */
  public static postAlerts(
    alerts: postableAlerts,
  ): CancelablePromise<any> {
    return __request(OpenAPI, {
      method: 'POST',
      url: '/alerts',
      body: alerts,
      errors: {
        400: `Bad request`,
        500: `Internal server error`,
      },
    });
  }

}