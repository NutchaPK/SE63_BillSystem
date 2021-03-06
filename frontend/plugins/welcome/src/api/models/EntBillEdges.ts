/* tslint:disable */
/* eslint-disable */
/**
 * SUT SA Example API
 * This is a sample server for SUT SE 2563
 *
 * The version of the OpenAPI document: 1.0
 * Contact: support@swagger.io
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import {
    EntFinancier,
    EntFinancierFromJSON,
    EntFinancierFromJSONTyped,
    EntFinancierToJSON,
    EntPaytype,
    EntPaytypeFromJSON,
    EntPaytypeFromJSONTyped,
    EntPaytypeToJSON,
    EntUnpaybill,
    EntUnpaybillFromJSON,
    EntUnpaybillFromJSONTyped,
    EntUnpaybillToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntBillEdges
 */
export interface EntBillEdges {
    /**
     * 
     * @type {EntFinancier}
     * @memberof EntBillEdges
     */
    officer?: EntFinancier;
    /**
     * 
     * @type {EntPaytype}
     * @memberof EntBillEdges
     */
    paytype?: EntPaytype;
    /**
     * 
     * @type {EntUnpaybill}
     * @memberof EntBillEdges
     */
    treatment?: EntUnpaybill;
}

export function EntBillEdgesFromJSON(json: any): EntBillEdges {
    return EntBillEdgesFromJSONTyped(json, false);
}

export function EntBillEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntBillEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'officer': !exists(json, 'officer') ? undefined : EntFinancierFromJSON(json['officer']),
        'paytype': !exists(json, 'paytype') ? undefined : EntPaytypeFromJSON(json['paytype']),
        'treatment': !exists(json, 'treatment') ? undefined : EntUnpaybillFromJSON(json['treatment']),
    };
}

export function EntBillEdgesToJSON(value?: EntBillEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'officer': EntFinancierToJSON(value.officer),
        'paytype': EntPaytypeToJSON(value.paytype),
        'treatment': EntUnpaybillToJSON(value.treatment),
    };
}


