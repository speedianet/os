package cliController

import (
	"strings"

	"github.com/speedianet/os/src/domain/dto"
	"github.com/speedianet/os/src/domain/useCase"
	"github.com/speedianet/os/src/domain/valueObject"
	internalDbInfra "github.com/speedianet/os/src/infra/internalDatabase"
	marketplaceInfra "github.com/speedianet/os/src/infra/marketplace"
	vhostInfra "github.com/speedianet/os/src/infra/vhost"
	cliHelper "github.com/speedianet/os/src/presentation/cli/helper"
	"github.com/spf13/cobra"
)

type MarketplaceController struct {
	persistentDbSvc *internalDbInfra.PersistentDatabaseService
}

func NewMarketplaceController(
	persistentDbSvc *internalDbInfra.PersistentDatabaseService,
) *MarketplaceController {
	return &MarketplaceController{
		persistentDbSvc: persistentDbSvc,
	}
}

func (controller MarketplaceController) GetCatalog() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-catalog",
		Short: "GetCatalogItems",
		Run: func(cmd *cobra.Command, args []string) {
			marketplaceQueryRepo := marketplaceInfra.NewMarketplaceQueryRepo(
				controller.persistentDbSvc,
			)

			catalogItems, err := useCase.GetMarketplaceCatalog(marketplaceQueryRepo)
			if err != nil {
				cliHelper.ResponseWrapper(false, err.Error())
			}

			cliHelper.ResponseWrapper(true, catalogItems)
		},
	}
	return cmd
}

func parseDataFields(
	dataFieldsStr []string,
) []valueObject.MarketplaceInstallableItemDataField {
	dataFields := []valueObject.MarketplaceInstallableItemDataField{}

	for _, dataFieldStr := range dataFieldsStr {
		dataFieldsParts := strings.Split(dataFieldStr, ":")
		if len(dataFieldsParts) < 2 {
			panic("InvalidDataFields")
		}

		dataField := valueObject.NewMarketplaceInstallableItemDataFieldPanic(
			valueObject.NewDataFieldKeyPanic(dataFieldsParts[0]),
			valueObject.NewDataFieldValuePanic(dataFieldsParts[1]),
		)
		dataFields = append(dataFields, dataField)
	}

	return dataFields
}

func (controller MarketplaceController) InstallCatalogItem() *cobra.Command {
	var catalogIdInt int
	var hostnameStr string
	var installDirStr string
	var dataFieldsStr []string

	cmd := &cobra.Command{
		Use:   "install",
		Short: "InstallCatalogItem",
		Run: func(cmd *cobra.Command, args []string) {
			catalogId := valueObject.NewMarketplaceCatalogItemIdPanic(catalogIdInt)
			hostname := valueObject.NewFqdnPanic(hostnameStr)

			var installDirPtr *valueObject.UnixFilePath
			if installDirStr != "" {
				installDir := valueObject.NewUnixFilePathPanic(installDirStr)
				installDirPtr = &installDir
			}

			// Format: key:value,key:value
			dataFields := parseDataFields(dataFieldsStr)

			marketplaceQueryRepo := marketplaceInfra.NewMarketplaceQueryRepo(controller.persistentDbSvc)
			marketplaceCmdRepo := marketplaceInfra.NewMarketplaceCmdRepo(controller.persistentDbSvc)
			vhostQueryRepo := vhostInfra.VirtualHostQueryRepo{}
			vhostCmdRepo := vhostInfra.VirtualHostCmdRepo{}

			dto := dto.NewInstallMarketplaceCatalogItem(
				catalogId,
				hostname,
				installDirPtr,
				dataFields,
			)
			err := useCase.InstallMarketplaceCatalogItem(
				marketplaceQueryRepo,
				marketplaceCmdRepo,
				vhostQueryRepo,
				vhostCmdRepo,
				dto,
			)
			if err != nil {
				cliHelper.ResponseWrapper(false, err.Error())
			}

			cliHelper.ResponseWrapper(true, "MarketplaceCatalogItemInstalled")
		},
	}

	cmd.Flags().IntVarP(
		&catalogIdInt, "catalogId", "i", 0, "Catalog item ID",
	)
	cmd.MarkFlagRequired("catalogId")
	cmd.Flags().StringVarP(
		&hostnameStr, "hostname", "n", "", "Hostname on which it will be installed",
	)
	cmd.MarkFlagRequired("hostname")
	cmd.Flags().StringVarP(
		&installDirStr, "installDir", "d", "", "Directory that stores installed files",
	)
	cmd.Flags().StringSliceVarP(
		&dataFieldsStr, "dataFields", "f", []string{}, "Installation data fields (key:value)",
	)
	return cmd
}

func (controller MarketplaceController) DeleteInstalledItem() *cobra.Command {
	var installedIdInt int
	var shouldUninstallServices bool
	var shouldRemoveFiles bool

	cmd := &cobra.Command{
		Use:   "uninstall",
		Short: "DeleteInstalledItem",
		Run: func(cmd *cobra.Command, args []string) {
			installedId := valueObject.NewMarketplaceInstalledItemIdPanic(installedIdInt)

			deleteMarketplaceInstalledItem := dto.NewDeleteMarketplaceInstalledItem(
				installedId, shouldUninstallServices, shouldRemoveFiles,
			)

			marketplaceQueryRepo := marketplaceInfra.NewMarketplaceQueryRepo(controller.persistentDbSvc)
			marketplaceCmdRepo := marketplaceInfra.NewMarketplaceCmdRepo(controller.persistentDbSvc)

			err := useCase.DeleteMarketplaceInstalledItem(
				marketplaceQueryRepo,
				marketplaceCmdRepo,
				deleteMarketplaceInstalledItem,
			)
			if err != nil {
				cliHelper.ResponseWrapper(false, err.Error())
			}

			cliHelper.ResponseWrapper(true, "MarketplaceInstalledItemDeleted")
		},
	}

	cmd.Flags().IntVarP(&installedIdInt, "installedId", "i", 0, "Installed item ID")
	cmd.MarkFlagRequired("installedId")
	cmd.Flags().BoolVarP(
		&shouldUninstallServices,
		"shouldUninstallServices",
		"s",
		true,
		"Should uninstall installed item services",
	)
	cmd.Flags().BoolVarP(
		&shouldRemoveFiles,
		"shouldRemoveFiles",
		"f",
		true,
		"Should remove installed item files",
	)
	return cmd
}
