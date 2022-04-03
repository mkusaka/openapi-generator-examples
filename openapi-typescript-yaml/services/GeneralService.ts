/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { alertmanagerStatus } from '../models/alertmanagerStatus';

import type { CancelablePromise } from '../core/CancelablePromise';
import { OpenAPI } from '../core/OpenAPI';
import { request as __request } from '../core/request';

export class GeneralService {

  /**
   * Get current status of an Alertmanager instance and its cluster
   * @returns alertmanagerStatus Get status response
   * @throws ApiError
   */
  public static getStatus(): CancelablePromise<alertmanagerStatus> {
    return __request(OpenAPI, {
      method: 'GET',
      url: '/status',
    });
  }

}