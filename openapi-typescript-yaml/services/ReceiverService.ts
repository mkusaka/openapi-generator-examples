/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { receiver } from '../models/receiver';

import type { CancelablePromise } from '../core/CancelablePromise';
import { OpenAPI } from '../core/OpenAPI';
import { request as __request } from '../core/request';

export class ReceiverService {

  /**
   * Get list of all receivers (name of notification integrations)
   * @returns receiver Get receivers response
   * @throws ApiError
   */
  public static getReceivers(): CancelablePromise<Array<receiver>> {
    return __request(OpenAPI, {
      method: 'GET',
      url: '/receivers',
    });
  }

}