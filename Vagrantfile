Vagrant.configure(2) do |config|

  config.vm.box = "terrywang/archlinux"
  config.vm.network "public_network", :bridge => "en0: Wi-Fi (AirPort)"
  config.vm.provision "shell", inline: <<-SHELL

    sudo pacman -Syu --noconfirm
    sudo pacman -S --noconfirm go nginx git postgresql

    sudo mkdir /var/lib/postgres/data
    sudo chown -c -R postgres:postgres /var/lib/postgres

    sudo -i -u postgres initdb -D '/var/lib/postgres/data'

    sudo systemctl start postgresql
    sudo systemctl enable postgresql

    sudo -i -u postgres psql -c "CREATE USER vagrant WITH PASSWORD 'vagrant';"
    sudo -i -u postgres createdb vagrant

    echo 'export GOPATH=/vagrant/go;' >> /home/vagrant/.bashrc
    echo 'export APP_PORT=8080;' >> /home/vagrant/.bashrc
    echo 'export APP_DBURL=postgres://postgres:postgres@localhost/vagrant?sslmode=disable;' >> /home/vagrant/.bashrc
    echo 'export APP_ENV=dev;' >> /home/vagrant/.bashrc

  SHELL

end
