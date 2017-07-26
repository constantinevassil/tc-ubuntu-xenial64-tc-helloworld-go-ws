# tc-ubuntu-xenial64-tc-helloworld-go-ws

Vagrant config to minimal web werver with Go from your Mac.

## Getting started

On Mac

## tc-helloworld-go-ws

```bash
git clone https://github.com/topconnector/tc-ubuntu-xenial64-tc-helloworld-go-ws.git 
cd tc-ubuntu-xenial64-tc-helloworld-go-ws
```

You must have the following installed:

* Virtualbox >= 5.1.22

  Download and install from https://www.virtualbox.org/.
    
* Vagrant >= 1.9.7

  Download and install from https://www.vagrantup.com/.
  
* vagrant-vbguest Vagrant plugin
  automatically installs the host's VirtualBox Guest Additions on the guest system.

  Install by running: 

```bash
   vagrant plugin install vagrant-vbguest
   vagrant vbguest --do install
```

 
* update Vagrant box

  Install by running: 
    
```bash
    vagrant box update --provider virtualbox
```
   
* run Virtual machine (VM)

  Install by running: 
  
  Virtualbox:
  
```bash
    vagrant up --provider virtualbox
```

You should be prompted for the network interface after you run the command. 
Choose whichever interface is providing your internet connection

For example:

1) en5: Apple USB Ethernet Adapter

## Connecting to the Vagrant

We first need to ensure that we can access the vagrant locally before we test across other sites on the network. We also need to determine the IP address that was assigned to the vagrant instance. From the tc-ubuntu-xenial64-tc-helloworld-go-ws directory, execute the following command in the terminal:

```bash
    vagrant ssh
    
```

Install latest Golang:

```bash
cd ..
ubuntu@tc-rocksdb:/vagrant$ sudo apt-get update && sudo apt-get dist-upgrade
```

Check Golang version:

```bash
ubuntu@tc-rocksdb:/vagrant$ go version
ubuntu@tc-rocksdb:/vagrant$ 
go version go1.8.3 linux/amd64
```

### Set the GOPATH environment variable

The GOPATH environment variable specifies the location of your workspace. 

```bash
export GOPATH=/vagrant/mygo
cd /vagrant/mygo/tc-helloworld-go-ws
CGO_ENABLED=0 go build -a -tags netgo -ldflags '-w' .
./tc-helloworld-go-ws
```
