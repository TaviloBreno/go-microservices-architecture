import React from 'react';
import { useQuery } from '@apollo/client';
import { GET_ORDERS } from '../apollo/queries';

const LoadingSpinner = () => (
  <div className="flex justify-center items-center p-8">
    <div className="loading-spinner"></div>
    <span className="ml-2 text-gray-600">Carregando pedidos...</span>
  </div>
);

const ErrorMessage = ({ error, onRetry }) => (
  <div className="p-6 text-center">
    <div className="text-red-600 mb-2">
      <svg className="w-8 h-8 mx-auto mb-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
        <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M12 8v4m0 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
      </svg>
      Erro ao carregar pedidos
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
      case 'completed':
      case 'aprovado':
      case 'paid':
        return 'status-completed';
      case 'pending':
      case 'processando':
      case 'processing':
        return 'status-processing';
      case 'failed':
      case 'cancelled':
      case 'cancelado':
        return 'status-failed';
      default:
        return 'status-pending';
    }
  };

  const getStatusText = (status) => {
    const statusLower = status?.toLowerCase();
    switch (statusLower) {
      case 'completed':
        return 'Concluído';
      case 'pending':
        return 'Pendente';
      case 'processing':
        return 'Processando';
      case 'failed':
        return 'Falhou';
      case 'cancelled':
        return 'Cancelado';
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

const OrdersTable = () => {
  const { data, loading, error, refetch } = useQuery(GET_ORDERS, {
    pollInterval: 5000, // Atualiza a cada 5 segundos
    errorPolicy: 'all',
  });

  if (loading) return <LoadingSpinner />;
  if (error) return <ErrorMessage error={error} onRetry={refetch} />;

  const orders = data?.orders || [];

  if (orders.length === 0) {
    return (
      <div className="text-center p-8">
        <div className="text-gray-400 mb-2">
          <svg className="w-12 h-12 mx-auto mb-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={1} d="M16 11V7a4 4 0 00-8 0v4M5 9h14l1 12H4L5 9z" />
          </svg>
        </div>
        <h3 className="text-lg font-medium text-gray-900 mb-1">Nenhum pedido encontrado</h3>
        <p className="text-gray-500">Os pedidos aparecerão aqui quando forem criados.</p>
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
                Produto
              </th>
              <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                Usuário ID
              </th>
              <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                Quantidade
              </th>
              <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                Preço
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
            {orders.map((order) => (
              <tr key={order.id} className="hover:bg-gray-50 transition-colors duration-150">
                <td className="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">
                  #{order.id}
                </td>
                <td className="px-6 py-4 whitespace-nowrap">
                  <div className="text-sm font-medium text-gray-900">{order.productName}</div>
                </td>
                <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                  #{order.userID}
                </td>
                <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                  {order.quantity}
                </td>
                <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-900 font-medium">
                  R$ {typeof order.price === 'number' ? order.price.toFixed(2) : order.price}
                </td>
                <td className="px-6 py-4 whitespace-nowrap">
                  <StatusBadge status={order.status} />
                </td>
                <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                  {order.createdAt ? new Date(order.createdAt).toLocaleDateString('pt-BR') : 'N/A'}
                </td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>
    </div>
  );
};

export default OrdersTable;