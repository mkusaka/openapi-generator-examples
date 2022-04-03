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
        '_1' = 1,
        '_2' = 2,
        '_3' = 3,
    }


}
