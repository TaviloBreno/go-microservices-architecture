import React from 'react';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';

// Componentes sem Apollo por enquanto
import Header from './components/HeaderSimple';

// Páginas simplificadas
const Dashboard = () => (
  <div className="min-h-screen bg-gray-50">
    <div className="max-w-7xl mx-auto py-6 sm:px-6 lg:px-8">
      <div className="px-4 py-6 sm:px-0">
        <div className="border-4 border-dashed border-gray-200 rounded-lg p-8">
          <h1 className="text-3xl font-bold text-gray-900 mb-4">🎯 Dashboard</h1>
          <p className="text-gray-600 mb-6">Frontend React funcionando! 🚀</p>
          
          <div className="grid grid-cols-1 md:grid-cols-3 gap-6">
            <div className="bg-white p-6 rounded-lg shadow">
              <h3 className="text-lg font-semibold text-gray-900 mb-2">📦 Pedidos</h3>
              <p className="text-3xl font-bold text-blue-600">--</p>
              <p className="text-sm text-gray-500">Aguardando conexão GraphQL</p>
            </div>
            
            <div className="bg-white p-6 rounded-lg shadow">
              <h3 className="text-lg font-semibold text-gray-900 mb-2">💳 Pagamentos</h3>
              <p className="text-3xl font-bold text-green-600">--</p>
              <p className="text-sm text-gray-500">Aguardando conexão GraphQL</p>
            </div>
            
            <div className="bg-white p-6 rounded-lg shadow">
              <h3 className="text-lg font-semibold text-gray-900 mb-2">🔔 Notificações</h3>
              <p className="text-3xl font-bold text-purple-600">--</p>
              <p className="text-sm text-gray-500">Aguardando conexão GraphQL</p>
            </div>
          </div>
          
          <div className="mt-8 p-4 bg-blue-50 border border-blue-200 rounded-lg">
            <h4 className="text-lg font-medium text-blue-900 mb-2">🔧 Status da Aplicação</h4>
            <ul className="text-sm text-blue-800">
              <li>✅ React 18 + Vite funcionando</li>
              <li>✅ Tailwind CSS carregado</li>
              <li>✅ React Router configurado</li>
              <li>⏳ Aguardando BFF GraphQL (porta 8080)</li>
            </ul>
          </div>
        </div>
      </div>
    </div>
  </div>
);

const OrdersPage = () => (
  <div className="min-h-screen bg-gray-50">
    <div className="max-w-7xl mx-auto py-6 sm:px-6 lg:px-8">
      <div className="px-4 py-6 sm:px-0">
        <h1 className="text-3xl font-bold text-gray-900 mb-6">📦 Pedidos</h1>
        <div className="bg-white shadow rounded-lg p-6">
          <p className="text-gray-600">Página de pedidos - Aguardando integração GraphQL</p>
        </div>
      </div>
    </div>
  </div>
);

const PaymentsPage = () => (
  <div className="min-h-screen bg-gray-50">
    <div className="max-w-7xl mx-auto py-6 sm:px-6 lg:px-8">
      <div className="px-4 py-6 sm:px-0">
        <h1 className="text-3xl font-bold text-gray-900 mb-6">💳 Pagamentos</h1>
        <div className="bg-white shadow rounded-lg p-6">
          <p className="text-gray-600">Página de pagamentos - Aguardando integração GraphQL</p>
        </div>
      </div>
    </div>
  </div>
);

const NotificationsPage = () => (
  <div className="min-h-screen bg-gray-50">
    <div className="max-w-7xl mx-auto py-6 sm:px-6 lg:px-8">
      <div className="px-4 py-6 sm:px-0">
        <h1 className="text-3xl font-bold text-gray-900 mb-6">🔔 Notificações</h1>
        <div className="bg-white shadow rounded-lg p-6">
          <p className="text-gray-600">Página de notificações - Aguardando integração GraphQL</p>
        </div>
      </div>
    </div>
  </div>
);

// Layout principal
const Layout = ({ children }) => (
  <div className="min-h-screen bg-gray-50">
    <Header />
    <main>{children}</main>
  </div>
);

// Página 404
const NotFound = () => (
  <div className="min-h-screen bg-gray-50 flex items-center justify-center">
    <div className="text-center">
      <div className="text-6xl font-bold text-gray-300 mb-4">404</div>
      <h1 className="text-2xl font-bold text-gray-900 mb-2">Página não encontrada</h1>
      <p className="text-gray-600 mb-8">A página que você está procurando não existe.</p>
      <a href="/" className="btn-primary">
        Voltar ao Dashboard
      </a>
    </div>
  </div>
);

function App() {
  return (
    <Router>
      <div className="App">
        <Routes>
          <Route path="/" element={
            <Layout>
              <Dashboard />
            </Layout>
          } />
          <Route path="/orders" element={
            <Layout>
              <OrdersPage />
            </Layout>
          } />
          <Route path="/payments" element={
            <Layout>
              <PaymentsPage />
            </Layout>
          } />
          <Route path="/notifications" element={
            <Layout>
              <NotificationsPage />
            </Layout>
          } />
          <Route path="*" element={<NotFound />} />
        </Routes>
      </div>
    </Router>
  );
}

export default App;
