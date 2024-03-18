
docker volume create kanidmd
docker create --name kanidmd \
  -p 443:8443 \
  -p 636:3636 \
  -v kanidmd:/data \
  kanidm/server:latest
docker cp ./test_idm_config.toml kanidmd:/data/server.toml

## generate needed certs :
docker run --rm -i -t -v kanidmd:/data \
  kanidm/server:latest \
  kanidmd cert-generate

docker start kanidmd


docker exec -i -t kanidmd kanidmd recover-account admin

## Assuming kandim client is already installed to setup an OAuth client, server address on localhost is on https://localhost:443

## Both the service accounts for init setup : 
kanidm login --name admin 
kanidm service-account credential generate -D admin idm_admin
kanidm login --name idm_admin


kanidm system oauth2 create SplitWiseLiteDev "SplitWiseLite Dev client" https://0.0.0.0:8080

## Creating groups and assiging it scopes

kanidm group create splitwise_dev_admin --name idm_admin
kanidm system oauth2 update-scope-map SplitWiseLiteDev splitwise_dev_admin dev_admin
kanidm system oauth2 warning-insecure-client-disable-pkce SplitWiseLiteDev
