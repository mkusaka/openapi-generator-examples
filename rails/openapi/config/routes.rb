=begin
Swagger Petstore - OpenAPI 3.0

This is a sample Pet Store Server based on the OpenAPI 3.0 specification.  You can find out more about Swagger at [http://swagger.io](http://swagger.io). In the third iteration of the pet store, we've switched to the design first approach! You can now help us improve the API whether it's by making changes to the definition itself or to the code. That way, with time, we can improve the API in general, and expose some of the new features in OAS3.  Some useful links: - [The Pet Store repository](https://github.com/swagger-api/swagger-petstore) - [The source API definition for the Pet Store](https://github.com/swagger-api/swagger-petstore/blob/master/src/main/resources/openapi.yaml)

The version of the OpenAPI document: 1.0.9
Contact: apiteam@swagger.io
Generated by: https://github.com/openapitools/openapi-generator.git

=end
Rails.application.routes.draw do

  def add_openapi_route http_method, path, opts = {}
    full_path = path.gsub(/{(.*?)}/, ':\1')
    match full_path, to: "#{opts.fetch(:controller_name)}##{opts[:action_name]}", via: http_method
  end

  add_openapi_route 'POST', '/api/v3/api/d/pet', controller_name: 'pet', action_name: 'add_pet'
  add_openapi_route 'DELETE', '/api/v3/pet/{petId}', controller_name: 'pet', action_name: 'destroy'
  add_openapi_route 'GET', '/api/v3/pet/findByStatus', controller_name: 'pet', action_name: 'find_pets_by_status'
  add_openapi_route 'GET', '/api/v3/pet/findByTags', controller_name: 'pet', action_name: 'find_pets_by_tags'
  add_openapi_route 'GET', '/api/v3/pet/{petId}', controller_name: 'pet', action_name: 'show'
  add_openapi_route 'PUT', '/api/v3/api/d/pet', controller_name: 'pet', action_name: 'update_pet'
  add_openapi_route 'POST', '/api/v3/pet/{petId}', controller_name: 'pet', action_name: 'update_pet_with_form'
  add_openapi_route 'POST', '/api/v3/pet/{petId}/uploadImage', controller_name: 'pet', action_name: 'upload_file'
  add_openapi_route 'DELETE', '/api/v3/store/order/{orderId}', controller_name: 'store', action_name: 'delete_order'
  add_openapi_route 'GET', '/api/v3/store/inventory', controller_name: 'store', action_name: 'get_inventory'
  add_openapi_route 'GET', '/api/v3/store/order/{orderId}', controller_name: 'store', action_name: 'get_order_by_id'
  add_openapi_route 'POST', '/api/v3/store/order', controller_name: 'store', action_name: 'place_order'
  add_openapi_route 'POST', '/api/v3/user', controller_name: 'user', action_name: 'create'
  add_openapi_route 'POST', '/api/v3/user/createWithList', controller_name: 'user', action_name: 'create_users_with_list_input'
  add_openapi_route 'DELETE', '/api/v3/user/{username}', controller_name: 'user', action_name: 'destroy'
  add_openapi_route 'GET', '/api/v3/user/{username}', controller_name: 'user', action_name: 'show'
  add_openapi_route 'GET', '/api/v3/user/login', controller_name: 'user', action_name: 'login_user'
  add_openapi_route 'GET', '/api/v3/user/logout', controller_name: 'user', action_name: 'logout_user'
  add_openapi_route 'PUT', '/api/v3/user/{username}', controller_name: 'user', action_name: 'update'
end
