import React from 'react';

type CardProps = {
  children: React.ReactNode;
};

const Card = ({ children }: CardProps) => (
  <div style={{ border: '1px solid #eee', borderRadius: 8, padding: 16, margin: 8, background: '#fff' }}>
    {children}
  </div>
);

export default Card;
