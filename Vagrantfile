# vi: set ft=ruby :
 
Vagrant.configure("2") do |config|
    config.vm.box = "bento/ubuntu-17.04"
    ##config.vm.synced_folder ENV['HOME'], "/myhome", type: "nfs"

    config.vm.provider "virtualbox" do |v|
      v.memory = 1024
      v.cpus = 1
      v.gui = true
    end
 
    config.vm.define "tc-vb-u-x-tc-helloworld-master" do |node|
      node.vm.hostname = "tc-vb-u-x-tc-helloworld-master"
      #node.vm.network "private_network", type: "dhcp"
      node.vm.network "public_network"
      #node.vm.network "public_network", type: "dhcp"
      #node.vm.network "public_network", ip: "192.168.0.99" #192.168.33.10
      #node.vm.network "public_network", ip: "192.168.0.17" #192.168.33.10
      #node.vm.network :forwarded_port, guest: 80, host: 4567

      # Use NFS for shared folders for better performance
      ##node.vm.synced_folder '.', '/vagrant', nfs: true
    end 
end
