# \SatellitesApi

All URIs are relative to *https://virtserver.swaggerhub.com/universe-sh/Houston/1.0.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddPools**](SatellitesApi.md#AddPools) | **Post** /v1/satellites/{cloud}/{satellite}/pools | 
[**AddSatellites**](SatellitesApi.md#AddSatellites) | **Post** /v1/satellites/{cloud} | 
[**CreateMetrics**](SatellitesApi.md#CreateMetrics) | **Post** /v1/satellites/{cloud}/{satellite}/{pool}/metrics | 
[**DeletePool**](SatellitesApi.md#DeletePool) | **Delete** /v1/satellites/{cloud}/{satellite}/{pool} | 
[**DeleteSatellite**](SatellitesApi.md#DeleteSatellite) | **Delete** /v1/satellites/{cloud}/{satellite} | 
[**ListMetrics**](SatellitesApi.md#ListMetrics) | **Get** /v1/satellites/{cloud}/{satellite}/{pool}/metrics | 
[**ListPools**](SatellitesApi.md#ListPools) | **Get** /v1/satellites/{cloud}/{satellite}/pools | 
[**ListSatellites**](SatellitesApi.md#ListSatellites) | **Get** /v1/satellites/{cloud} | 
[**ReadPool**](SatellitesApi.md#ReadPool) | **Get** /v1/satellites/{cloud}/{satellite}/{pool} | 
[**ReadSatellite**](SatellitesApi.md#ReadSatellite) | **Get** /v1/satellites/{cloud}/{satellite} | 
[**UpdatePool**](SatellitesApi.md#UpdatePool) | **Put** /v1/satellites/{cloud}/{satellite}/{pool} | 
[**UpdateSatellite**](SatellitesApi.md#UpdateSatellite) | **Put** /v1/satellites/{cloud}/{satellite} | 


# **AddPools**
> Generic AddPools(ctx, cloud, satellite)


Add properties of pools

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cloud** | **string**|  | 
  **satellite** | **string**|  | 

### Return type

[**Generic**](generic.md)

### Authorization

[okta_jwt](../README.md#okta_jwt)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddSatellites**
> Generic AddSatellites(ctx, cloud)


Add properties of satellites

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cloud** | **string**|  | 

### Return type

[**Generic**](generic.md)

### Authorization

[okta_jwt](../README.md#okta_jwt)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateMetrics**
> Generic CreateMetrics(ctx, cloud, satellite, pool)


Create properties of metrics

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cloud** | **string**|  | 
  **satellite** | **string**|  | 
  **pool** | **string**|  | 

### Return type

[**Generic**](generic.md)

### Authorization

[okta_jwt](../README.md#okta_jwt)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeletePool**
> Generic DeletePool(ctx, cloud, satellite, pool)


Delete properties of pool

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cloud** | **string**|  | 
  **satellite** | **string**|  | 
  **pool** | **string**|  | 

### Return type

[**Generic**](generic.md)

### Authorization

[okta_jwt](../README.md#okta_jwt)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSatellite**
> Generic DeleteSatellite(ctx, cloud, satellite)


Delete properties of satellite

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cloud** | **string**|  | 
  **satellite** | **string**|  | 

### Return type

[**Generic**](generic.md)

### Authorization

[okta_jwt](../README.md#okta_jwt)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListMetrics**
> Metrics ListMetrics(ctx, cloud, satellite, pool)


List properties of metrics

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cloud** | **string**|  | 
  **satellite** | **string**|  | 
  **pool** | **string**|  | 

### Return type

[**Metrics**](metrics.md)

### Authorization

[okta_jwt](../README.md#okta_jwt)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListPools**
> Pools ListPools(ctx, cloud, satellite)


List properties of pools

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cloud** | **string**|  | 
  **satellite** | **string**|  | 

### Return type

[**Pools**](pools.md)

### Authorization

[okta_jwt](../README.md#okta_jwt)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListSatellites**
> Satellites ListSatellites(ctx, cloud)


List properties of satellites

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cloud** | **string**|  | 

### Return type

[**Satellites**](satellites.md)

### Authorization

[okta_jwt](../README.md#okta_jwt)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadPool**
> Pool ReadPool(ctx, cloud, satellite, pool)


Read properties of pool

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cloud** | **string**|  | 
  **satellite** | **string**|  | 
  **pool** | **string**|  | 

### Return type

[**Pool**](pool.md)

### Authorization

[okta_jwt](../README.md#okta_jwt)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadSatellite**
> Satellite ReadSatellite(ctx, cloud, satellite)


Read properties of satellite

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cloud** | **string**|  | 
  **satellite** | **string**|  | 

### Return type

[**Satellite**](satellite.md)

### Authorization

[okta_jwt](../README.md#okta_jwt)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdatePool**
> Pool UpdatePool(ctx, cloud, satellite, pool)


Update properties of pool

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cloud** | **string**|  | 
  **satellite** | **string**|  | 
  **pool** | **string**|  | 

### Return type

[**Pool**](pool.md)

### Authorization

[okta_jwt](../README.md#okta_jwt)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSatellite**
> Satellite UpdateSatellite(ctx, cloud, satellite)


Update properties of satellite

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cloud** | **string**|  | 
  **satellite** | **string**|  | 

### Return type

[**Satellite**](satellite.md)

### Authorization

[okta_jwt](../README.md#okta_jwt)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

