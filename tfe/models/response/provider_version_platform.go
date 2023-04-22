package response

type ProviderVersionPlatform struct {
	Data providerVersionPlatformData `json:"data"`
}

type providerVersionPlatformData struct {
	Id            string                                   `json:"id"`
	Type          string                                   `json:"type"`
	Attributes    providerVersionPlatformDataAttributes    `json:"attributes"`
	Relationships providerVersionPlatformDataRelationships `json:"relationships"`
	Links         providerVersionRelationshipPlatformLinks `json:"links"`
}

type providerVersionPlatformDataAttributes struct {
	Os                     string                                           `json:"os"`
	Arch                   string                                           `json:"arch"`
	Filename               string                                           `json:"filename"`
	Shasum                 string                                           `json:"shasum"`
	Permissions            providerVersionPlatformDataAttributesPermissions `json:"permissions"`
	ProviderBinaryUploaded bool                                             `json:"provider-binary-uploaded"`
}

type providerVersionPlatformDataAttributesPermissions struct {
	CanDelete      bool `json:"can-delete"`
	CanUploadAsset bool `json:"can-upload-asset"`
}

type providerVersionPlatformDataRelationships struct {
	RegistryProviderVersion providerVersionPlatformVersion `json:"registry-provider-version"`
}

type providerVersionPlatformVersion struct {
	Data providerVersionPlatformVersionData `json:"data"`
}

type providerVersionPlatformVersionData struct {
	Id   string `json:"id"`
	Type string `json:"type"`
}

type providerVersionRelationshipPlatformLinks struct {
	ProviderBinaryUpload string `json:"provider-binary-upload"`
}
