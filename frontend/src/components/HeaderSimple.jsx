import React from 'react';
import { NavLink } from 'react-router-dom';

const HealthIndicator = () => {
  return (
    <div className="flex items-center">
      <div className="w-2 h-2 rounded-full mr-2 bg-warning-400 animate-pulse"></div>
      <span className="text-xs text-gray-500">
        Aguardando BFF GraphQL
      </span>
    </div>
  );
};

const Header = () => {
  const navLinkClass = ({ isActive }) =>
    `px-3 py-2 rounded-md text-sm font-medium transition-colors duration-200 ${
      isActive
        ? 'bg-primary-600 text-white'
        : 'text-gray-700 hover:text-primary-600 hover:bg-primary-50'
    }`;

  return (
    <header className="bg-white shadow-sm border-b border-gray-200">
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div className="flex justify-between items-center h-16">
          {/* Logo e título */}
          <div className="flex items-center">
            <div className="flex-shrink-0">
              <h1 className="text-xl font-bold text-gray-900">
                🏗️ <span className="text-primary-600">Microservices</span> Dashboard
              </h1>
            </div>
          </div>

          {/* Navegação principal */}
          <nav className="flex space-x-1">
            <NavLink to="/" className={navLinkClass}>
              📊 Dashboard
            </NavLink>
            <NavLink to="/orders" className={navLinkClass}>
              📦 Pedidos
            </NavLink>
            <NavLink to="/payments" className={navLinkClass}>
              💳 Pagamentos
            </NavLink>
            <NavLink to="/notifications" className={navLinkClass}>
              🔔 Notificações
            </NavLink>
          </nav>

          {/* Indicador de saúde */}
          <div className="flex items-center space-x-4">
            <HealthIndicator />
          </div>
        </div>
      </div>
    </header>
  );
};

export default Header;