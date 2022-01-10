# OpenapiClient::PetApi

All URIs are relative to */api/v3*

| Method | HTTP request | Description |
| ------ | ------------ | ----------- |
| [**add_pet**](PetApi.md#add_pet) | **POST** /api/d/pet | Add a new pet to the store |
| [**delete_pet**](PetApi.md#delete_pet) | **DELETE** /pet/{petId} | Deletes a pet |
| [**find_pets_by_status**](PetApi.md#find_pets_by_status) | **GET** /pet/findByStatus | Finds Pets by status |
| [**find_pets_by_tags**](PetApi.md#find_pets_by_tags) | **GET** /pet/findByTags | Finds Pets by tags |
| [**get_pet_by_id**](PetApi.md#get_pet_by_id) | **GET** /pet/{petId} | Find pet by ID |
| [**update_pet**](PetApi.md#update_pet) | **PUT** /api/d/pet | Update an existing pet |
| [**update_pet_with_form**](PetApi.md#update_pet_with_form) | **POST** /pet/{petId} | Updates a pet in the store with form data |
| [**upload_file**](PetApi.md#upload_file) | **POST** /pet/{petId}/uploadImage | uploads an image |


## add_pet

> <Pet> add_pet(pet)

Add a new pet to the store

Add a new pet to the store

### Examples

```ruby
require 'time'
require 'openapi_client'
# setup authorization
OpenapiClient.configure do |config|
  # Configure OAuth2 access token for authorization: petstore_auth
  config.access_token = 'YOUR ACCESS TOKEN'
end

api_instance = OpenapiClient::PetApi.new
pet = OpenapiClient::Pet.new({name: 'doggie', photo_urls: ['photo_urls_example']}) # Pet | Create a new pet in the store

begin
  # Add a new pet to the store
  result = api_instance.add_pet(pet)
  p result
rescue OpenapiClient::ApiError => e
  puts "Error when calling PetApi->add_pet: #{e}"
end
```

#### Using the add_pet_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<Pet>, Integer, Hash)> add_pet_with_http_info(pet)

```ruby
begin
  # Add a new pet to the store
  data, status_code, headers = api_instance.add_pet_with_http_info(pet)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <Pet>
rescue OpenapiClient::ApiError => e
  puts "Error when calling PetApi->add_pet_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **pet** | [**Pet**](Pet.md) | Create a new pet in the store |  |

### Return type

[**Pet**](Pet.md)

### Authorization

[petstore_auth](../README.md#petstore_auth)

### HTTP request headers

- **Content-Type**: application/json, application/xml, application/x-www-form-urlencoded
- **Accept**: application/xml, application/json


## delete_pet

> delete_pet(pet_id, opts)

Deletes a pet

### Examples

```ruby
require 'time'
require 'openapi_client'
# setup authorization
OpenapiClient.configure do |config|
  # Configure OAuth2 access token for authorization: petstore_auth
  config.access_token = 'YOUR ACCESS TOKEN'
end

api_instance = OpenapiClient::PetApi.new
pet_id = 789 # Integer | Pet id to delete
opts = {
  api_key: 'api_key_example' # String | 
}

begin
  # Deletes a pet
  api_instance.delete_pet(pet_id, opts)
rescue OpenapiClient::ApiError => e
  puts "Error when calling PetApi->delete_pet: #{e}"
end
```

#### Using the delete_pet_with_http_info variant

This returns an Array which contains the response data (`nil` in this case), status code and headers.

> <Array(nil, Integer, Hash)> delete_pet_with_http_info(pet_id, opts)

```ruby
begin
  # Deletes a pet
  data, status_code, headers = api_instance.delete_pet_with_http_info(pet_id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => nil
rescue OpenapiClient::ApiError => e
  puts "Error when calling PetApi->delete_pet_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **pet_id** | **Integer** | Pet id to delete |  |
| **api_key** | **String** |  | [optional] |

### Return type

nil (empty response body)

### Authorization

[petstore_auth](../README.md#petstore_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined


## find_pets_by_status

> <Array<Pet>> find_pets_by_status(opts)

Finds Pets by status

Multiple status values can be provided with comma separated strings

### Examples

```ruby
require 'time'
require 'openapi_client'
# setup authorization
OpenapiClient.configure do |config|
  # Configure OAuth2 access token for authorization: petstore_auth
  config.access_token = 'YOUR ACCESS TOKEN'
end

api_instance = OpenapiClient::PetApi.new
opts = {
  status: 'available' # String | Status values that need to be considered for filter
}

begin
  # Finds Pets by status
  result = api_instance.find_pets_by_status(opts)
  p result
rescue OpenapiClient::ApiError => e
  puts "Error when calling PetApi->find_pets_by_status: #{e}"
end
```

#### Using the find_pets_by_status_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<Array<Pet>>, Integer, Hash)> find_pets_by_status_with_http_info(opts)

```ruby
begin
  # Finds Pets by status
  data, status_code, headers = api_instance.find_pets_by_status_with_http_info(opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <Array<Pet>>
rescue OpenapiClient::ApiError => e
  puts "Error when calling PetApi->find_pets_by_status_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **status** | **String** | Status values that need to be considered for filter | [optional][default to &#39;available&#39;] |

### Return type

[**Array&lt;Pet&gt;**](Pet.md)

### Authorization

[petstore_auth](../README.md#petstore_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/xml, application/json


## find_pets_by_tags

> <Array<Pet>> find_pets_by_tags(opts)

Finds Pets by tags

Multiple tags can be provided with comma separated strings. Use tag1, tag2, tag3 for testing.

### Examples

```ruby
require 'time'
require 'openapi_client'
# setup authorization
OpenapiClient.configure do |config|
  # Configure OAuth2 access token for authorization: petstore_auth
  config.access_token = 'YOUR ACCESS TOKEN'
end

api_instance = OpenapiClient::PetApi.new
opts = {
  tags: ['inner_example'] # Array<String> | Tags to filter by
}

begin
  # Finds Pets by tags
  result = api_instance.find_pets_by_tags(opts)
  p result
rescue OpenapiClient::ApiError => e
  puts "Error when calling PetApi->find_pets_by_tags: #{e}"
end
```

#### Using the find_pets_by_tags_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<Array<Pet>>, Integer, Hash)> find_pets_by_tags_with_http_info(opts)

```ruby
begin
  # Finds Pets by tags
  data, status_code, headers = api_instance.find_pets_by_tags_with_http_info(opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <Array<Pet>>
rescue OpenapiClient::ApiError => e
  puts "Error when calling PetApi->find_pets_by_tags_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **tags** | [**Array&lt;String&gt;**](String.md) | Tags to filter by | [optional] |

### Return type

[**Array&lt;Pet&gt;**](Pet.md)

### Authorization

[petstore_auth](../README.md#petstore_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/xml, application/json


## get_pet_by_id

> <Pet> get_pet_by_id(pet_id)

Find pet by ID

Returns a single pet

### Examples

```ruby
require 'time'
require 'openapi_client'
# setup authorization
OpenapiClient.configure do |config|
  # Configure API key authorization: api_key
  config.api_key['api_key'] = 'YOUR API KEY'
  # Uncomment the following line to set a prefix for the API key, e.g. 'Bearer' (defaults to nil)
  # config.api_key_prefix['api_key'] = 'Bearer'

  # Configure OAuth2 access token for authorization: petstore_auth
  config.access_token = 'YOUR ACCESS TOKEN'
end

api_instance = OpenapiClient::PetApi.new
pet_id = 789 # Integer | ID of pet to return

begin
  # Find pet by ID
  result = api_instance.get_pet_by_id(pet_id)
  p result
rescue OpenapiClient::ApiError => e
  puts "Error when calling PetApi->get_pet_by_id: #{e}"
end
```

#### Using the get_pet_by_id_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<Pet>, Integer, Hash)> get_pet_by_id_with_http_info(pet_id)

```ruby
begin
  # Find pet by ID
  data, status_code, headers = api_instance.get_pet_by_id_with_http_info(pet_id)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <Pet>
rescue OpenapiClient::ApiError => e
  puts "Error when calling PetApi->get_pet_by_id_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **pet_id** | **Integer** | ID of pet to return |  |

### Return type

[**Pet**](Pet.md)

### Authorization

[api_key](../README.md#api_key), [petstore_auth](../README.md#petstore_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/xml, application/json


## update_pet

> <Pet> update_pet(pet)

Update an existing pet

Update an existing pet by Id

### Examples

```ruby
require 'time'
require 'openapi_client'
# setup authorization
OpenapiClient.configure do |config|
  # Configure OAuth2 access token for authorization: petstore_auth
  config.access_token = 'YOUR ACCESS TOKEN'
end

api_instance = OpenapiClient::PetApi.new
pet = OpenapiClient::Pet.new({name: 'doggie', photo_urls: ['photo_urls_example']}) # Pet | Update an existent pet in the store

begin
  # Update an existing pet
  result = api_instance.update_pet(pet)
  p result
rescue OpenapiClient::ApiError => e
  puts "Error when calling PetApi->update_pet: #{e}"
end
```

#### Using the update_pet_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<Pet>, Integer, Hash)> update_pet_with_http_info(pet)

```ruby
begin
  # Update an existing pet
  data, status_code, headers = api_instance.update_pet_with_http_info(pet)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <Pet>
rescue OpenapiClient::ApiError => e
  puts "Error when calling PetApi->update_pet_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **pet** | [**Pet**](Pet.md) | Update an existent pet in the store |  |

### Return type

[**Pet**](Pet.md)

### Authorization

[petstore_auth](../README.md#petstore_auth)

### HTTP request headers

- **Content-Type**: application/json, application/xml, application/x-www-form-urlencoded
- **Accept**: application/xml, application/json


## update_pet_with_form

> update_pet_with_form(pet_id, opts)

Updates a pet in the store with form data

### Examples

```ruby
require 'time'
require 'openapi_client'
# setup authorization
OpenapiClient.configure do |config|
  # Configure OAuth2 access token for authorization: petstore_auth
  config.access_token = 'YOUR ACCESS TOKEN'
end

api_instance = OpenapiClient::PetApi.new
pet_id = 789 # Integer | ID of pet that needs to be updated
opts = {
  name: 'name_example', # String | Name of pet that needs to be updated
  status: 'status_example' # String | Status of pet that needs to be updated
}

begin
  # Updates a pet in the store with form data
  api_instance.update_pet_with_form(pet_id, opts)
rescue OpenapiClient::ApiError => e
  puts "Error when calling PetApi->update_pet_with_form: #{e}"
end
```

#### Using the update_pet_with_form_with_http_info variant

This returns an Array which contains the response data (`nil` in this case), status code and headers.

> <Array(nil, Integer, Hash)> update_pet_with_form_with_http_info(pet_id, opts)

```ruby
begin
  # Updates a pet in the store with form data
  data, status_code, headers = api_instance.update_pet_with_form_with_http_info(pet_id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => nil
rescue OpenapiClient::ApiError => e
  puts "Error when calling PetApi->update_pet_with_form_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **pet_id** | **Integer** | ID of pet that needs to be updated |  |
| **name** | **String** | Name of pet that needs to be updated | [optional] |
| **status** | **String** | Status of pet that needs to be updated | [optional] |

### Return type

nil (empty response body)

### Authorization

[petstore_auth](../README.md#petstore_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined


## upload_file

> <ApiResponse> upload_file(pet_id, opts)

uploads an image

### Examples

```ruby
require 'time'
require 'openapi_client'
# setup authorization
OpenapiClient.configure do |config|
  # Configure OAuth2 access token for authorization: petstore_auth
  config.access_token = 'YOUR ACCESS TOKEN'
end

api_instance = OpenapiClient::PetApi.new
pet_id = 789 # Integer | ID of pet to update
opts = {
  additional_metadata: 'additional_metadata_example', # String | Additional Metadata
  body: File.new('/path/to/some/file') # File | 
}

begin
  # uploads an image
  result = api_instance.upload_file(pet_id, opts)
  p result
rescue OpenapiClient::ApiError => e
  puts "Error when calling PetApi->upload_file: #{e}"
end
```

#### Using the upload_file_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<ApiResponse>, Integer, Hash)> upload_file_with_http_info(pet_id, opts)

```ruby
begin
  # uploads an image
  data, status_code, headers = api_instance.upload_file_with_http_info(pet_id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <ApiResponse>
rescue OpenapiClient::ApiError => e
  puts "Error when calling PetApi->upload_file_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **pet_id** | **Integer** | ID of pet to update |  |
| **additional_metadata** | **String** | Additional Metadata | [optional] |
| **body** | **File** |  | [optional] |

### Return type

[**ApiResponse**](ApiResponse.md)

### Authorization

[petstore_auth](../README.md#petstore_auth)

### HTTP request headers

- **Content-Type**: application/octet-stream
- **Accept**: application/json

