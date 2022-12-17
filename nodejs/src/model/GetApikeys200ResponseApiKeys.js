/**
 * Phishing Hunter API
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * The version of the OpenAPI document: 1.0
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 *
 */

import ApiClient from '../ApiClient';

/**
 * The GetApikeys200ResponseApiKeys model module.
 * @module model/GetApikeys200ResponseApiKeys
 * @version 1.0
 */
class GetApikeys200ResponseApiKeys {
    /**
     * Constructs a new <code>GetApikeys200ResponseApiKeys</code>.
     * @alias module:model/GetApikeys200ResponseApiKeys
     */
    constructor() { 
        
        GetApikeys200ResponseApiKeys.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>GetApikeys200ResponseApiKeys</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/GetApikeys200ResponseApiKeys} obj Optional instance to populate.
     * @return {module:model/GetApikeys200ResponseApiKeys} The populated <code>GetApikeys200ResponseApiKeys</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new GetApikeys200ResponseApiKeys();

            if (data.hasOwnProperty('urlscan')) {
                obj['urlscan'] = ApiClient.convertToType(data['urlscan'], 'String');
            }
            if (data.hasOwnProperty('shodan')) {
                obj['shodan'] = ApiClient.convertToType(data['shodan'], 'String');
            }
            if (data.hasOwnProperty('vt')) {
                obj['vt'] = ApiClient.convertToType(data['vt'], 'String');
            }
        }
        return obj;
    }

    /**
     * Validates the JSON data with respect to <code>GetApikeys200ResponseApiKeys</code>.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @return {boolean} to indicate whether the JSON data is valid with respect to <code>GetApikeys200ResponseApiKeys</code>.
     */
    static validateJSON(data) {
        // ensure the json data is a string
        if (data['urlscan'] && !(typeof data['urlscan'] === 'string' || data['urlscan'] instanceof String)) {
            throw new Error("Expected the field `urlscan` to be a primitive type in the JSON string but got " + data['urlscan']);
        }
        // ensure the json data is a string
        if (data['shodan'] && !(typeof data['shodan'] === 'string' || data['shodan'] instanceof String)) {
            throw new Error("Expected the field `shodan` to be a primitive type in the JSON string but got " + data['shodan']);
        }
        // ensure the json data is a string
        if (data['vt'] && !(typeof data['vt'] === 'string' || data['vt'] instanceof String)) {
            throw new Error("Expected the field `vt` to be a primitive type in the JSON string but got " + data['vt']);
        }

        return true;
    }


}



/**
 * @member {String} urlscan
 */
GetApikeys200ResponseApiKeys.prototype['urlscan'] = undefined;

/**
 * @member {String} shodan
 */
GetApikeys200ResponseApiKeys.prototype['shodan'] = undefined;

/**
 * @member {String} vt
 */
GetApikeys200ResponseApiKeys.prototype['vt'] = undefined;






export default GetApikeys200ResponseApiKeys;
