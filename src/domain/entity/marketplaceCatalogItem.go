package entity

import (
	"github.com/speedianet/os/src/domain/valueObject"
)

type MarketplaceCatalogItem struct {
	Id                 valueObject.MarketplaceItemId            `json:"id" yaml:"id"`
	Name               valueObject.MarketplaceItemName          `json:"name" yaml:"name"`
	Type               valueObject.MarketplaceItemType          `json:"type" yaml:"type"`
	Description        valueObject.MarketplaceItemDescription   `json:"description" yaml:"description"`
	ServiceNames       []valueObject.ServiceName                `json:"services" yaml:"serviceNames"`
	Mappings           []MarketplaceMapping                     `json:"mappings" yaml:"mappings"`
	DataFields         []valueObject.MarketplaceItemDataField   `json:"dataFields" yaml:"dataFields"`
	CmdSteps           []valueObject.MarketplaceItemInstallStep `json:"cmdSteps" yaml:"cmdSteps"`
	EstimatedSizeBytes valueObject.Byte                         `json:"estimatedSizeBytes" yaml:"estimatedSizeBytes"`
	AvatarUrl          valueObject.Url                          `json:"avatarUrl" yaml:"avatarUrl"`
	ScreenshotUrls     []valueObject.Url                        `json:"screenshotUrls" yaml:"screenshotUrls"`
}

func NewMarketplaceCatalogItem(
	id valueObject.MarketplaceItemId,
	itemName valueObject.MarketplaceItemName,
	itemType valueObject.MarketplaceItemType,
	description valueObject.MarketplaceItemDescription,
	serviceNames []valueObject.ServiceName,
	mappings []MarketplaceMapping,
	dataFields []valueObject.MarketplaceItemDataField,
	cmdSteps []valueObject.MarketplaceItemInstallStep,
	estimatedSizeBytes valueObject.Byte,
	avatarUrl valueObject.Url,
	screenshotUrls []valueObject.Url,
) MarketplaceCatalogItem {
	return MarketplaceCatalogItem{
		Id:                 id,
		Name:               itemName,
		Type:               itemType,
		Description:        description,
		ServiceNames:       serviceNames,
		Mappings:           mappings,
		DataFields:         dataFields,
		CmdSteps:           cmdSteps,
		EstimatedSizeBytes: estimatedSizeBytes,
		AvatarUrl:          avatarUrl,
		ScreenshotUrls:     screenshotUrls,
	}
}
