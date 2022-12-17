# \ScannerApi

All URIs are relative to *https://api.phishing-hunter.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetScanner**](ScannerApi.md#GetScanner) | **Get** /scanner | スキャン対象のドメイン一覧の取得
[**PostScannerAdd**](ScannerApi.md#PostScannerAdd) | **Post** /scanner/add | スキャン対象のドメインの追加
[**PostScannerRemove**](ScannerApi.md#PostScannerRemove) | **Post** /scanner/remove | スキャン対象のドメイン一覧の削除



## GetScanner

> GetScanner200Response GetScanner(ctx).Execute()

スキャン対象のドメイン一覧の取得

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
    resp, r, err := apiClient.ScannerApi.GetScanner(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScannerApi.GetScanner``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetScanner`: GetScanner200Response
    fmt.Fprintf(os.Stdout, "Response from `ScannerApi.GetScanner`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetScannerRequest struct via the builder pattern


### Return type

[**GetScanner200Response**](GetScanner200Response.md)

### Authorization

[custom_authorizer](../README.md#custom_authorizer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostScannerAdd

> PostScannerAdd200Response PostScannerAdd(ctx).PostScannerAddRequest(postScannerAddRequest).Execute()

スキャン対象のドメインの追加

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
    postScannerAddRequest := *openapiclient.NewPostScannerAddRequest() // PostScannerAddRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ScannerApi.PostScannerAdd(context.Background()).PostScannerAddRequest(postScannerAddRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScannerApi.PostScannerAdd``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostScannerAdd`: PostScannerAdd200Response
    fmt.Fprintf(os.Stdout, "Response from `ScannerApi.PostScannerAdd`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostScannerAddRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postScannerAddRequest** | [**PostScannerAddRequest**](PostScannerAddRequest.md) |  | 

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


## PostScannerRemove

> PostScannerAdd200Response PostScannerRemove(ctx).PostScannerAddRequest(postScannerAddRequest).Execute()

スキャン対象のドメイン一覧の削除

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
    postScannerAddRequest := *openapiclient.NewPostScannerAddRequest() // PostScannerAddRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ScannerApi.PostScannerRemove(context.Background()).PostScannerAddRequest(postScannerAddRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScannerApi.PostScannerRemove``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostScannerRemove`: PostScannerAdd200Response
    fmt.Fprintf(os.Stdout, "Response from `ScannerApi.PostScannerRemove`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostScannerRemoveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postScannerAddRequest** | [**PostScannerAddRequest**](PostScannerAddRequest.md) |  | 

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

