```
This project is under active development and is not ready for production use.
```

# [Speedia OS](https://speedia.net/os/) &middot; [![Roadmap](https://img.shields.io/badge/roadmap-014737)](https://github.com/orgs/speedianet/projects/1) [![Demo](https://img.shields.io/badge/read--only_demo-233876)](https://os.demo.speedia.net:1618/_/) [![Community](https://img.shields.io/badge/community-751A3D)](https://github.com/orgs/speedianet/discussions) [![Report Card](https://img.shields.io/badge/go%20report-A%2B-brightgreen)](https://goreportcard.com/report/github.com/speedianet/os) [![License](https://img.shields.io/badge/license-EPL-blue.svg)](https://github.com/speedianet/os/blob/main/LICENSE.md)

Speedia OS is a container operating system designed so you never have to write a Dockerfile again. It comes with a user-friendly dashboard, REST API and CLI for seamless container management.

A read-only demo of the dashboard is available at [https://os.demo.speedia.net:1618/\_/](https://os.demo.speedia.net:1618/_/). The default credentials are `demo` and `abc123`.

## Running

To run Speedia OS, you can pull the image from DockerHub and use the following command:

```
docker run --rm --name myapp-container \
  --env 'PRIMARY_VHOST=myapp.net' \
  -p 8080:80 -p 8443:443 -p 1618:1618 \
  -it docker.io/speedianet/os:latest
```

In this example, the container ports 80, 443, and 1618 are mapped to host ports 8080, 8443, and 1618, respectively. If you are running multiple containers on the same host, consider using a reverse proxy to manage traffic.

You can customize the container name, vhost, and host ports as needed. The `--rm` flag ensures the container is removed upon stopping. To retain the container, simply omit this flag.

After deploying the container, access the shell to create a new account with the following command:

```
docker exec -it myapp-container /bin/bash
os account create -u admin -p admin
```

Once the account is created, you can access the dashboard at `https://localhost:1618/_/` and log in with the credentials you just set up. Note that you may encounter an SSL warning due to the self-signed certificate, which you can ignore or replace with your own certificate later.

Through the dashboard, you can deploy applications using the Marketplace feature with just a few clicks. You can also use the CLI for deployments, such as:

```
os mktplace install -s wp \
  -f 'adminUsername:admin' \
  -f 'adminPassword:abc123' \
  -f 'adminMailAddress:user@example.com'
```

The API Swagger documentation is available at `https://localhost:1618/api/swagger/`.

Speedia OS is compatible with Docker, Podman, Docker Swarm, Rancher, Kubernetes, Portainer, and any other tool that supports OCI-compliant containers.

## Development

The public roadmap for Speedia OS is available [here](https://github.com/orgs/speedianet/projects/1). You may create issues or pull requests to contribute to the project.

In this repository you'll find the REST API and CLI code plus the dashboard assets. The API and CLI uses Clean Architecture, DDD, TDD, CQRS, Object Calisthenics, etc. Understand how these concepts works before proceeding is advised.

To run this project during development you must install [Air](https://github.com/cosmtrek/air). Air is a tool that will watch for changes in the project and recompile it automatically.

### Environment Variables

You must have an `.env` file in the root of the git directory **during development**. You can use the `.env.example` file as a template. Air will read the `.env` file and use it to run the project during development.

If you add a new env var that is required to run the apis, please add it to the `src/presentation/cli/checkEnvs.go` file.

When running in production, the `/speedia/.env` file is only used if the environment variables weren't set in the system.

### Unit Testing

Speedia OS commands can harm your system, so it's important to run the unit tests in a proper container:

```
podman build -t os-unit-test:latest -f Containerfile.test .
podman run --rm -it os-unit-test:latest
```

Make sure you have the `.env` file in the root of the git directory before running the tests.

Some tests can run in your local machine, although it's not recommended. However, if you to give it a go, make sure to create the `/speedia/` directory before running the tests:

```
sudo mkdir /speedia
sudo chown $(whoami):$(whoami) /speedia
```

### Dev Utils

The `src/devUtils` folder is not a Clean Architecture layer, it's there to help you during development. You can add any file you want there, but it's not recommended to add any file that is not related to development since the code there is meant to be ignored by the build process.

For instance there you'll find a `testHelpers.go` file that is used to read the `.env` during tests.

### Building

#### Simple Build

To build the project, run the command below. It takes two minutes to build the project at first. After that, it takes less than 10 seconds to build.

```
podman build -t os:latest .
```

To run the project you may use the following command:

```
podman run --name os --env 'PRIMARY_VHOST=speedia.net' --rm -p 1618:1618 -it os:latest
```

When testing, consider publishing port 80 and 443 to the host so that you don't need to use a reverse proxy. You should also consider using `--env 'LOG_LEVEL=debug'` to increase the log verbosity.

#### Development Build

When developing the project, you may want to use the following steps for the best experience:

1. Add this to your `.bashrc` (or equivalent) file if you don't have it yet:

```bash
function os-build() {
    ports=(-p 1618:1618 -p 7080:7080)
    case ${1} in
    http)
        sudo sysctl net.ipv4.ip_unprivileged_port_start=80
        ports+=(-p 80:80 -p 443:443)
        ;;
    no-cache)
        podman image prune -a
        podman rmi localhost/os -f
        ;;
    esac

    make build
    podman build -t os:latest .
    podman run --name os \
        --env 'LOG_LEVEL=debug' --env 'DEV_MODE=true' --env 'PRIMARY_VHOST=speedia.cloud' \
        --hostname=speedia.cloud --cpus=2 --memory=2g --rm \
        --volume "$(pwd)/bin:/speedia/bin:Z,ro,bind,slave" \
        "${ports[@]}" -it os:latest
}
```

Read the script above and understand what it does. The `os-build` function will build the project, run the container, and expose the ports 1618 and 7080.

If you pass the `http` argument, it will also expose the ports 80 and 443 to the host. If you pass the `no-cache` argument, it will remove the image cache and rebuild the image from scratch.

Port 1618 is used for the dashboard and port 7080 is used for the OpenLiteSpeed admin panel which may come in handy during development related to the PHP features, but isn't necessary, so you can remove it if you want.

2. Run `source ~/.bashrc` (or equivalent) to reload the terminal or close and open the terminal;

3. Open a new terminal and run `os-build` to build and run the container on that window. You could add the `-d` flag to run the container in the background on the `os-build` script, but to easily stop the container with CTRL+C instead of using `podman stop os`, it's better to run in a second terminal, but you can do it as you prefer;

4. Back on the first terminal, run `air` to monitor any changes in the project and recompile it automatically. Since we're using the `DEV_MODE=true` environment variable, the frontend will also automatically reload when the `os-api` service is restarted by Air (check the Makefile to understand how it works);

5. On the very first time you build the container, you must run the following command to symlink the binary Air will generate to the one used by the container entrypoint:

```
podman exec os /bin/bash -c 'rm -f os && ln -s bin/os os && supervisorctl restart os-api'
```

If you look closely at the `os-build` function, you'll see that it mounts the `bin` directory to the container. This is necessary because Air will generate the binary in the `bin` directory and the container will look for the binary in the root of the container. The symlink command above will make sure the container is using the binary you're altering during development.

Note: all the commands above are meant to be run in the root of the project (before src/).

### Web UIs

This project has two web UIs, the previous Vue.js frontend and the new [Templ](https://templ.guide/) + [Alpine.js](https://alpinejs.dev/) + [HTMX](https://htmx.org/docs/) frontend. The Vue.js frontend is deprecated and will be removed in the future. It's available at `/_/` and the [Templ](https://templ.guide/) + [Alpine.js](https://alpinejs.dev/) + [HTMX](https://htmx.org/docs/) frontend is available at `/`.

The new frontend based on the [Templ](https://templ.guide/) + [Alpine.js](https://alpinejs.dev/) + [HTMX](https://htmx.org/docs/) combo mentioned was developed as a proof of concept to create an interface without needing to leave Go. To understand the entire conceptual and theoretical foundation behind using these technologies to create a new architecture, [access this article](https://ntorga.com/full-stack-go-app-with-htmx-and-alpinejs/). However, to grasp the practical basis of how to apply this new architecture, [refer to the proof of concept](https://github.com/ntorga/clean-ddd-full-stack-go-poc) used to develop it.

For the interface code to be read and rendered by Go, we need to convert all `.templ` files into `.go` files. To do this, run the following command at the root of the application:

```
templ generate -path src/presentation/api
```

With this, Go will be able to provide the entire application interface at the previously indicated route (`/`).

**NOTE:** It is important that this is done before using Air to create the binary; otherwise, the Web UI will not be embedded, and you will not be able to use it.

### VSCode Extensions

The following extensions are highly encouraged to be used during development:

```
EditorConfig.EditorConfig
GitHub.copilot
GitHub.vscode-pull-request-github
esbenp.prettier-vscode
foxundermoon.shell-format
golang.go
hbenl.vscode-test-explorer
ms-vscode.test-adapter-converter
redhat.vscode-yaml
streetsidesoftware.code-spell-checker
streetsidesoftware.code-spell-checker-portuguese-brazilian
timonwong.shellcheck
```

## REST API

### Authentication

The API accepts two types of tokens and uses the standard "Authorization: Bearer \<token\>" header:

- **sessionToken**: is a JWT, used for dashboard access and generated with the account login credentials. The token contains the accountId, IP address and expiration date. It expires in 3 hours and only the IP address used on the token generation is allowed to use it.

- **accountApiKey**: is a token meant for M2M communication. The token is a _AES-256-CTR-Encrypted-Base64-Encoded_ string, but only the SHA3-256 hash of the key is stored in the server. The accountId is retrieved during key decoding, thus you don't need to provide it. The token never expires, but the user can update it at any time.

### OpenApi // Swagger

To generate the swagger documentation, you must use the following command:

```
swag init -g src/presentation/api/api.go -o src/presentation/api/docs
```

The annotations are in the controller files. The reference file can be found [here](https://github.com/swaggo/swag#attribute).
