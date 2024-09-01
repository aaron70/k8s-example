import requests, random, string, time

def randomword(length):
   letters = string.ascii_lowercase
   return ''.join(random.choice(letters) for i in range(length))

def send_request(id):
    # The API endpoint
    url = f"http://localhost:5000/api/send?client={id}"
    
    # A GET request to the API
    response = requests.get(url)
    
    # Print the response
    response.json()

if __name__ == "__main__":
    
    for i in range(3):
        print(send_request(randomword()))
        time.sleep(3)
