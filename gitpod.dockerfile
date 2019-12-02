FROM gitpod/workspace-mongodb

RUN mongod --dbpath /workspace/data --fork --logpath /workspace/log &&
RUN git clone https://github.com/jjpera/estudo-spring-historico.git /workspace/dependence
RUN java -jar /workspace/dependence/estudo-spring-historico/target/swagger-spring-1.0.0.jar
