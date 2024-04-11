# Theme

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ID** | Pointer to **string** |  | [optional] 
**Heading** | **string** |  | 
**Visibility** | **string** |  | 
**LessonIDs** | Pointer to **[]string** | Lessons in theme. | [optional] 

## Methods

### NewTheme

`func NewTheme(heading string, visibility string, ) *Theme`

NewTheme instantiates a new Theme object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThemeWithDefaults

`func NewThemeWithDefaults() *Theme`

NewThemeWithDefaults instantiates a new Theme object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetID

`func (o *Theme) GetID() string`

GetID returns the ID field if non-nil, zero value otherwise.

### GetIDOk

`func (o *Theme) GetIDOk() (*string, bool)`

GetIDOk returns a tuple with the ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetID

`func (o *Theme) SetID(v string)`

SetID sets ID field to given value.

### HasID

`func (o *Theme) HasID() bool`

HasID returns a boolean if a field has been set.

### GetHeading

`func (o *Theme) GetHeading() string`

GetHeading returns the Heading field if non-nil, zero value otherwise.

### GetHeadingOk

`func (o *Theme) GetHeadingOk() (*string, bool)`

GetHeadingOk returns a tuple with the Heading field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeading

`func (o *Theme) SetHeading(v string)`

SetHeading sets Heading field to given value.


### GetVisibility

`func (o *Theme) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *Theme) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *Theme) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.


### GetLessonIDs

`func (o *Theme) GetLessonIDs() []string`

GetLessonIDs returns the LessonIDs field if non-nil, zero value otherwise.

### GetLessonIDsOk

`func (o *Theme) GetLessonIDsOk() (*[]string, bool)`

GetLessonIDsOk returns a tuple with the LessonIDs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLessonIDs

`func (o *Theme) SetLessonIDs(v []string)`

SetLessonIDs sets LessonIDs field to given value.

### HasLessonIDs

`func (o *Theme) HasLessonIDs() bool`

HasLessonIDs returns a boolean if a field has been set.


[[Back to Model list]](../API_README.md#documentation-for-models) [[Back to API list]](../API_README.md#documentation-for-api-endpoints) [[Back to README]](../API_README.md)


