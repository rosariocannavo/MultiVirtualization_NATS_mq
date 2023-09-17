# Create only one vm and run docker containers inside
Vagrant.configure("2") do |config|
    # Use the official Ubuntu 20.04 box
    config.vm.box = "ubuntu/focal64"
    config.vm.synced_folder "./vagrant", "/vagrant"
    # Forward port from vm to host
    config.vm.network "forwarded_port", guest: 8080, host: 8080  
  

    # Update and install Docker and Docker Compose
    config.vm.provision "shell", inline: <<-SHELL
      sudo apt-get update
      sudo apt-get install -y docker.io docker-compose
    SHELL
  
    # Specify the shell script to run when the VM boots
    config.vm.provision "shell", path: "bootstrap.sh"
  end
  