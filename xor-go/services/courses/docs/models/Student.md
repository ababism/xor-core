# Student

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountID** | Pointer to **string** | ID of the student&#39;s account. | [optional] 

## Methods

### NewStudent

`func NewStudent() *Student`

NewStudent instantiates a new Student object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStudentWithDefaults

`func NewStudentWithDefaults() *Student`

NewStudentWithDefaults instantiates a new Student object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountID

`func (o *Student) GetAccountID() string`

GetAccountID returns the AccountID field if non-nil, zero value otherwise.

### GetAccountIDOk

`func (o *Student) GetAccountIDOk() (*string, bool)`

GetAccountIDOk returns a tuple with the AccountID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountID

`func (o *Student) SetAccountID(v string)`

SetAccountID sets AccountID field to given value.

### HasAccountID

`func (o *Student) HasAccountID() bool`

HasAccountID returns a boolean if a field has been set.


[[Back to Model list]](../API_README.md#documentation-for-models) [[Back to API list]](../API_README.md#documentation-for-api-endpoints) [[Back to README]](../API_README.md)


