package cliController

import (
	"github.com/speedianet/os/src/domain/dto"
	"github.com/speedianet/os/src/domain/useCase"
	"github.com/speedianet/os/src/domain/valueObject"
	infraHelper "github.com/speedianet/os/src/infra/helper"
	internalDbInfra "github.com/speedianet/os/src/infra/internalDatabase"
	sslInfra "github.com/speedianet/os/src/infra/ssl"
	cliHelper "github.com/speedianet/os/src/presentation/cli/helper"
	"github.com/speedianet/os/src/presentation/service"
	"github.com/spf13/cobra"
)

type SslController struct {
	persistentDbSvc *internalDbInfra.PersistentDatabaseService
	transientDbSvc  *internalDbInfra.TransientDatabaseService
	sslService      *service.SslService
}

func NewSslController(
	persistentDbSvc *internalDbInfra.PersistentDatabaseService,
	transientDbSvc *internalDbInfra.TransientDatabaseService,
) *SslController {
	return &SslController{
		persistentDbSvc: persistentDbSvc,
		transientDbSvc:  transientDbSvc,
		sslService:      service.NewSslService(persistentDbSvc, transientDbSvc),
	}
}

func (controller *SslController) Read() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get",
		Short: "ReadSslPairs",
		Run: func(cmd *cobra.Command, args []string) {
			cliHelper.ServiceResponseWrapper(controller.sslService.Read())
		},
	}

	return cmd
}

func (controller *SslController) Create() *cobra.Command {
	var virtualHostsSlice []string
	var certFilePathStr, keyFilePathStr string

	cmd := &cobra.Command{
		Use:   "create",
		Short: "CreateSslPair",
		Run: func(cmd *cobra.Command, args []string) {
			requestBody := map[string]interface{}{
				"virtualHosts": virtualHostsSlice,
			}

			certFilePath, err := valueObject.NewUnixFilePath(certFilePathStr)
			if err != nil {
				cliHelper.ResponseWrapper(false, "InvalidCertificateFilePath")
			}
			certContentStr, err := infraHelper.GetFileContent(certFilePath.String())
			if err != nil {
				cliHelper.ResponseWrapper(false, "OpenSslCertificateFileError")
			}
			requestBody["certificate"] = certContentStr

			privateKeyFilePath, err := valueObject.NewUnixFilePath(keyFilePathStr)
			if err != nil {
				cliHelper.ResponseWrapper(false, "InvalidSslPrivateKeyFilePath")
			}
			privateKeyContentStr, err := infraHelper.GetFileContent(
				privateKeyFilePath.String(),
			)
			if err != nil {
				cliHelper.ResponseWrapper(false, "OpenSslPrivateKeyFileError")
			}
			requestBody["key"] = privateKeyContentStr

			cliHelper.ServiceResponseWrapper(controller.sslService.Create(requestBody))
		},
	}

	cmd.Flags().StringSliceVarP(
		&virtualHostsSlice, "virtualHosts", "v", []string{}, "VirtualHosts",
	)
	cmd.MarkFlagRequired("virtualHosts")
	cmd.Flags().StringVarP(
		&certFilePathStr, "certFilePath", "c", "", "SslCertificateFilePath",
	)
	cmd.MarkFlagRequired("certFilePath")
	cmd.Flags().StringVarP(&keyFilePathStr, "keyFilePath", "k", "", "SslKeyFilePath")
	cmd.MarkFlagRequired("keyFilePath")
	return cmd
}

func (controller *SslController) Delete() *cobra.Command {
	var sslPairIdStr string

	cmd := &cobra.Command{
		Use:   "delete",
		Short: "DeleteSslPair",
		Run: func(cmd *cobra.Command, args []string) {
			sslId := valueObject.NewSslIdPanic(sslPairIdStr)

			cronQueryRepo := sslInfra.SslQueryRepo{}
			cronCmdRepo := sslInfra.NewSslCmdRepo(
				controller.persistentDbSvc, controller.transientDbSvc,
			)

			err := useCase.DeleteSslPair(
				cronQueryRepo,
				cronCmdRepo,
				sslId,
			)
			if err != nil {
				cliHelper.ResponseWrapper(false, err.Error())
			}

			cliHelper.ResponseWrapper(true, "SslPairDeleted")
		},
	}

	cmd.Flags().StringVarP(&sslPairIdStr, "id", "i", "", "SslPairId")
	cmd.MarkFlagRequired("sslPairId")
	return cmd
}

func (controller *SslController) DeleteVhosts() *cobra.Command {
	var sslPairIdStr string
	var virtualHostsSlice []string

	cmd := &cobra.Command{
		Use:   "remove-vhosts",
		Short: "RemoveSslPairVhosts",
		Run: func(cmd *cobra.Command, args []string) {
			sslPairId := valueObject.NewSslIdPanic(sslPairIdStr)

			var virtualHosts []valueObject.Fqdn
			for _, vhost := range virtualHostsSlice {
				virtualHosts = append(virtualHosts, valueObject.NewFqdnPanic(vhost))
			}

			dto := dto.NewDeleteSslPairVhosts(sslPairId, virtualHosts)

			sslQueryRepo := sslInfra.SslQueryRepo{}
			sslCmdRepo := sslInfra.NewSslCmdRepo(
				controller.persistentDbSvc, controller.transientDbSvc,
			)

			err := useCase.DeleteSslPairVhosts(
				sslQueryRepo,
				sslCmdRepo,
				dto,
			)
			if err != nil {
				cliHelper.ResponseWrapper(false, err.Error())
			}

			cliHelper.ResponseWrapper(true, "SslPairVhostsRemoved")
		},
	}

	cmd.Flags().StringVarP(&sslPairIdStr, "id", "i", "", "SslPairId")
	cmd.MarkFlagRequired("sslPairId")
	cmd.Flags().StringSliceVarP(&virtualHostsSlice, "virtualHosts", "v", []string{}, "VirtualHosts")
	cmd.MarkFlagRequired("virtualHosts")
	return cmd
}