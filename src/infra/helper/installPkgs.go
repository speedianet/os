package infraHelper

import (
	"errors"
	"log"
	"os"
)

func InstallPkgs(packages []string) error {
	_, err := RunCmd("apt-get", "update", "-qq")
	if err != nil {
		return errors.New("UpdateRepositoriesFailed")
	}

	installPackages := append(
		[]string{"install", "-y", "--no-install-recommends"},
		packages...,
	)

	var installErr error
	nAttempts := 3
	for i := 0; i < nAttempts; i++ {
		_, err := RunCmd("apt-get", installPackages...)
		if err == nil {
			break
		}

		log.Printf("InstallPkgError: %s", err.Error())

		if i == nAttempts-1 {
			installErr = errors.New("InstallAttemptsFailed")
		}
	}

	os.RemoveAll("/var/lib/apt/lists")
	os.RemoveAll("/var/cache/apt/archives")

	return installErr
}
