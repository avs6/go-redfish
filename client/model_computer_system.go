/*
 * Redfish OAPI specification
 *
 * Partial Redfish OAPI specification for a limited client
 *
 * API version: 0.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package client

// Root redfish path.
type ComputerSystem struct {
	// The name of the resource.
	Id string `json:"Id,omitempty"`
	// The name of the resource.
	Name string `json:"Name"`
	// redfish version
	RedfishVersion string `json:"RedfishVersion,omitempty"`
	UUID string `json:"UUID,omitempty"`
	// The type of a resource.
	OdataType string `json:"@odata.type"`
	// The unique identifier for a resource.
	OdataId string `json:"@odata.id"`
	// The OData description of a payload.
	OdataContext string `json:"@odata.context,omitempty"`
	// redfish copyright
	RedfishCopyright string `json:"@Redfish.Copyright,omitempty"`
	Bios IdRef `json:"Bios,omitempty"`
	Processors IdRef `json:"Processors,omitempty"`
	Memory IdRef `json:"Memory,omitempty"`
	EthernetInterfaces IdRef `json:"EthernetInterfaces,omitempty"`
	SimpleStorage IdRef `json:"SimpleStorage,omitempty"`
	PowerState PowerState `json:"PowerState,omitempty"`
	Status Status `json:"Status,omitempty"`
	Boot Boot `json:"Boot,omitempty"`
	ProcessorSummary ProcessorSummary `json:"ProcessorSummary,omitempty"`
	MemorySummary MemorySummary `json:"MemorySummary,omitempty"`
	IndicatorLED IndicatorLed `json:"IndicatorLED,omitempty"`
	Links SystemLinks `json:"Links,omitempty"`
	Actions ComputerSystemActions `json:"Actions,omitempty"`
}
