FROM maven:3-jdk-8
RUN mkdir -p /usr/src/app
WORKDIR /usr/src/app
ADD . /usr/src/app
RUN mvn install
RUN mvn package
WORKDIR /usr/src/app/target
ENTRYPOINT ["java", "-jar", "javaconsumer-0.0.1-SNAPSHOT.jar"]
