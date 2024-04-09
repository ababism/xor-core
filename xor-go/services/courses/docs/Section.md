# Section

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ID** | Pointer to **string** |  | [optional] 
**Heading** | **string** |  | 
**Description** | **string** |  | 
**Visibility** | **string** |  | 
**Themes** | Pointer to [**[]Theme**](Theme.md) |  | [optional] 

## Methods

### NewSection

`func NewSection(heading string, description string, visibility string, ) *Section`

NewSection instantiates a new Section object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSectionWithDefaults

`func NewSectionWithDefaults() *Section`

NewSectionWithDefaults instantiates a new Section object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetID

`func (o *Section) GetID() string`

GetID returns the ID field if non-nil, zero value otherwise.

### GetIDOk

`func (o *Section) GetIDOk() (*string, bool)`

GetIDOk returns a tuple with the ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetID

`func (o *Section) SetID(v string)`

SetID sets ID field to given value.

### HasID

`func (o *Section) HasID() bool`

HasID returns a boolean if a field has been set.

### GetHeading

`func (o *Section) GetHeading() string`

GetHeading returns the Heading field if non-nil, zero value otherwise.

### GetHeadingOk

`func (o *Section) GetHeadingOk() (*string, bool)`

GetHeadingOk returns a tuple with the Heading field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeading

`func (o *Section) SetHeading(v string)`

SetHeading sets Heading field to given value.


### GetDescription

`func (o *Section) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Section) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Section) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetVisibility

`func (o *Section) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *Section) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *Section) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.


### GetThemes

`func (o *Section) GetThemes() []Theme`

GetThemes returns the Themes field if non-nil, zero value otherwise.

### GetThemesOk

`func (o *Section) GetThemesOk() (*[]Theme, bool)`

GetThemesOk returns a tuple with the Themes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThemes

`func (o *Section) SetThemes(v []Theme)`

SetThemes sets Themes field to given value.

### HasThemes

`func (o *Section) HasThemes() bool`

HasThemes returns a boolean if a field has been set.


[[Back to Model list]](../API_README.md#documentation-for-models) [[Back to API list]](../API_README.md#documentation-for-api-endpoints) [[Back to README]](../API_README.md)


