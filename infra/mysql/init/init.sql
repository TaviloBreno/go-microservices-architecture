-- ============================================================================
-- MICROSERVICES DATABASE INITIALIZATION
-- ============================================================================
-- Creates databases, tables and inserts sample data for all microservices
-- ============================================================================

-- Create Databases
CREATE DATABASE IF NOT EXISTS user_service;
CREATE DATABASE IF NOT EXISTS order_service;
CREATE DATABASE IF NOT EXISTS catalog_service;
CREATE DATABASE IF NOT EXISTS payment_service;
CREATE DATABASE IF NOT EXISTS notification_service;

-- ============================================================================
-- USER SERVICE
-- ============================================================================
USE user_service;

CREATE TABLE IF NOT EXISTS users (
    id VARCHAR(36) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    phone VARCHAR(20),
    address TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX idx_email (email),
    INDEX idx_created_at (created_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Insert sample users
INSERT INTO users (id, name, email, password_hash, phone, address) VALUES
('user-001', 'João Silva', 'joao.silva@email.com', '$2a$10$rWN7Y9Z8Y9Z8Y9Z8Y9Z8YuXxXxXxXxXxXxXxXxXxXxXxXxXxX', '+55 11 98765-4321', 'Rua das Flores, 123 - São Paulo, SP'),
('user-002', 'Maria Santos', 'maria.santos@email.com', '$2a$10$rWN7Y9Z8Y9Z8Y9Z8Y9Z8YuXxXxXxXxXxXxXxXxXxXxXxXxXxX', '+55 21 98765-1234', 'Av. Atlântica, 456 - Rio de Janeiro, RJ'),
('user-003', 'Pedro Oliveira', 'pedro.oliveira@email.com', '$2a$10$rWN7Y9Z8Y9Z8Y9Z8Y9Z8YuXxXxXxXxXxXxXxXxXxXxXxXxXxX', '+55 31 98765-5678', 'Rua dos Goitacazes, 789 - Belo Horizonte, MG'),
('user-004', 'Ana Costa', 'ana.costa@email.com', '$2a$10$rWN7Y9Z8Y9Z8Y9Z8Y9Z8YuXxXxXxXxXxXxXxXxXxXxXxXxXxX', '+55 41 98765-8765', 'Rua XV de Novembro, 321 - Curitiba, PR'),
('user-005', 'Carlos Souza', 'carlos.souza@email.com', '$2a$10$rWN7Y9Z8Y9Z8Y9Z8Y9Z8YuXxXxXxXxXxXxXxXxXxXxXxXxXxX', '+55 51 98765-4567', 'Av. Independência, 654 - Porto Alegre, RS'),
('user-006', 'Beatriz Lima', 'beatriz.lima@email.com', '$2a$10$rWN7Y9Z8Y9Z8Y9Z8Y9Z8YuXxXxXxXxXxXxXxXxXxXxXxXxXxX', '+55 85 98765-3456', 'Rua Dragão do Mar, 987 - Fortaleza, CE'),
('user-007', 'Rafael Alves', 'rafael.alves@email.com', '$2a$10$rWN7Y9Z8Y9Z8Y9Z8Y9Z8YuXxXxXxXxXxXxXxXxXxXxXxXxXxX', '+55 71 98765-2345', 'Av. Sete de Setembro, 147 - Salvador, BA'),
('user-008', 'Juliana Ferreira', 'juliana.ferreira@email.com', '$2a$10$rWN7Y9Z8Y9Z8Y9Z8Y9Z8YuXxXxXxXxXxXxXxXxXxXxXxXxXxX', '+55 61 98765-6789', 'SQN 312, Bloco A - Brasília, DF'),
('user-009', 'Lucas Rodrigues', 'lucas.rodrigues@email.com', '$2a$10$rWN7Y9Z8Y9Z8Y9Z8Y9Z8YuXxXxXxXxXxXxXxXxXxXxXxXxXxX', '+55 81 98765-9876', 'Rua da Aurora, 258 - Recife, PE'),
('user-010', 'Fernanda Martins', 'fernanda.martins@email.com', '$2a$10$rWN7Y9Z8Y9Z8Y9Z8Y9Z8YuXxXxXxXxXxXxXxXxXxXxXxXxXxX', '+55 19 98765-1357', 'Av. Francisco Glicério, 369 - Campinas, SP');

-- ============================================================================
-- CATALOG SERVICE
-- ============================================================================
USE catalog_service;

CREATE TABLE IF NOT EXISTS categories (
    id VARCHAR(36) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX idx_name (name)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE IF NOT EXISTS products (
    id VARCHAR(36) PRIMARY KEY,
    category_id VARCHAR(36) NOT NULL,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    price DECIMAL(10, 2) NOT NULL,
    stock_quantity INT NOT NULL DEFAULT 0,
    image_url VARCHAR(500),
    sku VARCHAR(100) UNIQUE,
    is_active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (category_id) REFERENCES categories(id),
    INDEX idx_category (category_id),
    INDEX idx_name (name),
    INDEX idx_price (price),
    INDEX idx_active (is_active)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Insert sample categories
INSERT INTO categories (id, name, description) VALUES
('cat-001', 'Eletrônicos', 'Produtos eletrônicos e gadgets'),
('cat-002', 'Livros', 'Livros físicos e digitais'),
('cat-003', 'Roupas', 'Vestuário e acessórios'),
('cat-004', 'Casa e Decoração', 'Itens para casa e decoração'),
('cat-005', 'Esportes', 'Artigos esportivos e fitness'),
('cat-006', 'Alimentos', 'Alimentos e bebidas'),
('cat-007', 'Beleza', 'Produtos de beleza e cuidados pessoais'),
('cat-008', 'Brinquedos', 'Brinquedos e jogos');

-- Insert sample products
INSERT INTO products (id, category_id, name, description, price, stock_quantity, image_url, sku, is_active) VALUES
-- Eletrônicos
('prod-001', 'cat-001', 'Smartphone XYZ Pro', 'Smartphone de última geração com câmera tripla de 108MP', 2999.99, 50, 'https://via.placeholder.com/300', 'SMART-XYZ-001', TRUE),
('prod-002', 'cat-001', 'Notebook Ultra 15"', 'Notebook com processador i7, 16GB RAM, SSD 512GB', 4599.90, 30, 'https://via.placeholder.com/300', 'NOTE-ULT-002', TRUE),
('prod-003', 'cat-001', 'Fone Bluetooth Premium', 'Fone de ouvido com cancelamento de ruído ativo', 599.00, 100, 'https://via.placeholder.com/300', 'FONE-BT-003', TRUE),
('prod-004', 'cat-001', 'Smart Watch Fit', 'Relógio inteligente com monitor cardíaco e GPS', 899.90, 75, 'https://via.placeholder.com/300', 'WATCH-FIT-004', TRUE),
('prod-005', 'cat-001', 'Tablet 10" WiFi', 'Tablet Android com tela Full HD e 128GB', 1299.00, 40, 'https://via.placeholder.com/300', 'TAB-10-005', TRUE),

-- Livros
('prod-006', 'cat-002', 'Clean Code - Robert Martin', 'Guia essencial para escrever código limpo', 89.90, 200, 'https://via.placeholder.com/300', 'BOOK-CC-006', TRUE),
('prod-007', 'cat-002', 'Domain-Driven Design', 'Atacando as complexidades no coração do software', 95.00, 150, 'https://via.placeholder.com/300', 'BOOK-DDD-007', TRUE),
('prod-008', 'cat-002', 'The Pragmatic Programmer', 'Seu caminho para a maestria', 79.90, 180, 'https://via.placeholder.com/300', 'BOOK-PP-008', TRUE),
('prod-009', 'cat-002', 'Design Patterns - GoF', 'Padrões de projeto reutilizáveis', 110.00, 120, 'https://via.placeholder.com/300', 'BOOK-DP-009', TRUE),
('prod-010', 'cat-002', 'Microservices Patterns', 'Padrões de microsserviços com exemplos', 115.90, 90, 'https://via.placeholder.com/300', 'BOOK-MS-010', TRUE),

-- Roupas
('prod-011', 'cat-003', 'Camiseta Básica - Preta', 'Camiseta 100% algodão tamanho M', 49.90, 500, 'https://via.placeholder.com/300', 'CAM-BAS-011', TRUE),
('prod-012', 'cat-003', 'Calça Jeans Slim', 'Calça jeans masculina corte slim', 159.90, 200, 'https://via.placeholder.com/300', 'CALCA-JS-012', TRUE),
('prod-013', 'cat-003', 'Jaqueta de Couro', 'Jaqueta de couro legítimo preta', 599.00, 80, 'https://via.placeholder.com/300', 'JAQ-COU-013', TRUE),
('prod-014', 'cat-003', 'Tênis Esportivo Pro', 'Tênis para corrida com amortecimento', 349.90, 150, 'https://via.placeholder.com/300', 'TEN-ESP-014', TRUE),

-- Casa e Decoração
('prod-015', 'cat-004', 'Luminária LED Moderna', 'Luminária de mesa com controle de intensidade', 129.90, 60, 'https://via.placeholder.com/300', 'LUM-LED-015', TRUE),
('prod-016', 'cat-004', 'Quadro Decorativo Abstrato', 'Quadro 60x80cm com moldura', 189.00, 45, 'https://via.placeholder.com/300', 'QUA-ABS-016', TRUE),
('prod-017', 'cat-004', 'Jogo de Cama Casal', 'Jogo de cama 100% algodão 4 peças', 199.90, 100, 'https://via.placeholder.com/300', 'CAMA-JG-017', TRUE),

-- Esportes
('prod-018', 'cat-005', 'Bola de Futebol Oficial', 'Bola oficial tamanho regulamentar', 129.90, 120, 'https://via.placeholder.com/300', 'BOL-FUT-018', TRUE),
('prod-019', 'cat-005', 'Halteres 5kg (par)', 'Par de halteres revestidos', 89.90, 80, 'https://via.placeholder.com/300', 'HALT-5K-019', TRUE),
('prod-020', 'cat-005', 'Tapete de Yoga Premium', 'Tapete antiderrapante 180x60cm', 119.00, 95, 'https://via.placeholder.com/300', 'TAP-YOG-020', TRUE),

-- Alimentos
('prod-021', 'cat-006', 'Café Premium Torrado 500g', 'Café 100% arábica torrado e moído', 29.90, 300, 'https://via.placeholder.com/300', 'CAF-PRM-021', TRUE),
('prod-022', 'cat-006', 'Chocolate Belga 70% Cacau', 'Barra de chocolate belga 100g', 19.90, 250, 'https://via.placeholder.com/300', 'CHOC-BEL-022', TRUE),
('prod-023', 'cat-006', 'Azeite Extra Virgem 500ml', 'Azeite de oliva extra virgem português', 39.90, 150, 'https://via.placeholder.com/300', 'AZE-EV-023', TRUE),

-- Beleza
('prod-024', 'cat-007', 'Perfume Masculino 100ml', 'Fragrância amadeirada intensa', 249.90, 70, 'https://via.placeholder.com/300', 'PERF-M-024', TRUE),
('prod-025', 'cat-007', 'Kit Skin Care Facial', 'Kit completo para cuidados faciais', 179.00, 85, 'https://via.placeholder.com/300', 'SKIN-KIT-025', TRUE),

-- Brinquedos
('prod-026', 'cat-008', 'Lego Creator 500 peças', 'Set de construção criativo', 199.90, 110, 'https://via.placeholder.com/300', 'LEGO-CR-026', TRUE),
('prod-027', 'cat-008', 'Boneca Fashion Doll', 'Boneca articulada com acessórios', 89.90, 150, 'https://via.placeholder.com/300', 'BON-FD-027', TRUE),
('prod-028', 'cat-008', 'Carrinho Controle Remoto', 'Carrinho de corrida 1:16 com controle', 159.90, 90, 'https://via.placeholder.com/300', 'CAR-RC-028', TRUE);

-- ============================================================================
-- ORDER SERVICE
-- ============================================================================
USE order_service;

CREATE TABLE IF NOT EXISTS orders (
    id VARCHAR(36) PRIMARY KEY,
    user_id VARCHAR(36) NOT NULL,
    status ENUM('pending', 'confirmed', 'processing', 'shipped', 'delivered', 'cancelled') NOT NULL DEFAULT 'pending',
    total_amount DECIMAL(10, 2) NOT NULL,
    payment_status ENUM('pending', 'paid', 'failed', 'refunded') NOT NULL DEFAULT 'pending',
    shipping_address TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX idx_user (user_id),
    INDEX idx_status (status),
    INDEX idx_created_at (created_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE IF NOT EXISTS order_items (
    id VARCHAR(36) PRIMARY KEY,
    order_id VARCHAR(36) NOT NULL,
    product_id VARCHAR(36) NOT NULL,
    product_name VARCHAR(255) NOT NULL,
    quantity INT NOT NULL,
    unit_price DECIMAL(10, 2) NOT NULL,
    subtotal DECIMAL(10, 2) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (order_id) REFERENCES orders(id) ON DELETE CASCADE,
    INDEX idx_order (order_id),
    INDEX idx_product (product_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Insert sample orders
INSERT INTO orders (id, user_id, status, total_amount, payment_status, shipping_address) VALUES
('order-001', 'user-001', 'delivered', 3599.89, 'paid', 'Rua das Flores, 123 - São Paulo, SP'),
('order-002', 'user-002', 'shipped', 189.90, 'paid', 'Av. Atlântica, 456 - Rio de Janeiro, RJ'),
('order-003', 'user-003', 'processing', 4599.90, 'paid', 'Rua dos Goitacazes, 789 - Belo Horizonte, MG'),
('order-004', 'user-004', 'confirmed', 349.90, 'paid', 'Rua XV de Novembro, 321 - Curitiba, PR'),
('order-005', 'user-005', 'pending', 1498.90, 'pending', 'Av. Independência, 654 - Porto Alegre, RS'),
('order-006', 'user-001', 'delivered', 599.00, 'paid', 'Rua das Flores, 123 - São Paulo, SP'),
('order-007', 'user-006', 'shipped', 899.90, 'paid', 'Rua Dragão do Mar, 987 - Fortaleza, CE'),
('order-008', 'user-007', 'processing', 279.70, 'paid', 'Av. Sete de Setembro, 147 - Salvador, BA'),
('order-009', 'user-008', 'confirmed', 199.90, 'paid', 'SQN 312, Bloco A - Brasília, DF'),
('order-010', 'user-009', 'pending', 2999.99, 'pending', 'Rua da Aurora, 258 - Recife, PE');

-- Insert sample order items
INSERT INTO order_items (id, order_id, product_id, product_name, quantity, unit_price, subtotal) VALUES
-- Order 1
('item-001', 'order-001', 'prod-001', 'Smartphone XYZ Pro', 1, 2999.99, 2999.99),
('item-002', 'order-001', 'prod-003', 'Fone Bluetooth Premium', 1, 599.00, 599.00),

-- Order 2
('item-003', 'order-002', 'prod-006', 'Clean Code - Robert Martin', 1, 89.90, 89.90),
('item-004', 'order-002', 'prod-008', 'The Pragmatic Programmer', 1, 79.90, 79.90),

-- Order 3
('item-005', 'order-003', 'prod-002', 'Notebook Ultra 15"', 1, 4599.90, 4599.90),

-- Order 4
('item-006', 'order-004', 'prod-014', 'Tênis Esportivo Pro', 1, 349.90, 349.90),

-- Order 5
('item-007', 'order-005', 'prod-005', 'Tablet 10" WiFi', 1, 1299.00, 1299.00),
('item-008', 'order-005', 'prod-015', 'Luminária LED Moderna', 1, 129.90, 129.90),

-- Order 6
('item-009', 'order-006', 'prod-003', 'Fone Bluetooth Premium', 1, 599.00, 599.00),

-- Order 7
('item-010', 'order-007', 'prod-004', 'Smart Watch Fit', 1, 899.90, 899.90),

-- Order 8
('item-011', 'order-008', 'prod-006', 'Clean Code - Robert Martin', 1, 89.90, 89.90),
('item-012', 'order-008', 'prod-007', 'Domain-Driven Design', 1, 95.00, 95.00),
('item-013', 'order-008', 'prod-009', 'Design Patterns - GoF', 1, 110.00, 110.00),

-- Order 9
('item-014', 'order-009', 'prod-017', 'Jogo de Cama Casal', 1, 199.90, 199.90),

-- Order 10
('item-015', 'order-010', 'prod-001', 'Smartphone XYZ Pro', 1, 2999.99, 2999.99);

-- ============================================================================
-- PAYMENT SERVICE
-- ============================================================================
USE payment_service;

CREATE TABLE IF NOT EXISTS payments (
    id VARCHAR(36) PRIMARY KEY,
    order_id VARCHAR(36) NOT NULL,
    user_id VARCHAR(36) NOT NULL,
    amount DECIMAL(10, 2) NOT NULL,
    payment_method ENUM('credit_card', 'debit_card', 'pix', 'boleto', 'paypal') NOT NULL,
    status ENUM('pending', 'processing', 'approved', 'failed', 'refunded', 'cancelled') NOT NULL DEFAULT 'pending',
    transaction_id VARCHAR(255),
    card_last_digits VARCHAR(4),
    card_brand VARCHAR(50),
    installments INT DEFAULT 1,
    failure_reason TEXT,
    processed_at TIMESTAMP NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX idx_order (order_id),
    INDEX idx_user (user_id),
    INDEX idx_status (status),
    INDEX idx_created_at (created_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Insert sample payments
INSERT INTO payments (id, order_id, user_id, amount, payment_method, status, transaction_id, card_last_digits, card_brand, installments, processed_at) VALUES
('pay-001', 'order-001', 'user-001', 3599.89, 'credit_card', 'approved', 'TXN-2024-001234', '1234', 'Visa', 3, '2024-01-15 10:30:00'),
('pay-002', 'order-002', 'user-002', 189.90, 'pix', 'approved', 'PIX-2024-005678', NULL, NULL, 1, '2024-01-16 14:20:00'),
('pay-003', 'order-003', 'user-003', 4599.90, 'credit_card', 'approved', 'TXN-2024-009876', '5678', 'Mastercard', 6, '2024-01-17 09:15:00'),
('pay-004', 'order-004', 'user-004', 349.90, 'debit_card', 'approved', 'TXN-2024-011223', '9012', 'Elo', 1, '2024-01-18 16:45:00'),
('pay-005', 'order-005', 'user-005', 1498.90, 'credit_card', 'pending', NULL, '3456', 'Visa', 2, NULL),
('pay-006', 'order-006', 'user-001', 599.00, 'credit_card', 'approved', 'TXN-2024-013344', '1234', 'Visa', 1, '2024-01-19 11:30:00'),
('pay-007', 'order-007', 'user-006', 899.90, 'pix', 'approved', 'PIX-2024-007788', NULL, NULL, 1, '2024-01-20 13:25:00'),
('pay-008', 'order-008', 'user-007', 279.70, 'credit_card', 'approved', 'TXN-2024-015566', '7890', 'Mastercard', 2, '2024-01-21 10:10:00'),
('pay-009', 'order-009', 'user-008', 199.90, 'boleto', 'approved', 'BOL-2024-009988', NULL, NULL, 1, '2024-01-22 15:00:00'),
('pay-010', 'order-010', 'user-009', 2999.99, 'credit_card', 'pending', NULL, '2468', 'Visa', 4, NULL);

-- ============================================================================
-- NOTIFICATION SERVICE
-- ============================================================================
USE notification_service;

CREATE TABLE IF NOT EXISTS notifications (
    id VARCHAR(36) PRIMARY KEY,
    user_id VARCHAR(36) NOT NULL,
    type ENUM('email', 'sms', 'push', 'in_app') NOT NULL,
    channel VARCHAR(50) NOT NULL,
    subject VARCHAR(255),
    message TEXT NOT NULL,
    status ENUM('pending', 'sent', 'failed', 'read') NOT NULL DEFAULT 'pending',
    sent_at TIMESTAMP NULL,
    read_at TIMESTAMP NULL,
    error_message TEXT,
    metadata JSON,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX idx_user (user_id),
    INDEX idx_type (type),
    INDEX idx_status (status),
    INDEX idx_created_at (created_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Insert sample notifications
INSERT INTO notifications (id, user_id, type, channel, subject, message, status, sent_at, metadata) VALUES
('notif-001', 'user-001', 'email', 'order_confirmation', 'Pedido Confirmado #order-001', 'Seu pedido foi confirmado e está sendo processado.', 'sent', '2024-01-15 10:31:00', '{"order_id": "order-001", "total": 3599.89}'),
('notif-002', 'user-001', 'push', 'payment_approved', 'Pagamento Aprovado', 'Seu pagamento foi aprovado com sucesso!', 'sent', '2024-01-15 10:30:30', '{"payment_id": "pay-001"}'),
('notif-003', 'user-001', 'email', 'order_shipped', 'Pedido Enviado #order-001', 'Seu pedido foi enviado e está a caminho.', 'sent', '2024-01-16 08:00:00', '{"order_id": "order-001", "tracking_code": "BR123456789BR"}'),
('notif-004', 'user-002', 'email', 'order_confirmation', 'Pedido Confirmado #order-002', 'Seu pedido foi confirmado e está sendo processado.', 'sent', '2024-01-16 14:21:00', '{"order_id": "order-002", "total": 189.90}'),
('notif-005', 'user-003', 'sms', 'payment_approved', 'Pagamento Aprovado', 'Pagamento de R$ 4599.90 aprovado!', 'sent', '2024-01-17 09:16:00', '{"payment_id": "pay-003"}'),
('notif-006', 'user-004', 'email', 'order_confirmation', 'Pedido Confirmado #order-004', 'Seu pedido foi confirmado e está sendo processado.', 'sent', '2024-01-18 16:46:00', '{"order_id": "order-004", "total": 349.90}'),
('notif-007', 'user-005', 'in_app', 'order_pending', 'Pedido Aguardando Pagamento', 'Seu pedido está aguardando confirmação de pagamento.', 'sent', '2024-01-19 10:00:00', '{"order_id": "order-005"}'),
('notif-008', 'user-006', 'email', 'order_confirmation', 'Pedido Confirmado #order-007', 'Seu pedido foi confirmado e está sendo processado.', 'sent', '2024-01-20 13:26:00', '{"order_id": "order-007", "total": 899.90}'),
('notif-009', 'user-007', 'push', 'order_processing', 'Pedido em Processamento', 'Seu pedido está sendo preparado para envio.', 'sent', '2024-01-21 10:11:00', '{"order_id": "order-008"}'),
('notif-010', 'user-008', 'email', 'payment_confirmed', 'Pagamento Confirmado - Boleto', 'Recebemos a confirmação do pagamento do seu boleto.', 'sent', '2024-01-22 15:01:00', '{"payment_id": "pay-009"}');

-- ============================================================================
-- GRANT PERMISSIONS
-- ============================================================================
GRANT ALL PRIVILEGES ON user_service.* TO 'microservices'@'%';
GRANT ALL PRIVILEGES ON order_service.* TO 'microservices'@'%';
GRANT ALL PRIVILEGES ON catalog_service.* TO 'microservices'@'%';
GRANT ALL PRIVILEGES ON payment_service.* TO 'microservices'@'%';
GRANT ALL PRIVILEGES ON notification_service.* TO 'microservices'@'%';
FLUSH PRIVILEGES;

-- ============================================================================
-- SUMMARY
-- ============================================================================
-- Databases created: 5
-- Tables created: 9
-- Sample data inserted:
--   - Users: 10
--   - Categories: 8
--   - Products: 28
--   - Orders: 10
--   - Order Items: 15
--   - Payments: 10
--   - Notifications: 10
-- ============================================================================
