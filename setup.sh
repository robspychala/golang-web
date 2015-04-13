echo "Setting up directories..."

mkdir -p go/src
mkdir -p go/bin
mkdir -p go/pkg

curl -L https://github.com/robspychala/golang-web/archive/master.tar.gz | tar -xz app
