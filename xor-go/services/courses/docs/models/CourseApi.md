# \CourseApi

All URIs are relative to *http://localhost:8080/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CoursesCourseIDBuyPost**](CourseApi.md#CoursesCourseIDBuyPost) | **Post** /courses/{courseID}/buy | Buy Course
[**CoursesCourseIDGet**](CourseApi.md#CoursesCourseIDGet) | **Get** /courses/{courseID} | Read published Course
[**CoursesEditCourseIDGet**](CourseApi.md#CoursesEditCourseIDGet) | **Get** /courses/edit/{courseID} | Get Course by ID
[**CoursesEditCourseIDPut**](CourseApi.md#CoursesEditCourseIDPut) | **Put** /courses/edit/{courseID} | Update Course
[**CoursesEditGet**](CourseApi.md#CoursesEditGet) | **Get** /courses/edit | Get Courses



## CoursesCourseIDBuyPost

> PaymentRedirect CoursesCourseIDBuyPost(ctx, courseID).Actor(actor).Execute()

Buy Course



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
    courseID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ID of the course to buy.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CourseApi.CoursesCourseIDBuyPost(context.Background(), courseID).Actor(actor).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CourseApi.CoursesCourseIDBuyPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CoursesCourseIDBuyPost`: PaymentRedirect
    fmt.Fprintf(os.Stdout, "Response from `CourseApi.CoursesCourseIDBuyPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**courseID** | **string** | ID of the course to buy. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCoursesCourseIDBuyPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **actor** | [**Actor**](Actor.md) |  | 


### Return type

[**PaymentRedirect**](PaymentRedirect.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../API_README.md#documentation-for-api-endpoints)
[[Back to Model list]](../API_README.md#documentation-for-models)
[[Back to README]](../API_README.md)


## CoursesCourseIDGet

> Course CoursesCourseIDGet(ctx, courseID).Actor(actor).Execute()

Read published Course



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
    courseID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ID of the course to retrieve
    actor := map[string][]openapiclient.OptionalActor{ ... } // OptionalActor |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CourseApi.CoursesCourseIDGet(context.Background(), courseID).Actor(actor).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CourseApi.CoursesCourseIDGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CoursesCourseIDGet`: Course
    fmt.Fprintf(os.Stdout, "Response from `CourseApi.CoursesCourseIDGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**courseID** | **string** | ID of the course to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiCoursesCourseIDGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **actor** | [**OptionalActor**](OptionalActor.md) |  | 

### Return type

[**Course**](Course.md)

### Authorization

[actorAuth](../API_README.md#actorAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../API_README.md#documentation-for-api-endpoints)
[[Back to Model list]](../API_README.md#documentation-for-models)
[[Back to README]](../API_README.md)


## CoursesEditCourseIDGet

> Course CoursesEditCourseIDGet(ctx, courseID).Actor(actor).Execute()

Get Course by ID



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
    courseID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ID of the course to retrieve
    actor := map[string][]openapiclient.Actor{ ... } // Actor | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CourseApi.CoursesEditCourseIDGet(context.Background(), courseID).Actor(actor).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CourseApi.CoursesEditCourseIDGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CoursesEditCourseIDGet`: Course
    fmt.Fprintf(os.Stdout, "Response from `CourseApi.CoursesEditCourseIDGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**courseID** | **string** | ID of the course to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiCoursesEditCourseIDGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **actor** | [**Actor**](Actor.md) |  | 

### Return type

[**Course**](Course.md)

### Authorization

[actorAuth](../API_README.md#actorAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../API_README.md#documentation-for-api-endpoints)
[[Back to Model list]](../API_README.md#documentation-for-models)
[[Back to README]](../API_README.md)


## CoursesEditCourseIDPut

> Course CoursesEditCourseIDPut(ctx, courseID).Actor(actor).Course(course).Execute()

Update Course



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
    courseID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ID of the course to update
    actor := map[string][]openapiclient.Actor{ ... } // Actor | 
    course := *openapiclient.NewCourse("Name_example", "Discipline_example", "Landing_example", "Visibility_example") // Course | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CourseApi.CoursesEditCourseIDPut(context.Background(), courseID).Actor(actor).Course(course).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CourseApi.CoursesEditCourseIDPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CoursesEditCourseIDPut`: Course
    fmt.Fprintf(os.Stdout, "Response from `CourseApi.CoursesEditCourseIDPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**courseID** | **string** | ID of the course to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiCoursesEditCourseIDPutRequest struct via the builder pattern


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


## CoursesEditGet

> []Course CoursesEditGet(ctx).Actor(actor).Execute()

Get Courses



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CourseApi.CoursesEditGet(context.Background()).Actor(actor).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CourseApi.CoursesEditGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CoursesEditGet`: []Course
    fmt.Fprintf(os.Stdout, "Response from `CourseApi.CoursesEditGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCoursesEditGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **actor** | [**Actor**](Actor.md) |  | 

### Return type

[**[]Course**](Course.md)

### Authorization

[actorAuth](../API_README.md#actorAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../API_README.md#documentation-for-api-endpoints)
[[Back to Model list]](../API_README.md#documentation-for-models)
[[Back to README]](../API_README.md)

