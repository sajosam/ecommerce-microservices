from flask import Flask, jsonify

app = Flask(__name__)

payments = {
    "1": [{"id": 1001, "user_id": 1, "amount": 100000, "status": "Completed"}],
    "2": [{"id": 1002, "user_id": 2, "amount": 50000, "status": "Pending"}]
}

@app.route("/payments/<user_id>")
def get_payments(user_id):
    return jsonify(payments.get(user_id, {"error": "No payments found"}))

if __name__ == "__main__":
    app.run(host="0.0.0.0", port=5003)
