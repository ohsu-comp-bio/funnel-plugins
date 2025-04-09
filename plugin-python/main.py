from concurrent import futures
import sys
import time
import grpc
import requests

from proto import auth_pb2
from proto import auth_pb2_grpc
from grpc_health.v1.health import HealthServicer
from grpc_health.v1 import health_pb2, health_pb2_grpc

class AuthorizeServicer(auth_pb2_grpc.AuthorizeServicer):
    """Python version of the Go plugin's Authorize.Get implementation."""

    def Get(self, request, context):
        user = request.user
        host = request.host

        if not user:
            context.set_code(grpc.StatusCode.INVALID_ARGUMENT)
            context.set_details("User is required.")
            return auth_pb2.GetResponse()

        try:
            url = f"{host}/token?user={user}"
            print(f"Calling: {url}")
            resp = requests.get(url)

            if resp.status_code != 200:
                context.set_code(grpc.StatusCode.PERMISSION_DENIED)
                context.set_details(f"Auth server returned {resp.status_code}")
                return auth_pb2.GetResponse()

            body = resp.content
            print("Response body:", body.decode())

            return auth_pb2.GetResponse(value=body)

        except Exception as e:
            context.set_code(grpc.StatusCode.INTERNAL)
            context.set_details(f"Error during request: {str(e)}")
            return auth_pb2.GetResponse()

def serve():
    # Required health service for go-plugin handshake
    health = HealthServicer()
    health.set("plugin", health_pb2.HealthCheckResponse.SERVING)

    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    auth_pb2_grpc.add_AuthorizeServicer_to_server(AuthorizeServicer(), server)
    health_pb2_grpc.add_HealthServicer_to_server(health, server)

    server.add_insecure_port('127.0.0.1:1234')
    server.start()

    # Required handshake output for go-plugin
    print("1|1|tcp|127.0.0.1:1234|grpc")
    sys.stdout.flush()

    try:
        while True:
            time.sleep(60 * 60 * 24)
    except KeyboardInterrupt:
        server.stop(0)

if __name__ == '__main__':
    serve()
