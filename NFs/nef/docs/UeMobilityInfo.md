# UeMobilityInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Supi** | [**Object**](Object.md) |  | 
**AppId** | Pointer to [**Object**](Object.md) |  | [optional] 
**UeTrajs** | [**[]UeTrajectoryInfo**](UeTrajectoryInfo.md) |  | 

## Methods

### NewUeMobilityInfo

`func NewUeMobilityInfo(supi Object, ueTrajs []UeTrajectoryInfo, ) *UeMobilityInfo`

NewUeMobilityInfo instantiates a new UeMobilityInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUeMobilityInfoWithDefaults

`func NewUeMobilityInfoWithDefaults() *UeMobilityInfo`

NewUeMobilityInfoWithDefaults instantiates a new UeMobilityInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSupi

`func (o *UeMobilityInfo) GetSupi() Object`

GetSupi returns the Supi field if non-nil, zero value otherwise.

### GetSupiOk

`func (o *UeMobilityInfo) GetSupiOk() (*Object, bool)`

GetSupiOk returns a tuple with the Supi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupi

`func (o *UeMobilityInfo) SetSupi(v Object)`

SetSupi sets Supi field to given value.


### GetAppId

`func (o *UeMobilityInfo) GetAppId() Object`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *UeMobilityInfo) GetAppIdOk() (*Object, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *UeMobilityInfo) SetAppId(v Object)`

SetAppId sets AppId field to given value.

### HasAppId

`func (o *UeMobilityInfo) HasAppId() bool`

HasAppId returns a boolean if a field has been set.

### GetUeTrajs

`func (o *UeMobilityInfo) GetUeTrajs() []UeTrajectoryInfo`

GetUeTrajs returns the UeTrajs field if non-nil, zero value otherwise.

### GetUeTrajsOk

`func (o *UeMobilityInfo) GetUeTrajsOk() (*[]UeTrajectoryInfo, bool)`

GetUeTrajsOk returns a tuple with the UeTrajs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeTrajs

`func (o *UeMobilityInfo) SetUeTrajs(v []UeTrajectoryInfo)`

SetUeTrajs sets UeTrajs field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


