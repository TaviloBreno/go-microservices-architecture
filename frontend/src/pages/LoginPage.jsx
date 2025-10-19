import React, { useState } from 'react';
import { useAuth } from '../contexts/AuthContext';
import { Button } from '../components/ui/button';
import { cn } from '../lib/utils';

const LoginPage = () => {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [error, setError] = useState('');
  const [showCredentials, setShowCredentials] = useState(false);
  const { login, loading, mockUsers } = useAuth();

  const handleSubmit = async (e) => {
    e.preventDefault();
    setError('');
    
    const result = await login(email, password);
    if (!result.success) {
      setError(result.error);
    }
  };

  const handleDemoLogin = (userCredentials) => {
    setEmail(userCredentials.email);
    setPassword(userCredentials.password);
    setError('');
  };

  return (
    <div className="min-h-screen bg-gradient-to-br from-primary/10 via-background to-secondary/10 flex items-center justify-center p-4">
      <div className="w-full max-w-md">
        {/* Header */}
        <div className="text-center mb-8">
          <div className="mb-4">
            <h1 className="text-3xl font-bold text-foreground">
              🏗️ <span className="text-primary">Microservices</span>
            </h1>
            <p className="text-muted-foreground mt-2">
              Dashboard de Gerenciamento
            </p>
          </div>
        </div>

        {/* Card de Login */}
        <div className="card">
          <div className="mb-6">
            <h2 className="text-2xl font-semibold text-center text-foreground">
              Entrar
            </h2>
            <p className="text-muted-foreground text-center mt-2">
              Acesse sua conta para continuar
            </p>
          </div>

          {/* Formulário */}
          <form onSubmit={handleSubmit} className="space-y-4">
            {error && (
              <div className="p-3 rounded-md bg-destructive/10 border border-destructive/20">
                <p className="text-sm text-destructive">{error}</p>
              </div>
            )}

            <div>
              <label htmlFor="email" className="block text-sm font-medium text-foreground mb-2">
                Email
              </label>
              <input
                id="email"
                type="email"
                value={email}
                onChange={(e) => setEmail(e.target.value)}
                className="w-full px-3 py-2 border border-input rounded-md bg-background text-foreground placeholder:text-muted-foreground focus:outline-none focus:ring-2 focus:ring-ring focus:ring-offset-2"
                placeholder="seu@email.com"
                required
              />
            </div>

            <div>
              <label htmlFor="password" className="block text-sm font-medium text-foreground mb-2">
                Senha
              </label>
              <input
                id="password"
                type="password"
                value={password}
                onChange={(e) => setPassword(e.target.value)}
                className="w-full px-3 py-2 border border-input rounded-md bg-background text-foreground placeholder:text-muted-foreground focus:outline-none focus:ring-2 focus:ring-ring focus:ring-offset-2"
                placeholder="Sua senha"
                required
              />
            </div>

            <Button
              type="submit"
              className="w-full"
              disabled={loading}
            >
              {loading ? (
                <div className="flex items-center">
                  <div className="loading-spinner w-4 h-4 mr-2"></div>
                  Entrando...
                </div>
              ) : (
                'Entrar'
              )}
            </Button>
          </form>

          {/* Divisor */}
          <div className="relative my-6">
            <div className="absolute inset-0 flex items-center">
              <div className="w-full border-t border-border"></div>
            </div>
            <div className="relative flex justify-center text-xs uppercase">
              <span className="bg-background px-2 text-muted-foreground">
                Ou use credenciais demo
              </span>
            </div>
          </div>

          {/* Botão para mostrar credenciais */}
          <Button
            type="button"
            variant="outline"
            className="w-full mb-4"
            onClick={() => setShowCredentials(!showCredentials)}
          >
            {showCredentials ? 'Ocultar' : 'Mostrar'} Credenciais Demo
          </Button>

          {/* Credenciais Demo */}
          {showCredentials && (
            <div className="space-y-2">
              {mockUsers.map((user, index) => (
                <div
                  key={index}
                  className="p-3 border border-border rounded-md bg-muted/50 hover:bg-muted cursor-pointer transition-colors"
                  onClick={() => handleDemoLogin(user)}
                >
                  <div className="flex items-center justify-between">
                    <div>
                      <p className="text-sm font-medium text-foreground">
                        {user.name}
                      </p>
                      <p className="text-xs text-muted-foreground">
                        {user.email}
                      </p>
                    </div>
                    <div className="text-xs bg-primary/10 text-primary px-2 py-1 rounded">
                      {user.role}
                    </div>
                  </div>
                  <p className="text-xs text-muted-foreground mt-1">
                    Senha: {user.password}
                  </p>
                </div>
              ))}
            </div>
          )}
        </div>

        {/* Footer */}
        <div className="text-center mt-8">
          <p className="text-xs text-muted-foreground">
            © 2025 Microservices Architecture Demo
          </p>
        </div>
      </div>
    </div>
  );
};

export default LoginPage;