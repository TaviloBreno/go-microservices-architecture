import React from 'react';
import { useQuery } from '@apollo/client';
import { GET_NOTIFICATIONS } from '../apollo/queries';

const LoadingSpinner = () => (
  <div className="flex justify-center items-center p-8">
    <div className="loading-spinner"></div>
    <span className="ml-2 text-gray-600">Carregando notificações...</span>
  </div>
);

const ErrorMessage = ({ error, onRetry }) => (
  <div className="p-6 text-center">
    <div className="text-red-600 mb-2">
      <svg className="w-8 h-8 mx-auto mb-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
        <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M12 8v4m0 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
      </svg>
      Erro ao carregar notificações
    </div>
    <p className="text-gray-600 text-sm mb-4">{error.message}</p>
    <button 
      onClick={onRetry}
      className="btn-primary text-sm"
    >
      Tentar novamente
    </button>
  </div>
);

const StatusBadge = ({ status }) => {
  const getStatusClass = (status) => {
    const statusLower = status?.toLowerCase();
    switch (statusLower) {
      case 'sent':
      case 'delivered':
      case 'read':
      case 'enviado':
        return 'status-completed';
      case 'pending':
      case 'pendente':
        return 'status-pending';
      case 'processing':
      case 'sending':
      case 'enviando':
        return 'status-processing';
      case 'failed':
      case 'error':
      case 'falhou':
        return 'status-failed';
      default:
        return 'status-pending';
    }
  };

  const getStatusText = (status) => {
    const statusLower = status?.toLowerCase();
    switch (statusLower) {
      case 'sent':
      case 'enviado':
        return 'Enviado';
      case 'delivered':
        return 'Entregue';
      case 'read':
        return 'Lido';
      case 'pending':
      case 'pendente':
        return 'Pendente';
      case 'processing':
      case 'sending':
        return 'Enviando';
      case 'failed':
      case 'error':
        return 'Falhou';
      default:
        return status || 'Desconhecido';
    }
  };

  return (
    <span className={`status-badge ${getStatusClass(status)}`}>
      {getStatusText(status)}
    </span>
  );
};

const NotificationTypeBadge = ({ type }) => {
  const getTypeIcon = (type) => {
    const typeLower = type?.toLowerCase();
    switch (typeLower) {
      case 'email':
        return '📧';
      case 'sms':
        return '📱';
      case 'push':
        return '🔔';
      case 'webhook':
        return '🔗';
      case 'order':
      case 'pedido':
        return '📦';
      case 'payment':
      case 'pagamento':
        return '💳';
      default:
        return '📢';
    }
  };

  const getTypeText = (type) => {
    const typeLower = type?.toLowerCase();
    switch (typeLower) {
      case 'email':
        return 'E-mail';
      case 'sms':
        return 'SMS';
      case 'push':
        return 'Push';
      case 'webhook':
        return 'Webhook';
      case 'order':
        return 'Pedido';
      case 'payment':
        return 'Pagamento';
      default:
        return type || 'Geral';
    }
  };

  return (
    <span className="inline-flex items-center text-sm text-gray-700">
      <span className="mr-1">{getTypeIcon(type)}</span>
      {getTypeText(type)}
    </span>
  );
};

const NotificationsTable = () => {
  const { data, loading, error, refetch } = useQuery(GET_NOTIFICATIONS, {
    pollInterval: 5000, // Atualiza a cada 5 segundos
    errorPolicy: 'all',
  });

  if (loading) return <LoadingSpinner />;
  if (error) return <ErrorMessage error={error} onRetry={refetch} />;

  const notifications = data?.notifications || [];

  if (notifications.length === 0) {
    return (
      <div className="text-center p-8">
        <div className="text-gray-400 mb-2">
          <svg className="w-12 h-12 mx-auto mb-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={1} d="M15 17h5l-5 5v-5zM4 7h6v10H4V7z" />
          </svg>
        </div>
        <h3 className="text-lg font-medium text-gray-900 mb-1">Nenhuma notificação encontrada</h3>
        <p className="text-gray-500">As notificações aparecerão aqui quando forem enviadas.</p>
      </div>
    );
  }

  return (
    <div className="table-container">
      <div className="overflow-x-auto">
        <table className="min-w-full divide-y divide-gray-200">
          <thead className="bg-gray-50">
            <tr>
              <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                ID
              </th>
              <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                Pedido ID
              </th>
              <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                Usuário ID
              </th>
              <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                Mensagem
              </th>
              <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                Tipo
              </th>
              <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                Status
              </th>
              <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                Data
              </th>
            </tr>
          </thead>
          <tbody className="bg-white divide-y divide-gray-200">
            {notifications.map((notification) => (
              <tr key={notification.id} className="hover:bg-gray-50 transition-colors duration-150">
                <td className="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">
                  #{notification.id}
                </td>
                <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                  #{notification.orderID}
                </td>
                <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                  #{notification.userID}
                </td>
                <td className="px-6 py-4 max-w-xs">
                  <div className="text-sm text-gray-900 truncate" title={notification.message}>
                    {notification.message}
                  </div>
                </td>
                <td className="px-6 py-4 whitespace-nowrap">
                  <NotificationTypeBadge type={notification.type} />
                </td>
                <td className="px-6 py-4 whitespace-nowrap">
                  <StatusBadge status={notification.status} />
                </td>
                <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                  {notification.createdAt ? new Date(notification.createdAt).toLocaleDateString('pt-BR') : 'N/A'}
                </td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>
    </div>
  );
};

export default NotificationsTable;