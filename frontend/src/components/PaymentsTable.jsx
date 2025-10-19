import React from 'react';
import { useQuery } from '@apollo/client';
import { GET_PAYMENTS } from '../apollo/queries';

const LoadingSpinner = () => (
  <div className="flex justify-center items-center p-8">
    <div className="loading-spinner"></div>
    <span className="ml-2 text-gray-600">Carregando pagamentos...</span>
  </div>
);

const ErrorMessage = ({ error, onRetry }) => (
  <div className="p-6 text-center">
    <div className="text-red-600 mb-2">
      <svg className="w-8 h-8 mx-auto mb-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
        <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M12 8v4m0 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
      </svg>
      Erro ao carregar pagamentos
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
      case 'approved':
      case 'paid':
      case 'success':
      case 'completed':
        return 'status-completed';
      case 'pending':
      case 'processing':
        return 'status-processing';
      case 'failed':
      case 'rejected':
      case 'cancelled':
        return 'status-failed';
      default:
        return 'status-pending';
    }
  };

  const getStatusText = (status) => {
    const statusLower = status?.toLowerCase();
    switch (statusLower) {
      case 'approved':
      case 'paid':
      case 'success':
        return 'Aprovado';
      case 'completed':
        return 'Concluído';
      case 'pending':
        return 'Pendente';
      case 'processing':
        return 'Processando';
      case 'failed':
        return 'Falhou';
      case 'rejected':
        return 'Rejeitado';
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

const PaymentMethodBadge = ({ method }) => {
  const getMethodIcon = (method) => {
    const methodLower = method?.toLowerCase();
    switch (methodLower) {
      case 'card':
      case 'credit_card':
      case 'cartao':
        return '💳';
      case 'pix':
        return '🇧🇷';
      case 'bank_transfer':
      case 'transferencia':
        return '🏦';
      case 'boleto':
        return '📄';
      default:
        return '💰';
    }
  };

  const getMethodText = (method) => {
    const methodLower = method?.toLowerCase();
    switch (methodLower) {
      case 'card':
      case 'credit_card':
        return 'Cartão';
      case 'pix':
        return 'PIX';
      case 'bank_transfer':
        return 'Transferência';
      case 'boleto':
        return 'Boleto';
      default:
        return method || 'N/A';
    }
  };

  return (
    <span className="inline-flex items-center text-sm text-gray-700">
      <span className="mr-1">{getMethodIcon(method)}</span>
      {getMethodText(method)}
    </span>
  );
};

const PaymentsTable = () => {
  const { data, loading, error, refetch } = useQuery(GET_PAYMENTS, {
    pollInterval: 5000, // Atualiza a cada 5 segundos
    errorPolicy: 'all',
  });

  if (loading) return <LoadingSpinner />;
  if (error) return <ErrorMessage error={error} onRetry={refetch} />;

  const payments = data?.payments || [];

  if (payments.length === 0) {
    return (
      <div className="text-center p-8">
        <div className="text-gray-400 mb-2">
          <svg className="w-12 h-12 mx-auto mb-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={1} d="M3 10h18M7 15h1m4 0h1m-7 4h12a3 3 0 003-3V8a3 3 0 00-3-3H6a3 3 0 00-3 3v8a3 3 0 003 3z" />
          </svg>
        </div>
        <h3 className="text-lg font-medium text-gray-900 mb-1">Nenhum pagamento encontrado</h3>
        <p className="text-gray-500">Os pagamentos aparecerão aqui quando forem processados.</p>
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
                Valor
              </th>
              <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                Método
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
            {payments.map((payment) => (
              <tr key={payment.id} className="hover:bg-gray-50 transition-colors duration-150">
                <td className="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">
                  #{payment.id}
                </td>
                <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                  #{payment.orderID}
                </td>
                <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                  #{payment.userID}
                </td>
                <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-900 font-medium">
                  R$ {typeof payment.amount === 'number' ? payment.amount.toFixed(2) : payment.amount}
                </td>
                <td className="px-6 py-4 whitespace-nowrap">
                  <PaymentMethodBadge method={payment.paymentMethod} />
                </td>
                <td className="px-6 py-4 whitespace-nowrap">
                  <StatusBadge status={payment.status} />
                </td>
                <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                  {payment.createdAt ? new Date(payment.createdAt).toLocaleDateString('pt-BR') : 'N/A'}
                </td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>
    </div>
  );
};

export default PaymentsTable;