package servicesInfra

import (
	"embed"
	"errors"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/speedianet/sam/src/domain/valueObject"
	infraHelper "github.com/speedianet/sam/src/infra/helper"
)

//go:embed assets/*
var assets embed.FS

func copyAssets(srcPath string, dstPath string) error {
	srcPath = "assets/" + srcPath
	srcFile, err := assets.Open(srcPath)
	if err != nil {
		log.Printf("OpenSourceFileError: %s", err)
		return errors.New("OpenSourceFileError")
	}
	defer srcFile.Close()

	dstFile, err := os.OpenFile(dstPath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Printf("OpenDestinationFileError: %s", err)
		return errors.New("OpenDestinationFileError")
	}
	defer dstFile.Close()

	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		log.Printf("CopyFileError: %s", err)
		return errors.New("CopyFileError")
	}

	return nil
}

func appendSupervisorConf(svcName string, svcBin string) error {
	supervisorConf := `
[program:` + svcName + `]
command=` + svcBin + `
user=root
directory=/speedia
autostart=true
autorestart=true
startretries=3
startsecs=5
stdout_logfile=/dev/stdout
stdout_logfile_maxbytes=0
stderr_logfile=/dev/stderr
stderr_logfile_maxbytes=0
`

	f, err := os.OpenFile("/speedia/supervisord.conf", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Printf("OpenSupervisorConfError: %s", err)
		return errors.New("OpenSupervisorConfError")
	}
	defer f.Close()

	if _, err := f.WriteString(supervisorConf); err != nil {
		log.Printf("WriteSupervisorConfError: %s", err)
		return errors.New("WriteSupervisorConfError")
	}

	return nil
}

func installOLS() error {
	err := infraHelper.DownloadFile(
		"https://repo.litespeed.sh",
		"/speedia/repo.litespeed.sh",
	)
	if err != nil {
		log.Printf("DownloadRepoFileError: %s", err)
		return errors.New("DownloadRepoFileError")
	}

	_, err = infraHelper.RunCmd(
		"bash",
		"/speedia/repo.litespeed.sh",
	)
	if err != nil {
		log.Printf("RepoAddError: %s", err)
		return errors.New("RepoAddError")
	}

	err = os.Remove("/speedia/repo.litespeed.sh")
	if err != nil {
		log.Printf("RemoveRepoFileError: %s", err)
		return errors.New("RemoveRepoFileError")
	}

	_, err = infraHelper.RunCmd(
		"install_packages",
		"openlitespeed",
		"lsphp74",
		"lsphp74-common",
		"lsphp74-curl",
		"lsphp74-intl",
		"lsphp74-mysql",
		"lsphp74-opcache",
		"lsphp74-sqlite3",
		"lsphp81",
		"lsphp81-common",
		"lsphp81-curl",
		"lsphp81-intl",
		"lsphp81-mysql",
		"lsphp81-opcache",
		"lsphp81-sqlite3",
		"lsphp82",
		"lsphp82-common",
		"lsphp82-curl",
		"lsphp82-intl",
		"lsphp82-mysql",
		"lsphp82-opcache",
		"lsphp82-sqlite3",
	)
	if err != nil {
		log.Printf("InstallServiceError: %s", err)
		return errors.New("InstallServiceError")
	}

	defaultDirs := []string{
		"logs",
		"conf",
		"html",
	}
	for _, dir := range defaultDirs {
		err = os.MkdirAll("/app/"+dir, 0755)
		if err != nil {
			log.Printf("CreateAppDirError: %s", err)
			return errors.New("CreateAppDirError")
		}
	}

	_, err = infraHelper.RunCmd(
		"chown",
		"-R",
		"nobody:nogroup",
		"/app",
	)
	if err != nil {
		log.Printf("ChownAppDirError: %s", err)
		return errors.New("ChownAppDirError")
	}

	err = copyAssets(
		"httpd_config.conf",
		"/usr/local/lsws/conf/httpd_config.conf",
	)
	if err != nil {
		log.Printf("CopyAssetsError: %s", err)
		return errors.New("CopyAssetsError")
	}

	err = copyAssets(
		"vhconf.conf",
		"/app/conf/vhconf.conf",
	)
	if err != nil {
		log.Printf("CopyAssetsError: %s", err)
		return errors.New("CopyAssetsError")
	}

	virtualHost := os.Getenv("VIRTUAL_HOST")
	_, err = infraHelper.RunCmd(
		"sed",
		"-i",
		"s/speedia.net/"+virtualHost+"/g",
		"/app/conf/vhconf.conf",
	)
	if err != nil {
		log.Printf("RenameVHostError: %s", err)
		return errors.New("RenameVHostError")
	}

	_, err = infraHelper.RunCmd(
		"chown",
		"-R",
		"lsadm:nogroup",
		"/app/conf",
	)
	if err != nil {
		log.Printf("ChownConfDirError: %s", err)
		return errors.New("ChownConfDirError")
	}

	err = copyAssets(
		"ols-entrypoint.sh",
		"/speedia/ols-entrypoint.sh",
	)
	if err != nil {
		log.Printf("CopyAssetsError: %s", err)
		return errors.New("CopyAssetsError")
	}

	err = appendSupervisorConf(
		"openlitespeed",
		"bash /speedia/ols-entrypoint.sh",
	)
	if err != nil {
		return errors.New("AddSupervisorConfError")
	}

	return nil
}

func installMysql(version *valueObject.ServiceVersion) error {
	err := infraHelper.DownloadFile(
		"https://r.mariadb.com/downloads/mariadb_repo_setup",
		"/speedia/repo.mariadb.sh",
	)
	if err != nil {
		log.Printf("DownloadRepoFileError: %s", err)
		return errors.New("DownloadRepoFileError")
	}

	versionFlag := ""
	if version != nil {
		versionFlag = "--mariadb-server-version=" + version.String()
	}

	_, err = infraHelper.RunCmd(
		"bash",
		"/speedia/repo.mariadb.sh",
		versionFlag,
	)
	if err != nil {
		log.Printf("RepoAddError: %s", err)
		return errors.New("RepoAddError")
	}

	err = os.Remove("/speedia/repo.mariadb.sh")
	if err != nil {
		log.Printf("RemoveRepoFileError: %s", err)
		return errors.New("RemoveRepoFileError")
	}

	_, err = infraHelper.RunCmd(
		"install_packages",
		"mariadb-server",
	)
	if err != nil {
		log.Printf("InstallServiceError: %s", err)
		return errors.New("InstallServiceError")
	}

	os.Symlink("/usr/bin/mariadb", "/usr/bin/mysql")
	os.Symlink("/usr/bin/mariadb-admin", "/usr/bin/mysqladmin")
	os.Symlink("/usr/bin/mariadbd-safe", "/usr/bin/mysqld_safe")

	_, err = infraHelper.RunCmd(
		"/usr/bin/mysqld_safe",
		"--no-watch",
	)
	if err != nil {
		log.Printf("StartMysqldSafeError: %s", err)
		return errors.New("StartMysqldSafeError")
	}

	time.Sleep(5 * time.Second)

	rootPass := infraHelper.GenPass(16)
	postInstallQueries := []string{
		"ALTER USER 'root'@'localhost' IDENTIFIED BY '" + rootPass + "';",
		"DELETE FROM mysql.user WHERE User='';",
		"DELETE FROM mysql.user WHERE User='root' AND Host NOT IN ('localhost', '127.0.0.1', '::1');",
		"DROP DATABASE IF EXISTS test;",
		"DELETE FROM mysql.db WHERE Db='test' OR Db='test\\_%';",
		"FLUSH PRIVILEGES;",
	}
	postInstallQueriesJoined := strings.Join(postInstallQueries, "; ")
	_, err = infraHelper.RunCmd(
		"mysql",
		"-e",
		postInstallQueriesJoined,
	)
	if err != nil {
		log.Printf("PostInstallQueryError: %s", err)
		return errors.New("PostInstallQueryError")
	}

	err = infraHelper.UpdateFile(
		"/root/.my.cnf",
		"[client]\nuser=root\npassword="+rootPass+"\n",
		true,
	)
	if err != nil {
		log.Printf("CreateMyCnfError: %s", err)
		return errors.New("CreateMyCnfError")
	}

	err = os.Chmod("/root/.my.cnf", 0400)
	if err != nil {
		log.Printf("ChmodMyCnfError: %s", err)
		return errors.New("ChmodMyCnfError")
	}

	_, err = infraHelper.RunCmd(
		"mysqladmin",
		"shutdown",
	)
	if err != nil {
		log.Printf("StopMysqldSafeError: %s", err)
		return errors.New("StopMysqldSafeError")
	}

	err = appendSupervisorConf(
		"mysql",
		"/usr/bin/mysqld_safe",
	)
	if err != nil {
		return errors.New("AddSupervisorConfError")
	}

	return nil
}

func Install(
	name valueObject.ServiceName,
	version *valueObject.ServiceVersion,
) error {
	switch name.String() {
	case "openlitespeed", "litespeed":
		return installOLS()
	case "mysql", "maria", "mariadb", "percona":
		return installMysql(version)
	}

	installCmd := exec.Command(
		"install_packages",
		name.String(),
	)

	err := installCmd.Run()
	if err != nil {
		log.Printf("InstallServiceError: %s", err)
		return errors.New("InstallServiceError")
	}

	return nil
}
