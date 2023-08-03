# mvc-lms

Clone the repo. From the root directory run the following commands:
```
- go mod vendor
- go mod tidy
```
## MYSQL:
1. `mysql -u root -p` : and enter password
2. Import the sql dump file to your sql database: `mysql -u root -p books < dump.sql`
3. Replace previously written credentials in sampleConfig.yaml with your database credentials and rename that file with config.yaml. Do the same thing with sampleConfig.yaml in `./pkg/models`

## Run the test function:
1. run command `go test ./pkg/models` .
2. Ensure that should be OK 

## Running the server:
1. `go build -o mvc ./cmd/main.go`
2.  Run the binary file: `./mvc`

## Accessing the website
1. Open localhost:9000 on your browser

The username as well as password of the first admin is `admin` 
You can make new admins after logging in using the above mentioned credentials

Or just skip these steps and run command make in root directory 

Or you can also run commands 
```
chmod +x ./scripts/setup.sh
 ./scripts/setup.sh
```