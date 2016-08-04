# \DefaultApi

All URIs are relative to *https://www.orkiv.com/i/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AllGet**](DefaultApi.md#AllGet) | **Get** /all/ | 
[**CategoriesDelete**](DefaultApi.md#CategoriesDelete) | **Delete** /categories/ | 
[**CategoriesPost**](DefaultApi.md#CategoriesPost) | **Post** /categories/ | 
[**CategoriesPut**](DefaultApi.md#CategoriesPut) | **Put** /categories/ | 
[**ItemAddPost**](DefaultApi.md#ItemAddPost) | **Post** /item/add/ | 
[**ItemAddbulkPost**](DefaultApi.md#ItemAddbulkPost) | **Post** /item/addbulk/ | 
[**ItemDelete**](DefaultApi.md#ItemDelete) | **Delete** /item/ | 
[**ItemPut**](DefaultApi.md#ItemPut) | **Put** /item/ | 
[**ItemsCountPost**](DefaultApi.md#ItemsCountPost) | **Post** /items/count/ | 
[**ItemsPost**](DefaultApi.md#ItemsPost) | **Post** /items/ | 
[**ItemsallfieldsPost**](DefaultApi.md#ItemsallfieldsPost) | **Post** /items/?allfields | 
[**OrdersPost**](DefaultApi.md#OrdersPost) | **Post** /orders/ | 
[**QueryPost**](DefaultApi.md#QueryPost) | **Post** /query/ | 
[**QueryallfieldsPost**](DefaultApi.md#QueryallfieldsPost) | **Post** /query/?allfields | 
[**ServicesDelete**](DefaultApi.md#ServicesDelete) | **Delete** /services/ | 
[**ServicesGet**](DefaultApi.md#ServicesGet) | **Get** /services/ | 
[**ServicesPost**](DefaultApi.md#ServicesPost) | **Post** /services/ | 
[**ServicesPut**](DefaultApi.md#ServicesPut) | **Put** /services/ | 
[**WriteDelete**](DefaultApi.md#WriteDelete) | **Delete** /write/ | 
[**WritePost**](DefaultApi.md#WritePost) | **Post** /write/ | 


# **AllGet**
> []InventoryGroup AllGet()




### Parameters
This endpoint does not need any parameter.

### Return type

[**[]InventoryGroup**](InventoryGroup.md)

### Authorization

[APIKey](../README.md#APIKey), [AccountID](../README.md#AccountID)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CategoriesDelete**
> Response CategoriesDelete($id)




### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| Id of category to remove | 

### Return type

[**Response**](Response.md)

### Authorization

[APIKey](../README.md#APIKey), [AccountID](../README.md#AccountID)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CategoriesPost**
> []Category CategoriesPost($query)




### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | [**Dictionary**](Dictionary.md)| Category to query against system | [optional] 

### Return type

[**[]Category**](Category.md)

### Authorization

[APIKey](../README.md#APIKey), [AccountID](../README.md#AccountID)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CategoriesPut**
> Category CategoriesPut($id, $category)



If no ID is specified a new category will be created!


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| category id to update. | 
 **category** | [**Category**](Category.md)| New category information. | 

### Return type

[**Category**](Category.md)

### Authorization

[APIKey](../README.md#APIKey), [AccountID](../README.md#AccountID)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ItemAddPost**
> Item ItemAddPost($item)




### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **item** | [**Item**](Item.md)| Item to create. | 

### Return type

[**Item**](Item.md)

### Authorization

[APIKey](../README.md#APIKey), [AccountID](../README.md#AccountID)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ItemAddbulkPost**
> Response ItemAddbulkPost($items)




### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **items** | [**[]Item**](Item.md)| Items to create. | 

### Return type

[**Response**](Response.md)

### Authorization

[APIKey](../README.md#APIKey), [AccountID](../README.md#AccountID)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ItemDelete**
> Response ItemDelete($id)




### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| item id to remove | 

### Return type

[**Response**](Response.md)

### Authorization

[APIKey](../README.md#APIKey), [AccountID](../README.md#AccountID)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ItemPut**
> Response ItemPut($id, $item)




### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| item id to update. | 
 **item** | [**Dictionary**](Dictionary.md)| New item information. | 

### Return type

[**Response**](Response.md)

### Authorization

[APIKey](../README.md#APIKey), [AccountID](../README.md#AccountID)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ItemsCountPost**
> float32 ItemsCountPost($query)




### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | [**Dictionary**](Dictionary.md)| Item to query against system. | [optional] 

### Return type

**float32**

### Authorization

[APIKey](../README.md#APIKey), [AccountID](../README.md#AccountID)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ItemsPost**
> []Item ItemsPost($query)




### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | [**Dictionary**](Dictionary.md)| Item to query against system. | [optional] 

### Return type

[**[]Item**](Item.md)

### Authorization

[APIKey](../README.md#APIKey), [AccountID](../README.md#AccountID)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ItemsallfieldsPost**
> []Dictionary ItemsallfieldsPost($query)




### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | [**Dictionary**](Dictionary.md)| Item to query against system. | [optional] 

### Return type

[**[]Dictionary**](Dictionary.md)

### Authorization

[APIKey](../README.md#APIKey), [AccountID](../README.md#AccountID)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrdersPost**
> []Order OrdersPost($query)




### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | [**Dictionary**](Dictionary.md)| Order to query against system. | [optional] 

### Return type

[**[]Order**](Order.md)

### Authorization

[APIKey](../README.md#APIKey), [AccountID](../README.md#AccountID)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **QueryPost**
> []Item QueryPost($page, $categoryid, $sort, $search, $minprice, $maxprice, $query)




### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **float32**| Current page index. | [optional] 
 **categoryid** | **string**| Get items under specified category id. | [optional] 
 **sort** | **string**| Comma delimited Sort string. ie ; +ordprice. Please use number based fields only | [optional] 
 **search** | **string**| Performs a regex pattern match against the items within your account | [optional] 
 **minprice** | **float32**| Min price in hundreds. | [optional] 
 **maxprice** | **float32**| Max price in hudreds. | [optional] 
 **query** | [**Dictionary**](Dictionary.md)| Custom parameters to query against system. | [optional] 

### Return type

[**[]Item**](Item.md)

### Authorization

[APIKey](../README.md#APIKey), [AccountID](../README.md#AccountID)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **QueryallfieldsPost**
> []Dictionary QueryallfieldsPost($page, $categoryid, $sort, $search, $minprice, $maxprice, $query)




### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **float32**| Current page index. | [optional] 
 **categoryid** | **string**| Get items under specified category id. | [optional] 
 **sort** | **string**| Comma delimited Sort string. ie ; +ordprice. Please use number based fields only | [optional] 
 **search** | **string**| Performs a regex pattern match against the items within your account | [optional] 
 **minprice** | **float32**| Min price in hundreds. | [optional] 
 **maxprice** | **float32**| Max price in hudreds. | [optional] 
 **query** | [**Dictionary**](Dictionary.md)| Custom parameters to query against system. | [optional] 

### Return type

[**[]Dictionary**](Dictionary.md)

### Authorization

[APIKey](../README.md#APIKey), [AccountID](../README.md#AccountID)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ServicesDelete**
> Response ServicesDelete($id)




### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| ID of the service to update | 

### Return type

[**Response**](Response.md)

### Authorization

[APIKey](../README.md#APIKey), [AccountID](../README.md#AccountID)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ServicesGet**
> []Service ServicesGet()




### Parameters
This endpoint does not need any parameter.

### Return type

[**[]Service**](Service.md)

### Authorization

[APIKey](../README.md#APIKey), [AccountID](../README.md#AccountID)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ServicesPost**
> Service ServicesPost($service)




### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **service** | [**Service**](Service.md)| Service to create. | 

### Return type

[**Service**](Service.md)

### Authorization

[APIKey](../README.md#APIKey), [AccountID](../README.md#AccountID)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ServicesPut**
> Response ServicesPut($id, $service)




### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| ID of the service to update | 
 **service** | [**Service**](Service.md)| New service data to set. | 

### Return type

[**Response**](Response.md)

### Authorization

[APIKey](../README.md#APIKey), [AccountID](../README.md#AccountID)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WriteDelete**
> Response WriteDelete($id)




### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| Will delete event attached to this serviceid | [optional] 

### Return type

[**Response**](Response.md)

### Authorization

[APIKey](../README.md#APIKey), [AccountID](../README.md#AccountID)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WritePost**
> Response WritePost($eventRequest)



Will ovveride the current event of the specified service.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **eventRequest** | [**EventRequest**](EventRequest.md)| Event to upload | 

### Return type

[**Response**](Response.md)

### Authorization

[APIKey](../README.md#APIKey), [AccountID](../README.md#AccountID)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

