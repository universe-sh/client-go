# \TeamsApi

All URIs are relative to *http://api-houston-$.endpoints.$.cloud.goog*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddTeams**](TeamsApi.md#AddTeams) | **Post** /v1/teams | 
[**DeleteInvitations**](TeamsApi.md#DeleteInvitations) | **Delete** /v1/teams/invitations/{invitation} | 
[**DeleteTeam**](TeamsApi.md#DeleteTeam) | **Delete** /v1/teams/{team} | 
[**ListInvitations**](TeamsApi.md#ListInvitations) | **Get** /v1/teams/{team}/invitations | 
[**ListInvoices**](TeamsApi.md#ListInvoices) | **Get** /v1/teams/{team}/invoices | 
[**ListTeams**](TeamsApi.md#ListTeams) | **Get** /v1/teams | 
[**ReadInvitations**](TeamsApi.md#ReadInvitations) | **Get** /v1/teams/invitations/{invitation} | 
[**ReadInvitationsAccept**](TeamsApi.md#ReadInvitationsAccept) | **Get** /v1/teams/invitations/{invitation}/accept | 
[**ReadInvoice**](TeamsApi.md#ReadInvoice) | **Get** /v1/teams/{team}/invoices/{invoice} | 
[**ReadPreferences**](TeamsApi.md#ReadPreferences) | **Get** /v1/teams/{team}/preferences | 
[**ReadTeam**](TeamsApi.md#ReadTeam) | **Get** /v1/teams/{team} | 
[**UpdateInvitations**](TeamsApi.md#UpdateInvitations) | **Put** /v1/teams/invitations/{invitation} | 
[**UpdatePreferences**](TeamsApi.md#UpdatePreferences) | **Put** /v1/teams/{team}/preferences | 
[**UpdateTeam**](TeamsApi.md#UpdateTeam) | **Put** /v1/teams/{team} | 



## AddTeams

> Generic AddTeams(ctx, )



Add properties of teams

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


## DeleteInvitations

> Generic DeleteInvitations(ctx, invitation)



Delete properties of invitations

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**invitation** | **string**|  | 

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


## DeleteTeam

> Generic DeleteTeam(ctx, team)



Delete properties of team

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**team** | **string**|  | 

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


## ListInvitations

> []Invitation ListInvitations(ctx, team)



List properties of invitations

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**team** | **string**|  | 

### Return type

[**[]Invitation**](invitation.md)

### Authorization

[okta_jwt](../README.md#okta_jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListInvoices

> []Invoice ListInvoices(ctx, team)



List properties of invoices

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**team** | **string**|  | 

### Return type

[**[]Invoice**](invoice.md)

### Authorization

[okta_jwt](../README.md#okta_jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTeams

> []Team ListTeams(ctx, )



List properties of teams

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]Team**](team.md)

### Authorization

[okta_jwt](../README.md#okta_jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadInvitations

> []Invitation ReadInvitations(ctx, invitation)



Read properties of invitations

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**invitation** | **string**|  | 

### Return type

[**[]Invitation**](invitation.md)

### Authorization

[okta_jwt](../README.md#okta_jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadInvitationsAccept

> Invitation ReadInvitationsAccept(ctx, invitation)



Read properties of invitationsaccept

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**invitation** | **string**|  | 

### Return type

[**Invitation**](invitation.md)

### Authorization

[okta_jwt](../README.md#okta_jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadInvoice

> Invoice ReadInvoice(ctx, team, invoice)



Read properties of invoice

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**team** | **string**|  | 
**invoice** | **string**|  | 

### Return type

[**Invoice**](invoice.md)

### Authorization

[okta_jwt](../README.md#okta_jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadPreferences

> Preferences ReadPreferences(ctx, team)



Read properties of preferences

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**team** | **string**|  | 

### Return type

[**Preferences**](preferences.md)

### Authorization

[okta_jwt](../README.md#okta_jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadTeam

> Team ReadTeam(ctx, team)



Read properties of team

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**team** | **string**|  | 

### Return type

[**Team**](team.md)

### Authorization

[okta_jwt](../README.md#okta_jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateInvitations

> []Invitation UpdateInvitations(ctx, invitation)



Update properties of invitations

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**invitation** | **string**|  | 

### Return type

[**[]Invitation**](invitation.md)

### Authorization

[okta_jwt](../README.md#okta_jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePreferences

> Preferences UpdatePreferences(ctx, team)



Update properties of preferences

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**team** | **string**|  | 

### Return type

[**Preferences**](preferences.md)

### Authorization

[okta_jwt](../README.md#okta_jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTeam

> Team UpdateTeam(ctx, team)



Update properties of team

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**team** | **string**|  | 

### Return type

[**Team**](team.md)

### Authorization

[okta_jwt](../README.md#okta_jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

