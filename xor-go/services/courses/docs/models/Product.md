# Product

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ID** | Pointer to **string** | Unique identifier for the product. | [optional] 
**Owner** | **string** | ID of the owner of the product. | 
**Price** | **float32** | Price of the product. | 
**Item** | **string** | ID of the item associated with the product. | 

## Methods

### NewProduct

`func NewProduct(owner string, price float32, item string, ) *Product`

NewProduct instantiates a new Product object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductWithDefaults

`func NewProductWithDefaults() *Product`

NewProductWithDefaults instantiates a new Product object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetID

`func (o *Product) GetID() string`

GetID returns the ID field if non-nil, zero value otherwise.

### GetIDOk

`func (o *Product) GetIDOk() (*string, bool)`

GetIDOk returns a tuple with the ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetID

`func (o *Product) SetID(v string)`

SetID sets ID field to given value.

### HasID

`func (o *Product) HasID() bool`

HasID returns a boolean if a field has been set.

### GetOwner

`func (o *Product) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *Product) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *Product) SetOwner(v string)`

SetOwner sets Owner field to given value.


### GetPrice

`func (o *Product) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *Product) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *Product) SetPrice(v float32)`

SetPrice sets Price field to given value.


### GetItem

`func (o *Product) GetItem() string`

GetItem returns the Item field if non-nil, zero value otherwise.

### GetItemOk

`func (o *Product) GetItemOk() (*string, bool)`

GetItemOk returns a tuple with the Item field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItem

`func (o *Product) SetItem(v string)`

SetItem sets Item field to given value.



[[Back to Model list]](../API_README.md#documentation-for-models) [[Back to API list]](../API_README.md#documentation-for-api-endpoints) [[Back to README]](../API_README.md)


