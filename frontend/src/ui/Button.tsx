import React from 'react';

type ButtonProps = React.ButtonHTMLAttributes<HTMLButtonElement> & {
  children: React.ReactNode;
};

const Button = ({ children, ...props }: ButtonProps) => (
  <button style={{ padding: 8, borderRadius: 4, background: '#2d72d9', color: '#fff', border: 'none' }} {...props}>
    {children}
  </button>
);

export default Button;
