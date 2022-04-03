/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */

export type Pet = {
    id?: number;
    name: string;
    /**
     * pet status in the store
     */
    status?: Pet.status;
};

export namespace Pet {

    /**
     * pet status in the store
     */
    export enum status {
        ONE = 'one',
        TWO = 'two',
        THREE = 'three',
    }


}
