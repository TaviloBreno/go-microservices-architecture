import React from 'react';
import NotificationsTable from '../components/NotificationsTable';

const NotificationsPage = () => {
  return (
    <div className="min-h-screen bg-gray-50">
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
        {/* Header */}
        <div className="mb-8">
          <div className="flex items-center justify-between">
            <div>
              <h1 className="text-3xl font-bold text-gray-900">Notificações</h1>
              <p className="text-gray-600 mt-2">Acompanhe todas as notificações enviadas</p>
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

        {/* Estatísticas das Notificações */}
        <div className="grid grid-cols-1 md:grid-cols-4 gap-6 mb-8">
          <div className="card">
            <div className="flex items-center">
              <div className="w-10 h-10 bg-success-100 rounded-lg flex items-center justify-center mr-3">
                <svg className="w-6 h-6 text-success-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z" />
                </svg>
              </div>
              <div>
                <p className="text-sm font-medium text-gray-600">Enviadas</p>
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
                <p className="text-sm font-medium text-gray-600">Falharam</p>
                <p className="text-xl font-bold text-gray-900">-</p>
              </div>
            </div>
          </div>

          <div className="card">
            <div className="flex items-center">
              <div className="w-10 h-10 bg-primary-100 rounded-lg flex items-center justify-center mr-3">
                <svg className="w-6 h-6 text-primary-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M15 17h5l-5 5v-5zM4.021 18.425l7.14-7.14" />
                </svg>
              </div>
              <div>
                <p className="text-sm font-medium text-gray-600">Total</p>
                <p className="text-xl font-bold text-gray-900">-</p>
              </div>
            </div>
          </div>
        </div>

        {/* Tipos de Notificação */}
        <div className="card mb-6">
          <h3 className="text-lg font-semibold text-gray-900 mb-4">Tipos de Notificação Suportados</h3>
          <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-4">
            <div className="flex items-center p-3 bg-gray-50 rounded-lg">
              <span className="text-2xl mr-3">📧</span>
              <div>
                <p className="font-medium text-gray-900">E-mail</p>
                <p className="text-sm text-gray-600">Notificações por e-mail</p>
              </div>
            </div>
            <div className="flex items-center p-3 bg-gray-50 rounded-lg">
              <span className="text-2xl mr-3">📱</span>
              <div>
                <p className="font-medium text-gray-900">SMS</p>
                <p className="text-sm text-gray-600">Mensagens de texto</p>
              </div>
            </div>
            <div className="flex items-center p-3 bg-gray-50 rounded-lg">
              <span className="text-2xl mr-3">🔔</span>
              <div>
                <p className="font-medium text-gray-900">Push</p>
                <p className="text-sm text-gray-600">Notificações push</p>
              </div>
            </div>
            <div className="flex items-center p-3 bg-gray-50 rounded-lg">
              <span className="text-2xl mr-3">🔗</span>
              <div>
                <p className="font-medium text-gray-900">Webhook</p>
                <p className="text-sm text-gray-600">Integrações externas</p>
              </div>
            </div>
          </div>
        </div>

        {/* Informações do Fluxo */}
        <div className="card mb-6">
          <div className="flex items-start">
            <div className="w-10 h-10 bg-primary-100 rounded-lg flex items-center justify-center mr-4 mt-1">
              <svg className="w-6 h-6 text-primary-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
              </svg>
            </div>
            <div>
              <h3 className="text-lg font-medium text-gray-900 mb-2">Fluxo de Notificações</h3>
              <div className="text-gray-600 text-sm space-y-2">
                <p><strong>1. Pedido Criado:</strong> Notificação de confirmação é enviada ao cliente</p>
                <p><strong>2. Pagamento Processado:</strong> Cliente recebe notificação do status do pagamento</p>
                <p><strong>3. Pedido Atualizado:</strong> Notificações de mudança de status são enviadas</p>
                <p>As notificações são processadas automaticamente em tempo real e aparecem nesta página.</p>
              </div>
            </div>
          </div>
        </div>

        {/* Tabela de Notificações */}
        <div className="card">
          <div className="mb-4">
            <h2 className="text-xl font-semibold text-gray-900">Lista de Notificações</h2>
            <p className="text-gray-600 text-sm">Todas as notificações enviadas pelo sistema aparecem abaixo.</p>
          </div>
          <NotificationsTable />
        </div>
      </div>
    </div>
  );
};

export default NotificationsPage;