# Welcome to the Jot API

Jot is a journaling application with an emphasis of tracking how you're doing over time

## Set Up Instructions
### Initial Set up
1. [Download and install golang](https://go.dev/doc/install)
   1. [Instructions for setting up GOPATH and GOROOT](https://www.geeksforgeeks.org/golang-gopath-and-goroot/)
2. Install Dependencies ```go mod download```
3. Run application ```go run main.go```
4. Application will run on port ```8080``` by default. Add optional ENV variable, ```SERVER_PORT```, to run command to set up different port number
5. Navigate to endpoint ```/api/healthz```. If all is set up correctly you should receive a 200 response.

### Database Setup
The project uses a postgres database with [GORM](https://gorm.io/docs/) for the orm. Migrations are handled using [Goose](https://github.com/pressly/goose?tab=readme-ov-file#goose) Variables should be set accordingly inside a ```dev.env``` file.

Database connection requires the following:
* DB_HOST=```your_host_here```
* DB_PORT=```your_port_here```
* DB_USERNAME=```your_username_here```
* DB_PASSWORD=```your_password_here```
* DB_NAME=```your_db_name_here```
* DB_SSLMODE=```your_ssl_mode_here```

*Host, port, and ssl mode will default to localhost, 5432, and disable respectively but username, password, and db name MUST have a value set*

Goose requires the following environment variables to run properly:
* GOOSE_DRIVER=```postgres```
* GOOSE_DBSTRING="```host=your_host_here user=your_username_here password=your_password_here dbname=db_name_here port=your_port_here sslmode=your_ssl_mode_here```"
* GOOSE_MIGRATION_DIR="```./db/migrations/```"

*You need the goose db string values to match your DB env values*