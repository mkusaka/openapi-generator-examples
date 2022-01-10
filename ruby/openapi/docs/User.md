# OpenapiClient::User

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **Integer** |  | [optional] |
| **username** | **String** |  | [optional] |
| **first_name** | **String** |  | [optional] |
| **last_name** | **String** |  | [optional] |
| **email** | **String** |  | [optional] |
| **password** | **String** |  | [optional] |
| **phone** | **String** |  | [optional] |
| **user_status** | **Integer** | User Status | [optional] |

## Example

```ruby
require 'openapi_client'

instance = OpenapiClient::User.new(
  id: 10,
  username: theUser,
  first_name: John,
  last_name: James,
  email: john@email.com,
  password: 12345,
  phone: 12345,
  user_status: 1
)
```

