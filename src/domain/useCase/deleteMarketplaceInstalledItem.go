package useCase

import (
	"errors"
	"log/slog"

	"github.com/speedianet/os/src/domain/dto"
	"github.com/speedianet/os/src/domain/repository"
)

func DeleteMarketplaceInstalledItem(
	marketplaceQueryRepo repository.MarketplaceQueryRepo,
	marketplaceCmdRepo repository.MarketplaceCmdRepo,
	dto dto.DeleteMarketplaceInstalledItem,
) error {
	_, err := marketplaceQueryRepo.ReadInstalledItemById(dto.InstalledId)
	if err != nil {
		return errors.New("MarketplaceInstalledItemNotFound")
	}

	err = marketplaceCmdRepo.UninstallItem(dto)
	if err != nil {
		slog.Info("UninstallMarketplaceItemError", slog.Any("err", err))
		return errors.New("UninstallMarketplaceItemInfraError")
	}

	return nil
}
