import React from "react";

const Sidebar: React.FC = () => {
  return (
    <aside className="w-64 bg-white shadow-md h-full flex flex-col">
      <div className="h-16 flex items-center justify-center font-bold text-xl border-b">Admin</div>
      <nav className="flex-1 p-4">
        <ul className="space-y-2">
          <li><a href="#" className="block px-4 py-2 rounded hover:bg-gray-200">Dashboard</a></li>
          <li><a href="#" className="block px-4 py-2 rounded hover:bg-gray-200">Users</a></li>
          <li><a href="#" className="block px-4 py-2 rounded hover:bg-gray-200">Settings</a></li>
        </ul>
      </nav>
      <div className="p-4 border-t">
        <button className="w-full bg-red-500 text-white py-2 rounded hover:bg-red-600">Logout</button>
      </div>
    </aside>
  );
};

export default Sidebar;

