import React from 'react';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import { ApolloProvider } from '@apollo/client';
import client from './apollo/client';

// Componentes
import Header from './components/Header';

// Páginas
import Dashboard from './pages/Dashboard';
import OrdersPage from './pages/OrdersPage';
import PaymentsPage from './pages/PaymentsPage';
import NotificationsPage from './pages/NotificationsPage';

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
    <ApolloProvider client={client}>
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
    </ApolloProvider>
  );
}

export default App;