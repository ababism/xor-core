# \UserApi

All URIs are relative to *http://localhost:8080/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CoursesCourseIDBuyPost**](UserApi.md#CoursesCourseIDBuyPost) | **Post** /courses/{courseID}/buy | Buy Course
[**LessonsLessonIDBuyPost**](UserApi.md#LessonsLessonIDBuyPost) | **Post** /lessons/{lessonID}/buy | Buy Lesson
[**StudentsRegisterPost**](UserApi.md#StudentsRegisterPost) | **Post** /students/register | Register Student Profile
[**TeachersRegisterPost**](UserApi.md#TeachersRegisterPost) | **Post** /teachers/register | Register Teacher Profile
[**UserAccessConfirmBuyerIDPost**](UserApi.md#UserAccessConfirmBuyerIDPost) | **Post** /user/access/confirm/{buyerID} | Confirm user access to products
[**UserAccessLessonsLessonIDGet**](UserApi.md#UserAccessLessonsLessonIDGet) | **Get** /user/access/lessons/{lessonID} | Get actor&#39;s lesson access
[**UserAccessLessonsPut**](UserApi.md#UserAccessLessonsPut) | **Put** /user/access/lessons | Change student&#39;s access to lesson



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
    resp, r, err := apiClient.UserApi.CoursesCourseIDBuyPost(context.Background(), courseID).Actor(actor).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.CoursesCourseIDBuyPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CoursesCourseIDBuyPost`: PaymentRedirect
    fmt.Fprintf(os.Stdout, "Response from `UserApi.CoursesCourseIDBuyPost`: %v\n", resp)
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


## LessonsLessonIDBuyPost

> LessonsLessonIDBuyPost200Response LessonsLessonIDBuyPost(ctx, lessonID).Actor(actor).Execute()

Buy Lesson



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
    lessonID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ID of the lesson.
    actor := map[string][]openapiclient.Actor{ ... } // Actor | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserApi.LessonsLessonIDBuyPost(context.Background(), lessonID).Actor(actor).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.LessonsLessonIDBuyPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LessonsLessonIDBuyPost`: LessonsLessonIDBuyPost200Response
    fmt.Fprintf(os.Stdout, "Response from `UserApi.LessonsLessonIDBuyPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lessonID** | **string** | ID of the lesson. | 

### Other Parameters

Other parameters are passed through a pointer to a apiLessonsLessonIDBuyPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **actor** | [**Actor**](Actor.md) |  | 

### Return type

[**LessonsLessonIDBuyPost200Response**](LessonsLessonIDBuyPost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../API_README.md#documentation-for-api-endpoints)
[[Back to Model list]](../API_README.md#documentation-for-models)
[[Back to README]](../API_README.md)


## StudentsRegisterPost

> StudentsRegisterPost(ctx).Student(student).Actor(actor).Execute()

Register Student Profile



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
    student := *openapiclient.NewStudent() // Student | 
    actor := map[string][]openapiclient.OptionalActor{ ... } // OptionalActor |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserApi.StudentsRegisterPost(context.Background()).Student(student).Actor(actor).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.StudentsRegisterPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStudentsRegisterPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **student** | [**Student**](Student.md) |  | 
 **actor** | [**OptionalActor**](OptionalActor.md) |  | 

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


## TeachersRegisterPost

> TeachersRegisterPost(ctx).Teacher(teacher).Actor(actor).Execute()

Register Teacher Profile



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
    teacher := *openapiclient.NewTeacher() // Teacher | 
    actor := map[string][]openapiclient.OptionalActor{ ... } // OptionalActor |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserApi.TeachersRegisterPost(context.Background()).Teacher(teacher).Actor(actor).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.TeachersRegisterPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTeachersRegisterPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **teacher** | [**Teacher**](Teacher.md) |  | 
 **actor** | [**OptionalActor**](OptionalActor.md) |  | 

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


## UserAccessConfirmBuyerIDPost

> UserAccessConfirmBuyerIDPost(ctx, buyerID).UserAccessConfirmBuyerIDPostRequest(userAccessConfirmBuyerIDPostRequest).Execute()

Confirm user access to products



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
    buyerID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ID of the confirmed owner.
    userAccessConfirmBuyerIDPostRequest := *openapiclient.NewUserAccessConfirmBuyerIDPostRequest() // UserAccessConfirmBuyerIDPostRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserApi.UserAccessConfirmBuyerIDPost(context.Background(), buyerID).UserAccessConfirmBuyerIDPostRequest(userAccessConfirmBuyerIDPostRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.UserAccessConfirmBuyerIDPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**buyerID** | **string** | ID of the confirmed owner. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserAccessConfirmBuyerIDPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userAccessConfirmBuyerIDPostRequest** | [**UserAccessConfirmBuyerIDPostRequest**](UserAccessConfirmBuyerIDPostRequest.md) |  | 

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


## UserAccessLessonsLessonIDGet

> LessonAccess UserAccessLessonsLessonIDGet(ctx, lessonID).Actor(actor).Execute()

Get actor's lesson access



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
    lessonID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ID of the lesson to get access to.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserApi.UserAccessLessonsLessonIDGet(context.Background(), lessonID).Actor(actor).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.UserAccessLessonsLessonIDGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UserAccessLessonsLessonIDGet`: LessonAccess
    fmt.Fprintf(os.Stdout, "Response from `UserApi.UserAccessLessonsLessonIDGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lessonID** | **string** | ID of the lesson to get access to. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserAccessLessonsLessonIDGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **actor** | [**Actor**](Actor.md) |  | 


### Return type

[**LessonAccess**](LessonAccess.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../API_README.md#documentation-for-api-endpoints)
[[Back to Model list]](../API_README.md#documentation-for-models)
[[Back to README]](../API_README.md)


## UserAccessLessonsPut

> LessonAccess UserAccessLessonsPut(ctx).Actor(actor).LessonAccess(lessonAccess).Execute()

Change student's access to lesson



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
    lessonAccess := *openapiclient.NewLessonAccess() // LessonAccess | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserApi.UserAccessLessonsPut(context.Background()).Actor(actor).LessonAccess(lessonAccess).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.UserAccessLessonsPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UserAccessLessonsPut`: LessonAccess
    fmt.Fprintf(os.Stdout, "Response from `UserApi.UserAccessLessonsPut`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUserAccessLessonsPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **actor** | [**Actor**](Actor.md) |  | 
 **lessonAccess** | [**LessonAccess**](LessonAccess.md) |  | 

### Return type

[**LessonAccess**](LessonAccess.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../API_README.md#documentation-for-api-endpoints)
[[Back to Model list]](../API_README.md#documentation-for-models)
[[Back to README]](../API_README.md)

