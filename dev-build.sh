#!/bin/bash

# Define the execution arguments.
ports=(-p 1618:1618)
case ${1} in
http)
  sudo sysctl net.ipv4.ip_unprivileged_port_start=80
  ports+=(-p 80:80 -p 443:443)
  ;;
ols)
  ports+=(-p 7080:7080)
  ;;
no-cache)
  podman image prune -a
  podman rmi localhost/os -f
  ;;
esac

echo "=> Building the container..."
make build
podman build -t os:latest .
podman run --name os -d \
  --env 'LOG_LEVEL=debug' --env 'DEV_MODE=true' --env 'PRIMARY_VHOST=speedia.cloud' \
  --hostname=speedia.cloud --cpus=2 --memory=2g --rm \
  --volume "$(pwd)/bin:/speedia/bin:Z,ro,bind,slave" \
  "${ports[@]}" -it os:latest

echo "=> Waiting for the container to start..."
sleep 5

echo "=> Replacing the standard binary with the development binary..."
podman exec os /bin/bash -c 'rm -f os && ln -s bin/os os && supervisorctl restart os-api'

echo
echo "<<>><<>><<>><<>><<>><<>><<>><<>><<>><<>><<>><<>><<>><<>><<>>"
echo
echo "=> Starting the development build..."
echo "Any changes to the code will trigger a rebuild automatically."
echo "Please, ignore the 'Only root can run SOS' message."
echo
echo "<<>><<>><<>><<>><<>><<>><<>><<>><<>><<>><<>><<>><<>><<>><<>>"
echo
sleep 3

stopDevBuild() {
  kill $airPid
  kill $podmanPid
  echo
  echo "=> Development build stopped."
  echo
  exit
}

trap stopDevBuild SIGINT

air &
airPid=$!
podman attach os &
podmanPid=$!

wait $airPid
wait $podmanPid
