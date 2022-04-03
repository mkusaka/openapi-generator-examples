/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { Pet } from '../models/Pet';

import type { CancelablePromise } from '../core/CancelablePromise';
import { OpenAPI } from '../core/OpenAPI';
import { request as __request } from '../core/request';

export class PetService {

    /**
     * Finds Pets by status
     * Multiple status values can be provided with comma separated strings
     * @param status Status values that need to be considered for filter
     * @returns Pet successful operation
     * @throws ApiError
     */
    public static findPetsByStatus(
        status: Array<string>,
    ): CancelablePromise<Array<Pet>> {
        return __request(OpenAPI, {
            method: 'GET',
            url: '/pet/findByStatus',
            query: {
                'status': status,
            },
            errors: {
                400: `Invalid status value`,
            },
        });
    }

}