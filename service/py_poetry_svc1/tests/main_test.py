# content of test_sample.py
from app import main
from fastapi.testclient import TestClient


client = TestClient(main.app)

def test_sum():
  assert main.sum(1, 2) == 3

def test_main_route():
    response = client.get("/")
    assert response.status_code == 200
    assert response.json() == {"message": "Hello World"}