package sslInfra

import (
	"errors"
	"log"
	"os"
	"regexp"

	"github.com/speedianet/os/src/domain/dto"
	"github.com/speedianet/os/src/domain/valueObject"
	infraHelper "github.com/speedianet/os/src/infra/helper"
)

type SslCmdRepo struct{}

func (repo SslCmdRepo) Add(addSslPair dto.AddSslPair) error {
	sslQueryRepo := SslQueryRepo{}

	vhostSymlinkOf := addSslPair.VirtualHosts[0]
	for vhostIndex, vhost := range addSslPair.VirtualHosts {
		sslPair, err := sslQueryRepo.GetSslPairByVirtualHost(vhost)
		if err != nil && err.Error() != "SslPairNotFound" {
			log.Printf("FailedToValidateSslPairExistence (%s): %s", vhost.String(), err.Error())
			continue
		}

		sslPairExists := sslPair.Id != ""
		if sslPairExists {
			err := repo.Delete(sslPair.Id)
			if err != nil {
				log.Printf("FailedToDeleteTheOldSslPair (%s): %s", vhost.String(), err.Error())
				continue
			}
		}

		isSymlink := vhostIndex != 0
		if isSymlink {
			vhostCertToSymlinkPath := "/app/conf/pki/" + vhostSymlinkOf.String() + ".crt"
			vhostCertSymlinkPath := "/app/conf/pki/" + vhost.String() + ".crt"
			err = os.Symlink(vhostCertToSymlinkPath, vhostCertSymlinkPath)
			if err != nil {
				log.Printf("FailedToAddSslCertSymlink (%s): %s", vhost.String(), err.Error())
				continue
			}

			vhostKeyToSymlinkPath := "/app/conf/pki/" + vhostSymlinkOf.String() + ".key"
			vhostCertKeySymlinkPath := "/app/conf/pki/" + vhost.String() + ".key"
			err = os.Symlink(vhostKeyToSymlinkPath, vhostCertKeySymlinkPath)
			if err != nil {
				log.Printf("FailedToAddSslKeySymlink (%s): %s", vhost.String(), err.Error())
				continue
			}
		}

		if !isSymlink {
			shouldOverwrite := true

			vhostCertFilePath := "/app/conf/pki/" + vhost.String() + ".crt"
			err = infraHelper.UpdateFile(vhostCertFilePath, addSslPair.Certificate.String(), shouldOverwrite)
			if err != nil {
				return err
			}

			vhostCertKeyFilePath := "/app/conf/pki/" + vhost.String() + ".key"
			err = infraHelper.UpdateFile(vhostCertKeyFilePath, addSslPair.Key.String(), shouldOverwrite)
			if err != nil {
				return err
			}
		}

		vhostConfFilePath, err := sslQueryRepo.GetVhostConfFilePath(vhost)
		if err != nil {
			log.Printf("FailedToGetVhostConfFilePath (%s): %s", vhost.String(), err.Error())
			continue
		}

		_, err = infraHelper.RunCmd(
			"sed",
			"-i",
			"/root \\/app\\/html\\/"+vhost.String()+";/a\\\\n"+
				"    ssl_certificate /app/conf/pki/"+vhost.String()+".crt;\\n"+
				"    ssl_certificate_key /app/conf/pki/"+vhost.String()+".key;\\n",
			vhostConfFilePath.String(),
		)
		if err != nil {
			log.Printf("AddSslPairError (%s): %s", vhost.String(), err.Error())
			continue
		}

		log.Printf(
			"SSL '%v' added in '%v' virtual host.",
			addSslPair.Certificate.Id.String(),
			vhost.String(),
		)
	}

	return nil
}

func (repo SslCmdRepo) Delete(sslId valueObject.SslId) error {
	sslQueryRepo := SslQueryRepo{}
	sslPairsToDelete, err := sslQueryRepo.GetSslPairById(sslId)
	if err != nil {
		return errors.New("SslNotFound")
	}

	for _, sslPairVhostToDelete := range sslPairsToDelete.VirtualHosts {
		sslPairVhostToDeleteStr := sslPairVhostToDelete.String()

		vhostCertFilePath := "/app/conf/pki/" + sslPairVhostToDeleteStr + ".crt"
		err = os.Remove(vhostCertFilePath)
		if err != nil {
			log.Printf(
				"FailedToDeleteCertFile (%s): %s", sslPairVhostToDelete.String(), err.Error(),
			)
			continue
		}

		vhostCertKeyFilePath := "/app/conf/pki/" + sslPairVhostToDeleteStr + ".key"
		err = os.Remove(vhostCertKeyFilePath)
		if err != nil {
			log.Printf(
				"FailedToDeleteCertKeyFile (%s): %s", sslPairVhostToDelete.String(), err.Error(),
			)
			continue
		}

		vhostConfFilePath, err := sslQueryRepo.GetVhostConfFilePath(sslPairVhostToDelete)
		if err != nil {
			log.Printf("DeleteSslError (%s): %s", sslPairVhostToDelete.String(), err.Error())
			continue
		}

		vhostConfContentStr, err := infraHelper.GetFileContent(vhostConfFilePath.String())
		if err != nil {
			log.Printf("DeleteSslError (%s): %s", sslPairVhostToDelete.String(), err.Error())
			continue
		}

		vhostSslPortConfRegex := regexp.MustCompile(`\s*listen 443 ssl;`)
		vhostConfWithoutSslPort := vhostSslPortConfRegex.ReplaceAllString(vhostConfContentStr, "")
		vhostSslConfRegex := regexp.MustCompile(
			`\s*ssl_certificate\s+[^\n]*\n\s*ssl_certificate_key\s+[^\n]*\n`,
		)
		vhostConfWithoutSslConf := vhostSslConfRegex.ReplaceAllString(vhostConfWithoutSslPort, "")
		err = infraHelper.UpdateFile(vhostConfFilePath.String(), vhostConfWithoutSslConf, true)
		if err != nil {
			log.Printf("DeleteSslError (%s): %s", sslPairVhostToDelete.String(), err.Error())
			continue
		}

		log.Printf(
			"SSL '%s' of '%s' virtual host deleted.",
			sslId.String(),
			sslPairVhostToDelete.String(),
		)
	}

	return nil
}
