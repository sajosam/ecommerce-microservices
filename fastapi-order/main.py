from fastapi import FastAPI
import requests
import os

app = FastAPI()

orders = {
    "1": [{"id": 101, "user_id": 1, "item": "Laptop", "product_id": 11}],
    "2": [{"id": 102, "user_id": 2, "item": "Phone", "product_id": 22}]
}

USER_SERVICE_URL = f"http://{os.getenv('USER_SERVICE')}/users"
PAYMENT_SERVICE_URL = f"http://{os.getenv('PAYMENT_SERVICE')}/payments"
PRODUCT_SERVICE_URL = f"http://{os.getenv('PRODUCT_SERVICE')}/products"

@app.get("/orders/{user_id}")
def get_orders(user_id: str):
    user_orders = orders.get(user_id, [])
    if not user_orders:
        return {"error": "No orders found"}

    result = []
    for order in user_orders:
        # ✅ Correct paths
        user = requests.get(f"{USER_SERVICE_URL}/{order['user_id']}").json()
        payments = requests.get(f"{PAYMENT_SERVICE_URL}/{order['user_id']}").json()
        product = requests.get(f"{PRODUCT_SERVICE_URL}/{order['product_id']}").json()

        result.append({
            "order": order,
            "user": user,
            "payments": payments,
            "product": product
        })

    return result

if __name__ == "__main__":
    import uvicorn
    uvicorn.run(app, host="0.0.0.0", port=5002)
