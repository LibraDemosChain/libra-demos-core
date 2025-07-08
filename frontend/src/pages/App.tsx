import React from 'react';
import { BrowserRouter, Routes, Route } from 'react-router-dom';
import Home from './Home';
import Auth from './Auth';
import Governance from './Governance';
import Documentation from './Documentation';

const App = () => (
  <BrowserRouter>
    <Routes>
      <Route path="/" element={<Home />} />
      <Route path="/auth" element={<Auth />} />
      <Route path="/governance" element={<Governance />} />
      <Route path="/docs" element={<Documentation />} />
    </Routes>
  </BrowserRouter>
);

export default App;
