# OptionalActor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ID** | Pointer to **string** | Unique identifier for the actor. | [optional] 
**Roles** | Pointer to **string** | Role of the actor. | [optional] 

## Methods

### NewOptionalActor

`func NewOptionalActor() *OptionalActor`

NewOptionalActor instantiates a new OptionalActor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOptionalActorWithDefaults

`func NewOptionalActorWithDefaults() *OptionalActor`

NewOptionalActorWithDefaults instantiates a new OptionalActor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetID

`func (o *OptionalActor) GetID() string`

GetID returns the ID field if non-nil, zero value otherwise.

### GetIDOk

`func (o *OptionalActor) GetIDOk() (*string, bool)`

GetIDOk returns a tuple with the ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetID

`func (o *OptionalActor) SetID(v string)`

SetID sets ID field to given value.

### HasID

`func (o *OptionalActor) HasID() bool`

HasID returns a boolean if a field has been set.

### GetRoles

`func (o *OptionalActor) GetRoles() string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *OptionalActor) GetRolesOk() (*string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *OptionalActor) SetRoles(v string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *OptionalActor) HasRoles() bool`

HasRoles returns a boolean if a field has been set.


[[Back to Model list]](../API_README.md#documentation-for-models) [[Back to API list]](../API_README.md#documentation-for-api-endpoints) [[Back to README]](../API_README.md)


