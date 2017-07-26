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
    vagrant box update
```
   
* run Virtual machine (VM)

  Install by running: 
  
  Virtualbox:
  
```bash
    vagrant up
```
