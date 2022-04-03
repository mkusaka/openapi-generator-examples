/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */

export type Order = {
    id?: number | null;
    petId?: number;
    quantity?: number;
    shipDate?: string;
    /**
     * Order Status
     */
    status?: 'placed' | 'approved' | 'delivered';
    complete?: boolean;
};
