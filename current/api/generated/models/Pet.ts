/* tslint:disable */
/* eslint-disable */
/**
 * Swagger Petstore - OpenAPI 3.0
 * This is a sample Pet Store Server based on the OpenAPI 3.0 specification.  You can find out more about Swagger at [http://swagger.io](http://swagger.io). In the third iteration of the pet store, we\'ve switched to the design first approach! You can now help us improve the API whether it\'s by making changes to the definition itself or to the code. That way, with time, we can improve the API in general, and expose some of the new features in OAS3.  Some useful links: - [The Pet Store repository](https://github.com/swagger-api/swagger-petstore) - [The source API definition for the Pet Store](https://github.com/swagger-api/swagger-petstore/blob/master/src/main/resources/openapi.yaml)
 *
 * The version of the OpenAPI document: 1.0.9
 * Contact: apiteam@swagger.io
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import {
    Category,
    CategoryFromJSON,
    CategoryFromJSONTyped,
    CategoryToJSON,
} from './Category';
import {
    Tag,
    TagFromJSON,
    TagFromJSONTyped,
    TagToJSON,
} from './Tag';

/**
 * 
 * @export
 * @interface Pet
 */
export interface Pet {
    /**
     * 
     * @type {number}
     * @memberof Pet
     */
    id?: number;
    /**
     * 
     * @type {string}
     * @memberof Pet
     */
    name: string;
    /**
     * 
     * @type {Category}
     * @memberof Pet
     */
    category?: Category;
    /**
     * 
     * @type {Array<string>}
     * @memberof Pet
     */
    photoUrls: Array<string>;
    /**
     * 
     * @type {Array<Tag>}
     * @memberof Pet
     */
    tags?: Array<Tag>;
    /**
     * pet status in the store
     * @type {string}
     * @memberof Pet
     */
    status?: PetStatusEnum;
    /**
     * 
     * @type {string}
     * @memberof Pet
     */
    stringRank?: PetStringRankEnum;
    /**
     * 
     * @type {number}
     * @memberof Pet
     */
    numberRank?: PetNumberRankEnum;
}


/**
 * @export
 */
export const PetStatusEnum = {
    Available: 'available',
    Pending: 'pending',
    Sold: 'sold'
} as const;

export type PetStatusEnum = typeof PetStatusEnum[keyof typeof PetStatusEnum];

/**
 * @export
 */
export const PetStringRankEnum = {
    _1: '1',
    _2: '2'
} as const;

export type PetStringRankEnum = typeof PetStringRankEnum[keyof typeof PetStringRankEnum];

/**
 * @export
 */
export const PetNumberRankEnum = {
    NUMBER_1: 1,
    NUMBER_2: 2
} as const;

export type PetNumberRankEnum = typeof PetNumberRankEnum[keyof typeof PetNumberRankEnum];


export function PetFromJSON(json: any): Pet {
    return PetFromJSONTyped(json, false);
}

export function PetFromJSONTyped(json: any, ignoreDiscriminator: boolean): Pet {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'id': !exists(json, 'id') ? undefined : json['id'],
        'name': json['name'],
        'category': !exists(json, 'category') ? undefined : CategoryFromJSON(json['category']),
        'photoUrls': json['photoUrls'],
        'tags': !exists(json, 'tags') ? undefined : ((json['tags'] as Array<any>).map(TagFromJSON)),
        'status': !exists(json, 'status') ? undefined : json['status'],
        'stringRank': !exists(json, 'stringRank') ? undefined : json['stringRank'],
        'numberRank': !exists(json, 'numberRank') ? undefined : json['numberRank'],
    };
}

export function PetToJSON(value?: Pet | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'id': value.id,
        'name': value.name,
        'category': CategoryToJSON(value.category),
        'photoUrls': value.photoUrls,
        'tags': value.tags === undefined ? undefined : ((value.tags as Array<any>).map(TagToJSON)),
        'status': value.status,
        'stringRank': value.stringRank,
        'numberRank': value.numberRank,
    };
}

