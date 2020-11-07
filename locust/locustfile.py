import time, random
from locust import HttpUser, task, between

class QuickstartUser(HttpUser):
    wait_time = between(1, 2)

    @task
    def nodejs_get(self):
        l = random.randint(1, 100)
        s = "/nodejs/write?line=%d" % (l)
        self.client.get(s)
    @task
    def nodejs_post(self):
        num1 = random.randint(0, 1000)
        num2 = random.randint(0, 1000)
        self.client.post("/nodejs/sha256", data = {"num1": num1, "num2": num2})
    @task
    def go_get(self):
        l = random.randint(1, 100)
        s = "/go/write?line=%d" % (l)
        self.client.get(s)
    @task
    def go_post(self):
        num1 = random.randint(0, 1000)
        num2 = random.randint(0, 1000)
        self.client.post("/go/sha256", data = {"num1": num1, "num2": num2})
