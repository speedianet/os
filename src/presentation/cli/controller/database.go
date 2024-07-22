package cliController

import (
	internalDbInfra "github.com/speedianet/os/src/infra/internalDatabase"
	cliHelper "github.com/speedianet/os/src/presentation/cli/helper"
	"github.com/speedianet/os/src/presentation/service"
	"github.com/spf13/cobra"
)

type DatabaseController struct {
	persistentDbSvc *internalDbInfra.PersistentDatabaseService
	databaseService *service.DatabaseService
}

func NewDatabaseController(
	persistentDbSvc *internalDbInfra.PersistentDatabaseService,
) *DatabaseController {
	return &DatabaseController{
		persistentDbSvc: persistentDbSvc,
		databaseService: service.NewDatabaseService(persistentDbSvc),
	}
}

func (controller *DatabaseController) Read() *cobra.Command {
	var dbTypeStr string

	cmd := &cobra.Command{
		Use:   "list",
		Short: "GetDatabases",
		Run: func(cmd *cobra.Command, args []string) {
			requestBody := map[string]interface{}{
				"dbType": dbTypeStr,
			}

			cliHelper.ServiceResponseWrapper(
				controller.databaseService.Read(requestBody),
			)
		},
	}

	cmd.Flags().StringVarP(&dbTypeStr, "db-type", "t", "", "DatabaseType")
	cmd.MarkFlagRequired("db-type")
	return cmd
}

func (controller *DatabaseController) Create() *cobra.Command {
	var dbTypeStr, dbNameStr string

	cmd := &cobra.Command{
		Use:   "create",
		Short: "CreateNewDatabase",
		Run: func(cmd *cobra.Command, args []string) {
			requestBody := map[string]interface{}{
				"dbType": dbTypeStr,
				"dbName": dbNameStr,
			}

			cliHelper.ServiceResponseWrapper(
				controller.databaseService.Create(requestBody),
			)
		},
	}

	cmd.Flags().StringVarP(&dbTypeStr, "db-type", "t", "", "DatabaseType")
	cmd.MarkFlagRequired("db-type")
	cmd.Flags().StringVarP(&dbNameStr, "db-name", "n", "", "DatabaseName")
	cmd.MarkFlagRequired("db-name")
	return cmd
}

func (controller *DatabaseController) Delete() *cobra.Command {
	var dbTypeStr, dbNameStr string

	cmd := &cobra.Command{
		Use:   "delete",
		Short: "DeleteDatabase",
		Run: func(cmd *cobra.Command, args []string) {
			requestBody := map[string]interface{}{
				"dbType": dbTypeStr,
				"dbName": dbNameStr,
			}

			cliHelper.ServiceResponseWrapper(
				controller.databaseService.Delete(requestBody),
			)
		},
	}

	cmd.Flags().StringVarP(&dbTypeStr, "db-type", "t", "", "DatabaseType")
	cmd.MarkFlagRequired("db-type")
	cmd.Flags().StringVarP(&dbNameStr, "db-name", "n", "", "DatabaseName")
	cmd.MarkFlagRequired("db-name")
	return cmd
}

func (controller *DatabaseController) CreateUser() *cobra.Command {
	var dbTypeStr, dbNameStr, dbUserStr, dbPassStr string
	var privilegesSlice []string

	cmd := &cobra.Command{
		Use:   "create-user",
		Short: "CreateNewDatabaseUser",
		Run: func(cmd *cobra.Command, args []string) {
			requestBody := map[string]interface{}{
				"dbType":     dbTypeStr,
				"dbName":     dbNameStr,
				"username":   dbUserStr,
				"password":   dbPassStr,
				"privileges": privilegesSlice,
			}

			cliHelper.ServiceResponseWrapper(
				controller.databaseService.CreateUser(requestBody),
			)
		},
	}

	cmd.Flags().StringVarP(&dbTypeStr, "db-type", "t", "", "DatabaseType")
	cmd.MarkFlagRequired("db-type")
	cmd.Flags().StringVarP(&dbNameStr, "db-name", "n", "", "DatabaseName")
	cmd.MarkFlagRequired("db-name")
	cmd.Flags().StringVarP(&dbUserStr, "username", "u", "", "Username")
	cmd.MarkFlagRequired("username")
	cmd.Flags().StringVarP(&dbPassStr, "password", "p", "", "Password")
	cmd.MarkFlagRequired("password")
	cmd.Flags().StringSliceVarP(
		&privilegesSlice,
		"privileges",
		"r",
		[]string{},
		"DatabasePrivileges (Comma-separated)",
	)

	return cmd
}

func (controller *DatabaseController) DeleteUser() *cobra.Command {
	var dbTypeStr, dbNameStr, dbUserStr string

	cmd := &cobra.Command{
		Use:   "delete-user",
		Short: "DeleteDatabaseUser",
		Run: func(cmd *cobra.Command, args []string) {
			requestBody := map[string]interface{}{
				"dbType":   dbTypeStr,
				"dbName":   dbNameStr,
				"username": dbUserStr,
			}

			cliHelper.ServiceResponseWrapper(
				controller.databaseService.DeleteUser(requestBody),
			)
		},
	}

	cmd.Flags().StringVarP(&dbTypeStr, "db-type", "t", "", "DatabaseType")
	cmd.MarkFlagRequired("db-type")
	cmd.Flags().StringVarP(&dbNameStr, "db-name", "n", "", "DatabaseName")
	cmd.MarkFlagRequired("db-name")
	cmd.Flags().StringVarP(&dbUserStr, "username", "u", "", "Username")
	cmd.MarkFlagRequired("username")
	return cmd
}
