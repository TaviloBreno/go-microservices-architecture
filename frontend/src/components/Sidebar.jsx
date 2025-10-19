import React, { useState } from 'react';
import { NavLink, useLocation } from 'react-router-dom';
import { 
  LayoutDashboard, 
  Package, 
  CreditCard, 
  Bell, 
  Menu, 
  X,
  LogOut,
  User,
  ChevronLeft,
  ChevronRight,
  ShoppingCart,
  Users
} from 'lucide-react';
import { useAuth } from '../contexts/AuthContext';
import { Button } from '../components/ui/button';
import { cn } from '../lib/utils';
import ThemeToggle from './ThemeToggle';

const Sidebar = ({ isCollapsed, setIsCollapsed }) => {
  const { user, logout } = useAuth();
  const location = useLocation();
  const [isMobileOpen, setIsMobileOpen] = useState(false);

  const navigationItems = [
    {
      title: 'Dashboard',
      href: '/',
      icon: LayoutDashboard,
      description: 'Visão geral do sistema'
    },
    {
      title: 'Produtos',
      href: '/products',
      icon: ShoppingCart,
      description: 'Catálogo de produtos'
    },
    {
      title: 'Usuários',
      href: '/users',
      icon: Users,
      description: 'Gerenciar usuários'
    },
    {
      title: 'Pedidos',
      href: '/orders',
      icon: Package,
      description: 'Gerenciar pedidos'
    },
    {
      title: 'Pagamentos',
      href: '/payments',
      icon: CreditCard,
      description: 'Transações financeiras'
    },
    {
      title: 'Notificações',
      href: '/notifications',
      icon: Bell,
      description: 'Central de notificações'
    }
  ];

  const handleLogout = () => {
    logout();
    setIsMobileOpen(false);
  };

  const SidebarContent = () => (
    <>
      {/* Header */}
      <div className={cn(
        "flex items-center px-4 py-6 border-b border-border",
        isCollapsed ? "justify-center" : "justify-between"
      )}>
        {!isCollapsed && (
          <div className="flex items-center space-x-2">
            <div className="text-2xl">🏗️</div>
            <div>
              <h1 className="text-lg font-bold text-foreground">
                Microservices
              </h1>
              <p className="text-xs text-muted-foreground">Dashboard</p>
            </div>
          </div>
        )}
        
        {isCollapsed && (
          <div className="text-2xl">🏗️</div>
        )}

        {/* Botão de expandir/recolher - apenas desktop */}
        <Button
          variant="ghost"
          size="icon"
          onClick={() => setIsCollapsed(!isCollapsed)}
          className="hidden lg:flex h-8 w-8"
        >
          {isCollapsed ? (
            <ChevronRight className="h-4 w-4" />
          ) : (
            <ChevronLeft className="h-4 w-4" />
          )}
        </Button>

        {/* Botão de fechar - apenas mobile */}
        <Button
          variant="ghost"
          size="icon"
          onClick={() => setIsMobileOpen(false)}
          className="lg:hidden h-8 w-8"
        >
          <X className="h-4 w-4" />
        </Button>
      </div>

      {/* Navegação */}
      <nav className="flex-1 px-4 py-6 space-y-2">
        {navigationItems.map((item) => {
          const isActive = location.pathname === item.href;
          const Icon = item.icon;
          
          return (
            <NavLink
              key={item.href}
              to={item.href}
              onClick={() => setIsMobileOpen(false)}
              className={cn(
                "flex items-center space-x-3 px-3 py-2 rounded-md text-sm font-medium transition-all duration-200",
                isActive
                  ? "bg-primary text-primary-foreground"
                  : "text-muted-foreground hover:text-foreground hover:bg-accent"
              )}
              title={isCollapsed ? item.title : undefined}
            >
              <Icon className="h-5 w-5 flex-shrink-0" />
              {!isCollapsed && (
                <div className="flex-1 min-w-0">
                  <div>{item.title}</div>
                  {!isActive && (
                    <div className="text-xs opacity-70">{item.description}</div>
                  )}
                </div>
              )}
            </NavLink>
          );
        })}
      </nav>

      {/* Footer com informações do usuário */}
      <div className="p-4 border-t border-border">
        {!isCollapsed && (
          <div className="mb-3">
            <div className="flex items-center space-x-3 p-2 rounded-md bg-muted/50">
              <div className="text-2xl">{user?.avatar || '👤'}</div>
              <div className="flex-1 min-w-0">
                <p className="text-sm font-medium text-foreground truncate">
                  {user?.name}
                </p>
                <p className="text-xs text-muted-foreground truncate">
                  {user?.email}
                </p>
                <div className="text-xs bg-primary/10 text-primary px-2 py-0.5 rounded mt-1 inline-block">
                  {user?.role}
                </div>
              </div>
            </div>
          </div>
        )}

        {/* Theme Toggle */}
        <div className={cn(
          "mb-3",
          isCollapsed ? "flex justify-center" : ""
        )}>
          <ThemeToggle 
            showText={!isCollapsed}
            className={cn(
              isCollapsed ? "w-10 h-10" : "w-full"
            )} 
          />
        </div>

        <Button
          variant="outline"
          onClick={handleLogout}
          className={cn(
            "w-full",
            isCollapsed ? "px-2" : "justify-start"
          )}
        >
          <LogOut className="h-4 w-4" />
          {!isCollapsed && <span className="ml-2">Sair</span>}
        </Button>
      </div>
    </>
  );

  return (
    <>
      {/* Mobile Menu Button */}
      <Button
        variant="ghost"
        size="icon"
        onClick={() => setIsMobileOpen(true)}
        className="lg:hidden fixed top-4 left-4 z-50 h-10 w-10"
      >
        <Menu className="h-5 w-5" />
      </Button>

      {/* Mobile Overlay */}
      {isMobileOpen && (
        <div 
          className="lg:hidden fixed inset-0 bg-background/80 backdrop-blur-sm z-40"
          onClick={() => setIsMobileOpen(false)}
        />
      )}

      {/* Desktop Sidebar */}
      <aside className={cn(
        "hidden lg:flex flex-col bg-card border-r border-border transition-all duration-300",
        isCollapsed ? "w-16" : "w-64"
      )}>
        <SidebarContent />
      </aside>

      {/* Mobile Sidebar */}
      <aside className={cn(
        "lg:hidden fixed left-0 top-0 z-50 h-full w-64 bg-card border-r border-border transform transition-transform duration-300",
        isMobileOpen ? "translate-x-0" : "-translate-x-full"
      )}>
        <div className="flex flex-col h-full">
          <SidebarContent />
        </div>
      </aside>
    </>
  );
};

export default Sidebar;