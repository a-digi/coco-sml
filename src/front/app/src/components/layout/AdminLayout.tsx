import React from "react";
import Sidebar from "./sidebar/Sidebar";
import Topbar from "./top/Topbar";
import Content from "./content/Content";

interface AdminLayoutProps {
  children: React.ReactNode;
}

const AdminLayout: React.FC<AdminLayoutProps> = ({ children }) => {
  return (
    <div className="flex h-screen bg-gray-100">
      <Sidebar />
      <div className="flex flex-col flex-1">
        <Topbar />
        <Content>{children}</Content>
      </div>
    </div>
  );
};

export default AdminLayout;
