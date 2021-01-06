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
    EntTypetreatmentEdges,
    EntTypetreatmentEdgesFromJSON,
    EntTypetreatmentEdgesFromJSONTyped,
    EntTypetreatmentEdgesToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntTypetreatment
 */
export interface EntTypetreatment {
    /**
     * 
     * @type {EntTypetreatmentEdges}
     * @memberof EntTypetreatment
     */
    edges?: EntTypetreatmentEdges;
    /**
     * ID of the ent.
     * @type {number}
     * @memberof EntTypetreatment
     */
    id?: number;
    /**
     * Type holds the value of the "type" field.
     * @type {string}
     * @memberof EntTypetreatment
     */
    type?: string;
}

export function EntTypetreatmentFromJSON(json: any): EntTypetreatment {
    return EntTypetreatmentFromJSONTyped(json, false);
}

export function EntTypetreatmentFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntTypetreatment {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'edges': !exists(json, 'edges') ? undefined : EntTypetreatmentEdgesFromJSON(json['edges']),
        'id': !exists(json, 'id') ? undefined : json['id'],
        'type': !exists(json, 'type') ? undefined : json['type'],
    };
}

export function EntTypetreatmentToJSON(value?: EntTypetreatment | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'edges': EntTypetreatmentEdgesToJSON(value.edges),
        'id': value.id,
        'type': value.type,
    };
}


