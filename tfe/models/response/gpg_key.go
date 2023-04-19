package response

import (
	"time"
)

type GpgKey struct {
	Data struct {
		Id         string `json:"id"`
		Type       string `json:"type"`
		Attributes struct {
			AsciiArmor     string    `json:"ascii-armor"`
			KeyId          string    `json:"key-id"`
			Namespace      string    `json:"namespace"`
			Source         string    `json:"source"`
			SourceUrl      string    `json:"source-url"`
			TrustSignature string    `json:"trust-signature"`
			Created        time.Time `json:"created-at"`
			Updated        time.Time `json:"updated-at"`
		} `json:"attributes"`
	} `json:"data"`
	Links struct {
		Self string `json:"self"`
	} `json:"links"`
}
