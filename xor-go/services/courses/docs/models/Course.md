# Course

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ID** | Pointer to **string** |  | [optional] 
**FeedbackID** | Pointer to **string** |  | [optional] 
**TeacherID** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**Discipline** | **string** |  | 
**Landing** | **string** |  | 
**Visibility** | **string** |  | 
**Sections** | Pointer to [**[]Section**](Section.md) |  | [optional] 

## Methods

### NewCourse

`func NewCourse(name string, discipline string, landing string, visibility string, ) *Course`

NewCourse instantiates a new Course object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCourseWithDefaults

`func NewCourseWithDefaults() *Course`

NewCourseWithDefaults instantiates a new Course object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetID

`func (o *Course) GetID() string`

GetID returns the ID field if non-nil, zero value otherwise.

### GetIDOk

`func (o *Course) GetIDOk() (*string, bool)`

GetIDOk returns a tuple with the ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetID

`func (o *Course) SetID(v string)`

SetID sets ID field to given value.

### HasID

`func (o *Course) HasID() bool`

HasID returns a boolean if a field has been set.

### GetFeedbackID

`func (o *Course) GetFeedbackID() string`

GetFeedbackID returns the FeedbackID field if non-nil, zero value otherwise.

### GetFeedbackIDOk

`func (o *Course) GetFeedbackIDOk() (*string, bool)`

GetFeedbackIDOk returns a tuple with the FeedbackID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeedbackID

`func (o *Course) SetFeedbackID(v string)`

SetFeedbackID sets FeedbackID field to given value.

### HasFeedbackID

`func (o *Course) HasFeedbackID() bool`

HasFeedbackID returns a boolean if a field has been set.

### GetTeacherID

`func (o *Course) GetTeacherID() string`

GetTeacherID returns the TeacherID field if non-nil, zero value otherwise.

### GetTeacherIDOk

`func (o *Course) GetTeacherIDOk() (*string, bool)`

GetTeacherIDOk returns a tuple with the TeacherID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeacherID

`func (o *Course) SetTeacherID(v string)`

SetTeacherID sets TeacherID field to given value.

### HasTeacherID

`func (o *Course) HasTeacherID() bool`

HasTeacherID returns a boolean if a field has been set.

### GetName

`func (o *Course) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Course) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Course) SetName(v string)`

SetName sets Name field to given value.


### GetDiscipline

`func (o *Course) GetDiscipline() string`

GetDiscipline returns the Discipline field if non-nil, zero value otherwise.

### GetDisciplineOk

`func (o *Course) GetDisciplineOk() (*string, bool)`

GetDisciplineOk returns a tuple with the Discipline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscipline

`func (o *Course) SetDiscipline(v string)`

SetDiscipline sets Discipline field to given value.


### GetLanding

`func (o *Course) GetLanding() string`

GetLanding returns the Landing field if non-nil, zero value otherwise.

### GetLandingOk

`func (o *Course) GetLandingOk() (*string, bool)`

GetLandingOk returns a tuple with the Landing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanding

`func (o *Course) SetLanding(v string)`

SetLanding sets Landing field to given value.


### GetVisibility

`func (o *Course) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *Course) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *Course) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.


### GetSections

`func (o *Course) GetSections() []Section`

GetSections returns the Sections field if non-nil, zero value otherwise.

### GetSectionsOk

`func (o *Course) GetSectionsOk() (*[]Section, bool)`

GetSectionsOk returns a tuple with the Sections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSections

`func (o *Course) SetSections(v []Section)`

SetSections sets Sections field to given value.

### HasSections

`func (o *Course) HasSections() bool`

HasSections returns a boolean if a field has been set.


[[Back to Model list]](../API_README.md#documentation-for-models) [[Back to API list]](../API_README.md#documentation-for-api-endpoints) [[Back to README]](../API_README.md)


