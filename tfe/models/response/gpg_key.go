package response

import (
	"time"
)

type GpgKey struct {
	Data gpgKeyData `json:"data"`
}

type gpgKeyData struct {
	Type       string              `json:"type"`
	Id         string              `json:"id"`
	Attributes gpgKeyDataAttribute `json:"attributes"`
	Links      gpgKeyDataLinks     `json:"links"`
}

type gpgKeyDataAttribute struct {
	AsciiArmor     string    `json:"ascii-armor"`
	CreatedAt      time.Time `json:"created-at"`
	KeyId          string    `json:"key-id"`
	Namespace      string    `json:"namespace"`
	Source         string    `json:"source"`
	SourceUrl      string    `json:"source-url"`
	TrustSignature string    `json:"trust-signature"`
	UpdatedAt      time.Time `json:"updated-at"`
}

type gpgKeyDataLinks struct {
	Self string `json:"self"`
}
