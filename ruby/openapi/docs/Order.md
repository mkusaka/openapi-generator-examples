# OpenapiClient::Order

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **Integer** |  | [optional] |
| **pet_id** | **Integer** |  | [optional] |
| **quantity** | **Integer** |  | [optional] |
| **ship_date** | **Time** |  | [optional] |
| **status** | **String** | Order Status | [optional] |
| **complete** | **Boolean** |  | [optional] |

## Example

```ruby
require 'openapi_client'

instance = OpenapiClient::Order.new(
  id: 10,
  pet_id: 198772,
  quantity: 7,
  ship_date: null,
  status: approved,
  complete: null
)
```

