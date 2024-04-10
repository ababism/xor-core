# PublicationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ID** | Pointer to **string** | Unique identifier for the publication request. | [optional] 
**CourseID** | Pointer to **string** | ID of the course associated with the request. | [optional] 
**AssigneeID** | Pointer to **string** | ID of the assignee (teacher or moderator). | [optional] 
**RequestStatus** | Pointer to **string** | Status of the publication request. | [optional] 
**Comment** | Pointer to **string** | Optional comment provided by the assignee. | [optional] 
**UpdatedAt** | Pointer to **time.Time** | Date and time when the request was last updated. | [optional] 

## Methods

### NewPublicationRequest

`func NewPublicationRequest() *PublicationRequest`

NewPublicationRequest instantiates a new PublicationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicationRequestWithDefaults

`func NewPublicationRequestWithDefaults() *PublicationRequest`

NewPublicationRequestWithDefaults instantiates a new PublicationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetID

`func (o *PublicationRequest) GetID() string`

GetID returns the ID field if non-nil, zero value otherwise.

### GetIDOk

`func (o *PublicationRequest) GetIDOk() (*string, bool)`

GetIDOk returns a tuple with the ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetID

`func (o *PublicationRequest) SetID(v string)`

SetID sets ID field to given value.

### HasID

`func (o *PublicationRequest) HasID() bool`

HasID returns a boolean if a field has been set.

### GetCourseID

`func (o *PublicationRequest) GetCourseID() string`

GetCourseID returns the CourseID field if non-nil, zero value otherwise.

### GetCourseIDOk

`func (o *PublicationRequest) GetCourseIDOk() (*string, bool)`

GetCourseIDOk returns a tuple with the CourseID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCourseID

`func (o *PublicationRequest) SetCourseID(v string)`

SetCourseID sets CourseID field to given value.

### HasCourseID

`func (o *PublicationRequest) HasCourseID() bool`

HasCourseID returns a boolean if a field has been set.

### GetAssigneeID

`func (o *PublicationRequest) GetAssigneeID() string`

GetAssigneeID returns the AssigneeID field if non-nil, zero value otherwise.

### GetAssigneeIDOk

`func (o *PublicationRequest) GetAssigneeIDOk() (*string, bool)`

GetAssigneeIDOk returns a tuple with the AssigneeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssigneeID

`func (o *PublicationRequest) SetAssigneeID(v string)`

SetAssigneeID sets AssigneeID field to given value.

### HasAssigneeID

`func (o *PublicationRequest) HasAssigneeID() bool`

HasAssigneeID returns a boolean if a field has been set.

### GetRequestStatus

`func (o *PublicationRequest) GetRequestStatus() string`

GetRequestStatus returns the RequestStatus field if non-nil, zero value otherwise.

### GetRequestStatusOk

`func (o *PublicationRequest) GetRequestStatusOk() (*string, bool)`

GetRequestStatusOk returns a tuple with the RequestStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestStatus

`func (o *PublicationRequest) SetRequestStatus(v string)`

SetRequestStatus sets RequestStatus field to given value.

### HasRequestStatus

`func (o *PublicationRequest) HasRequestStatus() bool`

HasRequestStatus returns a boolean if a field has been set.

### GetComment

`func (o *PublicationRequest) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *PublicationRequest) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *PublicationRequest) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *PublicationRequest) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *PublicationRequest) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PublicationRequest) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PublicationRequest) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *PublicationRequest) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../API_README.md#documentation-for-models) [[Back to API list]](../API_README.md#documentation-for-api-endpoints) [[Back to README]](../API_README.md)


