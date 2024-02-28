package cli

import (
	"fmt"
	"os"
	"path/filepath"

	cliInit "github.com/speedianet/os/src/presentation/cli/init"
	cliMiddleware "github.com/speedianet/os/src/presentation/cli/middleware"
	sharedMiddleware "github.com/speedianet/os/src/presentation/shared/middleware"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   filepath.Base(os.Args[0]),
	Short: "Speedia OS CLI",
	CompletionOptions: cobra.CompletionOptions{
		DisableDefaultCmd: true,
	},
}

func CliInit() {
	defer cliMiddleware.PanicHandler()
	cliMiddleware.PreventRootless()

	sharedMiddleware.CheckEnvs()

	transientDbSvc := cliInit.TransientDatabaseService()

	router := NewRouter(transientDbSvc)
	router.registerCliRoutes()

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
