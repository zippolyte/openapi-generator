/**
 * OpenAPI Petstore
 * This is a sample server Petstore server. For this sample, you can use the api key `special-key` to test the authorization filters.
 *
 * The version of the OpenAPI document: 1.0.0
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 *
 */


import ApiClient from './ApiClient';
import ApiResponse from './model/ApiResponse';
import Category from './model/Category';
import Order from './model/Order';
import Pet from './model/Pet';
import Tag from './model/Tag';
import User from './model/User';
import PetApi from './api/PetApi';
import StoreApi from './api/StoreApi';
import UserApi from './api/UserApi';


/**
* This_is_a_sample_server_Petstore_server__For_this_sample_you_can_use_the_api_key_special_key_to_test_the_authorization_filters_.<br>
* The <code>index</code> module provides access to constructors for all the classes which comprise the public API.
* <p>
* An AMD (recommended!) or CommonJS application will generally do something equivalent to the following:
* <pre>
* var OpenApiPetstore = require('index'); // See note below*.
* var xxxSvc = new OpenApiPetstore.XxxApi(); // Allocate the API class we're going to use.
* var yyyModel = new OpenApiPetstore.Yyy(); // Construct a model instance.
* yyyModel.someProperty = 'someValue';
* ...
* var zzz = xxxSvc.doSomething(yyyModel); // Invoke the service.
* ...
* </pre>
* <em>*NOTE: For a top-level AMD script, use require(['index'], function(){...})
* and put the application logic within the callback function.</em>
* </p>
* <p>
* A non-AMD browser application (discouraged) might do something like this:
* <pre>
* var xxxSvc = new OpenApiPetstore.XxxApi(); // Allocate the API class we're going to use.
* var yyy = new OpenApiPetstore.Yyy(); // Construct a model instance.
* yyyModel.someProperty = 'someValue';
* ...
* var zzz = xxxSvc.doSomething(yyyModel); // Invoke the service.
* ...
* </pre>
* </p>
* @module index
* @version 1.0.0
*/
export {
    /**
     * The ApiClient constructor.
     * @property {module:ApiClient}
     */
    ApiClient,

    /**
     * The ApiResponse model constructor.
     * @property {module:model/ApiResponse}
     */
    ApiResponse,

    /**
     * The Category model constructor.
     * @property {module:model/Category}
     */
    Category,

    /**
     * The Order model constructor.
     * @property {module:model/Order}
     */
    Order,

    /**
     * The Pet model constructor.
     * @property {module:model/Pet}
     */
    Pet,

    /**
     * The Tag model constructor.
     * @property {module:model/Tag}
     */
    Tag,

    /**
     * The User model constructor.
     * @property {module:model/User}
     */
    User,

    /**
    * The PetApi service constructor.
    * @property {module:api/PetApi}
    */
    PetApi,

    /**
    * The StoreApi service constructor.
    * @property {module:api/StoreApi}
    */
    StoreApi,

    /**
    * The UserApi service constructor.
    * @property {module:api/UserApi}
    */
    UserApi
};
