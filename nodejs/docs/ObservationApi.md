# PhishingHunterApi.ObservationApi

All URIs are relative to *https://api.phishing-hunter.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**getObservations**](ObservationApi.md#getObservations) | **GET** /observation | 監視対象ドメイン一覧の取得
[**postObservationAdd**](ObservationApi.md#postObservationAdd) | **POST** /observation/add | 監視対象ドメインの追加
[**postObservationRemove**](ObservationApi.md#postObservationRemove) | **POST** /observation/remove | 監視対象ドメインの削除



## getObservations

> GetObservations200Response getObservations()

監視対象ドメイン一覧の取得

### Example

```javascript
import PhishingHunterApi from 'phishing_hunter_api';
let defaultClient = PhishingHunterApi.ApiClient.instance;
// Configure API key authorization: custom_authorizer
let custom_authorizer = defaultClient.authentications['custom_authorizer'];
custom_authorizer.apiKey = 'YOUR API KEY';
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//custom_authorizer.apiKeyPrefix = 'Token';

let apiInstance = new PhishingHunterApi.ObservationApi();
apiInstance.getObservations((error, data, response) => {
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

[**GetObservations200Response**](GetObservations200Response.md)

### Authorization

[custom_authorizer](../README.md#custom_authorizer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## postObservationAdd

> GetScanner500Response postObservationAdd(opts)

監視対象ドメインの追加

### Example

```javascript
import PhishingHunterApi from 'phishing_hunter_api';
let defaultClient = PhishingHunterApi.ApiClient.instance;
// Configure API key authorization: custom_authorizer
let custom_authorizer = defaultClient.authentications['custom_authorizer'];
custom_authorizer.apiKey = 'YOUR API KEY';
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//custom_authorizer.apiKeyPrefix = 'Token';

let apiInstance = new PhishingHunterApi.ObservationApi();
let opts = {
  'postObservationAddRequest': {"url":"http://example-hoge.com","interval":"hour","period":1,"comment":"","expire":1671580800,"api_config":{"urlscan":{"visibility":"unlisted"}}} // PostObservationAddRequest | 
};
apiInstance.postObservationAdd(opts, (error, data, response) => {
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
 **postObservationAddRequest** | [**PostObservationAddRequest**](PostObservationAddRequest.md)|  | [optional] 

### Return type

[**GetScanner500Response**](GetScanner500Response.md)

### Authorization

[custom_authorizer](../README.md#custom_authorizer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## postObservationRemove

> GetScanner500Response postObservationRemove(opts)

監視対象ドメインの削除



### Example

```javascript
import PhishingHunterApi from 'phishing_hunter_api';
let defaultClient = PhishingHunterApi.ApiClient.instance;
// Configure API key authorization: custom_authorizer
let custom_authorizer = defaultClient.authentications['custom_authorizer'];
custom_authorizer.apiKey = 'YOUR API KEY';
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//custom_authorizer.apiKeyPrefix = 'Token';

let apiInstance = new PhishingHunterApi.ObservationApi();
let opts = {
  'postObservationRemoveRequest': {"url":"http://example-hoge.com"} // PostObservationRemoveRequest | 
};
apiInstance.postObservationRemove(opts, (error, data, response) => {
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
 **postObservationRemoveRequest** | [**PostObservationRemoveRequest**](PostObservationRemoveRequest.md)|  | [optional] 

### Return type

[**GetScanner500Response**](GetScanner500Response.md)

### Authorization

[custom_authorizer](../README.md#custom_authorizer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

