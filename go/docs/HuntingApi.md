# \HuntingApi

All URIs are relative to *https://api.phishing-hunter.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetHunting**](HuntingApi.md#GetHunting) | **Get** /hunting | ハンティング設定の取得
[**PostHunting**](HuntingApi.md#PostHunting) | **Post** /hunting | ハンティング設定の更新



## GetHunting

> GetHunting200Response GetHunting(ctx).Execute()

ハンティング設定の取得

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
    resp, r, err := apiClient.HuntingApi.GetHunting(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HuntingApi.GetHunting``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetHunting`: GetHunting200Response
    fmt.Fprintf(os.Stdout, "Response from `HuntingApi.GetHunting`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetHuntingRequest struct via the builder pattern


### Return type

[**GetHunting200Response**](GetHunting200Response.md)

### Authorization

[custom_authorizer](../README.md#custom_authorizer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostHunting

> PostScannerAdd200Response PostHunting(ctx).GetHunting200Response(getHunting200Response).Execute()

ハンティング設定の更新

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
    getHunting200Response := *openapiclient.NewGetHunting200Response() // GetHunting200Response |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HuntingApi.PostHunting(context.Background()).GetHunting200Response(getHunting200Response).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HuntingApi.PostHunting``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostHunting`: PostScannerAdd200Response
    fmt.Fprintf(os.Stdout, "Response from `HuntingApi.PostHunting`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostHuntingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getHunting200Response** | [**GetHunting200Response**](GetHunting200Response.md) |  | 

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

