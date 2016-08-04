# Order

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderId** | **string** | Order ID | [optional] [default to null]
**InfoEmail** | **string** | Customer email | [optional] [default to null]
**InfoFirst** | **string** | Customer first name | [optional] [default to null]
**InfoLast** | **string** | Customer last name | [optional] [default to null]
**Phone** | **string** | Customer phone number | [optional] [default to null]
**Shipset** | **bool** | Customer billing address matches shipping address | [optional] [default to null]
**InfoAdr1** | **string** | Customer billing address line &#39;1&#39; | [optional] [default to null]
**InfoAdr2** | **string** | Customer billing address line &#39;2&#39; | [optional] [default to null]
**InfoCty** | **string** | Customer billing city | [optional] [default to null]
**InfoZip** | **string** | Customer billing zip code | [optional] [default to null]
**State** | **string** | Customer billing state | [optional] [default to null]
**InfoSadr1** | **string** | Customer shipping address line &#39;1&#39; | [optional] [default to null]
**InfoSadr2** | **string** | Customer shipping address line &#39;2&#39; | [optional] [default to null]
**InfoScty** | **string** | Customer shipping city | [optional] [default to null]
**InfoSzip** | **string** | Customer shipping zip code | [optional] [default to null]
**Sstate** | **string** | Customer shipping state | [optional] [default to null]
**TaxAmount** | **float32** | Tax amount in hundreds. Must divide by &#39;100&#39; for USD value | [optional] [default to null]
**ShippingAmount** | **float32** | Shipping total in USD | [optional] [default to null]
**AmountTotal** | **float32** | Checkout total in USD | [optional] [default to null]
**ItemIDs** | **[]string** | Array of items purchased at checkout | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


