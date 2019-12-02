FROM gitpod/workspace-mongodb

RUN mongod --dbpath /workspace/data --fork --logpath /workspace/log &&
RUN git clone https://github.com/jjpera/estudo-spring-historico.git /workspace/dependence
