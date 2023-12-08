package cliController

import (
	"github.com/speedianet/os/src/domain/dto"
	"github.com/speedianet/os/src/domain/useCase"
	"github.com/speedianet/os/src/domain/valueObject"
	"github.com/speedianet/os/src/infra"
	cliHelper "github.com/speedianet/os/src/presentation/cli/helper"
	"github.com/spf13/cobra"
)

func GetServicesController() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get",
		Short: "GetServices",
		Run: func(cmd *cobra.Command, args []string) {
			servicesQueryRepo := infra.ServicesQueryRepo{}
			servicesList, err := useCase.GetServices(servicesQueryRepo)
			if err != nil {
				cliHelper.ResponseWrapper(false, err.Error())
			}

			cliHelper.ResponseWrapper(true, servicesList)
		},
	}

	return cmd
}

func GetInstallableServicesController() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get-installables",
		Short: "GetInstallableServices",
		Run: func(cmd *cobra.Command, args []string) {
			servicesQueryRepo := infra.ServicesQueryRepo{}
			servicesList, err := useCase.GetInstallableServices(servicesQueryRepo)
			if err != nil {
				cliHelper.ResponseWrapper(false, err.Error())
			}

			cliHelper.ResponseWrapper(true, servicesList)
		},
	}

	return cmd
}

func AddInstallableServiceController() *cobra.Command {
	var nameStr string
	var versionStr string
	var startupFileStr string
	var portsSlice []uint

	cmd := &cobra.Command{
		Use:   "add-installable",
		Short: "AddInstallableService",
		Run: func(cmd *cobra.Command, args []string) {
			svcName := valueObject.NewServiceNamePanic(nameStr)

			var svcVersionPtr *valueObject.ServiceVersion
			if versionStr != "" {
				svcVersion := valueObject.NewServiceVersionPanic(versionStr)
				svcVersionPtr = &svcVersion
			}

			var startupFilePtr *valueObject.UnixFilePath
			if startupFileStr != "" {
				startupFile := valueObject.NewUnixFilePathPanic(startupFileStr)
				startupFilePtr = &startupFile
			}

			var ports []valueObject.NetworkPort
			for _, port := range portsSlice {
				ports = append(ports, valueObject.NewNetworkPortPanic(port))
			}

			addInstallableServiceDto := dto.NewAddInstallableService(
				svcName,
				svcVersionPtr,
				startupFilePtr,
				ports,
			)

			servicesQueryRepo := infra.ServicesQueryRepo{}
			servicesCmdRepo := infra.ServicesCmdRepo{}

			err := useCase.AddInstallableService(
				servicesQueryRepo,
				servicesCmdRepo,
				addInstallableServiceDto,
			)
			if err != nil {
				cliHelper.ResponseWrapper(false, err.Error())
			}

			cliHelper.ResponseWrapper(true, "InstallableServiceAdded")
		},
	}

	cmd.Flags().StringVarP(&nameStr, "name", "n", "", "ServiceName")
	cmd.MarkFlagRequired("name")
	cmd.Flags().StringVarP(&versionStr, "version", "v", "", "ServiceVersion")
	cmd.Flags().StringVarP(&startupFileStr, "startup-file", "f", "", "StartupFile")
	cmd.Flags().UintSliceVarP(&portsSlice, "ports", "p", []uint{}, "Ports")
	return cmd
}

func AddCustomServiceController() *cobra.Command {
	var nameStr string
	var typeStr string
	var commandStr string
	var versionStr string
	var portsSlice []uint

	cmd := &cobra.Command{
		Use:   "add-custom",
		Short: "AddCustomService",
		Run: func(cmd *cobra.Command, args []string) {
			svcName := valueObject.NewServiceNamePanic(nameStr)
			svcType := valueObject.NewServiceTypePanic(typeStr)
			svcCommand := valueObject.NewUnixCommandPanic(commandStr)

			var svcVersionPtr *valueObject.ServiceVersion
			if versionStr != "" {
				svcVersion := valueObject.NewServiceVersionPanic(versionStr)
				svcVersionPtr = &svcVersion
			}

			var ports []valueObject.NetworkPort
			for _, port := range portsSlice {
				ports = append(ports, valueObject.NewNetworkPortPanic(port))
			}

			addCustomServiceDto := dto.NewAddCustomService(
				svcName,
				svcType,
				svcCommand,
				svcVersionPtr,
				ports,
			)

			servicesQueryRepo := infra.ServicesQueryRepo{}
			servicesCmdRepo := infra.ServicesCmdRepo{}

			err := useCase.AddCustomService(
				servicesQueryRepo,
				servicesCmdRepo,
				addCustomServiceDto,
			)
			if err != nil {
				cliHelper.ResponseWrapper(false, err.Error())
			}

			cliHelper.ResponseWrapper(true, "CustomServiceAdded")
		},
	}

	cmd.Flags().StringVarP(&nameStr, "name", "n", "", "ServiceName")
	cmd.MarkFlagRequired("name")
	cmd.Flags().StringVarP(&typeStr, "type", "t", "", "ServiceType")
	cmd.MarkFlagRequired("type")
	cmd.Flags().StringVarP(&commandStr, "command", "c", "", "UnixCommand")
	cmd.MarkFlagRequired("command")
	cmd.Flags().StringVarP(&versionStr, "version", "v", "", "ServiceVersion")
	cmd.Flags().UintSliceVarP(&portsSlice, "ports", "p", []uint{}, "Ports")
	return cmd
}

func UpdateServiceController() *cobra.Command {
	var nameStr string
	var statusStr string
	var versionStr string

	cmd := &cobra.Command{
		Use:   "update",
		Short: "UpdateServiceStatus",
		Run: func(cmd *cobra.Command, args []string) {

			svcName := valueObject.NewServiceNamePanic(nameStr)
			svcStatus := valueObject.NewServiceStatusPanic(statusStr)
			var svcVersionPtr *valueObject.ServiceVersion
			if versionStr != "" {
				svcVersion := valueObject.NewServiceVersionPanic(versionStr)
				svcVersionPtr = &svcVersion
			}

			updateSvcStatusDto := dto.NewUpdateSvcStatus(
				svcName,
				svcStatus,
				svcVersionPtr,
			)

			servicesQueryRepo := infra.ServicesQueryRepo{}
			servicesCmdRepo := infra.ServicesCmdRepo{}

			err := useCase.UpdateServiceStatus(
				servicesQueryRepo,
				servicesCmdRepo,
				updateSvcStatusDto,
			)
			if err != nil {
				cliHelper.ResponseWrapper(false, err.Error())
			}

			cliHelper.ResponseWrapper(true, "ServiceStatusUpdated")
		},
	}

	cmd.Flags().StringVarP(&nameStr, "name", "n", "", "ServiceName")
	cmd.MarkFlagRequired("name")
	cmd.Flags().StringVarP(&statusStr, "status", "s", "", "ServiceStatus")
	cmd.MarkFlagRequired("status")
	cmd.Flags().StringVarP(&versionStr, "version", "v", "", "ServiceVersion")
	return cmd
}
