"use client";

import React, { useState } from 'react';

export default function Sidebar() {
  return (
    <div className="flex">
      {/* Sidebar */}
      <div
        // Conditional class based on isOpen 
        // state to control width and visibility
        className={`text-black 
                    fixed h-screen transition-all 
                    duration-300 z-10 bg-white border-r-1 border-gray-100
                    w-64`}>
        {/* Sidebar content */}
        <div className="flex flex-col items-center py-6 px-4">
          <div className="mb-6">
            <h2 className="text-xl font-semibold text-black">Admin Dashboard</h2>
          </div>
          <div className="mt-2 w-full text-center">
            <a href="/admin/product"
              className="block py-1 px-4 text-lg text-black hover:bg-gray-100 rounded-md transition-colors duration-200">
              Product
            </a>
          </div>
          <div className="mt-2 w-full text-center">
            <a href="#"
              className="block py-1 px-4 text-lg text-black hover:bg-gray-100 rounded-md transition-colors duration-200">
              Cart
            </a>
          </div>
          <div className="mt-2 w-full text-center">
            <a href="#"
              className="block py-1 px-4 text-lg text-black hover:bg-gray-100 rounded-md transition-colors duration-200">
              Brand
            </a>
          </div>
          <div className="mt-2 w-full text-center">
            <a href="#"
              className="block py-1 px-4 text-lg text-black hover:bg-gray-100 rounded-md transition-colors duration-200">
              Transaction
            </a>
          </div>
          <div className="mt-2 w-full text-center">
            <a href="#"
              className="block py-1 px-4 text-lg text-black hover:bg-gray-100 rounded-md transition-colors duration-200">
              Payment
            </a>
          </div>
          {/* Add more sidebar items here */}
        </div>
      </div>

      {/* Main content */}
      <div className={`flex-1 transition-all duration-300 ml-64`}>
        {/* Button to toggle sidebar */}
        
      </div>
    </div>
  );
};