import React from 'react';
import PaymentsTable from '../components/PaymentsTable';

const PaymentsPage = () => {
  return (
    <div className="min-h-screen bg-gray-50">
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
        {/* Header */}
        <div className="mb-8">
          <div className="flex items-center justify-between">
            <div>
              <h1 className="text-3xl font-bold text-gray-900">Pagamentos</h1>
              <p className="text-gray-600 mt-2">Monitore todos os pagamentos processados</p>
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

        {/* Estatísticas dos Pagamentos */}
        <div className="grid grid-cols-1 md:grid-cols-4 gap-6 mb-8">
          <div className="card">
            <div className="flex items-center">
              <div className="w-10 h-10 bg-success-100 rounded-lg flex items-center justify-center mr-3">
                <svg className="w-6 h-6 text-success-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z" />
                </svg>
              </div>
              <div>
                <p className="text-sm font-medium text-gray-600">Aprovados</p>
                <p className="text-xl font-bold text-gray-900">-</p>
              </div>
            </div>
          </div>

          <div className="card">
            <div className="flex items-center">
              <div className="w-10 h-10 bg-warning-100 rounded-lg flex items-center justify-center mr-3">
                <svg className="w-6 h-6 text-warning-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z" />
                </svg>
              </div>
              <div>
                <p className="text-sm font-medium text-gray-600">Pendentes</p>
                <p className="text-xl font-bold text-gray-900">-</p>
              </div>
            </div>
          </div>

          <div className="card">
            <div className="flex items-center">
              <div className="w-10 h-10 bg-danger-100 rounded-lg flex items-center justify-center mr-3">
                <svg className="w-6 h-6 text-danger-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M6 18L18 6M6 6l12 12" />
                </svg>
              </div>
              <div>
                <p className="text-sm font-medium text-gray-600">Rejeitados</p>
                <p className="text-xl font-bold text-gray-900">-</p>
              </div>
            </div>
          </div>

          <div className="card">
            <div className="flex items-center">
              <div className="w-10 h-10 bg-primary-100 rounded-lg flex items-center justify-center mr-3">
                <svg className="w-6 h-6 text-primary-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M12 8c-1.657 0-3 .895-3 2s1.343 2 3 2 3 .895 3 2-1.343 2-3 2m0-8c1.11 0 2.08.402 2.599 1M12 8V7m0 1v8m0 0v1m0-1c-1.11 0-2.08-.402-2.599-1" />
                </svg>
              </div>
              <div>
                <p className="text-sm font-medium text-gray-600">Total</p>
                <p className="text-xl font-bold text-gray-900">R$ -</p>
              </div>
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
              <h3 className="text-lg font-medium text-gray-900">Sistema de Pagamentos</h3>
              <p className="text-gray-600 text-sm">Os pagamentos são processados automaticamente quando um pedido é criado. Esta página mostra todos os pagamentos em tempo real.</p>
            </div>
          </div>
        </div>

        {/* Tabela de Pagamentos */}
        <div className="card">
          <div className="mb-4">
            <h2 className="text-xl font-semibold text-gray-900">Lista de Pagamentos</h2>
            <p className="text-gray-600 text-sm">Todos os pagamentos processados no sistema aparecem abaixo.</p>
          </div>
          <PaymentsTable />
        </div>
      </div>
    </div>
  );
};

export default PaymentsPage;