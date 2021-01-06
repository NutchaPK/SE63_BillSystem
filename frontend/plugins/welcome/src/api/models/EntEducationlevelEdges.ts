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
    EntDoctorinfo,
    EntDoctorinfoFromJSON,
    EntDoctorinfoFromJSONTyped,
    EntDoctorinfoToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntEducationlevelEdges
 */
export interface EntEducationlevelEdges {
    /**
     * Educationlevel2doctorinfo holds the value of the educationlevel2doctorinfo edge.
     * @type {Array<EntDoctorinfo>}
     * @memberof EntEducationlevelEdges
     */
    educationlevel2doctorinfo?: Array<EntDoctorinfo>;
}

export function EntEducationlevelEdgesFromJSON(json: any): EntEducationlevelEdges {
    return EntEducationlevelEdgesFromJSONTyped(json, false);
}

export function EntEducationlevelEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntEducationlevelEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'educationlevel2doctorinfo': !exists(json, 'educationlevel2doctorinfo') ? undefined : ((json['educationlevel2doctorinfo'] as Array<any>).map(EntDoctorinfoFromJSON)),
    };
}

export function EntEducationlevelEdgesToJSON(value?: EntEducationlevelEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'educationlevel2doctorinfo': value.educationlevel2doctorinfo === undefined ? undefined : ((value.educationlevel2doctorinfo as Array<any>).map(EntDoctorinfoToJSON)),
    };
}


