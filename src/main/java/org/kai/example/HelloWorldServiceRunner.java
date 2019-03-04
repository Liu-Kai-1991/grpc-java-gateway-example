package org.kai.example;

import com.google.common.flogger.FluentLogger;
import io.grpc.Server;
import io.grpc.ServerBuilder;
import java.io.IOException;

public class HelloWorldServiceRunner {

  private int port = 9091;
  private Server server;
  private final FluentLogger logger = FluentLogger.forEnclosingClass();

  private void start() throws IOException {

    server = ServerBuilder.forPort(port)
        .addService(new HelloWorldServiceImpl())
        .build()
        .start();
    logger.atInfo().log("Server started, listening on " + port);
    Runtime.getRuntime().addShutdownHook(new Thread() {
      @Override
      public void run() {
        logger.atInfo().log("*** shutting down gRPC server since JVM is shutting down");
        stopService();
        logger.atInfo().log("*** server shut down");
      }
    });
  }

  private void stopService() {
    if (server != null) {
      server.shutdown();
    }
  }

  private void blockUntilShutdown() throws InterruptedException {
    if (server != null) {
      server.awaitTermination();
    }
  }

  public static void main(String[] args) throws IOException, InterruptedException {
    final HelloWorldServiceRunner server = new HelloWorldServiceRunner();
    server.start();
    server.blockUntilShutdown();
  }
}
