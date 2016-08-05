# InventoryClient\DefaultApi

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
[**ItemGet**](DefaultApi.md#ItemGet) | **Get** /item/ | 
[**ItemMediaDelete**](DefaultApi.md#ItemMediaDelete) | **Delete** /item-media/ | 
[**ItemMediaPost**](DefaultApi.md#ItemMediaPost) | **Post** /item-media/ | 
[**ItemPut**](DefaultApi.md#ItemPut) | **Put** /item/ | 
[**ItemsCountPost**](DefaultApi.md#ItemsCountPost) | **Post** /items/count/ | 
[**ItemsPost**](DefaultApi.md#ItemsPost) | **Post** /items/ | 
[**OrdersPost**](DefaultApi.md#OrdersPost) | **Post** /orders/ | 
[**OrdersServicesPost**](DefaultApi.md#OrdersServicesPost) | **Post** /orders/services/ | 
[**QueryPost**](DefaultApi.md#QueryPost) | **Post** /query/ | 
[**ServicesDelete**](DefaultApi.md#ServicesDelete) | **Delete** /services/ | 
[**ServicesGet**](DefaultApi.md#ServicesGet) | **Get** /services/ | 
[**ServicesOpenGet**](DefaultApi.md#ServicesOpenGet) | **Get** /services/open/ | 
[**ServicesPost**](DefaultApi.md#ServicesPost) | **Post** /services/ | 
[**ServicesPut**](DefaultApi.md#ServicesPut) | **Put** /services/ | 
[**VariationDelete**](DefaultApi.md#VariationDelete) | **Delete** /variation/ | 
[**VariationGet**](DefaultApi.md#VariationGet) | **Get** /variation/ | 
[**VariationPost**](DefaultApi.md#VariationPost) | **Post** /variation/ | 
[**VariationPut**](DefaultApi.md#VariationPut) | **Put** /variation/ | 
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
 **query** | [**Category**](Category.md)| Category to query against system | [optional] 

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
 **item** | [**ItemRequest**](ItemRequest.md)| Item to create. | 

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
 **items** | [**[]ItemRequest**](ItemRequest.md)| Items to create. | 

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

# **ItemGet**
> Item ItemGet($id)




### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| Item ID to open. | 

### Return type

[**Item**](Item.md)

### Authorization

[APIKey](../README.md#APIKey), [AccountID](../README.md#AccountID)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ItemMediaDelete**
> Response ItemMediaDelete($imageurl)




### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **imageurl** | **string**| URL of image to remove | 

### Return type

[**Response**](Response.md)

### Authorization

[APIKey](../README.md#APIKey), [AccountID](../README.md#AccountID)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ItemMediaPost**
> string ItemMediaPost($id, $image)



This endpoint is currently in testing.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| Valid item id to bind image to. | 
 **image** | ***os.File**| Image. | 

### Return type

**string**

### Authorization

[APIKey](../README.md#APIKey), [AccountID](../README.md#AccountID)

### HTTP request headers

 - **Content-Type**: multipart/form-data, application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ItemPut**
> Response ItemPut($id, $item)




### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| item id to update. | 
 **item** | [**ItemRequest**](ItemRequest.md)| New item information. | 

### Return type

[**Response**](Response.md)

### Authorization

[APIKey](../README.md#APIKey), [AccountID](../README.md#AccountID)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ItemsCountPost**
> float32 ItemsCountPost($minprice, $maxprice, $query)




### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **minprice** | **float32**| Min price of items to find | [optional] 
 **maxprice** | **float32**| Max price of items to find | [optional] 
 **query** | [**ItemRequest**](ItemRequest.md)| Item to query against system. | [optional] 

### Return type

**float32**

### Authorization

[APIKey](../README.md#APIKey), [AccountID](../README.md#AccountID)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ItemsPost**
> []Item ItemsPost($minprice, $maxprice, $query)




### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **minprice** | **float32**| Min price of items to find | [optional] 
 **maxprice** | **float32**| Max price of items to find | [optional] 
 **query** | [**ItemRequest**](ItemRequest.md)| Item to query against system. | [optional] 

### Return type

[**[]Item**](Item.md)

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
 **query** | [**OrderRequest**](OrderRequest.md)| Order to query against item invoices. | [optional] 

### Return type

[**[]Order**](Order.md)

### Authorization

[APIKey](../README.md#APIKey), [AccountID](../README.md#AccountID)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrdersServicesPost**
> []Order OrdersServicesPost($query)




### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | [**OrderRequest**](OrderRequest.md)| Order to query against service invoices. | [optional] 

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
 **minprice** | **float32**| Min price in hundreds (cents). | [optional] 
 **maxprice** | **float32**| Max price in hundreds (cents). | [optional] 
 **query** | [**ItemRequest**](ItemRequest.md)| Custom parameters to query against system. | [optional] 

### Return type

[**[]Item**](Item.md)

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

# **ServicesOpenGet**
> Service ServicesOpenGet($id)




### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| ID of service to open | 

### Return type

[**Service**](Service.md)

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
 **service** | [**ServiceRequest**](ServiceRequest.md)| Service to create. | 

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
 **service** | [**ServiceRequest**](ServiceRequest.md)| New service data to set. | 

### Return type

[**Response**](Response.md)

### Authorization

[APIKey](../README.md#APIKey), [AccountID](../README.md#AccountID)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VariationDelete**
> Response VariationDelete($id)




### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| variation id to remove | 

### Return type

[**Response**](Response.md)

### Authorization

[APIKey](../README.md#APIKey), [AccountID](../README.md#AccountID)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VariationGet**
> Variation VariationGet($id)




### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| Variation ID to open. | 

### Return type

[**Variation**](Variation.md)

### Authorization

[APIKey](../README.md#APIKey), [AccountID](../README.md#AccountID)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VariationPost**
> Response VariationPost($id, $item)




### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| Valid item id to bind variation to. | 
 **item** | [**Variation**](Variation.md)| Variation information. | 

### Return type

[**Response**](Response.md)

### Authorization

[APIKey](../README.md#APIKey), [AccountID](../README.md#AccountID)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VariationPut**
> Response VariationPut($id, $item)




### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| variation id to update. | 
 **item** | [**Variation**](Variation.md)| New variation information. | 

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

