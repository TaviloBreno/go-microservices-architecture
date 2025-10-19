import React, { useState, useEffect } from 'react';

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
    <p className="text-gray-600 text-sm mb-4">{error}</p>
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
      case 'delivered':
      case 'entregue':
        return 'status-completed';
      case 'pending':
      case 'processando':
      case 'processing':
      case 'confirmed':
      case 'confirmado':
        return 'status-processing';
      case 'failed':
      case 'cancelled':
      case 'cancelado':
        return 'status-failed';
      case 'shipped':
      case 'enviado':
        return 'status-pending';
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
      case 'confirmed':
        return 'Confirmado';
      case 'shipped':
        return 'Enviado';
      case 'delivered':
        return 'Entregue';
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
  const [orders, setOrders] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  const fetchOrders = async () => {
    try {
      setLoading(true);
      const response = await fetch('http://localhost:8080/api/orders');
      
      if (!response.ok) {
        throw new Error('Erro ao buscar pedidos');
      }

      const data = await response.json();
      setOrders(data || []);
      setLoading(false);
      setError(null);
    } catch (err) {
      setError(err.message);
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchOrders();
    
    // Atualiza a cada 5 segundos
    const interval = setInterval(fetchOrders, 5000);
    
    return () => clearInterval(interval);
  }, []);

  if (loading && orders.length === 0) return <LoadingSpinner />;
  if (error) return <ErrorMessage error={error} onRetry={fetchOrders} />;

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
                Usuário
              </th>
              <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                Total
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
                  <div className="text-sm font-medium text-gray-900">{order.user_name}</div>
                  <div className="text-xs text-gray-500">{order.user_email}</div>
                </td>
                <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-900 font-medium">
                  R$ {typeof order.total === 'number' ? order.total.toFixed(2) : order.total}
                </td>
                <td className="px-6 py-4 whitespace-nowrap">
                  <StatusBadge status={order.status} />
                </td>
                <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                  {new Date(order.created_at).toLocaleDateString('pt-BR')}
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
