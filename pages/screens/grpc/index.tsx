import { GrpcWebFetchTransport } from "@protobuf-ts/grpcweb-transport";
import { PingPongServiceClient } from "gen/node/pingpong/v1/pingpong.client";
import React from "react";

export const GrpcScreen = () => {
  const handleReq = () => {
    const client = new PingPongServiceClient(
      new GrpcWebFetchTransport({
        baseUrl: "http://localhost:3000",
        format: "text",
      })
    );
    const call = client.pingPong({ ping: "ping" }, {});
    call.response.then((value) => {
      console.log(value);
    });
  };

  return (
    <div>
      <button onClick={handleReq}>send request</button>
    </div>
  );
};
