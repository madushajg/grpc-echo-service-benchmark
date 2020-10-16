import ballerina/grpc;
import ballerina/log;

service HelloWorld on new grpc:Listener(9090) {
    resource function hello(grpc:Caller caller, string name, grpc:Headers headers) {
        log:printDebug(string `Server received hello from ${name}`);
        string message = string `Hello ${name}`;

        grpc:Error? err = caller->send(message, new);
        if (err is grpc:Error) {
            log:printError("Error from Connector: " + err.message());
        }

        grpc:Error? result = caller->complete();
        if (result is grpc:Error) {
            log:printError("Error in sending completed notification to caller",
                err = result);
        }
    }
}
