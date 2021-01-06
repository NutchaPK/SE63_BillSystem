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
    EntDepartment,
    EntDepartmentFromJSON,
    EntDepartmentFromJSONTyped,
    EntDepartmentToJSON,
    EntEducationlevel,
    EntEducationlevelFromJSON,
    EntEducationlevelFromJSONTyped,
    EntEducationlevelToJSON,
    EntOfficeroom,
    EntOfficeroomFromJSON,
    EntOfficeroomFromJSONTyped,
    EntOfficeroomToJSON,
    EntPrename,
    EntPrenameFromJSON,
    EntPrenameFromJSONTyped,
    EntPrenameToJSON,
    EntTreatment,
    EntTreatmentFromJSON,
    EntTreatmentFromJSONTyped,
    EntTreatmentToJSON,
    EntUser,
    EntUserFromJSON,
    EntUserFromJSONTyped,
    EntUserToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntDoctorinfoEdges
 */
export interface EntDoctorinfoEdges {
    /**
     * 
     * @type {EntDepartment}
     * @memberof EntDoctorinfoEdges
     */
    department?: EntDepartment;
    /**
     * 
     * @type {EntEducationlevel}
     * @memberof EntDoctorinfoEdges
     */
    educationlevel?: EntEducationlevel;
    /**
     * 
     * @type {EntOfficeroom}
     * @memberof EntDoctorinfoEdges
     */
    officeroom?: EntOfficeroom;
    /**
     * 
     * @type {EntPrename}
     * @memberof EntDoctorinfoEdges
     */
    prename?: EntPrename;
    /**
     * Treatment holds the value of the treatment edge.
     * @type {Array<EntTreatment>}
     * @memberof EntDoctorinfoEdges
     */
    treatment?: Array<EntTreatment>;
    /**
     * 
     * @type {EntUser}
     * @memberof EntDoctorinfoEdges
     */
    user?: EntUser;
}

export function EntDoctorinfoEdgesFromJSON(json: any): EntDoctorinfoEdges {
    return EntDoctorinfoEdgesFromJSONTyped(json, false);
}

export function EntDoctorinfoEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntDoctorinfoEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'department': !exists(json, 'department') ? undefined : EntDepartmentFromJSON(json['department']),
        'educationlevel': !exists(json, 'educationlevel') ? undefined : EntEducationlevelFromJSON(json['educationlevel']),
        'officeroom': !exists(json, 'officeroom') ? undefined : EntOfficeroomFromJSON(json['officeroom']),
        'prename': !exists(json, 'prename') ? undefined : EntPrenameFromJSON(json['prename']),
        'treatment': !exists(json, 'treatment') ? undefined : ((json['treatment'] as Array<any>).map(EntTreatmentFromJSON)),
        'user': !exists(json, 'user') ? undefined : EntUserFromJSON(json['user']),
    };
}

export function EntDoctorinfoEdgesToJSON(value?: EntDoctorinfoEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'department': EntDepartmentToJSON(value.department),
        'educationlevel': EntEducationlevelToJSON(value.educationlevel),
        'officeroom': EntOfficeroomToJSON(value.officeroom),
        'prename': EntPrenameToJSON(value.prename),
        'treatment': value.treatment === undefined ? undefined : ((value.treatment as Array<any>).map(EntTreatmentToJSON)),
        'user': EntUserToJSON(value.user),
    };
}


