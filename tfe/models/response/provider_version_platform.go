package response

type ProviderVersionPlatform struct {
	Data ProviderVersionPlatformData `json:"data"`
}

type ProviderVersionPlatformData struct {
	Id            string                                   `json:"id"`
	Type          string                                   `json:"type"`
	Attributes    ProviderVersionPlatformDataAttributes    `json:"attributes"`
	Relationships ProviderVersionPlatformDataRelationships `json:"relationships"`
	Links         ProviderVersionRelationshipPlatformLinks `json:"links"`
}

type ProviderVersionPlatformDataAttributes struct {
	Os                     string                                           `json:"os"`
	Arch                   string                                           `json:"arch"`
	Filename               string                                           `json:"filename"`
	Shasum                 string                                           `json:"shasum"`
	Permissions            ProviderVersionPlatformDataAttributesPermissions `json:"permissions"`
	ProviderBinaryUploaded bool                                             `json:"provider-binary-uploaded"`
}

type ProviderVersionPlatformDataAttributesPermissions struct {
	CanDelete      bool `json:"can-delete"`
	CanUploadAsset bool `json:"can-upload-asset"`
}

type ProviderVersionPlatformDataRelationships struct {
	RegistryProviderVersion ProviderVersionPlatformVersion `json:"registry-provider-version"`
}

type ProviderVersionPlatformVersion struct {
	Data ProviderVersionPlatformVersionData `json:"data"`
}

type ProviderVersionPlatformVersionData struct {
	Id   string `json:"id"`
	Type string `json:"type"`
}

type ProviderVersionRelationshipPlatformLinks struct {
	ProviderBinaryUpload string `json:"provider-binary-upload"`
}
