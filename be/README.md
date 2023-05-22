# Template Project for creating API

## How To Develop

### Onboarding
- Install Golang `1.11.2`
- Find and Replace all the import files into your project name
- In root project, run command `go mod init <project name>`
- Delete Both Gopkg.lock and Gopkg.toml files
- Run command `go mod tidy` to check unused dependencies
- Compile to check all over again

### Compiling
- For UNIX, build using `make build`
- For Windows, build using `go build ./cmd/api-temp` - `go build ./cmd/cron`

### Running
- Run docker, refer to [this document](../docker/README.md)

### Learning
- To learn how this service works, go to `cmd/api/api.main.go` for API and `cmd/cron/cron.main.go` for Cron
