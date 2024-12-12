import React from 'react';
import { Button, Link } from "@nextui-org/react";
import NavigationBar from './components/NavigationBar';

export default function App() {
  return (
    <>
      <NavigationBar />
      <div className="min-h-screen bg-gray-200 text-gray-800">
        {/* Hero Section */}
        <header className="relative h-[80vh] bg-cover bg-center text-white" style={{ backgroundImage: 'url("https://www.zenadrone.com/wp-content/uploads/2022/10/oil-and-gas-industry-2-1024x536.jpg")' }}>
          <div className="absolute inset-0 bg-black bg-opacity-50"></div>
          <div className="container mx-auto px-6 h-full flex flex-col justify-center items-center relative z-10">
            <h1 className="text-5xl font-bold mb-4 text-center">Welcome to RoboRent</h1>
            <p className="text-lg mb-6 text-center">
              the future of automation at your fingertips! We specialize in providing state-of-the-art robots for rent, tailored to meet your personal and business needs. Whether you're looking for advanced robotic assistants, cutting-edge industrial machines, or innovative solutions for events, we've got you covered. Our flexible rental plans, expert support, and commitment to quality ensure a seamless experience that empowers you to achieve more. Explore the possibilities today and let us bring the future to you!
            </p>
            <Button
              size="lg"
              style={{ backgroundColor: "#FFC107", color: "#000" }}
              as={Link}
              href="/catalog"
              className="font-semibold"
            >
              Explore Catalog
            </Button>
          </div>
        </header>



        <footer className="bg-black text-white py-4">
          <div className="container mx-auto text-center">
            <p className="text-sm">&copy; {new Date().getFullYear()} RoboRent. All rights reserved.</p>
          </div>
        </footer>
      </div>
    </>
  );
}
