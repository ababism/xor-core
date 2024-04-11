# Actor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ID** | **string** | Unique identifier for the actor. | 
**Roles** | **string** | Role of the actor. | 

## Methods

### NewActor

`func NewActor(iD string, roles string, ) *Actor`

NewActor instantiates a new Actor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActorWithDefaults

`func NewActorWithDefaults() *Actor`

NewActorWithDefaults instantiates a new Actor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetID

`func (o *Actor) GetID() string`

GetID returns the ID field if non-nil, zero value otherwise.

### GetIDOk

`func (o *Actor) GetIDOk() (*string, bool)`

GetIDOk returns a tuple with the ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetID

`func (o *Actor) SetID(v string)`

SetID sets ID field to given value.


### GetRoles

`func (o *Actor) GetRoles() string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *Actor) GetRolesOk() (*string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *Actor) SetRoles(v string)`

SetRoles sets Roles field to given value.



[[Back to Model list]](../API_README.md#documentation-for-models) [[Back to API list]](../API_README.md#documentation-for-api-endpoints) [[Back to README]](../API_README.md)


