import React from 'react';
import { useQuery } from '@apollo/client';
import { GET_DASHBOARD_DATA } from '../apollo/queries';

const StatCard = ({ title, value, icon, color = 'primary', trend }) => (
  <div className="card">
    <div className="flex items-center justify-between">
      <div>
        <p className="text-sm font-medium text-gray-600">{title}</p>
        <p className="text-2xl font-bold text-gray-900">{value}</p>
        {trend && (
          <p className={`text-sm ${trend > 0 ? 'text-success-600' : 'text-danger-600'}`}>
            {trend > 0 ? '↗' : '↘'} {Math.abs(trend)}%
          </p>
        )}
      </div>
      <div className={`w-12 h-12 bg-${color}-100 rounded-lg flex items-center justify-center`}>
        <div className={`text-${color}-600`}>
          {icon}
        </div>
      </div>
    </div>
  </div>
);

const RecentActivity = ({ activities }) => (
  <div className="card">
    <h3 className="text-lg font-semibold text-gray-900 mb-4">Atividade Recente</h3>
    <div className="space-y-3">
      {activities.length === 0 ? (
        <p className="text-gray-500 text-sm">Nenhuma atividade recente</p>
      ) : (
        activities.slice(0, 5).map((activity, index) => (
          <div key={index} className="flex items-center space-x-3 p-3 bg-gray-50 rounded-lg">
            <div className="flex-shrink-0">
              <div className={`w-8 h-8 bg-${activity.color}-100 rounded-full flex items-center justify-center`}>
                <span className="text-xs">{activity.icon}</span>
              </div>
            </div>
            <div className="flex-1 min-w-0">
              <p className="text-sm font-medium text-gray-900 truncate">
                {activity.title}
              </p>
              <p className="text-xs text-gray-500">{activity.description}</p>
            </div>
            <div className="text-xs text-gray-400">
              {activity.time}
            </div>
          </div>
        ))
      )}
    </div>
  </div>
);

const QuickStats = ({ orders, payments, notifications }) => {
  const completedOrders = orders?.filter(o => o.status?.toLowerCase() === 'completed').length || 0;
  const pendingOrders = orders?.filter(o => o.status?.toLowerCase() === 'pending').length || 0;
  
  const successfulPayments = payments?.filter(p => 
    ['approved', 'paid', 'success', 'completed'].includes(p.status?.toLowerCase())
  ).length || 0;
  
  const sentNotifications = notifications?.filter(n => 
    ['sent', 'delivered', 'enviado'].includes(n.status?.toLowerCase())
  ).length || 0;

  const totalRevenue = payments?.reduce((acc, payment) => {
    if (['approved', 'paid', 'success', 'completed'].includes(payment.status?.toLowerCase())) {
      return acc + (typeof payment.amount === 'number' ? payment.amount : parseFloat(payment.amount) || 0);
    }
    return acc;
  }, 0) || 0;

  const stats = [
    {
      title: 'Total de Pedidos',
      value: orders?.length || 0,
      icon: (
        <svg className="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M16 11V7a4 4 0 00-8 0v4M5 9h14l1 12H4L5 9z" />
        </svg>
      ),
      color: 'primary',
    },
    {
      title: 'Pedidos Concluídos',
      value: completedOrders,
      icon: (
        <svg className="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z" />
        </svg>
      ),
      color: 'success',
    },
    {
      title: 'Pagamentos Aprovados',
      value: successfulPayments,
      icon: (
        <svg className="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M3 10h18M7 15h1m4 0h1m-7 4h12a3 3 0 003-3V8a3 3 0 00-3-3H6a3 3 0 00-3 3v8a3 3 0 003 3z" />
        </svg>
      ),
      color: 'success',
    },
    {
      title: 'Receita Total',
      value: `R$ ${totalRevenue.toFixed(2)}`,
      icon: (
        <svg className="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M12 8c-1.657 0-3 .895-3 2s1.343 2 3 2 3 .895 3 2-1.343 2-3 2m0-8c1.11 0 2.08.402 2.599 1M12 8V7m0 1v8m0 0v1m0-1c-1.11 0-2.08-.402-2.599-1" />
        </svg>
      ),
      color: 'warning',
    },
    {
      title: 'Notificações Enviadas',
      value: sentNotifications,
      icon: (
        <svg className="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M15 17h5l-5 5v-5zM4.021 18.425l7.14-7.14" />
        </svg>
      ),
      color: 'primary',
    },
    {
      title: 'Pedidos Pendentes',
      value: pendingOrders,
      icon: (
        <svg className="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z" />
        </svg>
      ),
      color: 'warning',
    },
  ];

  return (
    <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6 mb-8">
      {stats.map((stat, index) => (
        <StatCard key={index} {...stat} />
      ))}
    </div>
  );
};

const Dashboard = () => {
  const { data, loading, error, refetch } = useQuery(GET_DASHBOARD_DATA, {
    pollInterval: 10000, // Atualiza a cada 10 segundos
    errorPolicy: 'all',
  });

  if (loading) {
    return (
      <div className="min-h-screen bg-gray-50">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
          <div className="flex justify-center items-center h-64">
            <div className="loading-spinner"></div>
            <span className="ml-2 text-gray-600">Carregando dashboard...</span>
          </div>
        </div>
      </div>
    );
  }

  if (error) {
    return (
      <div className="min-h-screen bg-gray-50">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
          <div className="card text-center">
            <div className="text-red-600 mb-4">
              <svg className="w-12 h-12 mx-auto mb-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M12 8v4m0 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
              </svg>
              <h3 className="text-lg font-medium text-gray-900 mb-2">Erro ao carregar dashboard</h3>
              <p className="text-gray-600 text-sm mb-4">{error.message}</p>
            </div>
            <button onClick={refetch} className="btn-primary">
              Tentar novamente
            </button>
          </div>
        </div>
      </div>
    );
  }

  const { orders = [], payments = [], notifications = [] } = data || {};

  // Gerar atividades recentes com base nos dados
  const recentActivities = [
    ...orders.slice(-3).map(order => ({
      title: `Novo pedido #${order.id}`,
      description: `${order.productName} - R$ ${order.price}`,
      time: 'há poucos minutos',
      icon: '📦',
      color: 'primary'
    })),
    ...payments.slice(-2).map(payment => ({
      title: `Pagamento processado #${payment.id}`,
      description: `R$ ${payment.amount} - ${payment.status}`,
      time: 'há poucos minutos',
      icon: '💳',
      color: 'success'
    })),
    ...notifications.slice(-2).map(notification => ({
      title: `Notificação enviada #${notification.id}`,
      description: notification.message?.substring(0, 50) + '...',
      time: 'há poucos minutos',
      icon: '🔔',
      color: 'warning'
    }))
  ].slice(0, 5);

  return (
    <div className="min-h-screen bg-gray-50">
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
        {/* Header */}
        <div className="mb-8">
          <h1 className="text-3xl font-bold text-gray-900">Dashboard</h1>
          <p className="text-gray-600 mt-2">Visão geral do sistema de microserviços</p>
        </div>

        {/* Estatísticas Rápidas */}
        <QuickStats orders={orders} payments={payments} notifications={notifications} />

        {/* Conteúdo Principal */}
        <div className="grid grid-cols-1 lg:grid-cols-3 gap-8">
          {/* Atividade Recente */}
          <div className="lg:col-span-2">
            <RecentActivity activities={recentActivities} />
          </div>

          {/* Status do Sistema */}
          <div className="card">
            <h3 className="text-lg font-semibold text-gray-900 mb-4">Status do Sistema</h3>
            <div className="space-y-4">
              <div className="flex items-center justify-between">
                <span className="text-sm text-gray-600">BFF GraphQL</span>
                <span className="status-badge status-completed">Online</span>
              </div>
              <div className="flex items-center justify-between">
                <span className="text-sm text-gray-600">Order Service</span>
                <span className="status-badge status-processing">Conectando</span>
              </div>
              <div className="flex items-center justify-between">
                <span className="text-sm text-gray-600">Payment Service</span>
                <span className="status-badge status-completed">Online</span>
              </div>
              <div className="flex items-center justify-between">
                <span className="text-sm text-gray-600">User Service</span>
                <span className="status-badge status-completed">Online</span>
              </div>
              <div className="flex items-center justify-between">
                <span className="text-sm text-gray-600">Notification Service</span>
                <span className="status-badge status-completed">Online</span>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};

export default Dashboard;