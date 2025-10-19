import React, { createContext, useContext, useState, useEffect } from 'react';

// Contexto de autenticação
const AuthContext = createContext();

// Hook para usar o contexto de autenticação
export const useAuth = () => {
  const context = useContext(AuthContext);
  if (!context) {
    throw new Error('useAuth deve ser usado dentro de um AuthProvider');
  }
  return context;
};

// Usuários mock para demonstração
const MOCK_USERS = [
  {
    id: 1,
    email: 'admin@microservices.com',
    password: 'admin123',
    name: 'Administrator',
    role: 'admin',
    avatar: '👨‍💼'
  },
  {
    id: 2,
    email: 'user@microservices.com',
    password: 'user123',
    name: 'User Demo',
    role: 'user',
    avatar: '👤'
  }
];

// Provider de autenticação
export const AuthProvider = ({ children }) => {
  const [user, setUser] = useState(null);
  const [loading, setLoading] = useState(true);

  // Verificar se há usuário logado no localStorage ao iniciar
  useEffect(() => {
    const savedUser = localStorage.getItem('user');
    if (savedUser) {
      try {
        setUser(JSON.parse(savedUser));
      } catch (error) {
        console.error('Erro ao carregar usuário:', error);
        localStorage.removeItem('user');
      }
    }
    setLoading(false);
  }, []);

  // Função de login
  const login = async (email, password) => {
    setLoading(true);
    
    // Simular delay de rede
    await new Promise(resolve => setTimeout(resolve, 1000));
    
    const foundUser = MOCK_USERS.find(
      u => u.email === email && u.password === password
    );
    
    if (foundUser) {
      const userWithoutPassword = { ...foundUser };
      delete userWithoutPassword.password;
      
      setUser(userWithoutPassword);
      localStorage.setItem('user', JSON.stringify(userWithoutPassword));
      setLoading(false);
      return { success: true };
    } else {
      setLoading(false);
      return { 
        success: false, 
        error: 'Credenciais inválidas' 
      };
    }
  };

  // Função de logout
  const logout = () => {
    setUser(null);
    localStorage.removeItem('user');
  };

  // Função para verificar se está autenticado
  const isAuthenticated = () => {
    return !!user;
  };

  // Função para verificar role
  const hasRole = (role) => {
    return user?.role === role;
  };

  const value = {
    user,
    login,
    logout,
    loading,
    isAuthenticated,
    hasRole,
    mockUsers: MOCK_USERS.map(u => ({
      email: u.email,
      password: u.password,
      name: u.name,
      role: u.role
    }))
  };

  return (
    <AuthContext.Provider value={value}>
      {children}
    </AuthContext.Provider>
  );
};