echo "Setting up directories..."

mkdir -p go/{src,bin,pkg}

echo "Downloading vagrant config..."

curl -s -L https://github.com/robspychala/golang-web/archive/master.tar.gz | tar -xvz --strip-components 1 --include *./Vagrantfile

echo "Downloading example go files..."

curl -s -L https://github.com/robspychala/golang-web/archive/master.tar.gz | tar -xvz -C go/src --strip-components 1 --include *./app

echo "Setting up vagrant..."

vagrant up

echo "Finished."
