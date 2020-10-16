FROM ballerina/ballerina:swan-lake-preview3 AS ballerina-builder
USER root
ADD . /src
WORKDIR /src

RUN ballerina build -a --skip-tests

FROM openjdk:8-jdk-alpine
WORKDIR /home/ballerina

RUN addgroup troupe \
    && adduser -S -s /bin/bash -g 'ballerina' -G troupe -D ballerina \
    && apk add --update --no-cache bash \
    && chown -R ballerina:troupe /usr/bin/java \
    && rm -rf /var/cache/apk/*

COPY --from=ballerina-builder /src/target/bin/service.jar /home/ballerina
RUN chown ballerina /home/ballerina/service.jar

EXPOSE  9090
USER ballerina

CMD java -XX:+UnlockExperimentalVMOptions -XX:+UseCGroupMemoryLimitForHeap -XX:MaxRAMFraction=3 -jar service.jar \
    || cat ballerina-internal.log
