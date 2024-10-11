# Welcome to the Jot API

Jot is a journaling application with an emphasis of tracking how you're doing over time

## Set Up Instructions
1. [Download and install golang](https://go.dev/doc/install)
   2. [Instructions for setting up GOPATH and GOROOT](https://www.geeksforgeeks.org/golang-gopath-and-goroot/)
2. Install Dependencies ```go mod download```
3. Run application ```go run main.go```
4. Application will run on port ```8080``` by default. Add optional ENV variable, ```SERVER_PORT```, to run command to set up different port number
5. Navigate to endpoint ```/api/healthz```. If all is setup correctly you should receive a 200 response.
