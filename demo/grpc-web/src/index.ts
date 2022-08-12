import { EchoServiceClient } from "./proto/EchoServiceClientPb";
import { EchoRequest } from "./proto/echo_pb";
var echoService = new EchoServiceClient("http://localhost:8080", null, null);

var request = new EchoRequest();
request.setMessage("Hello World!");

const call = echoService.echo(
  request,
  { "custom-header-1": "value1" },
  (err, response) => {
    console.log(response.getMessage());
  }
);
call.on("status", (status) => {
  console.log("log status:", status);
});
