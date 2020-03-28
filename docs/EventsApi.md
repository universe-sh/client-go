# \EventsApi

All URIs are relative to *https://virtserver.swaggerhub.com/universe-sh/Houston/1.0.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListEvents**](EventsApi.md#ListEvents) | **Get** /v1/events | 
[**ReadEvent**](EventsApi.md#ReadEvent) | **Get** /v1/events/{event} | 


# **ListEvents**
> Events ListEvents(ctx, )


List properties of events

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**Events**](events.md)

### Authorization

[okta_jwt](../README.md#okta_jwt)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadEvent**
> Event ReadEvent(ctx, event)


Read properties of event

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **event** | **string**|  | 

### Return type

[**Event**](event.md)

### Authorization

[okta_jwt](../README.md#okta_jwt)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

