from flask import Flask, request, jsonify
import csv

app = Flask(__name__)

# Load user tokens from the CSV file
def load_users(filename):
    users = {}
    with open(filename, newline='') as csvfile:
        reader = csv.DictReader(csvfile)
        for row in reader:
            users[row['user']] = row['token']
    return users

# Load the users at startup
USER_DB = load_users('example-users.csv')

@app.route('/', methods=['GET'])
def index():
    return jsonify({"message": "Hello, world! To get a token, send a GET request to /token?user=<user>"})

@app.route('/token', methods=['GET'])
def get_token():
    user = request.args.get('user')
    if not user:
        return jsonify({"error": "User is required"}), 400
    
    token = USER_DB.get(user)
    if token:
        return jsonify({"user": user, "token": token})
    else:
        return jsonify({"error": f"User {user} not found"}), 404

if __name__ == '__main__':
    app.run(host='localhost', port=8080, debug=True)

