# \ModerationApi

All URIs are relative to *http://localhost:8080/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PublicationRequestsPost**](ModerationApi.md#PublicationRequestsPost) | **Post** /publication-requests | Request Course Publication
[**PublicationRequestsRequestIDPut**](ModerationApi.md#PublicationRequestsRequestIDPut) | **Put** /publication-requests/{requestID} | Update Publication Request



## PublicationRequestsPost

> PublicationRequestsPost(ctx).Actor(actor).PublicationRequest(publicationRequest).Execute()

Request Course Publication



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
    actor := map[string][]openapiclient.Actor{ ... } // Actor | 
    publicationRequest := *openapiclient.NewPublicationRequest() // PublicationRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ModerationApi.PublicationRequestsPost(context.Background()).Actor(actor).PublicationRequest(publicationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ModerationApi.PublicationRequestsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPublicationRequestsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **actor** | [**Actor**](Actor.md) |  | 
 **publicationRequest** | [**PublicationRequest**](PublicationRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../API_README.md#documentation-for-api-endpoints)
[[Back to Model list]](../API_README.md#documentation-for-models)
[[Back to README]](../API_README.md)


## PublicationRequestsRequestIDPut

> PublicationRequest PublicationRequestsRequestIDPut(ctx, requestID).Actor(actor).Execute()

Update Publication Request



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
    requestID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ID of the publication request to update.
    actor := map[string][]openapiclient.Actor{ ... } // Actor | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ModerationApi.PublicationRequestsRequestIDPut(context.Background(), requestID).Actor(actor).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ModerationApi.PublicationRequestsRequestIDPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PublicationRequestsRequestIDPut`: PublicationRequest
    fmt.Fprintf(os.Stdout, "Response from `ModerationApi.PublicationRequestsRequestIDPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**requestID** | **string** | ID of the publication request to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPublicationRequestsRequestIDPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **actor** | [**Actor**](Actor.md) |  | 

### Return type

[**PublicationRequest**](PublicationRequest.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../API_README.md#documentation-for-api-endpoints)
[[Back to Model list]](../API_README.md#documentation-for-models)
[[Back to README]](../API_README.md)

