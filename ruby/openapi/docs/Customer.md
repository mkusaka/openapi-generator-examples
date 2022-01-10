# OpenapiClient::Customer

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **Integer** |  | [optional] |
| **username** | **String** |  | [optional] |
| **address** | [**Array&lt;Address&gt;**](Address.md) |  | [optional] |

## Example

```ruby
require 'openapi_client'

instance = OpenapiClient::Customer.new(
  id: 100000,
  username: fehguy,
  address: null
)
```

