# LessonAccess

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ID** | Pointer to **string** | Unique identifier for the lesson access. | [optional] 
**LessonID** | Pointer to **string** | ID of the lesson. | [optional] 
**StudentID** | Pointer to **string** | ID of the student. | [optional] 
**AccessStatus** | Pointer to **string** | Access status of the lesson. | [optional] 
**UpdatedAt** | Pointer to **time.Time** | Date and time when the request was last updated. | [optional] 

## Methods

### NewLessonAccess

`func NewLessonAccess() *LessonAccess`

NewLessonAccess instantiates a new LessonAccess object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLessonAccessWithDefaults

`func NewLessonAccessWithDefaults() *LessonAccess`

NewLessonAccessWithDefaults instantiates a new LessonAccess object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetID

`func (o *LessonAccess) GetID() string`

GetID returns the ID field if non-nil, zero value otherwise.

### GetIDOk

`func (o *LessonAccess) GetIDOk() (*string, bool)`

GetIDOk returns a tuple with the ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetID

`func (o *LessonAccess) SetID(v string)`

SetID sets ID field to given value.

### HasID

`func (o *LessonAccess) HasID() bool`

HasID returns a boolean if a field has been set.

### GetLessonID

`func (o *LessonAccess) GetLessonID() string`

GetLessonID returns the LessonID field if non-nil, zero value otherwise.

### GetLessonIDOk

`func (o *LessonAccess) GetLessonIDOk() (*string, bool)`

GetLessonIDOk returns a tuple with the LessonID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLessonID

`func (o *LessonAccess) SetLessonID(v string)`

SetLessonID sets LessonID field to given value.

### HasLessonID

`func (o *LessonAccess) HasLessonID() bool`

HasLessonID returns a boolean if a field has been set.

### GetStudentID

`func (o *LessonAccess) GetStudentID() string`

GetStudentID returns the StudentID field if non-nil, zero value otherwise.

### GetStudentIDOk

`func (o *LessonAccess) GetStudentIDOk() (*string, bool)`

GetStudentIDOk returns a tuple with the StudentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStudentID

`func (o *LessonAccess) SetStudentID(v string)`

SetStudentID sets StudentID field to given value.

### HasStudentID

`func (o *LessonAccess) HasStudentID() bool`

HasStudentID returns a boolean if a field has been set.

### GetAccessStatus

`func (o *LessonAccess) GetAccessStatus() string`

GetAccessStatus returns the AccessStatus field if non-nil, zero value otherwise.

### GetAccessStatusOk

`func (o *LessonAccess) GetAccessStatusOk() (*string, bool)`

GetAccessStatusOk returns a tuple with the AccessStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessStatus

`func (o *LessonAccess) SetAccessStatus(v string)`

SetAccessStatus sets AccessStatus field to given value.

### HasAccessStatus

`func (o *LessonAccess) HasAccessStatus() bool`

HasAccessStatus returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *LessonAccess) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *LessonAccess) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *LessonAccess) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *LessonAccess) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../API_README.md#documentation-for-models) [[Back to API list]](../API_README.md#documentation-for-api-endpoints) [[Back to README]](../API_README.md)


