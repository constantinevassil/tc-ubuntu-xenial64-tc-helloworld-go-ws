# vi: set ft=ruby :
 
Vagrant.configure("2") do |config|
    config.vm.box = "bento/ubuntu-17.04"

    config.vm.provider "virtualbox" do |v|
      v.memory = 1024
      v.cpus = 1
      v.gui = true
    end
 
    config.vm.define "tc-vb-u-x-tc-helloworld-master" do |node|
      node.vm.hostname = "tc-vb-u-x-tc-helloworld-master"
      node.vm.network "public_network"
    end 
end
