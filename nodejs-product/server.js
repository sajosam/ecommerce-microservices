const express = require("express");
const app = express();

const products = {
  "1": [{ id: 11, name: "Laptop", price: 100000 }],
  "2": [{ id: 22, name: "Phone", price: 50000 }],
};

app.get("/products/:product_id", (req, res) => {
  const productId = req.params.product_id;
  res.json(products[productId] || { error: "No products found" });
});

app.get("/health", (req, res) => {
  res.json({ status: "ok", message: "Product service is running on port 5004" });
});

app.listen(5004, () => console.log("Product service running on port 5004"));
