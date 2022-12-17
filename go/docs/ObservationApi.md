# \ObservationApi

All URIs are relative to *https://api.phishing-hunter.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetObservations**](ObservationApi.md#GetObservations) | **Get** /observation | 監視対象ドメイン一覧の取得
[**PostObservationAdd**](ObservationApi.md#PostObservationAdd) | **Post** /observation/add | 監視対象ドメインの追加
[**PostObservationRemove**](ObservationApi.md#PostObservationRemove) | **Post** /observation/remove | 監視対象ドメインの削除



## GetObservations

> GetObservations200Response GetObservations(ctx).Execute()

監視対象ドメイン一覧の取得

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
    resp, r, err := apiClient.ObservationApi.GetObservations(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObservationApi.GetObservations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetObservations`: GetObservations200Response
    fmt.Fprintf(os.Stdout, "Response from `ObservationApi.GetObservations`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetObservationsRequest struct via the builder pattern


### Return type

[**GetObservations200Response**](GetObservations200Response.md)

### Authorization

[custom_authorizer](../README.md#custom_authorizer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostObservationAdd

> GetScanner500Response PostObservationAdd(ctx).PostObservationAddRequest(postObservationAddRequest).Execute()

監視対象ドメインの追加

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
    postObservationAddRequest := *openapiclient.NewPostObservationAddRequest() // PostObservationAddRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ObservationApi.PostObservationAdd(context.Background()).PostObservationAddRequest(postObservationAddRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObservationApi.PostObservationAdd``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostObservationAdd`: GetScanner500Response
    fmt.Fprintf(os.Stdout, "Response from `ObservationApi.PostObservationAdd`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostObservationAddRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postObservationAddRequest** | [**PostObservationAddRequest**](PostObservationAddRequest.md) |  | 

### Return type

[**GetScanner500Response**](GetScanner500Response.md)

### Authorization

[custom_authorizer](../README.md#custom_authorizer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostObservationRemove

> GetScanner500Response PostObservationRemove(ctx).PostObservationRemoveRequest(postObservationRemoveRequest).Execute()

監視対象ドメインの削除



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
    postObservationRemoveRequest := *openapiclient.NewPostObservationRemoveRequest() // PostObservationRemoveRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ObservationApi.PostObservationRemove(context.Background()).PostObservationRemoveRequest(postObservationRemoveRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObservationApi.PostObservationRemove``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostObservationRemove`: GetScanner500Response
    fmt.Fprintf(os.Stdout, "Response from `ObservationApi.PostObservationRemove`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostObservationRemoveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postObservationRemoveRequest** | [**PostObservationRemoveRequest**](PostObservationRemoveRequest.md) |  | 

### Return type

[**GetScanner500Response**](GetScanner500Response.md)

### Authorization

[custom_authorizer](../README.md#custom_authorizer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

