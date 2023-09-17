# Golang - micro service

```
https://go.dev/dl/ << download golang version
```

# Now that you have installed the Go release on your system, you need to set environment variables. Running the following three commands will set the required environment variable paths:

```
$export GOROOT=/usr/local/go 
$export GOPATH=＄HOME/Projects/Proj1  
$export PATH=＄GOPATH/bin:＄GOROOT/bin:＄PATH
```

# With this, you have successfully installed and set up your Go workplace. Check if everything is working properly by checking the version, using the following command:

```
$go version
$mkdir library-api
$cd library-api
$go mod init library-api
$go run main.go
```

# Auto run golang when there is code changes

```
$npx nodemon --exec "go run" ./main.go --signal SIGTERM
```
