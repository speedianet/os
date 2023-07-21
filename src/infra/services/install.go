package servicesInfra

import (
	"embed"
	"errors"
	"io"
	"log"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/speedianet/sam/src/domain/valueObject"
	infraHelper "github.com/speedianet/sam/src/infra/helper"
)

var supportedServicesVersion = map[string]string{
	"mariadb": `^(10\.([6-9]|10|11)|11\.[0-9]{1,2})$`,
	"node":    `^(1[2-9]|20)$`,
	"redis":   `^6\.(0|2)|^7\.0$`,
}

var OlsPackages = []string{
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
}

var MariaDbPackages = []string{
	"mariadb-server",
}

var NodePackages = []string{
	"nodejs",
}

var RedisPackages = []string{
	"redis-server",
}

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

func installOls() error {
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

	err = infraHelper.InstallPkgs(OlsPackages)
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

	err = SupervisordFacade{}.AddConf(
		"openlitespeed",
		"bash /speedia/ols-entrypoint.sh",
	)
	if err != nil {
		return errors.New("AddSupervisorConfError")
	}

	return nil
}

func installMariaDb(version *valueObject.ServiceVersion) error {
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
		re := regexp.MustCompile(supportedServicesVersion["mariadb"])
		isVersionAllowed := re.MatchString(version.String())

		if !isVersionAllowed {
			log.Printf("InvalidMysqlVersion: %s", version.String())
			return errors.New("InvalidMysqlVersion")
		}

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

	err = infraHelper.InstallPkgs(MariaDbPackages)
	if err != nil {
		log.Printf("InstallServiceError: %s", err)
		return errors.New("InstallServiceError")
	}

	os.Symlink("/usr/bin/mariadb", "/usr/bin/mysql")
	os.Symlink("/usr/bin/mariadb-admin", "/usr/bin/mysqladmin")
	os.Symlink("/usr/bin/mariadbd-safe", "/usr/bin/mysqld_safe")

	return nil
}

func installMysql(version *valueObject.ServiceVersion) error {
	err := installMariaDb(version)
	if err != nil {
		return errors.New("InstallMariaDbError")
	}

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

	err = SupervisordFacade{}.AddConf(
		"mysql",
		"/usr/bin/mysqld_safe",
	)
	if err != nil {
		return errors.New("AddSupervisorConfError")
	}

	return nil
}

func installNode(version *valueObject.ServiceVersion) error {
	repoFilePath := "/speedia/repo.node.sh"

	repoUrl := "https://deb.nodesource.com/setup_lts.x"
	if version != nil {
		re := regexp.MustCompile(supportedServicesVersion["node"])
		isVersionAllowed := re.MatchString(version.String())

		if !isVersionAllowed {
			log.Printf("InvalidNodeVersion: %s", version.String())
			return errors.New("InvalidNodeVersion")
		}

		repoUrl = "https://deb.nodesource.com/setup_" + version.String() + ".x"
	}

	err := infraHelper.DownloadFile(
		repoUrl,
		repoFilePath,
	)
	if err != nil {
		log.Printf("DownloadRepoFileError: %s", err)
		return errors.New("DownloadRepoFileError")
	}

	_, err = infraHelper.RunCmd(
		"bash",
		repoFilePath,
	)
	if err != nil {
		log.Printf("RepoAddError: %s", err)
		return errors.New("RepoAddError")
	}

	err = os.Remove(repoFilePath)
	if err != nil {
		log.Printf("RemoveRepoFileError: %s", err)
		return errors.New("RemoveRepoFileError")
	}

	err = infraHelper.InstallPkgs(NodePackages)
	if err != nil {
		log.Printf("InstallServiceError: %s", err)
		return errors.New("InstallServiceError")
	}

	return nil
}

func installRedis(version *valueObject.ServiceVersion) error {
	versionFlag := ""
	if version != nil {
		re := regexp.MustCompile(supportedServicesVersion["redis"])
		isVersionAllowed := re.MatchString(version.String())

		if !isVersionAllowed {
			log.Printf("InvalidRedisVersion: %s", version.String())
			return errors.New("InvalidRedisVersion")
		}
	}

	err := infraHelper.InstallPkgs(
		[]string{"lsb-release", "gpg"},
	)
	if err != nil {
		log.Printf("InstallPackagesError: %s", err)
		return errors.New("InstallPackagesError")
	}

	osRelease, err := infraHelper.GetOsRelease()
	if err != nil {
		log.Printf("GetOsReleaseError: %s", err)
		return errors.New("GetOsReleaseError")
	}

	err = infraHelper.DownloadFile(
		"https://packages.redis.io/gpg",
		"/speedia/redis.gpg",
	)
	if err != nil {
		log.Printf("DownloadRepoFileError: %s", err)
		return errors.New("DownloadRepoFileError")
	}

	_, err = infraHelper.RunCmd(
		"gpg",
		"--dearmor",
		"-o",
		"/usr/share/keyrings/redis-archive-keyring.gpg",
		"/speedia/redis.gpg",
	)
	if err != nil {
		log.Printf("GpgImportError: %s", err)
		return errors.New("GpgImportError")
	}

	err = os.Remove("/speedia/redis.gpg")
	if err != nil {
		log.Printf("RemoveRepoFileError: %s", err)
		return errors.New("RemoveRepoFileError")
	}

	repoLine := "deb [signed-by=/usr/share/keyrings/redis-archive-keyring.gpg] https://packages.redis.io/deb " + osRelease + " main"
	err = infraHelper.UpdateFile(
		"/etc/apt/sources.list.d/redis.list",
		repoLine,
		true,
	)
	if err != nil {
		log.Printf("CreateRepoFileError: %s", err)
		return errors.New("CreateRepoFileError")
	}

	if version != nil {
		versionStr := version.String()
		latestVersion, err := infraHelper.GetPkgLatestVersion(
			"redis-server",
			&versionStr,
		)
		if err != nil {
			log.Print(err)
			return err
		}

		versionFlag = "=" + latestVersion
	}

	err = infraHelper.InstallPkgs(
		[]string{RedisPackages[0] + versionFlag},
	)
	if err != nil {
		log.Printf("InstallServiceError: %s", err)
		return errors.New("InstallServiceError")
	}

	err = SupervisordFacade{}.AddConf(
		"redis",
		"/usr/bin/redis-server /etc/redis/redis.conf",
	)
	if err != nil {
		return errors.New("AddSupervisorConfError")
	}

	_, err = infraHelper.RunCmd(
		"sed",
		"-i",
		"s/^daemonize yes/daemonize no/g",
		"/etc/redis/redis.conf",
	)
	if err != nil {
		log.Printf("UpdateRedisConfError: %s", err)
		return errors.New("UpdateRedisConfError")
	}

	return nil
}

func Install(
	name valueObject.ServiceName,
	version *valueObject.ServiceVersion,
) error {
	switch name.String() {
	case "openlitespeed", "litespeed":
		return installOls()
	case "mysql", "mysqld", "maria", "mariadb", "percona", "perconadb":
		return installMysql(version)
	case "node", "nodejs":
		return installNode(version)
	case "redis", "redis-server":
		return installRedis(version)
	default:
		log.Printf("ServiceNotImplemented: %s", name.String())
		return errors.New("ServiceNotImplemented")
	}
}
