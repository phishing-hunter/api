# PhishingHunterApi.SettingsApi

All URIs are relative to *https://api.phishing-hunter.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**getApikeys**](SettingsApi.md#getApikeys) | **GET** /apikeys | 外部サービス設定の取得
[**getNotifySlack**](SettingsApi.md#getNotifySlack) | **GET** /notify | 通知設定の一覧を取得
[**postApikeys**](SettingsApi.md#postApikeys) | **POST** /apikeys | 外部サービス設定の更新
[**postNotifySlack**](SettingsApi.md#postNotifySlack) | **POST** /notify | 通知設定の更新



## getApikeys

> GetApikeys200Response getApikeys()

外部サービス設定の取得

### Example

```javascript
import PhishingHunterApi from 'phishing_hunter_api';
let defaultClient = PhishingHunterApi.ApiClient.instance;
// Configure API key authorization: custom_authorizer
let custom_authorizer = defaultClient.authentications['custom_authorizer'];
custom_authorizer.apiKey = 'YOUR API KEY';
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//custom_authorizer.apiKeyPrefix = 'Token';

let apiInstance = new PhishingHunterApi.SettingsApi();
apiInstance.getApikeys((error, data, response) => {
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

[**GetApikeys200Response**](GetApikeys200Response.md)

### Authorization

[custom_authorizer](../README.md#custom_authorizer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## getNotifySlack

> GetNotifySlack200Response getNotifySlack()

通知設定の一覧を取得

### Example

```javascript
import PhishingHunterApi from 'phishing_hunter_api';
let defaultClient = PhishingHunterApi.ApiClient.instance;
// Configure API key authorization: custom_authorizer
let custom_authorizer = defaultClient.authentications['custom_authorizer'];
custom_authorizer.apiKey = 'YOUR API KEY';
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//custom_authorizer.apiKeyPrefix = 'Token';

let apiInstance = new PhishingHunterApi.SettingsApi();
apiInstance.getNotifySlack((error, data, response) => {
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

[**GetNotifySlack200Response**](GetNotifySlack200Response.md)

### Authorization

[custom_authorizer](../README.md#custom_authorizer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## postApikeys

> PostScannerAdd200Response postApikeys(opts)

外部サービス設定の更新



### Example

```javascript
import PhishingHunterApi from 'phishing_hunter_api';
let defaultClient = PhishingHunterApi.ApiClient.instance;
// Configure API key authorization: custom_authorizer
let custom_authorizer = defaultClient.authentications['custom_authorizer'];
custom_authorizer.apiKey = 'YOUR API KEY';
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//custom_authorizer.apiKeyPrefix = 'Token';

let apiInstance = new PhishingHunterApi.SettingsApi();
let opts = {
  'getApikeys200Response': {"api_keys":{"urlscan":"","shodan":"","vt":""}} // GetApikeys200Response | 
};
apiInstance.postApikeys(opts, (error, data, response) => {
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
 **getApikeys200Response** | [**GetApikeys200Response**](GetApikeys200Response.md)|  | [optional] 

### Return type

[**PostScannerAdd200Response**](PostScannerAdd200Response.md)

### Authorization

[custom_authorizer](../README.md#custom_authorizer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## postNotifySlack

> PostScannerAdd200Response postNotifySlack(opts)

通知設定の更新

### Example

```javascript
import PhishingHunterApi from 'phishing_hunter_api';
let defaultClient = PhishingHunterApi.ApiClient.instance;
// Configure API key authorization: custom_authorizer
let custom_authorizer = defaultClient.authentications['custom_authorizer'];
custom_authorizer.apiKey = 'YOUR API KEY';
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//custom_authorizer.apiKeyPrefix = 'Token';

let apiInstance = new PhishingHunterApi.SettingsApi();
let opts = {
  'postNotifySlackRequest': {"type":"email","webhook_url":"","service":"hunting"} // PostNotifySlackRequest | 
};
apiInstance.postNotifySlack(opts, (error, data, response) => {
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
 **postNotifySlackRequest** | [**PostNotifySlackRequest**](PostNotifySlackRequest.md)|  | [optional] 

### Return type

[**PostScannerAdd200Response**](PostScannerAdd200Response.md)

### Authorization

[custom_authorizer](../README.md#custom_authorizer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

