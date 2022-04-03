/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { gettableSilence } from '../models/gettableSilence';
import type { gettableSilences } from '../models/gettableSilences';
import type { postableSilence } from '../models/postableSilence';

import type { CancelablePromise } from '../core/CancelablePromise';
import { OpenAPI } from '../core/OpenAPI';
import { request as __request } from '../core/request';

export class SilenceService {

  /**
   * Get a list of silences
   * @param filter A list of matchers to filter silences by
   * @returns gettableSilences Get silences response
   * @throws ApiError
   */
  public static getSilences(
    filter?: Array<string>,
  ): CancelablePromise<gettableSilences> {
    return __request(OpenAPI, {
      method: 'GET',
      url: '/silences',
      query: {
        'filter': filter,
      },
      errors: {
        500: `Internal server error`,
      },
    });
  }

  /**
   * Post a new silence or update an existing one
   * @param silence The silence to create
   * @returns any Create / update silence response
   * @throws ApiError
   */
  public static postSilences(
    silence: postableSilence,
  ): CancelablePromise<{
    silenceID?: string;
  }> {
    return __request(OpenAPI, {
      method: 'POST',
      url: '/silences',
      body: silence,
      errors: {
        400: `Bad request`,
        404: `A silence with the specified ID was not found`,
      },
    });
  }

  /**
   * Get a silence by its ID
   * @param silenceId ID of the silence to get
   * @returns gettableSilence Get silence response
   * @throws ApiError
   */
  public static getSilence(
    silenceId: string,
  ): CancelablePromise<gettableSilence> {
    return __request(OpenAPI, {
      method: 'GET',
      url: '/silence/{silenceID}',
      path: {
        'silenceID': silenceId,
      },
      errors: {
        404: `A silence with the specified ID was not found`,
        500: `Internal server error`,
      },
    });
  }

  /**
   * Delete a silence by its ID
   * @param silenceId ID of the silence to get
   * @param silenceId ID of the silence to get
   * @returns any Delete silence response
   * @throws ApiError
   */
  public static deleteSilence(
    silenceId: string,
    silenceId: string,
  ): CancelablePromise<any> {
    return __request(OpenAPI, {
      method: 'DELETE',
      url: '/silence/{silenceID}',
      path: {
        'silenceID': silenceId,
        'silenceID': silenceId,
      },
      errors: {
        500: `Internal server error`,
      },
    });
  }

}