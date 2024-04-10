# UserAccessConfirmBuyerIDPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BuyerID** | Pointer to **string** | ID of the buyer. | [optional] 
**Products** | Pointer to [**[]Product**](Product.md) |  | [optional] 

## Methods

### NewUserAccessConfirmBuyerIDPostRequest

`func NewUserAccessConfirmBuyerIDPostRequest() *UserAccessConfirmBuyerIDPostRequest`

NewUserAccessConfirmBuyerIDPostRequest instantiates a new UserAccessConfirmBuyerIDPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserAccessConfirmBuyerIDPostRequestWithDefaults

`func NewUserAccessConfirmBuyerIDPostRequestWithDefaults() *UserAccessConfirmBuyerIDPostRequest`

NewUserAccessConfirmBuyerIDPostRequestWithDefaults instantiates a new UserAccessConfirmBuyerIDPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuyerID

`func (o *UserAccessConfirmBuyerIDPostRequest) GetBuyerID() string`

GetBuyerID returns the BuyerID field if non-nil, zero value otherwise.

### GetBuyerIDOk

`func (o *UserAccessConfirmBuyerIDPostRequest) GetBuyerIDOk() (*string, bool)`

GetBuyerIDOk returns a tuple with the BuyerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyerID

`func (o *UserAccessConfirmBuyerIDPostRequest) SetBuyerID(v string)`

SetBuyerID sets BuyerID field to given value.

### HasBuyerID

`func (o *UserAccessConfirmBuyerIDPostRequest) HasBuyerID() bool`

HasBuyerID returns a boolean if a field has been set.

### GetProducts

`func (o *UserAccessConfirmBuyerIDPostRequest) GetProducts() []Product`

GetProducts returns the Products field if non-nil, zero value otherwise.

### GetProductsOk

`func (o *UserAccessConfirmBuyerIDPostRequest) GetProductsOk() (*[]Product, bool)`

GetProductsOk returns a tuple with the Products field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducts

`func (o *UserAccessConfirmBuyerIDPostRequest) SetProducts(v []Product)`

SetProducts sets Products field to given value.

### HasProducts

`func (o *UserAccessConfirmBuyerIDPostRequest) HasProducts() bool`

HasProducts returns a boolean if a field has been set.


[[Back to Model list]](../API_README.md#documentation-for-models) [[Back to API list]](../API_README.md#documentation-for-api-endpoints) [[Back to README]](../API_README.md)


