/**
 * @fileoverview gRPC-Web generated client stub for Echo
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck


import * as grpcWeb from 'grpc-web';

import * as echo_pb from './echo_pb';


export class EchoServiceClient {
  client_: grpcWeb.AbstractClientBase;
  hostname_: string;
  credentials_: null | { [index: string]: string; };
  options_: null | { [index: string]: any; };

  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; }) {
    if (!options) options = {};
    if (!credentials) credentials = {};
    options['format'] = 'binary';

    this.client_ = new grpcWeb.GrpcWebClientBase(options);
    this.hostname_ = hostname;
    this.credentials_ = credentials;
    this.options_ = options;
  }

  methodDescriptorEcho = new grpcWeb.MethodDescriptor(
    '/Echo.EchoService/Echo',
    grpcWeb.MethodType.UNARY,
    echo_pb.EchoRequest,
    echo_pb.EchoResponse,
    (request: echo_pb.EchoRequest) => {
      return request.serializeBinary();
    },
    echo_pb.EchoResponse.deserializeBinary
  );

  echo(
    request: echo_pb.EchoRequest,
    metadata: grpcWeb.Metadata | null): Promise<echo_pb.EchoResponse>;

  echo(
    request: echo_pb.EchoRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: echo_pb.EchoResponse) => void): grpcWeb.ClientReadableStream<echo_pb.EchoResponse>;

  echo(
    request: echo_pb.EchoRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: echo_pb.EchoResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/Echo.EchoService/Echo',
        request,
        metadata || {},
        this.methodDescriptorEcho,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/Echo.EchoService/Echo',
    request,
    metadata || {},
    this.methodDescriptorEcho);
  }

  methodDescriptorHelloWorld = new grpcWeb.MethodDescriptor(
    '/Echo.EchoService/HelloWorld',
    grpcWeb.MethodType.UNARY,
    echo_pb.EchoRequest,
    echo_pb.EchoResponse,
    (request: echo_pb.EchoRequest) => {
      return request.serializeBinary();
    },
    echo_pb.EchoResponse.deserializeBinary
  );

  helloWorld(
    request: echo_pb.EchoRequest,
    metadata: grpcWeb.Metadata | null): Promise<echo_pb.EchoResponse>;

  helloWorld(
    request: echo_pb.EchoRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: echo_pb.EchoResponse) => void): grpcWeb.ClientReadableStream<echo_pb.EchoResponse>;

  helloWorld(
    request: echo_pb.EchoRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: echo_pb.EchoResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/Echo.EchoService/HelloWorld',
        request,
        metadata || {},
        this.methodDescriptorHelloWorld,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/Echo.EchoService/HelloWorld',
    request,
    metadata || {},
    this.methodDescriptorHelloWorld);
  }

}

