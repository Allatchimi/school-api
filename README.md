# To get started with the API, follow these steps:

### 1. Requirements

- Make installed for shortcuts

- Docker installed if you want to build and start postgres or redis containers

- Build and start Redis container with the command ```make docker-redis```

- Build and start postgres container with the command ```make docker-postgres```

- Rename .env.example to ```app.env```

- JWT .pem files with ES512(ECDSA SHA-512) algorithm: ```./assets/private/keys/jwt/private.pem``` ```./assets/private/keys/jwt/public.pem```
  You ca use this website to generate JWT keys for your tests [JWT online generator](https://jwt-keys.21no.de/)

- Password is hashed using Argon2id algorithm. If you want to customize salinity, you can edit the ```app.env``` file

Others information such configurations are on ```app.env```

### 2. Clone the repository

```
git clone https://github.com/EMENEC-FINANCE/school-api.git
```

```
cd go-api/
```

The entry point of the project is `cmd/` folder. In this folder the is the `main.go` file.

### 3. Install dependencies

```
make install
```

### 4. Run the API

```
make build
```

```
make run
```

API docs with openAPI v3.1(latest) is on

```
/api/v1/docs
```

If you want to scan vulnerabilities(security issues)

```
make scan
```


# Update GitHub Action Secrets for continuous integration(build and package)

Go to this link: [GitHub Action Secrets](https://github.com/EMENEC-FINANCE/school-api/settings/secrets/actions)

- ------------- On your GitHub Action Secrets page -------------

    - Set Secrets `GHCR_USERNAME` `GHCR_PASSWORD` with value your GitHub credentials. `GHCR_PASSWORD` is your personal access token with `write package` permission enabled
        

# Makefile Targets

- `docker-api`: Builds and starts the Docker container for local development.

- `docker-ghcr-push`: Builds the Docker image and pushes it to the GitHub Container Registry.

- `docker-ghcr-pull`: Pulls a specific image from the GitHub Container Registry.


# Additional Notes

By following these steps and customizing the Makefile to fit your specific needs, you can effectively manage your project using Docker and Make.