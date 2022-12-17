# \SettingsApi

All URIs are relative to *https://api.phishing-hunter.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetApikeys**](SettingsApi.md#GetApikeys) | **Get** /apikeys | 外部サービス設定の取得
[**GetNotifySlack**](SettingsApi.md#GetNotifySlack) | **Get** /notify | 通知設定の一覧を取得
[**PostApikeys**](SettingsApi.md#PostApikeys) | **Post** /apikeys | 外部サービス設定の更新
[**PostNotifySlack**](SettingsApi.md#PostNotifySlack) | **Post** /notify | 通知設定の更新



## GetApikeys

> GetApikeys200Response GetApikeys(ctx).Execute()

外部サービス設定の取得

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SettingsApi.GetApikeys(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsApi.GetApikeys``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApikeys`: GetApikeys200Response
    fmt.Fprintf(os.Stdout, "Response from `SettingsApi.GetApikeys`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetApikeysRequest struct via the builder pattern


### Return type

[**GetApikeys200Response**](GetApikeys200Response.md)

### Authorization

[custom_authorizer](../README.md#custom_authorizer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNotifySlack

> GetNotifySlack200Response GetNotifySlack(ctx).Execute()

通知設定の一覧を取得

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SettingsApi.GetNotifySlack(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsApi.GetNotifySlack``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNotifySlack`: GetNotifySlack200Response
    fmt.Fprintf(os.Stdout, "Response from `SettingsApi.GetNotifySlack`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetNotifySlackRequest struct via the builder pattern


### Return type

[**GetNotifySlack200Response**](GetNotifySlack200Response.md)

### Authorization

[custom_authorizer](../README.md#custom_authorizer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostApikeys

> PostScannerAdd200Response PostApikeys(ctx).GetApikeys200Response(getApikeys200Response).Execute()

外部サービス設定の更新



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    getApikeys200Response := *openapiclient.NewGetApikeys200Response() // GetApikeys200Response |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SettingsApi.PostApikeys(context.Background()).GetApikeys200Response(getApikeys200Response).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsApi.PostApikeys``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostApikeys`: PostScannerAdd200Response
    fmt.Fprintf(os.Stdout, "Response from `SettingsApi.PostApikeys`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostApikeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getApikeys200Response** | [**GetApikeys200Response**](GetApikeys200Response.md) |  | 

### Return type

[**PostScannerAdd200Response**](PostScannerAdd200Response.md)

### Authorization

[custom_authorizer](../README.md#custom_authorizer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostNotifySlack

> PostScannerAdd200Response PostNotifySlack(ctx).PostNotifySlackRequest(postNotifySlackRequest).Execute()

通知設定の更新

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    postNotifySlackRequest := *openapiclient.NewPostNotifySlackRequest() // PostNotifySlackRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SettingsApi.PostNotifySlack(context.Background()).PostNotifySlackRequest(postNotifySlackRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsApi.PostNotifySlack``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostNotifySlack`: PostScannerAdd200Response
    fmt.Fprintf(os.Stdout, "Response from `SettingsApi.PostNotifySlack`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostNotifySlackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postNotifySlackRequest** | [**PostNotifySlackRequest**](PostNotifySlackRequest.md) |  | 

### Return type

[**PostScannerAdd200Response**](PostScannerAdd200Response.md)

### Authorization

[custom_authorizer](../README.md#custom_authorizer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

