package cliController

import (
	"github.com/speedianet/os/src/domain/dto"
	"github.com/speedianet/os/src/domain/useCase"
	"github.com/speedianet/os/src/domain/valueObject"
	"github.com/speedianet/os/src/infra"
	infraHelper "github.com/speedianet/os/src/infra/helper"
	cliHelper "github.com/speedianet/os/src/presentation/cli/helper"
	"github.com/spf13/cobra"
)

func GetVirtualHostsController() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get",
		Short: "GetVirtualHosts",
		Run: func(cmd *cobra.Command, args []string) {
			vhostQueryRepo := infra.VirtualHostQueryRepo{}
			vhostsList, err := useCase.GetVirtualHosts(vhostQueryRepo)
			if err != nil {
				cliHelper.ResponseWrapper(false, err.Error())
			}

			cliHelper.ResponseWrapper(true, vhostsList)
		},
	}

	return cmd
}

func AddVirtualHostController() *cobra.Command {
	var hostnameStr string
	var typeStr string
	var parentHostnameStr string

	cmd := &cobra.Command{
		Use:   "add",
		Short: "AddVirtualHost",
		Run: func(cmd *cobra.Command, args []string) {
			hostname := valueObject.NewFqdnPanic(hostnameStr)

			vhostTypeStr := "top-level"
			if typeStr != "" {
				vhostTypeStr = typeStr
			}
			vhostType := valueObject.NewVirtualHostTypePanic(vhostTypeStr)

			var parentHostnamePtr *valueObject.Fqdn
			if parentHostnameStr != "" {
				parentHostname := valueObject.NewFqdnPanic(parentHostnameStr)
				parentHostnamePtr = &parentHostname
			}

			addVirtualHostDto := dto.NewAddVirtualHost(
				hostname,
				vhostType,
				parentHostnamePtr,
			)

			vhostQueryRepo := infra.VirtualHostQueryRepo{}
			vhostCmdRepo := infra.VirtualHostCmdRepo{}

			err := useCase.AddVirtualHost(
				vhostQueryRepo,
				vhostCmdRepo,
				addVirtualHostDto,
			)
			if err != nil {
				cliHelper.ResponseWrapper(false, err.Error())
			}

			cliHelper.ResponseWrapper(true, "VirtualHostAdded")
		},
	}

	cmd.Flags().StringVarP(&hostnameStr, "hostname", "n", "", "VirtualHostHostname")
	cmd.MarkFlagRequired("hostname")
	cmd.Flags().StringVarP(
		&typeStr, "type", "t", "", "VirtualHostType (top-level|subdomain|alias)",
	)
	cmd.Flags().StringVarP(
		&parentHostnameStr, "parent", "p", "", "ParentHostname",
	)
	return cmd
}

func DeleteVirtualHostController() *cobra.Command {
	var hostnameStr string

	cmd := &cobra.Command{
		Use:   "delete",
		Short: "DeleteVirtualHost",
		Run: func(cmd *cobra.Command, args []string) {
			hostname := valueObject.NewFqdnPanic(hostnameStr)

			vhostQueryRepo := infra.VirtualHostQueryRepo{}
			vhostCmdRepo := infra.VirtualHostCmdRepo{}

			primaryHostname, err := infraHelper.GetPrimaryHostname()
			if err != nil {
				panic("PrimaryHostnameNotFound")
			}

			err = useCase.DeleteVirtualHost(
				vhostQueryRepo,
				vhostCmdRepo,
				primaryHostname,
				hostname,
			)
			if err != nil {
				cliHelper.ResponseWrapper(false, err.Error())
			}

			cliHelper.ResponseWrapper(true, "VirtualHostDeleted")
		},
	}

	cmd.Flags().StringVarP(&hostnameStr, "hostname", "n", "", "VirtualHostHostname")
	cmd.MarkFlagRequired("hostname")
	return cmd
}

func GetVirtualHostsWithMappingsController() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get",
		Short: "GetVirtualHostsWithMappings",
		Run: func(cmd *cobra.Command, args []string) {
			vhostQueryRepo := infra.VirtualHostQueryRepo{}
			vhostsList, err := useCase.GetVirtualHostsWithMappings(vhostQueryRepo)
			if err != nil {
				cliHelper.ResponseWrapper(false, err.Error())
			}

			cliHelper.ResponseWrapper(true, vhostsList)
		},
	}

	return cmd
}

func AddVirtualHostMappingController() *cobra.Command {
	var hostnameStr string
	var pathStr string
	var matchPatternStr string
	var targetTypeStr string
	var targetServiceStr string
	var targetUrlStr string
	var targetHttpResponseCode uint

	cmd := &cobra.Command{
		Use:   "add",
		Short: "AddMapping",
		Run: func(cmd *cobra.Command, args []string) {
			hostname := valueObject.NewFqdnPanic(hostnameStr)
			path := valueObject.NewMappingPathPanic(pathStr)
			targetType := valueObject.NewMappingTargetTypePanic(targetTypeStr)

			matchPattern := valueObject.NewMappingMatchPatternPanic("begins-with")
			if matchPatternStr != "" {
				matchPattern = valueObject.NewMappingMatchPatternPanic(matchPatternStr)
			}

			var targetServicePtr *valueObject.ServiceName
			if targetServiceStr != "" {
				targetService := valueObject.NewServiceNamePanic(targetServiceStr)
				targetServicePtr = &targetService
			}

			var targetUrlPtr *valueObject.Url
			if targetUrlStr != "" {
				targetUrl := valueObject.NewUrlPanic(targetUrlStr)
				targetUrlPtr = &targetUrl
			}

			var targetHttpResponseCodePtr *valueObject.HttpResponseCode
			if targetHttpResponseCode != 0 {
				targetHttpResponseCode := valueObject.NewHttpResponseCodePanic(
					targetHttpResponseCode,
				)
				targetHttpResponseCodePtr = &targetHttpResponseCode
			}

			addMappingDto := dto.NewAddMapping(
				hostname,
				path,
				matchPattern,
				targetType,
				targetServicePtr,
				targetUrlPtr,
				targetHttpResponseCodePtr,
			)

			vhostQueryRepo := infra.VirtualHostQueryRepo{}
			vhostCmdRepo := infra.VirtualHostCmdRepo{}

			err := useCase.AddMapping(
				vhostQueryRepo,
				vhostCmdRepo,
				addMappingDto,
			)
			if err != nil {
				cliHelper.ResponseWrapper(false, err.Error())
			}

			cliHelper.ResponseWrapper(true, "MappingAdded")
		},
	}

	cmd.Flags().StringVarP(&hostnameStr, "hostname", "n", "", "Hostname")
	cmd.MarkFlagRequired("hostname")
	cmd.Flags().StringVarP(&pathStr, "path", "p", "", "MappingPath")
	cmd.MarkFlagRequired("path")
	cmd.Flags().StringVarP(&matchPatternStr, "match", "m", "", "MatchPattern (begins-with|contains|ends-with)")
	cmd.Flags().StringVarP(
		&targetTypeStr, "type", "t", "", "MappingTargetType (service|url|response-code)",
	)
	cmd.MarkFlagRequired("type")
	cmd.Flags().StringVarP(
		&targetServiceStr, "service", "s", "", "TargetServiceName",
	)
	cmd.Flags().StringVarP(
		&targetUrlStr, "url", "u", "", "TargetUrl",
	)
	cmd.Flags().UintVarP(
		&targetHttpResponseCode, "response-code", "r", 0, "TargetHttpResponseCode",
	)
	return cmd
}

func DeleteVirtualHostMappingController() *cobra.Command {
	var hostnameStr string
	var mappingIdUint uint

	cmd := &cobra.Command{
		Use:   "delete",
		Short: "DeleteMapping",
		Run: func(cmd *cobra.Command, args []string) {
			hostname := valueObject.NewFqdnPanic(hostnameStr)
			mappingId := valueObject.NewMappingIdPanic(mappingIdUint)

			vhostQueryRepo := infra.VirtualHostQueryRepo{}
			vhostCmdRepo := infra.VirtualHostCmdRepo{}

			err := useCase.DeleteMapping(
				vhostQueryRepo,
				vhostCmdRepo,
				hostname,
				mappingId,
			)
			if err != nil {
				cliHelper.ResponseWrapper(false, err.Error())
			}

			cliHelper.ResponseWrapper(true, "MappingDeleted")
		},
	}

	cmd.Flags().StringVarP(&hostnameStr, "hostname", "n", "", "VirtualHost Hostname")
	cmd.MarkFlagRequired("hostname")
	cmd.Flags().UintVarP(&mappingIdUint, "id", "i", 0, "MappingId")
	cmd.MarkFlagRequired("id")
	return cmd
}
