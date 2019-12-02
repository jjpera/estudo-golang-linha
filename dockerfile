FROM gitpod/workspace-full:latest

USER root

# install and start a mongodb
RUN sudo apt-get install gnupg
RUN wget -qO - https://www.mongodb.org/static/pgp/server-4.2.asc | sudo apt-key add -
RUN echo "deb [ arch=amd64 ] https://repo.mongodb.org/apt/ubuntu bionic/mongodb-org/4.2 multiverse" | sudo tee /etc/apt/sources.list.d/mongodb-org-4.2.list
RUN sudo apt-get update
RUN sudo apt-get install -y mongodb-org
RUN echo "mongodb-org hold" | sudo dpkg --set-selections
RUN echo "mongodb-org-server hold" | sudo dpkg --set-selections
RUN echo "mongodb-org-shell hold" | sudo dpkg --set-selections
RUN echo "mongodb-org-mongos hold" | sudo dpkg --set-selections
RUN echo "mongodb-org-tools hold" | sudo dpkg --set-selections
RUN sudo service mongod start

# install golang
# RUN sudo apt-get update
# RUN sudo apt-get -y upgrade
# RUN wget https://dl.google.com/go/go1.13.3.linux-amd64.tar.gz
# RUN sudo tar -xvf go1.13.3.linux-amd64.tar.gz
# RUN sudo mv go /usr/local

# verificar para substituir de RUN para ENV
# set 3 environment variables as GOROOT, GOPATH and PATH
# RUN export GOROOT=/usr/local/go
# verificar path para setar
# RUN export GOPATH=$HOME/Projects/Proj1
# RUN export PATH=$GOPATH/bin:$GOROOT/bin:$PATH


# start a mock for historico

# liberando a porta do projeto
EXPORT 8082
