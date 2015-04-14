Vagrant.configure(2) do |config|

  config.vm.box = "terrywang/archlinux"
  config.vm.network "public_network", :bridge => "en0: Wi-Fi (AirPort)"
  config.vm.provision "shell", inline: <<-SHELL

    sudo pacman -Syu --noconfirm
    sudo pacman -S --noconfirm go nginx git postgresql

    sudo mkdir /var/lib/postgres/data
    sudo chown -c -R postgres:postgres /var/lib/postgres

    sudo -i -u postgres bash -c 'initdb -D "/var/lib/postgres/data"'

    sudo systemctl start postgresql
    sudo systemctl enable postgresql

    psql -c "CREATE USER dev_application WITH PASSWORD 'dev_application';"

    createdb dev_application

  SHELL

end
