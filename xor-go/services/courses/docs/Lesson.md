# Lesson

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ID** | Pointer to **string** |  | [optional] 
**CourseID** | Pointer to **string** |  | [optional] 
**TeacherID** | Pointer to **string** |  | [optional] 
**Product** | Pointer to [**Product**](Product.md) |  | [optional] 
**Visibility** | **string** |  | 
**Transcript** | **string** |  | 
**VideoURI** | Pointer to **string** |  | [optional] 

## Methods

### NewLesson

`func NewLesson(visibility string, transcript string, ) *Lesson`

NewLesson instantiates a new Lesson object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLessonWithDefaults

`func NewLessonWithDefaults() *Lesson`

NewLessonWithDefaults instantiates a new Lesson object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetID

`func (o *Lesson) GetID() string`

GetID returns the ID field if non-nil, zero value otherwise.

### GetIDOk

`func (o *Lesson) GetIDOk() (*string, bool)`

GetIDOk returns a tuple with the ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetID

`func (o *Lesson) SetID(v string)`

SetID sets ID field to given value.

### HasID

`func (o *Lesson) HasID() bool`

HasID returns a boolean if a field has been set.

### GetCourseID

`func (o *Lesson) GetCourseID() string`

GetCourseID returns the CourseID field if non-nil, zero value otherwise.

### GetCourseIDOk

`func (o *Lesson) GetCourseIDOk() (*string, bool)`

GetCourseIDOk returns a tuple with the CourseID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCourseID

`func (o *Lesson) SetCourseID(v string)`

SetCourseID sets CourseID field to given value.

### HasCourseID

`func (o *Lesson) HasCourseID() bool`

HasCourseID returns a boolean if a field has been set.

### GetTeacherID

`func (o *Lesson) GetTeacherID() string`

GetTeacherID returns the TeacherID field if non-nil, zero value otherwise.

### GetTeacherIDOk

`func (o *Lesson) GetTeacherIDOk() (*string, bool)`

GetTeacherIDOk returns a tuple with the TeacherID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeacherID

`func (o *Lesson) SetTeacherID(v string)`

SetTeacherID sets TeacherID field to given value.

### HasTeacherID

`func (o *Lesson) HasTeacherID() bool`

HasTeacherID returns a boolean if a field has been set.

### GetProduct

`func (o *Lesson) GetProduct() Product`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *Lesson) GetProductOk() (*Product, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *Lesson) SetProduct(v Product)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *Lesson) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### GetVisibility

`func (o *Lesson) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *Lesson) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *Lesson) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.


### GetTranscript

`func (o *Lesson) GetTranscript() string`

GetTranscript returns the Transcript field if non-nil, zero value otherwise.

### GetTranscriptOk

`func (o *Lesson) GetTranscriptOk() (*string, bool)`

GetTranscriptOk returns a tuple with the Transcript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranscript

`func (o *Lesson) SetTranscript(v string)`

SetTranscript sets Transcript field to given value.


### GetVideoURI

`func (o *Lesson) GetVideoURI() string`

GetVideoURI returns the VideoURI field if non-nil, zero value otherwise.

### GetVideoURIOk

`func (o *Lesson) GetVideoURIOk() (*string, bool)`

GetVideoURIOk returns a tuple with the VideoURI field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoURI

`func (o *Lesson) SetVideoURI(v string)`

SetVideoURI sets VideoURI field to given value.

### HasVideoURI

`func (o *Lesson) HasVideoURI() bool`

HasVideoURI returns a boolean if a field has been set.


[[Back to Model list]](../API_README.md#documentation-for-models) [[Back to API list]](../API_README.md#documentation-for-api-endpoints) [[Back to README]](../API_README.md)


