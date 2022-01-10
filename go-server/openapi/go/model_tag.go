/*
 * Swagger Petstore - OpenAPI 3.0
 *
 * This is a sample Pet Store Server based on the OpenAPI 3.0 specification.  You can find out more about Swagger at [http://swagger.io](http://swagger.io). In the third iteration of the pet store, we've switched to the design first approach! You can now help us improve the API whether it's by making changes to the definition itself or to the code. That way, with time, we can improve the API in general, and expose some of the new features in OAS3.  Some useful links: - [The Pet Store repository](https://github.com/swagger-api/swagger-petstore) - [The source API definition for the Pet Store](https://github.com/swagger-api/swagger-petstore/blob/master/src/main/resources/openapi.yaml)
 *
 * API version: 1.0.9
 * Contact: apiteam@swagger.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type Tag struct {

	Id int64 `json:"id,omitempty"`

	Name string `json:"name,omitempty"`
}

// AssertTagRequired checks if the required fields are not zero-ed
func AssertTagRequired(obj Tag) error {
	return nil
}

// AssertRecurseTagRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of Tag (e.g. [][]Tag), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseTagRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aTag, ok := obj.(Tag)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertTagRequired(aTag)
	})
}
