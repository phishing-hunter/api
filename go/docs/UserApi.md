# \UserApi

All URIs are relative to *https://api.phishing-hunter.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetApikey**](UserApi.md#GetApikey) | **Get** /apikey | APIキーの一覧を取得
[**GetUsersInfo**](UserApi.md#GetUsersInfo) | **Get** /users/info | ユーザ情報の取得
[**PostApikeyCreate**](UserApi.md#PostApikeyCreate) | **Post** /apikey/create | APIキーの作成
[**PostApikeyDelete**](UserApi.md#PostApikeyDelete) | **Post** /apikey/delete | APIキーの削除



## GetApikey

> GetApikey200Response GetApikey(ctx).Execute()

APIキーの一覧を取得

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
    resp, r, err := apiClient.UserApi.GetApikey(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.GetApikey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApikey`: GetApikey200Response
    fmt.Fprintf(os.Stdout, "Response from `UserApi.GetApikey`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetApikeyRequest struct via the builder pattern


### Return type

[**GetApikey200Response**](GetApikey200Response.md)

### Authorization

[custom_authorizer](../README.md#custom_authorizer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsersInfo

> GetUsersInfo200Response GetUsersInfo(ctx).Execute()

ユーザ情報の取得



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
    resp, r, err := apiClient.UserApi.GetUsersInfo(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.GetUsersInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUsersInfo`: GetUsersInfo200Response
    fmt.Fprintf(os.Stdout, "Response from `UserApi.GetUsersInfo`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetUsersInfoRequest struct via the builder pattern


### Return type

[**GetUsersInfo200Response**](GetUsersInfo200Response.md)

### Authorization

[custom_authorizer](../README.md#custom_authorizer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostApikeyCreate

> PostApikeyCreate200Response PostApikeyCreate(ctx).PostApikeyCreateRequest(postApikeyCreateRequest).Execute()

APIキーの作成

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
    postApikeyCreateRequest := *openapiclient.NewPostApikeyCreateRequest() // PostApikeyCreateRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserApi.PostApikeyCreate(context.Background()).PostApikeyCreateRequest(postApikeyCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.PostApikeyCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostApikeyCreate`: PostApikeyCreate200Response
    fmt.Fprintf(os.Stdout, "Response from `UserApi.PostApikeyCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostApikeyCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postApikeyCreateRequest** | [**PostApikeyCreateRequest**](PostApikeyCreateRequest.md) |  | 

### Return type

[**PostApikeyCreate200Response**](PostApikeyCreate200Response.md)

### Authorization

[custom_authorizer](../README.md#custom_authorizer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostApikeyDelete

> PostApikeyDelete200Response PostApikeyDelete(ctx).PostApikeyCreate200Response(postApikeyCreate200Response).Execute()

APIキーの削除

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
    postApikeyCreate200Response := *openapiclient.NewPostApikeyCreate200Response() // PostApikeyCreate200Response |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserApi.PostApikeyDelete(context.Background()).PostApikeyCreate200Response(postApikeyCreate200Response).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.PostApikeyDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostApikeyDelete`: PostApikeyDelete200Response
    fmt.Fprintf(os.Stdout, "Response from `UserApi.PostApikeyDelete`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostApikeyDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postApikeyCreate200Response** | [**PostApikeyCreate200Response**](PostApikeyCreate200Response.md) |  | 

### Return type

[**PostApikeyDelete200Response**](PostApikeyDelete200Response.md)

### Authorization

[custom_authorizer](../README.md#custom_authorizer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

