# A "go" short-link service

## Installation
This tool is written in Go and can be easily installed  and started
with the following commands.

```
GOPATH=`pwd` go install github.com/quack79/golinks
bin/golinks
```

By default, the service will put all of its data in the directory `data` and will
listen to requests on the port `80`. Both of these, however, are easily configured
using the `--data=/path/to/data` and `--addr=:8067` command line flags.

## DNS Setup
To get the most benefit from the service, you should setup a DNS entry on your
local network, `go.corp.mycompany.com`. Make sure that corp.mycompany.com is in
the search domains for each user on the network. This is usually easily accomplished
by configuring your DHCP server. Now, simply typing "go" into your browser should
take you to the service, where you can register shortcuts. Obviously, those
shortcuts will also be available by typing "go/shortcut".

## Using the Service
Once you have it all setup, using it is pretty straight-forward.

#### Create a new shortcut
Type `go/edit/my-shortcut` and enter the URL.

#### Visit a shortcut
Type `go/my-shortcut` and you'll be redirected to the URL.

#### Shorten a URL
Type `go` and enter the URL.

#### View all go-links
Type `go/links` and you'll be redirected to a list of all your go-links.
