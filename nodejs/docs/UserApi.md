# PhishingHunterApi.UserApi

All URIs are relative to *https://api.phishing-hunter.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**getApikey**](UserApi.md#getApikey) | **GET** /apikey | APIキーの一覧を取得
[**getUsersInfo**](UserApi.md#getUsersInfo) | **GET** /users/info | ユーザ情報の取得
[**postApikeyCreate**](UserApi.md#postApikeyCreate) | **POST** /apikey/create | APIキーの作成
[**postApikeyDelete**](UserApi.md#postApikeyDelete) | **POST** /apikey/delete | APIキーの削除



## getApikey

> GetApikey200Response getApikey()

APIキーの一覧を取得

### Example

```javascript
import PhishingHunterApi from 'phishing_hunter_api';
let defaultClient = PhishingHunterApi.ApiClient.instance;
// Configure API key authorization: custom_authorizer
let custom_authorizer = defaultClient.authentications['custom_authorizer'];
custom_authorizer.apiKey = 'YOUR API KEY';
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//custom_authorizer.apiKeyPrefix = 'Token';

let apiInstance = new PhishingHunterApi.UserApi();
apiInstance.getApikey((error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters

This endpoint does not need any parameter.

### Return type

[**GetApikey200Response**](GetApikey200Response.md)

### Authorization

[custom_authorizer](../README.md#custom_authorizer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## getUsersInfo

> GetUsersInfo200Response getUsersInfo()

ユーザ情報の取得

ユーザを識別するための情報とAPI利用状況を返す

### Example

```javascript
import PhishingHunterApi from 'phishing_hunter_api';
let defaultClient = PhishingHunterApi.ApiClient.instance;
// Configure API key authorization: custom_authorizer
let custom_authorizer = defaultClient.authentications['custom_authorizer'];
custom_authorizer.apiKey = 'YOUR API KEY';
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//custom_authorizer.apiKeyPrefix = 'Token';

let apiInstance = new PhishingHunterApi.UserApi();
apiInstance.getUsersInfo((error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters

This endpoint does not need any parameter.

### Return type

[**GetUsersInfo200Response**](GetUsersInfo200Response.md)

### Authorization

[custom_authorizer](../README.md#custom_authorizer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## postApikeyCreate

> PostApikeyCreate200Response postApikeyCreate(opts)

APIキーの作成

### Example

```javascript
import PhishingHunterApi from 'phishing_hunter_api';
let defaultClient = PhishingHunterApi.ApiClient.instance;
// Configure API key authorization: custom_authorizer
let custom_authorizer = defaultClient.authentications['custom_authorizer'];
custom_authorizer.apiKey = 'YOUR API KEY';
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//custom_authorizer.apiKeyPrefix = 'Token';

let apiInstance = new PhishingHunterApi.UserApi();
let opts = {
  'postApikeyCreateRequest': {"description":"test api key","ip":["127.0.0.1","192.168.0.1/24"],"scope":["get:hunting","get:users_info"]} // PostApikeyCreateRequest | 
};
apiInstance.postApikeyCreate(opts, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postApikeyCreateRequest** | [**PostApikeyCreateRequest**](PostApikeyCreateRequest.md)|  | [optional] 

### Return type

[**PostApikeyCreate200Response**](PostApikeyCreate200Response.md)

### Authorization

[custom_authorizer](../README.md#custom_authorizer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## postApikeyDelete

> PostApikeyDelete200Response postApikeyDelete(opts)

APIキーの削除

### Example

```javascript
import PhishingHunterApi from 'phishing_hunter_api';
let defaultClient = PhishingHunterApi.ApiClient.instance;
// Configure API key authorization: custom_authorizer
let custom_authorizer = defaultClient.authentications['custom_authorizer'];
custom_authorizer.apiKey = 'YOUR API KEY';
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//custom_authorizer.apiKeyPrefix = 'Token';

let apiInstance = new PhishingHunterApi.UserApi();
let opts = {
  'postApikeyCreate200Response': {"api_key":"string"} // PostApikeyCreate200Response | 
};
apiInstance.postApikeyDelete(opts, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postApikeyCreate200Response** | [**PostApikeyCreate200Response**](PostApikeyCreate200Response.md)|  | [optional] 

### Return type

[**PostApikeyDelete200Response**](PostApikeyDelete200Response.md)

### Authorization

[custom_authorizer](../README.md#custom_authorizer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

