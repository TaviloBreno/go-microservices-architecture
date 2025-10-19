import React, { useState } from 'react';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import { AuthProvider } from './contexts/AuthContext';
import { ThemeProvider } from './contexts/ThemeContext';
import ProtectedRoute from './components/ProtectedRoute';
import Sidebar from './components/Sidebar';
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from './components/ui/card';
import { Badge } from './components/ui/badge';
import { cn } from './lib/utils';
import LoginPage from './pages/LoginPage';
import ProductsPage from './pages/ProductsPage';
import UsersPage from './pages/UsersPage';

// Páginas com shadcn/ui
const Dashboard = () => (
  <div className="p-6 space-y-6">
    <div className="mb-6">
      <h1 className="text-3xl font-bold text-foreground mb-2">🎯 Dashboard</h1>
      <p className="text-muted-foreground">Frontend React funcionando com autenticação e shadcn/ui! 🚀</p>
    </div>
    
    <div className="grid grid-cols-1 md:grid-cols-3 gap-6">
      <Card>
        <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
          <CardTitle className="text-sm font-medium">📦 Total de Pedidos</CardTitle>
          <div className="text-2xl">📦</div>
        </CardHeader>
        <CardContent>
          <div className="text-2xl font-bold text-primary">--</div>
          <p className="text-xs text-muted-foreground">
            Aguardando conexão GraphQL
          </p>
        </CardContent>
      </Card>
      
      <Card>
        <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
          <CardTitle className="text-sm font-medium">💳 Pagamentos</CardTitle>
          <div className="text-2xl">💳</div>
        </CardHeader>
        <CardContent>
          <div className="text-2xl font-bold text-green-600">--</div>
          <p className="text-xs text-muted-foreground">
            Aguardando conexão GraphQL
          </p>
        </CardContent>
      </Card>
      
      <Card>
        <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
          <CardTitle className="text-sm font-medium">🔔 Notificações</CardTitle>
          <div className="text-2xl">🔔</div>
        </CardHeader>
        <CardContent>
          <div className="text-2xl font-bold text-purple-600">--</div>
          <p className="text-xs text-muted-foreground">
            Aguardando conexão GraphQL
          </p>
        </CardContent>
      </Card>
    </div>
    
    <Card>
      <CardHeader>
        <CardTitle className="flex items-center gap-2">
          🔧 Status da Aplicação
        </CardTitle>
        <CardDescription>
          Estado atual dos componentes e serviços
        </CardDescription>
      </CardHeader>
      <CardContent>
        <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
          <div className="space-y-3">
            <div className="flex items-center justify-between">
              <span className="text-sm text-foreground">React 18 + Vite</span>
              <Badge variant="success">Funcionando</Badge>
            </div>
            <div className="flex items-center justify-between">
              <span className="text-sm text-foreground">Shadcn/ui</span>
              <Badge variant="success">Carregado</Badge>
            </div>
            <div className="flex items-center justify-between">
              <span className="text-sm text-foreground">Autenticação</span>
              <Badge variant="success">Ativo</Badge>
            </div>
            <div className="flex items-center justify-between">
              <span className="text-sm text-foreground">Dark Mode</span>
              <Badge variant="success">Disponível</Badge>
            </div>
          </div>
          <div className="space-y-3">
            <div className="flex items-center justify-between">
              <span className="text-sm text-foreground">Sidebar Responsiva</span>
              <Badge variant="success">Funcionando</Badge>
            </div>
            <div className="flex items-center justify-between">
              <span className="text-sm text-foreground">BFF GraphQL</span>
              <Badge variant="warning">Aguardando</Badge>
            </div>
            <div className="flex items-center justify-between">
              <span className="text-sm text-foreground">Microserviços</span>
              <Badge variant="warning">Offline</Badge>
            </div>
          </div>
        </div>
      </CardContent>
    </Card>
  </div>
);

const OrdersPage = () => (
  <div className="p-6 space-y-6">
    <div className="mb-6">
      <h1 className="text-3xl font-bold text-foreground mb-2">📦 Pedidos</h1>
      <p className="text-muted-foreground">Gerenciamento de pedidos do sistema</p>
    </div>
    <Card>
      <CardHeader>
        <CardTitle>Lista de Pedidos</CardTitle>
        <CardDescription>
          Acompanhe todos os pedidos do sistema em tempo real
        </CardDescription>
      </CardHeader>
      <CardContent>
        <p className="text-muted-foreground">Aguardando integração GraphQL para exibir dados reais...</p>
      </CardContent>
    </Card>
  </div>
);

const PaymentsPage = () => (
  <div className="p-6 space-y-6">
    <div className="mb-6">
      <h1 className="text-3xl font-bold text-foreground mb-2">💳 Pagamentos</h1>
      <p className="text-muted-foreground">Controle de transações financeiras</p>
    </div>
    <Card>
      <CardHeader>
        <CardTitle>Transações</CardTitle>
        <CardDescription>
          Monitore todas as transações e pagamentos
        </CardDescription>
      </CardHeader>
      <CardContent>
        <p className="text-muted-foreground">Aguardando integração GraphQL para exibir dados reais...</p>
      </CardContent>
    </Card>
  </div>
);

const NotificationsPage = () => (
  <div className="p-6 space-y-6">
    <div className="mb-6">
      <h1 className="text-3xl font-bold text-foreground mb-2">🔔 Notificações</h1>
      <p className="text-muted-foreground">Central de notificações do sistema</p>
    </div>
    <Card>
      <CardHeader>
        <CardTitle>Central de Notificações</CardTitle>
        <CardDescription>
          Todas as notificações enviadas pelo sistema
        </CardDescription>
      </CardHeader>
      <CardContent>
        <p className="text-muted-foreground">Aguardando integração GraphQL para exibir dados reais...</p>
      </CardContent>
    </Card>
  </div>
);

// Layout principal com sidebar e dark mode
const MainLayout = ({ children }) => {
  const [sidebarCollapsed, setSidebarCollapsed] = useState(false);

  return (
    <ThemeProvider>
      <div className="min-h-screen bg-background flex">
        <Sidebar 
          isCollapsed={sidebarCollapsed} 
          setIsCollapsed={setSidebarCollapsed}
        />
        <main className={cn(
          "flex-1 overflow-hidden transition-all duration-300",
          "lg:ml-0"
        )}>
          <div className="lg:pl-4 pt-16 lg:pt-0">
            {children}
          </div>
        </main>
      </div>
    </ThemeProvider>
  );
};

// Página 404
const NotFound = () => (
  <div className="min-h-screen bg-background flex items-center justify-center">
    <Card className="w-full max-w-md">
      <CardHeader className="text-center">
        <div className="text-6xl font-bold text-muted mb-4">404</div>
        <CardTitle>Página não encontrada</CardTitle>
        <CardDescription>
          A página que você está procurando não existe.
        </CardDescription>
      </CardHeader>
      <CardContent className="text-center">
        <a href="/" className="btn-primary">
          Voltar ao Dashboard
        </a>
      </CardContent>
    </Card>
  </div>
);

function App() {
  return (
    <AuthProvider>
      <Router>
        <div className="App">
          <Routes>
            {/* Login sem dark mode */}
            <Route path="/login" element={<LoginPage />} />
            
            {/* Dashboard com dark mode através do ProtectedRoute */}
            <Route path="/" element={
              <ProtectedRoute>
                <MainLayout>
                  <Dashboard />
                </MainLayout>
              </ProtectedRoute>
            } />
            <Route path="/products" element={
              <ProtectedRoute>
                <MainLayout>
                  <ProductsPage />
                </MainLayout>
              </ProtectedRoute>
            } />
            <Route path="/users" element={
              <ProtectedRoute>
                <MainLayout>
                  <UsersPage />
                </MainLayout>
              </ProtectedRoute>
            } />
            <Route path="/orders" element={
              <ProtectedRoute>
                <MainLayout>
                  <OrdersPage />
                </MainLayout>
              </ProtectedRoute>
            } />
            <Route path="/payments" element={
              <ProtectedRoute>
                <MainLayout>
                  <PaymentsPage />
                </MainLayout>
              </ProtectedRoute>
            } />
            <Route path="/notifications" element={
              <ProtectedRoute>
                <MainLayout>
                  <NotificationsPage />
                </MainLayout>
              </ProtectedRoute>
            } />
            <Route path="*" element={<NotFound />} />
          </Routes>
        </div>
      </Router>
    </AuthProvider>
  );
}

export default App;
