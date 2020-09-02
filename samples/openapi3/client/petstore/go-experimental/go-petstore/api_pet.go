/*
 * OpenAPI Petstore
 *
 * This spec is mainly for testing Petstore server and contains fake endpoints, models. Please do not use this for any other purpose. Special characters: \" \\
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package petstore

import (
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"
	"os"
)

// Linger please
var (
	_ _context.Context
)

type PetApi interface {

  /*
   * AddPet Add a new pet to the store
   * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
   * @return ApiAddPetRequest
   */
  AddPet(ctx _context.Context) ApiAddPetRequest

  /*
   * AddPetExecute executes the request
   */
  AddPetExecute(r ApiAddPetRequest) (*_nethttp.Response, error)

  /*
   * DeletePet Deletes a pet
   * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
   * @param petId Pet id to delete
   * @return ApiDeletePetRequest
   */
  DeletePet(ctx _context.Context, petId int64) ApiDeletePetRequest

  /*
   * DeletePetExecute executes the request
   */
  DeletePetExecute(r ApiDeletePetRequest) (*_nethttp.Response, error)

  /*
   * FindPetsByStatus Finds Pets by status
   * Multiple status values can be provided with comma separated strings
   * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
   * @return ApiFindPetsByStatusRequest
   */
  FindPetsByStatus(ctx _context.Context) ApiFindPetsByStatusRequest

  /*
   * FindPetsByStatusExecute executes the request
   * @return []Pet
   */
  FindPetsByStatusExecute(r ApiFindPetsByStatusRequest) ([]Pet, *_nethttp.Response, error)

  /*
   * FindPetsByTags Finds Pets by tags
   * Multiple tags can be provided with comma separated strings. Use tag1, tag2, tag3 for testing.
   * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
   * @return ApiFindPetsByTagsRequest
   */
  FindPetsByTags(ctx _context.Context) ApiFindPetsByTagsRequest

  /*
   * FindPetsByTagsExecute executes the request
   * @return []Pet
   */
  FindPetsByTagsExecute(r ApiFindPetsByTagsRequest) ([]Pet, *_nethttp.Response, error)

  /*
   * GetPetById Find pet by ID
   * Returns a single pet
   * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
   * @param petId ID of pet to return
   * @return ApiGetPetByIdRequest
   */
  GetPetById(ctx _context.Context, petId int64) ApiGetPetByIdRequest

  /*
   * GetPetByIdExecute executes the request
   * @return Pet
   */
  GetPetByIdExecute(r ApiGetPetByIdRequest) (Pet, *_nethttp.Response, error)

  /*
   * UpdatePet Update an existing pet
   * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
   * @return ApiUpdatePetRequest
   */
  UpdatePet(ctx _context.Context) ApiUpdatePetRequest

  /*
   * UpdatePetExecute executes the request
   */
  UpdatePetExecute(r ApiUpdatePetRequest) (*_nethttp.Response, error)

  /*
   * UpdatePetWithForm Updates a pet in the store with form data
   * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
   * @param petId ID of pet that needs to be updated
   * @return ApiUpdatePetWithFormRequest
   */
  UpdatePetWithForm(ctx _context.Context, petId int64) ApiUpdatePetWithFormRequest

  /*
   * UpdatePetWithFormExecute executes the request
   */
  UpdatePetWithFormExecute(r ApiUpdatePetWithFormRequest) (*_nethttp.Response, error)

  /*
   * UploadFile uploads an image
   * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
   * @param petId ID of pet to update
   * @return ApiUploadFileRequest
   */
  UploadFile(ctx _context.Context, petId int64) ApiUploadFileRequest

  /*
   * UploadFileExecute executes the request
   * @return ApiResponse
   */
  UploadFileExecute(r ApiUploadFileRequest) (ApiResponse, *_nethttp.Response, error)

  /*
   * UploadFileWithRequiredFile uploads an image (required)
   * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
   * @param petId ID of pet to update
   * @return ApiUploadFileWithRequiredFileRequest
   */
  UploadFileWithRequiredFile(ctx _context.Context, petId int64) ApiUploadFileWithRequiredFileRequest

  /*
   * UploadFileWithRequiredFileExecute executes the request
   * @return ApiResponse
   */
  UploadFileWithRequiredFileExecute(r ApiUploadFileWithRequiredFileRequest) (ApiResponse, *_nethttp.Response, error)
}

// PetApiService PetApi service
type PetApiService service

type ApiAddPetRequest struct {
	ctx _context.Context
	ApiService PetApi
	pet *Pet
}

func (r ApiAddPetRequest) Pet(pet Pet) ApiAddPetRequest {
	r.pet = &pet
	return r
}

func (r ApiAddPetRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.AddPetExecute(r)
}

/*
 * AddPet Add a new pet to the store
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiAddPetRequest
 */
func (a *PetApiService) AddPet(ctx _context.Context) ApiAddPetRequest {
	return ApiAddPetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 */
func (a *PetApiService) AddPetExecute(r ApiAddPetRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PetApiService.AddPet")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pet"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.pet == nil {
		return nil, reportError("pet is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json", "application/xml"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.pet
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiDeletePetRequest struct {
	ctx _context.Context
	ApiService PetApi
	petId int64
	apiKey *string
}

func (r ApiDeletePetRequest) ApiKey(apiKey string) ApiDeletePetRequest {
	r.apiKey = &apiKey
	return r
}

func (r ApiDeletePetRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.DeletePetExecute(r)
}

/*
 * DeletePet Deletes a pet
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param petId Pet id to delete
 * @return ApiDeletePetRequest
 */
func (a *PetApiService) DeletePet(ctx _context.Context, petId int64) ApiDeletePetRequest {
	return ApiDeletePetRequest{
		ApiService: a,
		ctx: ctx,
		petId: petId,
	}
}

/*
 * Execute executes the request
 */
func (a *PetApiService) DeletePetExecute(r ApiDeletePetRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PetApiService.DeletePet")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pet/{petId}"
	localVarPath = strings.Replace(localVarPath, "{"+"petId"+"}", _neturl.PathEscape(parameterToString(r.petId, "")) , -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.apiKey != nil {
		localVarHeaderParams["api_key"] = parameterToString(*r.apiKey, "")
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiFindPetsByStatusRequest struct {
	ctx _context.Context
	ApiService PetApi
	status *[]string
}

func (r ApiFindPetsByStatusRequest) Status(status []string) ApiFindPetsByStatusRequest {
	r.status = &status
	return r
}

func (r ApiFindPetsByStatusRequest) Execute() ([]Pet, *_nethttp.Response, error) {
	return r.ApiService.FindPetsByStatusExecute(r)
}

/*
 * FindPetsByStatus Finds Pets by status
 * Multiple status values can be provided with comma separated strings
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiFindPetsByStatusRequest
 */
func (a *PetApiService) FindPetsByStatus(ctx _context.Context) ApiFindPetsByStatusRequest {
	return ApiFindPetsByStatusRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return []Pet
 */
func (a *PetApiService) FindPetsByStatusExecute(r ApiFindPetsByStatusRequest) ([]Pet, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []Pet
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PetApiService.FindPetsByStatus")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pet/findByStatus"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.status == nil {
		return localVarReturnValue, nil, reportError("status is required and must be specified")
	}

	localVarQueryParams.Add("status", parameterToString(*r.status, "csv"))
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/xml", "application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiFindPetsByTagsRequest struct {
	ctx _context.Context
	ApiService PetApi
	tags *[]string
}

func (r ApiFindPetsByTagsRequest) Tags(tags []string) ApiFindPetsByTagsRequest {
	r.tags = &tags
	return r
}

func (r ApiFindPetsByTagsRequest) Execute() ([]Pet, *_nethttp.Response, error) {
	return r.ApiService.FindPetsByTagsExecute(r)
}

/*
 * FindPetsByTags Finds Pets by tags
 * Multiple tags can be provided with comma separated strings. Use tag1, tag2, tag3 for testing.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiFindPetsByTagsRequest
 */
func (a *PetApiService) FindPetsByTags(ctx _context.Context) ApiFindPetsByTagsRequest {
	return ApiFindPetsByTagsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return []Pet
 */
func (a *PetApiService) FindPetsByTagsExecute(r ApiFindPetsByTagsRequest) ([]Pet, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []Pet
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PetApiService.FindPetsByTags")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pet/findByTags"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.tags == nil {
		return localVarReturnValue, nil, reportError("tags is required and must be specified")
	}

	localVarQueryParams.Add("tags", parameterToString(*r.tags, "csv"))
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/xml", "application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetPetByIdRequest struct {
	ctx _context.Context
	ApiService PetApi
	petId int64
}


func (r ApiGetPetByIdRequest) Execute() (Pet, *_nethttp.Response, error) {
	return r.ApiService.GetPetByIdExecute(r)
}

/*
 * GetPetById Find pet by ID
 * Returns a single pet
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param petId ID of pet to return
 * @return ApiGetPetByIdRequest
 */
func (a *PetApiService) GetPetById(ctx _context.Context, petId int64) ApiGetPetByIdRequest {
	return ApiGetPetByIdRequest{
		ApiService: a,
		ctx: ctx,
		petId: petId,
	}
}

/*
 * Execute executes the request
 * @return Pet
 */
func (a *PetApiService) GetPetByIdExecute(r ApiGetPetByIdRequest) (Pet, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  Pet
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PetApiService.GetPetById")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pet/{petId}"
	localVarPath = strings.Replace(localVarPath, "{"+"petId"+"}", _neturl.PathEscape(parameterToString(r.petId, "")) , -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/xml", "application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["api_key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["api_key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiUpdatePetRequest struct {
	ctx _context.Context
	ApiService PetApi
	pet *Pet
}

func (r ApiUpdatePetRequest) Pet(pet Pet) ApiUpdatePetRequest {
	r.pet = &pet
	return r
}

func (r ApiUpdatePetRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.UpdatePetExecute(r)
}

/*
 * UpdatePet Update an existing pet
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiUpdatePetRequest
 */
func (a *PetApiService) UpdatePet(ctx _context.Context) ApiUpdatePetRequest {
	return ApiUpdatePetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 */
func (a *PetApiService) UpdatePetExecute(r ApiUpdatePetRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PetApiService.UpdatePet")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pet"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.pet == nil {
		return nil, reportError("pet is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json", "application/xml"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.pet
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiUpdatePetWithFormRequest struct {
	ctx _context.Context
	ApiService PetApi
	petId int64
	name *string
	status *string
}

func (r ApiUpdatePetWithFormRequest) Name(name string) ApiUpdatePetWithFormRequest {
	r.name = &name
	return r
}
func (r ApiUpdatePetWithFormRequest) Status(status string) ApiUpdatePetWithFormRequest {
	r.status = &status
	return r
}

func (r ApiUpdatePetWithFormRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.UpdatePetWithFormExecute(r)
}

/*
 * UpdatePetWithForm Updates a pet in the store with form data
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param petId ID of pet that needs to be updated
 * @return ApiUpdatePetWithFormRequest
 */
func (a *PetApiService) UpdatePetWithForm(ctx _context.Context, petId int64) ApiUpdatePetWithFormRequest {
	return ApiUpdatePetWithFormRequest{
		ApiService: a,
		ctx: ctx,
		petId: petId,
	}
}

/*
 * Execute executes the request
 */
func (a *PetApiService) UpdatePetWithFormExecute(r ApiUpdatePetWithFormRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PetApiService.UpdatePetWithForm")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pet/{petId}"
	localVarPath = strings.Replace(localVarPath, "{"+"petId"+"}", _neturl.PathEscape(parameterToString(r.petId, "")) , -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/x-www-form-urlencoded"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.name != nil {
		localVarFormParams.Add("name", parameterToString(*r.name, ""))
	}
	if r.status != nil {
		localVarFormParams.Add("status", parameterToString(*r.status, ""))
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiUploadFileRequest struct {
	ctx _context.Context
	ApiService PetApi
	petId int64
	additionalMetadata *string
	file **os.File
}

func (r ApiUploadFileRequest) AdditionalMetadata(additionalMetadata string) ApiUploadFileRequest {
	r.additionalMetadata = &additionalMetadata
	return r
}
func (r ApiUploadFileRequest) File(file *os.File) ApiUploadFileRequest {
	r.file = &file
	return r
}

func (r ApiUploadFileRequest) Execute() (ApiResponse, *_nethttp.Response, error) {
	return r.ApiService.UploadFileExecute(r)
}

/*
 * UploadFile uploads an image
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param petId ID of pet to update
 * @return ApiUploadFileRequest
 */
func (a *PetApiService) UploadFile(ctx _context.Context, petId int64) ApiUploadFileRequest {
	return ApiUploadFileRequest{
		ApiService: a,
		ctx: ctx,
		petId: petId,
	}
}

/*
 * Execute executes the request
 * @return ApiResponse
 */
func (a *PetApiService) UploadFileExecute(r ApiUploadFileRequest) (ApiResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ApiResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PetApiService.UploadFile")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pet/{petId}/uploadImage"
	localVarPath = strings.Replace(localVarPath, "{"+"petId"+"}", _neturl.PathEscape(parameterToString(r.petId, "")) , -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"multipart/form-data"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.additionalMetadata != nil {
		localVarFormParams.Add("additionalMetadata", parameterToString(*r.additionalMetadata, ""))
	}
	localVarFormFileName = "file"
	var localVarFile *os.File
	if r.file != nil {
		localVarFile = *r.file
	}
	if localVarFile != nil {
		fbs, _ := _ioutil.ReadAll(localVarFile)
		localVarFileBytes = fbs
		localVarFileName = localVarFile.Name()
		localVarFile.Close()
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiUploadFileWithRequiredFileRequest struct {
	ctx _context.Context
	ApiService PetApi
	petId int64
	requiredFile **os.File
	additionalMetadata *string
}

func (r ApiUploadFileWithRequiredFileRequest) RequiredFile(requiredFile *os.File) ApiUploadFileWithRequiredFileRequest {
	r.requiredFile = &requiredFile
	return r
}
func (r ApiUploadFileWithRequiredFileRequest) AdditionalMetadata(additionalMetadata string) ApiUploadFileWithRequiredFileRequest {
	r.additionalMetadata = &additionalMetadata
	return r
}

func (r ApiUploadFileWithRequiredFileRequest) Execute() (ApiResponse, *_nethttp.Response, error) {
	return r.ApiService.UploadFileWithRequiredFileExecute(r)
}

/*
 * UploadFileWithRequiredFile uploads an image (required)
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param petId ID of pet to update
 * @return ApiUploadFileWithRequiredFileRequest
 */
func (a *PetApiService) UploadFileWithRequiredFile(ctx _context.Context, petId int64) ApiUploadFileWithRequiredFileRequest {
	return ApiUploadFileWithRequiredFileRequest{
		ApiService: a,
		ctx: ctx,
		petId: petId,
	}
}

/*
 * Execute executes the request
 * @return ApiResponse
 */
func (a *PetApiService) UploadFileWithRequiredFileExecute(r ApiUploadFileWithRequiredFileRequest) (ApiResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ApiResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PetApiService.UploadFileWithRequiredFile")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/fake/{petId}/uploadImageWithRequiredFile"
	localVarPath = strings.Replace(localVarPath, "{"+"petId"+"}", _neturl.PathEscape(parameterToString(r.petId, "")) , -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.requiredFile == nil {
		return localVarReturnValue, nil, reportError("requiredFile is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"multipart/form-data"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.additionalMetadata != nil {
		localVarFormParams.Add("additionalMetadata", parameterToString(*r.additionalMetadata, ""))
	}
	localVarFormFileName = "requiredFile"
	localVarFile := *r.requiredFile
	if localVarFile != nil {
		fbs, _ := _ioutil.ReadAll(localVarFile)
		localVarFileBytes = fbs
		localVarFileName = localVarFile.Name()
		localVarFile.Close()
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
