# \DefaultApi

All URIs are relative to *http://localhost:8080/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CoursesEditPost**](DefaultApi.md#CoursesEditPost) | **Post** /courses/edit | Create Course



## CoursesEditPost

> Course CoursesEditPost(ctx).Actor(actor).Course(course).Execute()

Create Course



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
    course := *openapiclient.NewCourse("Name_example", "Discipline_example", "Landing_example", "Visibility_example") // Course | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.CoursesEditPost(context.Background()).Actor(actor).Course(course).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CoursesEditPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CoursesEditPost`: Course
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CoursesEditPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCoursesEditPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **actor** | [**Actor**](Actor.md) |  | 
 **course** | [**Course**](Course.md) |  | 

### Return type

[**Course**](Course.md)

### Authorization

[actorAuth](../API_README.md#actorAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../API_README.md#documentation-for-api-endpoints)
[[Back to Model list]](../API_README.md#documentation-for-models)
[[Back to README]](../API_README.md)

