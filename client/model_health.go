/*
 * Redfish OAPI specification
 *
 * Partial Redfish OAPI specification for a limited client
 *
 * API version: 0.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package client
type Health string

// List of Health
const (
	HEALTH_OK Health = "OK"
	HEALTH_WARNING Health = "Warning"
	HEALTH_CRITICAL Health = "Critical"
)