/*
 * examples/internal/proto/examplepb/response_body_service.proto
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package responsebody

type RuntimeStreamError struct {
	GrpcCode int32 `json:"grpc_code,omitempty"`
	HttpCode int32 `json:"http_code,omitempty"`
	Message string `json:"message,omitempty"`
	HttpStatus string `json:"http_status,omitempty"`
	Details []ProtobufAny `json:"details,omitempty"`
}