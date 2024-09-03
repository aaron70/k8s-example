import requests, time

if __name__ == "__main__":
    while True:
        URL = "http://server-svc.default.svc.cluster.local:80/api/send"
        requests.get(URL)
        time.sleep(5)
