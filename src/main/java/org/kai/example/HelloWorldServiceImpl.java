package org.kai.example;

import com.google.common.flogger.FluentLogger;
import io.grpc.stub.StreamObserver;
import org.kai.example.HelloWorldServiceGrpc.HelloWorldServiceImplBase;

public class HelloWorldServiceImpl extends HelloWorldServiceImplBase {
  private final FluentLogger logger = FluentLogger.forEnclosingClass();

  @Override
  public void greeting(GreetingRequest request, StreamObserver<GreetingResponse> responseObserver) {
    logger.atInfo().log(request.toString());
    responseObserver.onNext(
        GreetingResponse.newBuilder().setGreeting("Hello " + request.getName()).build());
  }
}
