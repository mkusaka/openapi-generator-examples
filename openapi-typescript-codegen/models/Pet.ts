/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */

import type { Category } from './Category';
import type { Tag } from './Tag';

export type Pet = {
    id?: number;
    name: string;
    category?: Category;
    photoUrls: Array<string>;
    tags?: Array<Tag>;
    /**
     * pet status in the store
     */
    status?: 'available' | 'pending' | 'sold';
    stringRank?: 1 | 2;
    numberRank?: 1 | 2;
};
