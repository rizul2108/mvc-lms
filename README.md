# mvc-lms

## Manual Setup
Clone the repo. FROM the root directory run the following commands:
```
go mod vendor
go mod tidy
```
### MySQL:
Run below commands after enterring your respective credentials
1. Run this command : `migrate -path database/migration/ -database "mysql://your_db_username:your_db_password@tcp(localhost:3306)/DB_NAME" -verbose up`
2. If this command gives error run command : `migrate -path database/migration/ -database "mysql://your_db_username:your_db_password@tcp(localhost:3306)/DB_NAME" force version`


### Run the test function:
1. run command `go test ./pkg/models` .
2. Ensure that should be OK 

### Running the server:
1. `go build -o mvc ./cmd/main.go`
2.  Run the binary file: `./mvc`

### Accessing the website
1. Open localhost:9000 on your browser

## Setup using Makefile
Run command `make` in root directory 

## Setup using bash script 
Run commands 
```
chmod +x ./scripts/setup.sh
 ./scripts/setup.sh
```

The username as well as password of the first admin is `admin` 
You can make new admins after logging in using the above mentioned credentials