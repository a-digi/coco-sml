import React from "react";

const Topbar: React.FC = () => {
  return (
    <header className="h-16 bg-white shadow flex items-center px-6 justify-between">
      <div className="font-semibold text-lg">Admin Dashboard</div>
      <div className="flex items-center space-x-4">
        <span className="text-gray-600">Welcome, Admin</span>
        <img
          src="https://ui-avatars.com/api/?name=Admin"
          alt="Admin Avatar"
          className="w-8 h-8 rounded-full border"
        />
      </div>
    </header>
  );
};

export default Topbar;

