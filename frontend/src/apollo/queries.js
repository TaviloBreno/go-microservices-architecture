import { gql } from '@apollo/client';

// Query para buscar todos os pedidos
export const GET_ORDERS = gql`
  query GetOrders {
    orders {
      id
      userID
      productName
      quantity
      price
      status
      createdAt
    }
  }
`;

// Query para buscar todos os usuários
export const GET_USERS = gql`
  query GetUsers {
    users {
      id
      name
      email
      createdAt
    }
  }
`;

// Query para buscar todos os pagamentos
export const GET_PAYMENTS = gql`
  query GetPayments {
    payments {
      id
      orderID
      userID
      amount
      status
      paymentMethod
      createdAt
    }
  }
`;

// Query para buscar todas as notificações
export const GET_NOTIFICATIONS = gql`
  query GetNotifications {
    notifications {
      id
      orderID
      userID
      message
      type
      status
      createdAt
    }
  }
`;

// Query para buscar um pedido específico por ID
export const GET_ORDER = gql`
  query GetOrder($id: ID!) {
    order(id: $id) {
      id
      userID
      productName
      quantity
      price
      status
      createdAt
    }
  }
`;

// Query para buscar um usuário específico por ID
export const GET_USER = gql`
  query GetUser($id: ID!) {
    user(id: $id) {
      id
      name
      email
      createdAt
    }
  }
`;

// Query para buscar um pagamento específico por ID
export const GET_PAYMENT = gql`
  query GetPayment($id: ID!) {
    payment(id: $id) {
      id
      orderID
      userID
      amount
      status
      paymentMethod
      createdAt
    }
  }
`;

// Query para buscar resumo completo do pedido
export const GET_ORDER_SUMMARY = gql`
  query GetOrderSummary($orderID: ID!) {
    orderSummary(orderID: $orderID) {
      order {
        id
        userID
        productName
        quantity
        price
        status
        createdAt
      }
      user {
        id
        name
        email
        createdAt
      }
      payment {
        id
        orderID
        userID
        amount
        status
        paymentMethod
        createdAt
      }
      notifications {
        id
        orderID
        userID
        message
        type
        status
        createdAt
      }
    }
  }
`;

// Query para buscar todos os dados da dashboard
export const GET_DASHBOARD_DATA = gql`
  query GetDashboardData {
    orders {
      id
      userID
      productName
      quantity
      price
      status
      createdAt
    }
    users {
      id
      name
      email
      createdAt
    }
    payments {
      id
      orderID
      userID
      amount
      status
      paymentMethod
      createdAt
    }
    notifications {
      id
      orderID
      userID
      message
      type
      status
      createdAt
    }
  }
`;

// Mutation para criar um novo pedido
export const CREATE_ORDER = gql`
  mutation CreateOrder($input: CreateOrderInput!) {
    createOrder(input: $input) {
      id
      userID
      productName
      quantity
      price
      status
      createdAt
    }
  }
`;

// Mutation para criar um novo usuário
export const CREATE_USER = gql`
  mutation CreateUser($input: CreateUserInput!) {
    createUser(input: $input) {
      id
      name
      email
      createdAt
    }
  }
`;

// Query para health check
export const GET_HEALTH = gql`
  query GetHealth {
    health
  }
`;