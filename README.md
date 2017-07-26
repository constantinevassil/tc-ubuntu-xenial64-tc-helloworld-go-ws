# tc-ubuntu-xenial64-tc-helloworld-go-ws

Vagrant config to minimal web werver with Go from your Mac.

## Getting started

Topics:

1. Mac computer
1. Virtualbox 
1. Vagrant
1. Ubuntu 16.04+
1. Go
1. helloworld web app.
1. Testing with browser on local host.
1. Testing with browser on remote machine in same network.


## tc-helloworld-go-ws

On Mac

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
    vagrant ssh tc-vb-u-x-tc-helloworld-master
    
```

Refresh Ubuntu:

```bash
vagrant@tc-vb-u-x-tc-helloworld-master:~$ sudo apt-get update && sudo apt-get dist-upgrade
```

Check Golang version:

```bash
vagrant@tc-vb-u-x-tc-helloworld-master:~$ sudo apt-get install golang-go
vagrant@tc-vb-u-x-tc-helloworld-master:~$ go version
go version go1.7.4 linux/amd64
```

### Set the GOPATH environment variable

The GOPATH environment variable specifies the location of your workspace. 

```bash
vagrant@tc-vb-u-x-tc-helloworld-master:~$export GOPATH=/vagrant/mygo
vagrant@tc-vb-u-x-tc-helloworld-master:~$cd /vagrant/mygo/src/tc-helloworld-go-ws
vagrant@tc-vb-u-x-tc-helloworld-master:/vagrant/mygo/src/tc-helloworld-go-ws$ CGO_ENABLED=0 go build -a -tags netgo -ldflags '-w' .
vagrant@tc-vb-u-x-tc-helloworld-master:/vagrant/mygo/src/tc-helloworld-go-ws$ ./tc-helloworld-go-ws
Started, serving at 8080

```

Find the Vagrants' public_network IP address:

```bash
ifconfig
...
enp0s8: flags=4163<UP,BROADCAST,RUNNING,MULTICAST>  mtu 1500
        inet 10.0.1.23  netmask 255.255.255.0  broadcast 10.0.1.255
        inet6 fe80::a00:27ff:feb8:4314  prefixlen 64  scopeid 0x20<link>
        ether 08:00:27:b8:43:14  txqueuelen 1000  (Ethernet)
        RX packets 439  bytes 37772 (37.7 KB)
        RX errors 0  dropped 0  overruns 0  frame 0
        TX packets 47  bytes 4895 (4.8 KB)
        TX errors 0  dropped 0 overruns 0  carrier 0  collisions 0
...
```

Open the browser on the host and enter this address: 10.0.1.23:8080

You should see:

Hello World from Go (4.28MB) - tc-helloworld-go-ws - v.1.0, it took 140ns to run

Open the browser remote machine in your network and enter this address: 10.0.1.23:8080

You should see:

Hello World from Go (4.28MB) - tc-helloworld-go-ws - v.1.0, it took 140ns to run
