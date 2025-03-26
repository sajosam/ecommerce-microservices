from fastapi import FastAPI

app = FastAPI()

orders = {
    "1": [{"id": 101, "user_id": 1, "item": "Laptop"}],
    "2": [{"id": 102, "user_id": 2, "item": "Phone"}]
}

@app.get("/orders/{user_id}")
def get_orders(user_id: str):
    return orders.get(user_id, {"error": "No orders found"})

if __name__ == "__main__":
    import uvicorn
    uvicorn.run(app, host="0.0.0.0", port=5002)
