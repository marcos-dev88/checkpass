#/bin/bash

YELLOW="\e[1;33m"
GREEN="${YELLOW}\e[32m"
WTOUTC="\e[0m"

go test -coverprofile cover.out;
go tool cover -html=cover.out -o cover.html;

echo -ne "\n${GREEN}***** to check the coverage html, copy this path:${WTOUTC} ${YELLOW}${PWD}/cover.html${WTOUTC}${GREEN} and paste in your browser ***** ${WTOUTC}\n"
