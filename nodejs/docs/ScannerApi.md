# PhishingHunterApi.ScannerApi

All URIs are relative to *https://api.phishing-hunter.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**getScanner**](ScannerApi.md#getScanner) | **GET** /scanner | スキャン対象のドメイン一覧の取得
[**postScannerAdd**](ScannerApi.md#postScannerAdd) | **POST** /scanner/add | スキャン対象のドメインの追加
[**postScannerRemove**](ScannerApi.md#postScannerRemove) | **POST** /scanner/remove | スキャン対象のドメイン一覧の削除



## getScanner

> GetScanner200Response getScanner()

スキャン対象のドメイン一覧の取得

### Example

```javascript
import PhishingHunterApi from 'phishing_hunter_api';
let defaultClient = PhishingHunterApi.ApiClient.instance;
// Configure API key authorization: custom_authorizer
let custom_authorizer = defaultClient.authentications['custom_authorizer'];
custom_authorizer.apiKey = 'YOUR API KEY';
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//custom_authorizer.apiKeyPrefix = 'Token';

let apiInstance = new PhishingHunterApi.ScannerApi();
apiInstance.getScanner((error, data, response) => {
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

[**GetScanner200Response**](GetScanner200Response.md)

### Authorization

[custom_authorizer](../README.md#custom_authorizer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## postScannerAdd

> PostScannerAdd200Response postScannerAdd(opts)

スキャン対象のドメインの追加

### Example

```javascript
import PhishingHunterApi from 'phishing_hunter_api';
let defaultClient = PhishingHunterApi.ApiClient.instance;
// Configure API key authorization: custom_authorizer
let custom_authorizer = defaultClient.authentications['custom_authorizer'];
custom_authorizer.apiKey = 'YOUR API KEY';
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//custom_authorizer.apiKeyPrefix = 'Token';

let apiInstance = new PhishingHunterApi.ScannerApi();
let opts = {
  'postScannerAddRequest': {"domain":"www.google.com"} // PostScannerAddRequest | 
};
apiInstance.postScannerAdd(opts, (error, data, response) => {
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
 **postScannerAddRequest** | [**PostScannerAddRequest**](PostScannerAddRequest.md)|  | [optional] 

### Return type

[**PostScannerAdd200Response**](PostScannerAdd200Response.md)

### Authorization

[custom_authorizer](../README.md#custom_authorizer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## postScannerRemove

> PostScannerAdd200Response postScannerRemove(opts)

スキャン対象のドメイン一覧の削除

### Example

```javascript
import PhishingHunterApi from 'phishing_hunter_api';
let defaultClient = PhishingHunterApi.ApiClient.instance;
// Configure API key authorization: custom_authorizer
let custom_authorizer = defaultClient.authentications['custom_authorizer'];
custom_authorizer.apiKey = 'YOUR API KEY';
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//custom_authorizer.apiKeyPrefix = 'Token';

let apiInstance = new PhishingHunterApi.ScannerApi();
let opts = {
  'postScannerAddRequest': {"keywords":["paypal","icloud","gmail"],"target_score":100,"webhook_url":"https://xxxx"} // PostScannerAddRequest | 
};
apiInstance.postScannerRemove(opts, (error, data, response) => {
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
 **postScannerAddRequest** | [**PostScannerAddRequest**](PostScannerAddRequest.md)|  | [optional] 

### Return type

[**PostScannerAdd200Response**](PostScannerAdd200Response.md)

### Authorization

[custom_authorizer](../README.md#custom_authorizer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

