## Setup the environment
```bash
podman play kube postgres.yml
# Wait for the database to finish loading
podman play kube openfga.yml
```

This will get you up and running the API in http://localhost:8080 and the playground in http://localhost:3000/playground

More info here: https://openfga.dev/docs/getting-started/setup-openfga/docker



## Useful commands 
You can take the store id from the playground
```bash
# Validate a model
fga model validate --file model.fga
# {"is_valid":true}

# Create a store
 fga store create --name pocstore
# {
#   "store": {
#     ...
#     "id":"01JEB6VHJ7V4A13B4Z8EAPW02W",
#     "name":"pocstore",
#     ...
#   }
# }

# Load tuples into the store
fga tuple write --store-id="01JEB6VHJ7V4A13B4Z8EAPW02W" --file tuples.yml
```
