Vagrant.configure(2) do |config|

  config.vm.box = "terrywang/archlinux"
  config.vm.network "public_network"
  config.vm.provision "shell", inline: <<-SHELL

    sudo pacman -Syu
    sudo pacman -S --noconfirm go nginx git postgresql

    sudo mkdir /var/lib/postgres/data
    sudo chown -c -R postgres:postgres /var/lib/postgres

  SHELL

end
