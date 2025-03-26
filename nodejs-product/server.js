const express = require("express");
const app = express();

const products = {
  "1": [{ id: 1, name: "Laptop", price: 1000 }],
  "2": [{ id: 2, name: "Phone", price: 500 }],
};

app.get("/products/:user_id", (req, res) => {
  const userId = req.params.user_id;
  res.json(products[userId] || { error: "No products found" });
});

app.listen(5004, () => console.log("Product service running on port 5004"));
