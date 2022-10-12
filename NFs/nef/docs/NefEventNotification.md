# NefEventNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Event** | [**NefEvent**](NefEvent.md) |  | 
**TimeStamp** | [**Object**](Object.md) |  | 
**SvcExprcInfos** | Pointer to [**[]ServiceExperienceInfo**](ServiceExperienceInfo.md) |  | [optional] 
**UeMobilityInfos** | Pointer to [**[]UeMobilityInfo**](UeMobilityInfo.md) |  | [optional] 
**UeCommInfos** | Pointer to [**[]UeCommunicationInfo**](UeCommunicationInfo.md) |  | [optional] 
**ExcepInfos** | Pointer to [**[]Object**](Object.md) |  | [optional] 
**CongestionInfos** | Pointer to [**[]Object**](Object.md) |  | [optional] 
**PerfDataInfos** | Pointer to [**[]PerformanceDataInfo**](PerformanceDataInfo.md) |  | [optional] 
**DispersionInfos** | Pointer to [**[]Object**](Object.md) |  | [optional] 
**CollBhvrInfs** | Pointer to [**[]Object**](Object.md) |  | [optional] 
**QoeMetrInfos** | Pointer to [**[]Object**](Object.md) |  | [optional] 
**ConsumpInfos** | Pointer to [**[]Object**](Object.md) |  | [optional] 
**NetAssInvInfos** | Pointer to [**[]Object**](Object.md) |  | [optional] 
**ChgPlyInvInfos** | Pointer to [**[]Object**](Object.md) |  | [optional] 
**MsAccActInfos** | Pointer to [**[]Object**](Object.md) |  | [optional] 

## Methods

### NewNefEventNotification

`func NewNefEventNotification(event NefEvent, timeStamp Object, ) *NefEventNotification`

NewNefEventNotification instantiates a new NefEventNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNefEventNotificationWithDefaults

`func NewNefEventNotificationWithDefaults() *NefEventNotification`

NewNefEventNotificationWithDefaults instantiates a new NefEventNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvent

`func (o *NefEventNotification) GetEvent() NefEvent`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *NefEventNotification) GetEventOk() (*NefEvent, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *NefEventNotification) SetEvent(v NefEvent)`

SetEvent sets Event field to given value.


### GetTimeStamp

`func (o *NefEventNotification) GetTimeStamp() Object`

GetTimeStamp returns the TimeStamp field if non-nil, zero value otherwise.

### GetTimeStampOk

`func (o *NefEventNotification) GetTimeStampOk() (*Object, bool)`

GetTimeStampOk returns a tuple with the TimeStamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeStamp

`func (o *NefEventNotification) SetTimeStamp(v Object)`

SetTimeStamp sets TimeStamp field to given value.


### GetSvcExprcInfos

`func (o *NefEventNotification) GetSvcExprcInfos() []ServiceExperienceInfo`

GetSvcExprcInfos returns the SvcExprcInfos field if non-nil, zero value otherwise.

### GetSvcExprcInfosOk

`func (o *NefEventNotification) GetSvcExprcInfosOk() (*[]ServiceExperienceInfo, bool)`

GetSvcExprcInfosOk returns a tuple with the SvcExprcInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSvcExprcInfos

`func (o *NefEventNotification) SetSvcExprcInfos(v []ServiceExperienceInfo)`

SetSvcExprcInfos sets SvcExprcInfos field to given value.

### HasSvcExprcInfos

`func (o *NefEventNotification) HasSvcExprcInfos() bool`

HasSvcExprcInfos returns a boolean if a field has been set.

### GetUeMobilityInfos

`func (o *NefEventNotification) GetUeMobilityInfos() []UeMobilityInfo`

GetUeMobilityInfos returns the UeMobilityInfos field if non-nil, zero value otherwise.

### GetUeMobilityInfosOk

`func (o *NefEventNotification) GetUeMobilityInfosOk() (*[]UeMobilityInfo, bool)`

GetUeMobilityInfosOk returns a tuple with the UeMobilityInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeMobilityInfos

`func (o *NefEventNotification) SetUeMobilityInfos(v []UeMobilityInfo)`

SetUeMobilityInfos sets UeMobilityInfos field to given value.

### HasUeMobilityInfos

`func (o *NefEventNotification) HasUeMobilityInfos() bool`

HasUeMobilityInfos returns a boolean if a field has been set.

### GetUeCommInfos

`func (o *NefEventNotification) GetUeCommInfos() []UeCommunicationInfo`

GetUeCommInfos returns the UeCommInfos field if non-nil, zero value otherwise.

### GetUeCommInfosOk

`func (o *NefEventNotification) GetUeCommInfosOk() (*[]UeCommunicationInfo, bool)`

GetUeCommInfosOk returns a tuple with the UeCommInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeCommInfos

`func (o *NefEventNotification) SetUeCommInfos(v []UeCommunicationInfo)`

SetUeCommInfos sets UeCommInfos field to given value.

### HasUeCommInfos

`func (o *NefEventNotification) HasUeCommInfos() bool`

HasUeCommInfos returns a boolean if a field has been set.

### GetExcepInfos

`func (o *NefEventNotification) GetExcepInfos() []Object`

GetExcepInfos returns the ExcepInfos field if non-nil, zero value otherwise.

### GetExcepInfosOk

`func (o *NefEventNotification) GetExcepInfosOk() (*[]Object, bool)`

GetExcepInfosOk returns a tuple with the ExcepInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcepInfos

`func (o *NefEventNotification) SetExcepInfos(v []Object)`

SetExcepInfos sets ExcepInfos field to given value.

### HasExcepInfos

`func (o *NefEventNotification) HasExcepInfos() bool`

HasExcepInfos returns a boolean if a field has been set.

### GetCongestionInfos

`func (o *NefEventNotification) GetCongestionInfos() []Object`

GetCongestionInfos returns the CongestionInfos field if non-nil, zero value otherwise.

### GetCongestionInfosOk

`func (o *NefEventNotification) GetCongestionInfosOk() (*[]Object, bool)`

GetCongestionInfosOk returns a tuple with the CongestionInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCongestionInfos

`func (o *NefEventNotification) SetCongestionInfos(v []Object)`

SetCongestionInfos sets CongestionInfos field to given value.

### HasCongestionInfos

`func (o *NefEventNotification) HasCongestionInfos() bool`

HasCongestionInfos returns a boolean if a field has been set.

### GetPerfDataInfos

`func (o *NefEventNotification) GetPerfDataInfos() []PerformanceDataInfo`

GetPerfDataInfos returns the PerfDataInfos field if non-nil, zero value otherwise.

### GetPerfDataInfosOk

`func (o *NefEventNotification) GetPerfDataInfosOk() (*[]PerformanceDataInfo, bool)`

GetPerfDataInfosOk returns a tuple with the PerfDataInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerfDataInfos

`func (o *NefEventNotification) SetPerfDataInfos(v []PerformanceDataInfo)`

SetPerfDataInfos sets PerfDataInfos field to given value.

### HasPerfDataInfos

`func (o *NefEventNotification) HasPerfDataInfos() bool`

HasPerfDataInfos returns a boolean if a field has been set.

### GetDispersionInfos

`func (o *NefEventNotification) GetDispersionInfos() []Object`

GetDispersionInfos returns the DispersionInfos field if non-nil, zero value otherwise.

### GetDispersionInfosOk

`func (o *NefEventNotification) GetDispersionInfosOk() (*[]Object, bool)`

GetDispersionInfosOk returns a tuple with the DispersionInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDispersionInfos

`func (o *NefEventNotification) SetDispersionInfos(v []Object)`

SetDispersionInfos sets DispersionInfos field to given value.

### HasDispersionInfos

`func (o *NefEventNotification) HasDispersionInfos() bool`

HasDispersionInfos returns a boolean if a field has been set.

### GetCollBhvrInfs

`func (o *NefEventNotification) GetCollBhvrInfs() []Object`

GetCollBhvrInfs returns the CollBhvrInfs field if non-nil, zero value otherwise.

### GetCollBhvrInfsOk

`func (o *NefEventNotification) GetCollBhvrInfsOk() (*[]Object, bool)`

GetCollBhvrInfsOk returns a tuple with the CollBhvrInfs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollBhvrInfs

`func (o *NefEventNotification) SetCollBhvrInfs(v []Object)`

SetCollBhvrInfs sets CollBhvrInfs field to given value.

### HasCollBhvrInfs

`func (o *NefEventNotification) HasCollBhvrInfs() bool`

HasCollBhvrInfs returns a boolean if a field has been set.

### GetQoeMetrInfos

`func (o *NefEventNotification) GetQoeMetrInfos() []Object`

GetQoeMetrInfos returns the QoeMetrInfos field if non-nil, zero value otherwise.

### GetQoeMetrInfosOk

`func (o *NefEventNotification) GetQoeMetrInfosOk() (*[]Object, bool)`

GetQoeMetrInfosOk returns a tuple with the QoeMetrInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQoeMetrInfos

`func (o *NefEventNotification) SetQoeMetrInfos(v []Object)`

SetQoeMetrInfos sets QoeMetrInfos field to given value.

### HasQoeMetrInfos

`func (o *NefEventNotification) HasQoeMetrInfos() bool`

HasQoeMetrInfos returns a boolean if a field has been set.

### GetConsumpInfos

`func (o *NefEventNotification) GetConsumpInfos() []Object`

GetConsumpInfos returns the ConsumpInfos field if non-nil, zero value otherwise.

### GetConsumpInfosOk

`func (o *NefEventNotification) GetConsumpInfosOk() (*[]Object, bool)`

GetConsumpInfosOk returns a tuple with the ConsumpInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumpInfos

`func (o *NefEventNotification) SetConsumpInfos(v []Object)`

SetConsumpInfos sets ConsumpInfos field to given value.

### HasConsumpInfos

`func (o *NefEventNotification) HasConsumpInfos() bool`

HasConsumpInfos returns a boolean if a field has been set.

### GetNetAssInvInfos

`func (o *NefEventNotification) GetNetAssInvInfos() []Object`

GetNetAssInvInfos returns the NetAssInvInfos field if non-nil, zero value otherwise.

### GetNetAssInvInfosOk

`func (o *NefEventNotification) GetNetAssInvInfosOk() (*[]Object, bool)`

GetNetAssInvInfosOk returns a tuple with the NetAssInvInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetAssInvInfos

`func (o *NefEventNotification) SetNetAssInvInfos(v []Object)`

SetNetAssInvInfos sets NetAssInvInfos field to given value.

### HasNetAssInvInfos

`func (o *NefEventNotification) HasNetAssInvInfos() bool`

HasNetAssInvInfos returns a boolean if a field has been set.

### GetChgPlyInvInfos

`func (o *NefEventNotification) GetChgPlyInvInfos() []Object`

GetChgPlyInvInfos returns the ChgPlyInvInfos field if non-nil, zero value otherwise.

### GetChgPlyInvInfosOk

`func (o *NefEventNotification) GetChgPlyInvInfosOk() (*[]Object, bool)`

GetChgPlyInvInfosOk returns a tuple with the ChgPlyInvInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChgPlyInvInfos

`func (o *NefEventNotification) SetChgPlyInvInfos(v []Object)`

SetChgPlyInvInfos sets ChgPlyInvInfos field to given value.

### HasChgPlyInvInfos

`func (o *NefEventNotification) HasChgPlyInvInfos() bool`

HasChgPlyInvInfos returns a boolean if a field has been set.

### GetMsAccActInfos

`func (o *NefEventNotification) GetMsAccActInfos() []Object`

GetMsAccActInfos returns the MsAccActInfos field if non-nil, zero value otherwise.

### GetMsAccActInfosOk

`func (o *NefEventNotification) GetMsAccActInfosOk() (*[]Object, bool)`

GetMsAccActInfosOk returns a tuple with the MsAccActInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsAccActInfos

`func (o *NefEventNotification) SetMsAccActInfos(v []Object)`

SetMsAccActInfos sets MsAccActInfos field to given value.

### HasMsAccActInfos

`func (o *NefEventNotification) HasMsAccActInfos() bool`

HasMsAccActInfos returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


