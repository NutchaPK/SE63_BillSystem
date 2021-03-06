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
    EntPatientrightstypeEdges,
    EntPatientrightstypeEdgesFromJSON,
    EntPatientrightstypeEdgesFromJSONTyped,
    EntPatientrightstypeEdgesToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntPatientrightstype
 */
export interface EntPatientrightstype {
    /**
     * Permission holds the value of the "Permission" field.
     * @type {string}
     * @memberof EntPatientrightstype
     */
    permission?: string;
    /**
     * PermissionArea holds the value of the "PermissionArea" field.
     * @type {string}
     * @memberof EntPatientrightstype
     */
    permissionArea?: string;
    /**
     * Responsible holds the value of the "Responsible" field.
     * @type {string}
     * @memberof EntPatientrightstype
     */
    responsible?: string;
    /**
     * 
     * @type {number}
     * @memberof EntPatientrightstype
     */
    abilitypatientrightsId?: number;
    /**
     * 
     * @type {EntPatientrightstypeEdges}
     * @memberof EntPatientrightstype
     */
    edges?: EntPatientrightstypeEdges;
    /**
     * ID of the ent.
     * @type {number}
     * @memberof EntPatientrightstype
     */
    id?: number;
}

export function EntPatientrightstypeFromJSON(json: any): EntPatientrightstype {
    return EntPatientrightstypeFromJSONTyped(json, false);
}

export function EntPatientrightstypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntPatientrightstype {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'permission': !exists(json, 'Permission') ? undefined : json['Permission'],
        'permissionArea': !exists(json, 'PermissionArea') ? undefined : json['PermissionArea'],
        'responsible': !exists(json, 'Responsible') ? undefined : json['Responsible'],
        'abilitypatientrightsId': !exists(json, 'abilitypatientrights_id') ? undefined : json['abilitypatientrights_id'],
        'edges': !exists(json, 'edges') ? undefined : EntPatientrightstypeEdgesFromJSON(json['edges']),
        'id': !exists(json, 'id') ? undefined : json['id'],
    };
}

export function EntPatientrightstypeToJSON(value?: EntPatientrightstype | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'Permission': value.permission,
        'PermissionArea': value.permissionArea,
        'Responsible': value.responsible,
        'abilitypatientrights_id': value.abilitypatientrightsId,
        'edges': EntPatientrightstypeEdgesToJSON(value.edges),
        'id': value.id,
    };
}


