# \LessonApi

All URIs are relative to *http://localhost:8080/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CoursesEditCourseIDDelete**](LessonApi.md#CoursesEditCourseIDDelete) | **Delete** /courses/edit/{courseID} | Delete Course
[**LessonsEditLessonIDDelete**](LessonApi.md#LessonsEditLessonIDDelete) | **Delete** /lessons/edit/{lessonID} | Delete Lesson
[**LessonsEditLessonIDGet**](LessonApi.md#LessonsEditLessonIDGet) | **Get** /lessons/edit/{lessonID} | Retrieve Lesson
[**LessonsEditLessonIDPut**](LessonApi.md#LessonsEditLessonIDPut) | **Put** /lessons/edit/{lessonID} | Update Lesson
[**LessonsEditPost**](LessonApi.md#LessonsEditPost) | **Post** /lessons/edit/ | Create Lesson
[**LessonsLessonIDBuyPost**](LessonApi.md#LessonsLessonIDBuyPost) | **Post** /lessons/{lessonID}/buy | Buy Lesson
[**LessonsLessonIDGet**](LessonApi.md#LessonsLessonIDGet) | **Get** /lessons/{lessonID} | Retrieve published Lesson



## CoursesEditCourseIDDelete

> CoursesEditCourseIDDelete(ctx, courseID).Actor(actor).Execute()

Delete Course



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
    courseID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ID of the course to delete

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LessonApi.CoursesEditCourseIDDelete(context.Background(), courseID).Actor(actor).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LessonApi.CoursesEditCourseIDDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**courseID** | **string** | ID of the course to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiCoursesEditCourseIDDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **actor** | [**Actor**](Actor.md) |  | 


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../API_README.md#documentation-for-api-endpoints)
[[Back to Model list]](../API_README.md#documentation-for-models)
[[Back to README]](../API_README.md)


## LessonsEditLessonIDDelete

> LessonsEditLessonIDDelete(ctx, lessonID).Actor(actor).Execute()

Delete Lesson



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
    lessonID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ID of the lesson to delete.
    actor := map[string][]openapiclient.Actor{ ... } // Actor | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LessonApi.LessonsEditLessonIDDelete(context.Background(), lessonID).Actor(actor).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LessonApi.LessonsEditLessonIDDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lessonID** | **string** | ID of the lesson to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiLessonsEditLessonIDDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **actor** | [**Actor**](Actor.md) |  | 

### Return type

 (empty response body)

### Authorization

[actorAuth](../API_README.md#actorAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../API_README.md#documentation-for-api-endpoints)
[[Back to Model list]](../API_README.md#documentation-for-models)
[[Back to README]](../API_README.md)


## LessonsEditLessonIDGet

> Lesson LessonsEditLessonIDGet(ctx, lessonID).Actor(actor).Execute()

Retrieve Lesson



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
    lessonID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ID of the lesson to retrieve.
    actor := map[string][]openapiclient.Actor{ ... } // Actor | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LessonApi.LessonsEditLessonIDGet(context.Background(), lessonID).Actor(actor).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LessonApi.LessonsEditLessonIDGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LessonsEditLessonIDGet`: Lesson
    fmt.Fprintf(os.Stdout, "Response from `LessonApi.LessonsEditLessonIDGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lessonID** | **string** | ID of the lesson to retrieve. | 

### Other Parameters

Other parameters are passed through a pointer to a apiLessonsEditLessonIDGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **actor** | [**Actor**](Actor.md) |  | 

### Return type

[**Lesson**](Lesson.md)

### Authorization

[actorAuth](../API_README.md#actorAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../API_README.md#documentation-for-api-endpoints)
[[Back to Model list]](../API_README.md#documentation-for-models)
[[Back to README]](../API_README.md)


## LessonsEditLessonIDPut

> Lesson LessonsEditLessonIDPut(ctx, lessonID).Actor(actor).Lesson(lesson).Execute()

Update Lesson



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
    lessonID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ID of the lesson to update.
    actor := map[string][]openapiclient.Actor{ ... } // Actor | 
    lesson := *openapiclient.NewLesson("Visibility_example", "Transcript_example") // Lesson | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LessonApi.LessonsEditLessonIDPut(context.Background(), lessonID).Actor(actor).Lesson(lesson).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LessonApi.LessonsEditLessonIDPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LessonsEditLessonIDPut`: Lesson
    fmt.Fprintf(os.Stdout, "Response from `LessonApi.LessonsEditLessonIDPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lessonID** | **string** | ID of the lesson to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiLessonsEditLessonIDPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **actor** | [**Actor**](Actor.md) |  | 
 **lesson** | [**Lesson**](Lesson.md) |  | 

### Return type

[**Lesson**](Lesson.md)

### Authorization

[actorAuth](../API_README.md#actorAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../API_README.md#documentation-for-api-endpoints)
[[Back to Model list]](../API_README.md#documentation-for-models)
[[Back to README]](../API_README.md)


## LessonsEditPost

> Lesson LessonsEditPost(ctx).Actor(actor).Lesson(lesson).Execute()

Create Lesson



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
    lesson := *openapiclient.NewLesson("Visibility_example", "Transcript_example") // Lesson | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LessonApi.LessonsEditPost(context.Background()).Actor(actor).Lesson(lesson).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LessonApi.LessonsEditPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LessonsEditPost`: Lesson
    fmt.Fprintf(os.Stdout, "Response from `LessonApi.LessonsEditPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLessonsEditPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **actor** | [**Actor**](Actor.md) |  | 
 **lesson** | [**Lesson**](Lesson.md) |  | 

### Return type

[**Lesson**](Lesson.md)

### Authorization

[actorAuth](../API_README.md#actorAuth)

### HTTP request headers

- **Content-Type**: application/json
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
    resp, r, err := apiClient.LessonApi.LessonsLessonIDBuyPost(context.Background(), lessonID).Actor(actor).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LessonApi.LessonsLessonIDBuyPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LessonsLessonIDBuyPost`: LessonsLessonIDBuyPost200Response
    fmt.Fprintf(os.Stdout, "Response from `LessonApi.LessonsLessonIDBuyPost`: %v\n", resp)
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


## LessonsLessonIDGet

> Lesson LessonsLessonIDGet(ctx, lessonID).Actor(actor).Execute()

Retrieve published Lesson



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
    lessonID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ID of the lesson to retrieve.
    actor := map[string][]openapiclient.OptionalActor{ ... } // OptionalActor |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LessonApi.LessonsLessonIDGet(context.Background(), lessonID).Actor(actor).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LessonApi.LessonsLessonIDGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LessonsLessonIDGet`: Lesson
    fmt.Fprintf(os.Stdout, "Response from `LessonApi.LessonsLessonIDGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lessonID** | **string** | ID of the lesson to retrieve. | 

### Other Parameters

Other parameters are passed through a pointer to a apiLessonsLessonIDGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **actor** | [**OptionalActor**](OptionalActor.md) |  | 

### Return type

[**Lesson**](Lesson.md)

### Authorization

[actorAuth](../API_README.md#actorAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../API_README.md#documentation-for-api-endpoints)
[[Back to Model list]](../API_README.md#documentation-for-models)
[[Back to README]](../API_README.md)

