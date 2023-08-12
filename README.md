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

### For Virtual Hosting on Ubuntu

Replace you_domain_name by the domain name you want to access the website.
1. Install apache2 : `sudo apt install apache2`
2. `sudo a2enmod proxy proxy_http`
3. `sudo nano /etc/apache2/sites-available/your_domain_name.conf` 
4. Copy and paste the virtual host file.
5. `sudo a2ensite your_domain_name.conf`
6. `sudo a2dissite 000-default.conf`
7. `sudo apache2ctl configtest`
8. `sudo nano /etc/hosts`
Add the following line:
```
127.0.0.1  your_domain_name
```
9. `sudo systemctl restart apache2`
10. `sudo systemctl status apache2`
Check your_domain_name on your browser

## Setup using Makefile
Run command `make` in root directory 
To host your website virtually on a custom domain name run these commands :
```
chmod +x ./scripts/virtualHostSetup.sh
 ./scripts/virtualHostSetup.sh
```
## Setup using bash script 
Run commands 
```
chmod +x ./scripts/setup.sh
 ./scripts/setup.sh
```

