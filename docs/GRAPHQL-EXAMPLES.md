# 🔷 GraphQL Examples - BFF API

Exemplos práticos de queries e mutations para testar o BFF GraphQL.

Acesse: http://localhost:8080/graphql

---

## 👤 User Operations

### Create User

```graphql
mutation {
  createUser(input: {
    name: "João Silva"
    email: "joao.silva@example.com"
    password: "senha123"
  }) {
    id
    name
    email
    createdAt
  }
}
```

### List Users

```graphql
query {
  users {
    id
    name
    email
    createdAt
    updatedAt
  }
}
```

### Get User by ID

```graphql
query {
  user(id: "1") {
    id
    name
    email
    createdAt
    orders {
      id
      totalAmount
      status
    }
  }
}
```

### Update User

```graphql
mutation {
  updateUser(
    id: "1"
    input: {
      name: "João Silva Santos"
      email: "joao.santos@example.com"
    }
  ) {
    id
    name
    email
    updatedAt
  }
}
```

---

## 🛒 Order Operations

### Create Order

```graphql
mutation {
  createOrder(input: {
    userId: "1"
    items: [
      {
        productId: "101"
        quantity: 2
        price: 99.90
      }
      {
        productId: "102"
        quantity: 1
        price: 149.90
      }
    ]
    totalAmount: 349.70
  }) {
    id
    userId
    totalAmount
    status
    createdAt
    items {
      productId
      quantity
      price
      subtotal
    }
  }
}
```

### List Orders

```graphql
query {
  orders {
    id
    userId
    totalAmount
    status
    createdAt
    items {
      productId
      quantity
      price
    }
  }
}
```

### List Orders by User

```graphql
query {
  orders(userId: "1") {
    id
    totalAmount
    status
    createdAt
    items {
      productId
      quantity
      price
      subtotal
    }
  }
}
```

### Get Order by ID

```graphql
query {
  order(id: "1") {
    id
    userId
    totalAmount
    status
    createdAt
    items {
      productId
      quantity
      price
      subtotal
    }
    payment {
      id
      amount
      paymentMethod
      status
    }
  }
}
```

### Cancel Order

```graphql
mutation {
  cancelOrder(id: "1") {
    id
    status
    updatedAt
  }
}
```

---

## 💳 Payment Operations

### Process Payment

```graphql
mutation {
  processPayment(input: {
    orderId: "1"
    amount: 349.70
    paymentMethod: "credit_card"
    cardNumber: "4111111111111111"
    cardHolder: "JOAO SILVA"
    expiryDate: "12/25"
    cvv: "123"
  }) {
    id
    orderId
    amount
    paymentMethod
    status
    transactionId
    processedAt
  }
}
```

### Process Payment - PIX

```graphql
mutation {
  processPayment(input: {
    orderId: "2"
    amount: 199.90
    paymentMethod: "pix"
    pixKey: "joao@example.com"
  }) {
    id
    orderId
    amount
    paymentMethod
    status
    transactionId
    processedAt
  }
}
```

### Process Payment - Boleto

```graphql
mutation {
  processPayment(input: {
    orderId: "3"
    amount: 299.90
    paymentMethod: "boleto"
  }) {
    id
    orderId
    amount
    paymentMethod
    status
    boletoCode
    expiryDate
  }
}
```

### Get Payment by ID

```graphql
query {
  payment(id: "1") {
    id
    orderId
    amount
    paymentMethod
    status
    transactionId
    processedAt
  }
}
```

### List Payments by Order

```graphql
query {
  payments(orderId: "1") {
    id
    amount
    paymentMethod
    status
    processedAt
  }
}
```

### Refund Payment

```graphql
mutation {
  refundPayment(id: "1") {
    id
    orderId
    amount
    status
    refundedAt
  }
}
```

---

## 📦 Product/Catalog Operations

### Create Product

```graphql
mutation {
  createProduct(input: {
    name: "Notebook Gamer"
    description: "Notebook de alta performance para jogos"
    price: 4999.90
    stockQuantity: 10
    category: "Electronics"
  }) {
    id
    name
    description
    price
    stockQuantity
    category
    createdAt
  }
}
```

### List Products

```graphql
query {
  products {
    id
    name
    price
    stockQuantity
    category
  }
}
```

### List Products by Category

```graphql
query {
  products(category: "Electronics") {
    id
    name
    price
    stockQuantity
  }
}
```

### Get Product by ID

```graphql
query {
  product(id: "101") {
    id
    name
    description
    price
    stockQuantity
    category
    createdAt
    updatedAt
  }
}
```

### Update Stock

```graphql
mutation {
  updateStock(id: "101", quantity: 5) {
    id
    name
    stockQuantity
    updatedAt
  }
}
```

---

## 📧 Notification Operations

### Send Notification

```graphql
mutation {
  sendNotification(input: {
    userId: "1"
    type: "email"
    subject: "Pedido Confirmado"
    message: "Seu pedido #1 foi confirmado com sucesso!"
  }) {
    id
    userId
    type
    subject
    status
    sentAt
  }
}
```

### List Notifications

```graphql
query {
  notifications(userId: "1") {
    id
    type
    subject
    message
    status
    sentAt
    createdAt
  }
}
```

---

## 🔄 Complex Queries (Multiple Services)

### Complete Order Flow

```graphql
# 1. Criar usuário
mutation CreateUser {
  createUser(input: {
    name: "Maria Santos"
    email: "maria@example.com"
    password: "senha123"
  }) {
    id
    name
    email
  }
}

# 2. Criar produto
mutation CreateProduct {
  createProduct(input: {
    name: "Mouse Gamer"
    description: "Mouse com RGB"
    price: 149.90
    stockQuantity: 50
    category: "Peripherals"
  }) {
    id
    name
    price
  }
}

# 3. Criar pedido
mutation CreateOrder {
  createOrder(input: {
    userId: "2"
    items: [
      {
        productId: "101"
        quantity: 1
        price: 149.90
      }
    ]
    totalAmount: 149.90
  }) {
    id
    totalAmount
    status
  }
}

# 4. Processar pagamento
mutation ProcessPayment {
  processPayment(input: {
    orderId: "1"
    amount: 149.90
    paymentMethod: "credit_card"
    cardNumber: "4111111111111111"
  }) {
    id
    status
    transactionId
  }
}

# 5. Consultar tudo
query CompleteOrder {
  user(id: "2") {
    name
    email
    orders {
      id
      totalAmount
      status
      items {
        productId
        quantity
        price
      }
      payment {
        paymentMethod
        status
      }
    }
  }
}
```

### Dashboard Query (Aggregate Data)

```graphql
query Dashboard {
  users {
    id
    name
    email
  }
  
  orders {
    id
    totalAmount
    status
  }
  
  products {
    id
    name
    price
    stockQuantity
  }
}
```

---

## 🧪 Test Scenarios

### Scenario 1: New Customer Purchase

```graphql
# Step 1: Create customer
mutation {
  createUser(input: {
    name: "Pedro Costa"
    email: "pedro@example.com"
    password: "senha123"
  }) {
    id
  }
}

# Step 2: Browse products
query {
  products(category: "Electronics") {
    id
    name
    price
    stockQuantity
  }
}

# Step 3: Create order
mutation {
  createOrder(input: {
    userId: "3"
    items: [
      { productId: "101", quantity: 1, price: 4999.90 }
    ]
    totalAmount: 4999.90
  }) {
    id
  }
}

# Step 4: Pay with PIX
mutation {
  processPayment(input: {
    orderId: "2"
    amount: 4999.90
    paymentMethod: "pix"
    pixKey: "pedro@example.com"
  }) {
    id
    status
  }
}
```

### Scenario 2: Order Tracking

```graphql
# Check order status
query {
  order(id: "1") {
    id
    status
    totalAmount
    createdAt
    items {
      productId
      quantity
      price
    }
    payment {
      paymentMethod
      status
      processedAt
    }
  }
}

# Check user's all orders
query {
  orders(userId: "1") {
    id
    totalAmount
    status
    createdAt
  }
}
```

### Scenario 3: Refund Process

```graphql
# Get payment details
query {
  payment(id: "1") {
    id
    orderId
    amount
    status
  }
}

# Process refund
mutation {
  refundPayment(id: "1") {
    id
    status
    refundedAt
  }
}

# Cancel order
mutation {
  cancelOrder(id: "1") {
    id
    status
  }
}
```

---

## 📊 Analytics Queries

### Revenue Analytics

```graphql
query {
  orders {
    id
    totalAmount
    status
    createdAt
  }
}
```

### Payment Methods Distribution

```graphql
query {
  payments(orderId: null) {
    paymentMethod
    amount
    status
  }
}
```

### Stock Status

```graphql
query {
  products {
    id
    name
    stockQuantity
    category
  }
}
```

---

## 🔍 Filtering and Pagination (Future)

```graphql
# Pagination example
query {
  orders(
    page: 1
    limit: 10
    sortBy: "createdAt"
    sortOrder: "DESC"
  ) {
    id
    totalAmount
    createdAt
  }
}

# Filtering example
query {
  products(
    category: "Electronics"
    priceMin: 100
    priceMax: 1000
    inStock: true
  ) {
    id
    name
    price
  }
}
```

---

## ⚠️ Error Handling Examples

### Invalid Input

```graphql
mutation {
  createUser(input: {
    name: "Test"
    email: "invalid-email"  # Invalid email format
    password: "123"         # Too short
  }) {
    id
  }
}

# Expected error:
# {
#   "errors": [{
#     "message": "Invalid email format",
#     "path": ["createUser"]
#   }]
# }
```

### Not Found

```graphql
query {
  user(id: "999999") {
    id
    name
  }
}

# Expected error:
# {
#   "errors": [{
#     "message": "User not found",
#     "path": ["user"]
#   }]
# }
```

### Insufficient Stock

```graphql
mutation {
  createOrder(input: {
    userId: "1"
    items: [
      {
        productId: "101"
        quantity: 1000  # More than available
        price: 99.90
      }
    ]
    totalAmount: 99900
  }) {
    id
  }
}

# Expected error:
# {
#   "errors": [{
#     "message": "Insufficient stock",
#     "path": ["createOrder"]
#   }]
# }
```

---

## 🚀 Performance Testing

### Batch Operations

```graphql
mutation {
  user1: createUser(input: { name: "User 1", email: "user1@test.com", password: "pass123" }) { id }
  user2: createUser(input: { name: "User 2", email: "user2@test.com", password: "pass123" }) { id }
  user3: createUser(input: { name: "User 3", email: "user3@test.com", password: "pass123" }) { id }
}
```

### Nested Query Performance

```graphql
query {
  users {
    id
    name
    orders {
      id
      totalAmount
      items {
        productId
        quantity
      }
      payment {
        status
        paymentMethod
      }
    }
  }
}
```

---

## 📝 Notes

- All IDs are strings (UUIDs in real implementation)
- Dates are in ISO 8601 format
- Amounts are in decimal format (e.g., 99.90)
- Payment methods: `credit_card`, `debit_card`, `pix`, `boleto`
- Order statuses: `pending`, `confirmed`, `processing`, `shipped`, `delivered`, `cancelled`
- Payment statuses: `pending`, `processing`, `completed`, `failed`, `refunded`

---

**Happy Testing! 🚀**
