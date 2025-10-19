import React from 'react';
import { Moon, Sun } from 'lucide-react';
import { useTheme } from '../contexts/ThemeContext';
import { Button } from './ui/button';
import { cn } from '../lib/utils';

const ThemeToggle = ({ className, showText = false }) => {
  const { toggleTheme, isDark } = useTheme();

  return (
    <Button
      variant="ghost"
      size={showText ? "default" : "icon"}
      onClick={toggleTheme}
      className={cn(showText ? "justify-start" : "h-9 w-9", className)}
      title={`Alternar para modo ${isDark ? 'claro' : 'escuro'}`}
    >
      {isDark ? (
        <Sun className="h-4 w-4 transition-all" />
      ) : (
        <Moon className="h-4 w-4 transition-all" />
      )}
      {showText && (
        <span className="ml-2">
          {isDark ? "Modo Claro" : "Modo Escuro"}
        </span>
      )}
      <span className="sr-only">Alternar tema</span>
    </Button>
  );
};

export default ThemeToggle;