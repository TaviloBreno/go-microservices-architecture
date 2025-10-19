import React from 'react';
import OrdersTable from '../components/OrdersTable';

const OrdersPage = () => {
  return (
    <div className="min-h-screen bg-gray-50">
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
        {/* Header */}
        <div className="mb-8">
          <div className="flex items-center justify-between">
            <div>
              <h1 className="text-3xl font-bold text-gray-900">Pedidos</h1>
              <p className="text-gray-600 mt-2">Gerencie todos os pedidos do sistema</p>
            </div>
            <div className="flex items-center space-x-3">
              <button
                onClick={() => window.location.reload()}
                className="btn-secondary"
              >
                <svg className="w-4 h-4 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15" />
                </svg>
                Atualizar
              </button>
            </div>
          </div>
        </div>

        {/* Informações */}
        <div className="card mb-6">
          <div className="flex items-center">
            <div className="w-10 h-10 bg-primary-100 rounded-lg flex items-center justify-center mr-4">
              <svg className="w-6 h-6 text-primary-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
              </svg>
            </div>
            <div>
              <h3 className="text-lg font-medium text-gray-900">Atualizações Automáticas</h3>
              <p className="text-gray-600 text-sm">Esta página é atualizada automaticamente a cada 5 segundos para mostrar os dados mais recentes.</p>
            </div>
          </div>
        </div>

        {/* Tabela de Pedidos */}
        <div className="card">
          <div className="mb-4">
            <h2 className="text-xl font-semibold text-gray-900">Lista de Pedidos</h2>
            <p className="text-gray-600 text-sm">Todos os pedidos criados no sistema aparecem abaixo.</p>
          </div>
          <OrdersTable />
        </div>
      </div>
    </div>
  );
};

export default OrdersPage;