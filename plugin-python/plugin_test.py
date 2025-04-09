import json
import pytest

from main import AuthorizeServicer
from proto import auth_pb2


class MockContext:
    def __init__(self):
        self.code = None
        self.details = None

    def set_code(self, code):
        self.code = code

    def set_details(self, details):
        self.details = details


@pytest.fixture(scope="module")
def service():
    return AuthorizeServicer()


@pytest.fixture
def context():
    return MockContext()


@pytest.fixture
def host():
    return "http://localhost:8080"


def test_authorized_user(service, context, host):
    req = auth_pb2.GetRequest(user="example", host=host)
    res = service.Get(req, context)

    assert context.code is None, f"gRPC error: {context.details}"
    actual = json.loads(res.value)

    assert actual["code"] == 200
    assert actual["config"]["AmazonS3"]["Key"] == "key1"
    assert actual["config"]["AmazonS3"]["Secret"] == "secret1"


def test_unauthorized_user(service, context, host):
    req = auth_pb2.GetRequest(user="error", host=host)
    res = service.Get(req, context)

    assert context.code is None or context.code.value == 0  # OK
    actual = json.loads(res.value)

    assert actual["code"] == 401
