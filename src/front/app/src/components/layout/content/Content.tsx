import React from "react";

interface ContentProps {
  children: React.ReactNode;
}

const Content: React.FC<ContentProps> = ({ children }) => {
  return (
    <main className="flex-1 p-6 overflow-y-auto">
      {children}
    </main>
  );
};

export default Content;

