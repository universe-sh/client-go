# \EventsApi

All URIs are relative to *http://api-houston-$.endpoints.$.cloud.goog*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddEvents**](EventsApi.md#AddEvents) | **Post** /v1/events | 
[**ListEvents**](EventsApi.md#ListEvents) | **Get** /v1/events | 
[**ReadEvent**](EventsApi.md#ReadEvent) | **Get** /v1/events/{event} | 



## AddEvents

> Generic AddEvents(ctx, )



Add properties of events

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**Generic**](generic.md)

### Authorization

[okta_jwt](../README.md#okta_jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEvents

> []Event ListEvents(ctx, )



List properties of events

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]Event**](event.md)

### Authorization

[okta_jwt](../README.md#okta_jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadEvent

> Event ReadEvent(ctx, event)



Read properties of event

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**event** | **string**| string event (name or id) of the event | 

### Return type

[**Event**](event.md)

### Authorization

[okta_jwt](../README.md#okta_jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

