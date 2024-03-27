package mktplaceInfra

import (
	"embed"
	"encoding/json"
	"errors"
	"io/fs"
	"log"

	"github.com/speedianet/os/src/domain/entity"
	"github.com/speedianet/os/src/domain/valueObject"
)

//go:embed assets/*
var assets embed.FS

type MktplaceCatalogQueryRepo struct{}

func (repo MktplaceCatalogQueryRepo) getMktCatalogItemFromFilePath(
	mktplaceItemFilePath valueObject.UnixFilePath,
) (entity.MarketplaceCatalogItem, error) {
	var mktplaceCatalogItem entity.MarketplaceCatalogItem

	mktplaceCatalogItemFile, err := assets.Open(mktplaceItemFilePath.String())
	if err != nil {
		return mktplaceCatalogItem, errors.New(
			"FailedToOpenMktCatalogItemFile: " + err.Error(),
		)
	}
	defer mktplaceCatalogItemFile.Close()

	mktplaceCatalogItemJsonDecoder := json.NewDecoder(mktplaceCatalogItemFile)
	err = mktplaceCatalogItemJsonDecoder.Decode(&mktplaceCatalogItem)
	if err != nil {
		return mktplaceCatalogItem, errors.New(
			"FailedToDecodeMktCatalogItemFile: " + err.Error(),
		)
	}

	return mktplaceCatalogItem, nil
}

func (repo MktplaceCatalogQueryRepo) GetItems() ([]entity.MarketplaceCatalogItem, error) {
	mktplaceCatalogItems := []entity.MarketplaceCatalogItem{}

	mktplaceItemFiles, err := fs.ReadDir(assets, "assets")
	if err != nil {
		return mktplaceCatalogItems, errors.New("FailedToGetMktAssetsFiles: " + err.Error())
	}

	if len(mktplaceItemFiles) == 0 {
		return mktplaceCatalogItems, errors.New("MktAssetsEmpty")
	}

	for mktplaceItemFileIndex, mktplaceItemFile := range mktplaceItemFiles {
		mktplaceItemFilePathStr := "assets/" + mktplaceItemFile.Name()
		mktplaceItemFilePath, err := valueObject.NewUnixFilePath(mktplaceItemFilePathStr)
		if err != nil {
			log.Printf("%s : %s", err.Error(), mktplaceItemFilePathStr)
			continue
		}

		mktplaceCatalogItem, err := repo.getMktCatalogItemFromFilePath(mktplaceItemFilePath)
		if err != nil {
			log.Printf("FailedToGetMktCatalogItem: %s", err.Error())
			continue
		}

		mktplaceItemIdInt := mktplaceItemFileIndex + 1
		mktplaceItemId, _ := valueObject.NewMktplaceItemId(mktplaceItemIdInt)
		mktplaceCatalogItem.Id = mktplaceItemId

		mktplaceCatalogItems = append(mktplaceCatalogItems, mktplaceCatalogItem)
	}

	return mktplaceCatalogItems, nil
}
