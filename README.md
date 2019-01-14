# sample-golang-webservice
This image is of a Sample Web Service which provides the localized response.

## Service Monitoring
Service has integrated New Relic APM.

Service requires establishment of trust certificate chain.  This is achieved via `./certs/newrelic.pem` certificate for `*.newrelic.com`.  This certificate will expire:  April 15, 2018.  Following is a helpful site for decoding public key properties:  https://www.sslshopper.com/certificate-decoder.html

Command to download updated `newrelic.pem`:
```
openssl s_client -connect collector.newrelic.com:443 </dev/null 2>/dev/null | openssl x509 -out newrelic.pem
```

### Golang Version
- [Golang 1.8](https://blog.golang.org/go1.8)

## Golang Development
Mac
```shell
# Golang Development is strongly encouraged out of the $GOPATH
# By default, Golang installs to your $HOME/go folder
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin
export PROJECTPATH=github.com/dukeluke16/sample-golang-webservice
git clone https://$PROJECTPATH $GOPATH/src/$PROJECTPATH

# Latest Golang 1.8 is supported
brew install go

# Change working director to the $PROJECTPATH
cd $GOPATH/src/$PROJECTPATH

# Preference is to explicitly manage your dependencies so you know when you are
#  going to break the Service Docker Image
# Optionally, one could dynamically get all dependencies
# go get ./...

# Build our Service
go build
```

Windows

[MSI installer](https://golang.org/dl/)

Open the MSI file and follow the prompts to install the Golang tools. By default, the installer puts the Golang distribution in c:\Go.

The installer should put the c:\Go\bin directory in your PATH environment variable. You may need to restart any open command prompts for the change to take effect.

_This script needs modified and verified on Windows._
```shell
# Golang Development is strongly encouraged out of the $GOPATH
export PROJECTPATH=github.com/dukeluke16/sample-golang-webservice
git clone https://$PROJECTPATH $GOPATH/src/$PROJECTPATH

# Change working director to the $PROJECTPATH
cd $GOPATH/src/$PROJECTPATH

# Preference is to explicitly manage your dependencies so you know when you are
#  going to break the Service Docker Image
# Optionally, one could dynamically get all dependencies
# go get ./...

# Go build our Service
go build
```
