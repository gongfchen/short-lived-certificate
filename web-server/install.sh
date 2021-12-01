# create the directory structure for compiling source
mkdir -p ~/.local/src
cd ~/.local/src

# download pcre (could be in repository if root will install for you)
wget http://kent.dl.sourceforge.net/sourceforge/pcre/pcre-7.8.tar.gz
tar -xzvf pcre-7.8.tar.gz

# download nginx
wget http://nginx.org/download/nginx-1.8.1.tar.gz
tar -xzvf nginx-1.8.1.tar.gz

# compile and install it
cd nginx-1.8.1
./configure --prefix=~/.local/opt/nginx --with-pcre=~/.local/src/pcre-7.8 --with-http_ssl_module
make
make install
