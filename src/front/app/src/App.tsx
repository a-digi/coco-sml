import React from 'react';
import AdminLayout from './components/layout/AdminLayout';
import './App.css';

function App() {
  return (
    <AdminLayout>
      {/* Place your dashboard content here */}
      <div className="text-2xl font-bold">Welcome to the Admin Dashboard</div>
    </AdminLayout>
  );
}

export default App;
