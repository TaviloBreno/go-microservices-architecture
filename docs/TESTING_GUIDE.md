# 🧪 Guia de Testes - Frontend com Dados Reais

## 📋 Checklist de Testes

Use este guia para validar que tudo está funcionando corretamente.

---

## ✅ Passo 1: Verificar Containers

```powershell
# Verificar se todos os containers estão rodando
docker-compose ps
```

**Resultado esperado:**
```
NAME              STATUS
mysql             Up (healthy)
rabbitmq          Up
bff-graphql       Up
catalog-service   Up
order-service     Up
payment-service   Up
user-service      Up
notification      Up
```

---

## ✅ Passo 2: Testar API REST do BFF

### 2.1 Testar Produtos

```powershell
curl http://localhost:8080/api/products
```

**Resultado esperado:**
```json
[
  {
    "id": 1,
    "name": "Smartphone XYZ Pro",
    "description": "Smartphone de última geração com câmera de 108MP",
    "price": 2999.99,
    "stock": 50,
    "image_url": "https://via.placeholder.com/300x200?text=Smartphone",
    "sku": "ELEC-001",
    "category": "Eletrônicos"
  },
  ...
]
```

✅ **Validações:**
- [ ] Retorna array com 28 produtos
- [ ] Cada produto tem todos os campos
- [ ] Preços são números decimais
- [ ] Categorias estão corretas

---

### 2.2 Testar Usuários

```powershell
curl http://localhost:8080/api/users
```

**Resultado esperado:**
```json
[
  {
    "id": 1,
    "name": "João Silva",
    "email": "joao.silva@email.com",
    "phone": "(11) 98765-4321",
    "address": "Rua das Flores, 123 - São Paulo/SP",
    "created_at": "2024-01-01T00:00:00Z",
    "updated_at": "2024-01-01T00:00:00Z"
  },
  ...
]
```

✅ **Validações:**
- [ ] Retorna array com 10 usuários
- [ ] Todos têm nome, email, telefone
- [ ] Endereços estão completos
- [ ] Datas estão no formato ISO

---

### 2.3 Testar Pedidos

```powershell
curl http://localhost:8080/api/orders
```

**Resultado esperado:**
```json
[
  {
    "id": 1,
    "user_id": 1,
    "user_name": "João Silva",
    "user_email": "joao.silva@email.com",
    "total": 2999.99,
    "status": "delivered",
    "created_at": "2024-01-15T10:30:00Z"
  },
  ...
]
```

✅ **Validações:**
- [ ] Retorna array com 10 pedidos
- [ ] Pedidos têm usuário associado
- [ ] Total é um número
- [ ] Status são válidos (delivered, shipped, processing, confirmed, pending)

---

### 2.4 Testar Pagamentos

```powershell
curl http://localhost:8080/api/payments
```

**Resultado esperado:**
```json
[
  {
    "id": 1,
    "user_id": 1,
    "user_name": "João Silva",
    "user_email": "joao.silva@email.com",
    "amount": 2999.99,
    "payment_method": "credit_card",
    "status": "approved",
    "created_at": "2024-01-15T10:35:00Z"
  },
  ...
]
```

✅ **Validações:**
- [ ] Retorna array com 10 pagamentos
- [ ] Pagamentos têm usuário associado
- [ ] Métodos são válidos (credit_card, pix, debit_card, boleto)
- [ ] Status são válidos (approved, pending)

---

## ✅ Passo 3: Iniciar Frontend

```powershell
cd frontend
npm run dev
```

**Resultado esperado:**
```
VITE v7.1.10  ready in 658 ms

➜  Local:   http://localhost:5173/
```

✅ **Validações:**
- [ ] Vite inicia sem erros
- [ ] Porta 5173 está disponível
- [ ] Console não mostra erros

---

## ✅ Passo 4: Testar Login

1. Abra o navegador em: http://localhost:5173
2. Você verá a tela de login
3. Faça login com:
   - **Email:** `admin@example.com`
   - **Senha:** `admin123`

✅ **Validações:**
- [ ] Tela de login aparece
- [ ] Formulário aceita input
- [ ] Login redireciona para dashboard
- [ ] Token é salvo (verificar localStorage)

---

## ✅ Passo 5: Testar Dashboard

Após login, você deve ver o Dashboard.

✅ **Validações:**
- [ ] Dashboard carrega sem erros
- [ ] Sidebar está visível à esquerda
- [ ] Cards de estatísticas aparecem
- [ ] Tema claro/escuro funciona (botão 🌙)

---

## ✅ Passo 6: Testar Página de Produtos

1. Clique em "🛍️ Produtos" na sidebar
2. Aguarde o carregamento

### Validações Visuais

✅ **Estatísticas (topo da página):**
- [ ] Total de Produtos: **28**
- [ ] Valor Total Estoque: **R$ 106.889,90** (aproximado)
- [ ] Categorias: **8**
- [ ] Estoque Baixo: **12**

✅ **Grid de Produtos:**
- [ ] Exibe 28 cards de produtos
- [ ] Cada card tem:
  - [ ] Imagem do produto
  - [ ] Nome do produto
  - [ ] Descrição
  - [ ] Preço formatado (R$ X.XXX,XX)
  - [ ] SKU
  - [ ] Badge de categoria
  - [ ] Badge de estoque (verde/amarelo/vermelho)

✅ **Produtos Específicos (verificar alguns):**
- [ ] **Smartphone XYZ Pro**
  - Preço: R$ 2.999,99
  - Estoque: 50 un. (🟢)
  - Categoria: Eletrônicos

- [ ] **Notebook Ultra 15"**
  - Preço: R$ 4.599,90
  - Estoque: 30 un. (🟡)
  - Categoria: Eletrônicos

- [ ] **Clean Code - Robert Martin**
  - Preço: R$ 89,90
  - Estoque: 200 un. (🟢)
  - Categoria: Livros

- [ ] **Camiseta Básica**
  - Preço: R$ 49,90
  - Estoque: 500 un. (🟢)
  - Categoria: Roupas

✅ **Responsividade:**
- [ ] Grid se adapta em telas menores
- [ ] Cards são legíveis em mobile
- [ ] Scroll funciona suavemente

---

## ✅ Passo 7: Testar Página de Usuários

1. Clique em "👥 Usuários" na sidebar
2. Aguarde o carregamento

### Validações Visuais

✅ **Estatísticas:**
- [ ] Total de Usuários: **10**
- [ ] Usuários Ativos: **10**
- [ ] Média de Cadastros: **1/mês**

✅ **Lista de Usuários:**
- [ ] Exibe 10 cards de usuários
- [ ] Cada card tem:
  - [ ] Nome completo
  - [ ] ID do usuário
  - [ ] Email com ícone 📧
  - [ ] Telefone com ícone 📱
  - [ ] Endereço completo com ícone 📍
  - [ ] Data de criação
  - [ ] Data de atualização

✅ **Usuários Específicos (verificar alguns):**
- [ ] **João Silva**
  - Email: joao.silva@email.com
  - Telefone: (11) 98765-4321
  - Endereço: Rua das Flores, 123 - São Paulo/SP

- [ ] **Maria Santos**
  - Email: maria.santos@email.com
  - Telefone: (11) 98765-4322
  - Endereço: Av. Paulista, 456 - São Paulo/SP

- [ ] **Pedro Oliveira**
  - Email: pedro.oliveira@email.com
  - Telefone: (11) 98765-4323
  - Endereço: Rua Augusta, 789 - São Paulo/SP

---

## ✅ Passo 8: Testar Página de Pedidos

1. Clique em "📦 Pedidos" na sidebar
2. Aguarde o carregamento

### Validações Visuais

✅ **Tabela de Pedidos:**
- [ ] Exibe 10 linhas de pedidos
- [ ] Colunas: ID, Usuário, Total, Status, Data
- [ ] Auto-atualiza a cada 5 segundos

✅ **Pedidos Específicos (verificar alguns):**
- [ ] **Pedido #1**
  - Usuário: João Silva
  - Email: joao.silva@email.com
  - Total: R$ 2.999,99
  - Status: 🟢 Entregue (delivered)
  - Data formatada (DD/MM/AA)

- [ ] **Pedido #2**
  - Usuário: Maria Santos
  - Status: 📦 Enviado (shipped)

- [ ] **Pedido #3**
  - Usuário: Pedro Oliveira
  - Status: 🟡 Processando (processing)

✅ **Status Badges:**
- [ ] Verde (🟢): delivered, confirmed
- [ ] Azul (🔵): shipped
- [ ] Amarelo (🟡): processing, pending
- [ ] Todos os status são traduzidos para português

✅ **Auto-Refresh:**
- [ ] Aguarde 5 segundos
- [ ] Observe se a tabela recarrega (pode ver no Network do DevTools)
- [ ] Dados permanecem consistentes

---

## ✅ Passo 9: Testar Navegação

### Sidebar

✅ **Links:**
- [ ] Clicar em cada item da sidebar
- [ ] Todos os links funcionam
- [ ] URL muda corretamente
- [ ] Página carrega sem erros

✅ **Itens da Sidebar:**
- [ ] 🎯 Dashboard → `/`
- [ ] 🛍️ Produtos → `/products`
- [ ] 👥 Usuários → `/users`
- [ ] 📦 Pedidos → `/orders`
- [ ] 💳 Pagamentos → `/payments`
- [ ] 🔔 Notificações → `/notifications`

✅ **Funcionalidades:**
- [ ] Sidebar pode ser recolhida (botão ◀️/▶️)
- [ ] Item ativo está destacado
- [ ] Hover mostra descrição
- [ ] Mobile menu funciona (☰)

---

## ✅ Passo 10: Testar Loading States

### Simular Conexão Lenta

1. Abra DevTools (F12)
2. Vá em Network
3. Selecione "Slow 3G"
4. Recarregue a página de produtos

✅ **Validações:**
- [ ] Skeleton loader aparece
- [ ] Cards de loading têm animação
- [ ] Não há flash de conteúdo
- [ ] Transição suave para dados reais

---

## ✅ Passo 11: Testar Error Handling

### Simular Erro da API

1. Pare o BFF: `docker-compose stop bff-graphql`
2. Tente acessar `/products`

✅ **Validações:**
- [ ] Mensagem de erro aparece
- [ ] Ícone de erro (⚠️) é exibido
- [ ] Botão "Tentar novamente" funciona
- [ ] Erro não quebra a página

3. Reinicie o BFF: `docker-compose start bff-graphql`
4. Clique em "Tentar novamente"

✅ **Validações:**
- [ ] Dados carregam normalmente
- [ ] Mensagem de erro desaparece

---

## ✅ Passo 12: Testar Formatações

### Valores Monetários

✅ **Verificar formatação:**
- [ ] R$ 2.999,99 (não R$ 2999.99)
- [ ] R$ 4.599,90 (não R$ 4599.9)
- [ ] R$ 89,90 (não R$ 89.9)
- [ ] Sempre 2 casas decimais
- [ ] Separador de milhar correto

### Datas

✅ **Verificar formatação:**
- [ ] DD/MM/AAAA (não YYYY-MM-DD)
- [ ] Exemplo: 15/01/2024
- [ ] Sem horas/minutos na lista

---

## ✅ Passo 13: Testar Responsividade

### Desktop (1920x1080)

✅ **Validações:**
- [ ] Sidebar fixa à esquerda
- [ ] Grid de produtos: 3 colunas
- [ ] Tabelas têm espaço adequado
- [ ] Estatísticas: 4 colunas

### Tablet (768x1024)

✅ **Validações:**
- [ ] Sidebar pode ser recolhida
- [ ] Grid de produtos: 2 colunas
- [ ] Tabelas são scrollable
- [ ] Estatísticas: 2 colunas

### Mobile (375x667)

✅ **Validações:**
- [ ] Sidebar vira menu hamburguer
- [ ] Grid de produtos: 1 coluna
- [ ] Cards ocupam tela inteira
- [ ] Estatísticas: 2 colunas
- [ ] Touch funciona corretamente

---

## ✅ Passo 14: Testar Dark Mode

1. Clique no botão 🌙 no canto superior direito

✅ **Validações Dark Mode:**
- [ ] Background fica escuro
- [ ] Texto fica claro
- [ ] Cards têm fundo escuro
- [ ] Badges mantêm cores
- [ ] Transição é suave

2. Clique novamente para voltar ao modo claro

✅ **Validações Light Mode:**
- [ ] Background fica claro
- [ ] Texto fica escuro
- [ ] Cards têm fundo branco
- [ ] Preferência é salva (localStorage)

---

## ✅ Passo 15: Testar Console

Abra DevTools (F12) → Console

✅ **Validações:**
- [ ] Sem erros no console
- [ ] Sem warnings de React
- [ ] Sem CORS errors
- [ ] Sem 404 errors
- [ ] Logs de conexão BFF aparecem (opcional)

---

## ✅ Passo 16: Testar Network

Abra DevTools (F12) → Network

✅ **Requests esperados:**
- [ ] GET http://localhost:8080/api/products → 200 OK
- [ ] GET http://localhost:8080/api/users → 200 OK
- [ ] GET http://localhost:8080/api/orders → 200 OK
- [ ] Response time < 100ms
- [ ] Response size < 20KB

✅ **Headers:**
- [ ] Content-Type: application/json
- [ ] Access-Control-Allow-Origin: *

---

## ✅ Passo 17: Testar Performance

### Lighthouse (DevTools → Lighthouse)

Execute Lighthouse na página `/products`

✅ **Scores esperados:**
- [ ] Performance: > 80
- [ ] Accessibility: > 90
- [ ] Best Practices: > 90
- [ ] SEO: > 80

---

## 🐛 Troubleshooting

### Problema: Produtos não carregam

**Sintomas:** Tela em branco ou erro "Failed to fetch"

**Soluções:**
1. Verificar se BFF está rodando: `docker-compose ps`
2. Testar API manualmente: `curl http://localhost:8080/api/products`
3. Verificar logs do BFF: `docker-compose logs bff-graphql`
4. Reiniciar BFF: `docker-compose restart bff-graphql`

---

### Problema: CORS Error

**Sintomas:** Erro no console: "blocked by CORS policy"

**Soluções:**
1. Verificar se BFF tem CORS habilitado no código
2. Reiniciar BFF: `docker-compose restart bff-graphql`
3. Limpar cache do navegador (Ctrl + Shift + Delete)

---

### Problema: Dados diferentes do esperado

**Sintomas:** Número de produtos diferente de 28

**Soluções:**
1. Verificar dados no MySQL:
   ```powershell
   docker exec -it mysql mysql -u microservices -pmicro123 -e "SELECT COUNT(*) FROM catalog_service.products"
   ```
2. Se vazio, recriar banco:
   ```powershell
   docker-compose down mysql
   Remove-Item -Recurse -Force ./data/mysql
   docker-compose up -d mysql
   ```

---

### Problema: Frontend não inicia

**Sintomas:** `npm run dev` falha

**Soluções:**
1. Reinstalar dependências:
   ```powershell
   cd frontend
   Remove-Item -Recurse -Force node_modules
   npm install
   npm run dev
   ```

---

### Problema: Página em branco após login

**Sintomas:** Login funciona mas página fica em branco

**Soluções:**
1. Verificar console (F12) por erros
2. Limpar localStorage:
   ```javascript
   localStorage.clear()
   ```
3. Relogar

---

## 📊 Checklist Final

Use este checklist para validação completa:

### Backend
- [ ] MySQL rodando e com dados
- [ ] BFF rodando na porta 8080
- [ ] 4 endpoints REST funcionando
- [ ] CORS configurado

### Frontend
- [ ] Vite rodando na porta 5173
- [ ] Login funciona
- [ ] ProductsPage exibe 28 produtos
- [ ] UsersPage exibe 10 usuários
- [ ] OrdersPage exibe 10 pedidos
- [ ] Navegação funciona
- [ ] Loading states funcionam
- [ ] Error handling funciona

### Visual
- [ ] Formatação de moeda correta
- [ ] Formatação de data correta
- [ ] Badges coloridos corretos
- [ ] Dark mode funciona
- [ ] Responsivo em todos os tamanhos

### Performance
- [ ] Sem erros no console
- [ ] Requests < 100ms
- [ ] Lighthouse > 80

---

## 🎉 Teste Aprovado!

Se todos os itens acima estão ✅, parabéns! 

Seu sistema está **100% funcional** com dados reais do banco de dados!

---

## 📞 Suporte

Se encontrar problemas:

1. Verifique os logs:
   ```powershell
   docker-compose logs bff-graphql
   docker-compose logs mysql
   ```

2. Consulte a documentação:
   - [INSTALLATION.md](./INSTALLATION.md)
   - [FRONTEND_INTEGRATION.md](./FRONTEND_INTEGRATION.md)

3. Reinicie tudo:
   ```powershell
   docker-compose down
   docker-compose up -d
   cd frontend ; npm run dev
   ```

---

**Criado:** 2024-01-20  
**Última atualização:** 2024-01-20  
**Versão:** 1.0.0
