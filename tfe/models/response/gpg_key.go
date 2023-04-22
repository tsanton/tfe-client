package response

import (
	"time"
)

type GpgKey struct {
	Data GpgKeyData `json:"data"`
}

type GpgKeyData struct {
	Type       string              `json:"type"`
	Id         string              `json:"id"`
	Attributes GpgKeyDataAttribute `json:"attributes"`
	Links      GpgKeyDataLinks     `json:"links"`
}

type GpgKeyDataAttribute struct {
	AsciiArmor     string    `json:"ascii-armor"`
	CreatedAt      time.Time `json:"created-at"`
	KeyId          string    `json:"key-id"`
	Namespace      string    `json:"namespace"`
	Source         string    `json:"source"`
	SourceUrl      string    `json:"source-url"`
	TrustSignature string    `json:"trust-signature"`
	UpdatedAt      time.Time `json:"updated-at"`
}

type GpgKeyDataLinks struct {
	Self string `json:"self"`
}
