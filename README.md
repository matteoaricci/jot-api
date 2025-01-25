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
The project uses a postgres database with [GORM](https://gorm.io/docs/) for the orm. Migrations are handled using [Goose](https://github.com/pressly/goose?tab=readme-ov-file#goose)

Goose requires environment variables to run properly so you will need a ```.env``` file with the following variables:
* GOOSE_DRIVER=```postgres```
* GOOSE_DBSTRING="```host=localhost user=your_username_here password=your_password_here dbname=jot_db port=5432 sslmode=disable```"
* GOOSE_MIGRATION_DIR="```./db/migrations/```"