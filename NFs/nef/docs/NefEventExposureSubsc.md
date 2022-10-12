# NefEventExposureSubsc

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventsSubs** | [**[]NefEventSubs**](NefEventSubs.md) |  | 
**EventsRepInfo** | Pointer to [**Object**](Object.md) |  | [optional] 
**NotifUri** | [**Object**](Object.md) |  | 
**NotifId** | **string** |  | 
**EventNotifs** | Pointer to [**[]NefEventNotification**](NefEventNotification.md) |  | [optional] 
**SuppFeat** | Pointer to [**Object**](Object.md) |  | [optional] 

## Methods

### NewNefEventExposureSubsc

`func NewNefEventExposureSubsc(eventsSubs []NefEventSubs, notifUri Object, notifId string, ) *NefEventExposureSubsc`

NewNefEventExposureSubsc instantiates a new NefEventExposureSubsc object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNefEventExposureSubscWithDefaults

`func NewNefEventExposureSubscWithDefaults() *NefEventExposureSubsc`

NewNefEventExposureSubscWithDefaults instantiates a new NefEventExposureSubsc object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventsSubs

`func (o *NefEventExposureSubsc) GetEventsSubs() []NefEventSubs`

GetEventsSubs returns the EventsSubs field if non-nil, zero value otherwise.

### GetEventsSubsOk

`func (o *NefEventExposureSubsc) GetEventsSubsOk() (*[]NefEventSubs, bool)`

GetEventsSubsOk returns a tuple with the EventsSubs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventsSubs

`func (o *NefEventExposureSubsc) SetEventsSubs(v []NefEventSubs)`

SetEventsSubs sets EventsSubs field to given value.


### GetEventsRepInfo

`func (o *NefEventExposureSubsc) GetEventsRepInfo() Object`

GetEventsRepInfo returns the EventsRepInfo field if non-nil, zero value otherwise.

### GetEventsRepInfoOk

`func (o *NefEventExposureSubsc) GetEventsRepInfoOk() (*Object, bool)`

GetEventsRepInfoOk returns a tuple with the EventsRepInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventsRepInfo

`func (o *NefEventExposureSubsc) SetEventsRepInfo(v Object)`

SetEventsRepInfo sets EventsRepInfo field to given value.

### HasEventsRepInfo

`func (o *NefEventExposureSubsc) HasEventsRepInfo() bool`

HasEventsRepInfo returns a boolean if a field has been set.

### GetNotifUri

`func (o *NefEventExposureSubsc) GetNotifUri() Object`

GetNotifUri returns the NotifUri field if non-nil, zero value otherwise.

### GetNotifUriOk

`func (o *NefEventExposureSubsc) GetNotifUriOk() (*Object, bool)`

GetNotifUriOk returns a tuple with the NotifUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifUri

`func (o *NefEventExposureSubsc) SetNotifUri(v Object)`

SetNotifUri sets NotifUri field to given value.


### GetNotifId

`func (o *NefEventExposureSubsc) GetNotifId() string`

GetNotifId returns the NotifId field if non-nil, zero value otherwise.

### GetNotifIdOk

`func (o *NefEventExposureSubsc) GetNotifIdOk() (*string, bool)`

GetNotifIdOk returns a tuple with the NotifId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifId

`func (o *NefEventExposureSubsc) SetNotifId(v string)`

SetNotifId sets NotifId field to given value.


### GetEventNotifs

`func (o *NefEventExposureSubsc) GetEventNotifs() []NefEventNotification`

GetEventNotifs returns the EventNotifs field if non-nil, zero value otherwise.

### GetEventNotifsOk

`func (o *NefEventExposureSubsc) GetEventNotifsOk() (*[]NefEventNotification, bool)`

GetEventNotifsOk returns a tuple with the EventNotifs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventNotifs

`func (o *NefEventExposureSubsc) SetEventNotifs(v []NefEventNotification)`

SetEventNotifs sets EventNotifs field to given value.

### HasEventNotifs

`func (o *NefEventExposureSubsc) HasEventNotifs() bool`

HasEventNotifs returns a boolean if a field has been set.

### GetSuppFeat

`func (o *NefEventExposureSubsc) GetSuppFeat() Object`

GetSuppFeat returns the SuppFeat field if non-nil, zero value otherwise.

### GetSuppFeatOk

`func (o *NefEventExposureSubsc) GetSuppFeatOk() (*Object, bool)`

GetSuppFeatOk returns a tuple with the SuppFeat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppFeat

`func (o *NefEventExposureSubsc) SetSuppFeat(v Object)`

SetSuppFeat sets SuppFeat field to given value.

### HasSuppFeat

`func (o *NefEventExposureSubsc) HasSuppFeat() bool`

HasSuppFeat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


