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
/**
 * 
 * @export
 * @interface ControllersPatientrecord
 */
export interface ControllersPatientrecord {
    /**
     * 
     * @type {number}
     * @memberof ControllersPatientrecord
     */
    age?: number;
    /**
     * 
     * @type {string}
     * @memberof ControllersPatientrecord
     */
    allergic?: string;
    /**
     * 
     * @type {string}
     * @memberof ControllersPatientrecord
     */
    birthday?: string;
    /**
     * 
     * @type {string}
     * @memberof ControllersPatientrecord
     */
    bloodtype?: string;
    /**
     * 
     * @type {string}
     * @memberof ControllersPatientrecord
     */
    date?: string;
    /**
     * 
     * @type {string}
     * @memberof ControllersPatientrecord
     */
    disease?: string;
    /**
     * 
     * @type {string}
     * @memberof ControllersPatientrecord
     */
    email?: string;
    /**
     * 
     * @type {number}
     * @memberof ControllersPatientrecord
     */
    gender?: number;
    /**
     * 
     * @type {string}
     * @memberof ControllersPatientrecord
     */
    home?: string;
    /**
     * 
     * @type {number}
     * @memberof ControllersPatientrecord
     */
    idcardnumber?: number;
    /**
     * 
     * @type {number}
     * @memberof ControllersPatientrecord
     */
    medicalrecordstaff?: number;
    /**
     * 
     * @type {string}
     * @memberof ControllersPatientrecord
     */
    name?: string;
    /**
     * 
     * @type {number}
     * @memberof ControllersPatientrecord
     */
    phonenumber?: number;
    /**
     * 
     * @type {number}
     * @memberof ControllersPatientrecord
     */
    prename?: number;
}

export function ControllersPatientrecordFromJSON(json: any): ControllersPatientrecord {
    return ControllersPatientrecordFromJSONTyped(json, false);
}

export function ControllersPatientrecordFromJSONTyped(json: any, ignoreDiscriminator: boolean): ControllersPatientrecord {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'age': !exists(json, 'age') ? undefined : json['age'],
        'allergic': !exists(json, 'allergic') ? undefined : json['allergic'],
        'birthday': !exists(json, 'birthday') ? undefined : json['birthday'],
        'bloodtype': !exists(json, 'bloodtype') ? undefined : json['bloodtype'],
        'date': !exists(json, 'date') ? undefined : json['date'],
        'disease': !exists(json, 'disease') ? undefined : json['disease'],
        'email': !exists(json, 'email') ? undefined : json['email'],
        'gender': !exists(json, 'gender') ? undefined : json['gender'],
        'home': !exists(json, 'home') ? undefined : json['home'],
        'idcardnumber': !exists(json, 'idcardnumber') ? undefined : json['idcardnumber'],
        'medicalrecordstaff': !exists(json, 'medicalrecordstaff') ? undefined : json['medicalrecordstaff'],
        'name': !exists(json, 'name') ? undefined : json['name'],
        'phonenumber': !exists(json, 'phonenumber') ? undefined : json['phonenumber'],
        'prename': !exists(json, 'prename') ? undefined : json['prename'],
    };
}

export function ControllersPatientrecordToJSON(value?: ControllersPatientrecord | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'age': value.age,
        'allergic': value.allergic,
        'birthday': value.birthday,
        'bloodtype': value.bloodtype,
        'date': value.date,
        'disease': value.disease,
        'email': value.email,
        'gender': value.gender,
        'home': value.home,
        'idcardnumber': value.idcardnumber,
        'medicalrecordstaff': value.medicalrecordstaff,
        'name': value.name,
        'phonenumber': value.phonenumber,
        'prename': value.prename,
    };
}


