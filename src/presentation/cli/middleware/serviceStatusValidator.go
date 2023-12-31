package cliMiddleware

import (
	cliHelper "github.com/speedianet/os/src/presentation/cli/helper"
	sharedHelper "github.com/speedianet/os/src/presentation/shared/helper"
	"github.com/spf13/cobra"
)

func ServiceStatusValidator(serviceNameStr string) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		err := sharedHelper.CheckServices(serviceNameStr)
		if err != nil {
			cliHelper.ResponseWrapper(false, err.Error())
		}
	}
}
