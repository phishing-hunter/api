# PhishingHunterApi.HuntingApi

All URIs are relative to *https://api.phishing-hunter.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**getHunting**](HuntingApi.md#getHunting) | **GET** /hunting | ハンティング設定の取得
[**postHunting**](HuntingApi.md#postHunting) | **POST** /hunting | ハンティング設定の更新



## getHunting

> GetHunting200Response getHunting()

ハンティング設定の取得

### Example

```javascript
import PhishingHunterApi from 'phishing_hunter_api';
let defaultClient = PhishingHunterApi.ApiClient.instance;
// Configure API key authorization: custom_authorizer
let custom_authorizer = defaultClient.authentications['custom_authorizer'];
custom_authorizer.apiKey = 'YOUR API KEY';
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//custom_authorizer.apiKeyPrefix = 'Token';

let apiInstance = new PhishingHunterApi.HuntingApi();
apiInstance.getHunting((error, data, response) => {
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

[**GetHunting200Response**](GetHunting200Response.md)

### Authorization

[custom_authorizer](../README.md#custom_authorizer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## postHunting

> PostScannerAdd200Response postHunting(opts)

ハンティング設定の更新

### Example

```javascript
import PhishingHunterApi from 'phishing_hunter_api';
let defaultClient = PhishingHunterApi.ApiClient.instance;
// Configure API key authorization: custom_authorizer
let custom_authorizer = defaultClient.authentications['custom_authorizer'];
custom_authorizer.apiKey = 'YOUR API KEY';
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//custom_authorizer.apiKeyPrefix = 'Token';

let apiInstance = new PhishingHunterApi.HuntingApi();
let opts = {
  'getHunting200Response': {"keywords":{"domain":["paypal","icloud","gmail"],"dom":["pay"]},"target_score":100} // GetHunting200Response | 
};
apiInstance.postHunting(opts, (error, data, response) => {
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
 **getHunting200Response** | [**GetHunting200Response**](GetHunting200Response.md)|  | [optional] 

### Return type

[**PostScannerAdd200Response**](PostScannerAdd200Response.md)

### Authorization

[custom_authorizer](../README.md#custom_authorizer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

