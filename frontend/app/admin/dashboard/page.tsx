"use client";

import React, { useEffect, useState } from 'react';
import LogoutButton from '@/app/components/LogoutButton';
import Sidebar from '@/app/components/Sidebar';
import { Navbar, NavbarContent, Dropdown, DropdownTrigger, DropdownMenu, DropdownItem, Avatar } from "@nextui-org/react";
import { useRouter } from 'next/navigation';

export default function Home() {
    const [user, setUser] = useState<{ data: { email: string; name: string } } | null>(null);
    const router = useRouter();
    
    useEffect(() => {
      const getTokenFromCookies = () => {
        const cookies = document.cookie.split('; ');
        const tokenCookie = cookies.find((cookie) => cookie.startsWith('auth-token='));
        return tokenCookie ? tokenCookie.split('=')[1] : null;
      };
  
      // Fetch user data from API
      const fetchUserData = async () => {
        try {
          const token = getTokenFromCookies();
  
          if (!token) {
            throw new Error('No authentication token found in cookies.');
          }
  
          console.log(token);
  
          const response = await fetch('http://localhost:3000/api/v1/auth/me', {
            method: 'GET',
            headers: {
              'Content-Type': 'application/json',
              'Authorization': `Bearer ${token}`,
            },
          });
          if (!response.ok) {
            throw new Error('Failed to fetch user data');
          }
          const data = await response.json();
          setUser(data);
        } catch (error) {
          console.error('Error fetching user data:', error);
        }
      };
  
      fetchUserData();
    }, []);
  
    const handleLogout = () => {
      document.cookie = 'auth-token=; Max-Age=0';
      setUser(null);
      router.push("/");
    };
  
  return (
    <>
    <Navbar maxWidth="full">
            <NavbarContent as="div" justify="end">
              <Dropdown placement="bottom-end">
                <DropdownTrigger>
                  <Avatar
                    showFallback
                    isBordered
                    as="button"
                    className="transition-transform"
                    color="warning"
                    name={user ? user.data.name : "Guest"}
                    size="sm"
                    src="https://images.unsplash.com/broken"
                  />
                </DropdownTrigger>
                <DropdownMenu aria-label="Profile Actions" variant="flat">
                  <DropdownItem key="profile" className="h-14 gap-2">
                    <p className="font-semibold">Signed in as</p>
                    <p className="font-semibold">{user?.data.email}</p>
                  </DropdownItem>
                  <DropdownItem key="logout" color="danger" onClick={handleLogout}>
                    Log Out
                  </DropdownItem>
                </DropdownMenu>
              </Dropdown>
            </NavbarContent>
          </Navbar>
    
          <div className="min-h-screen bg-gray-200 text-gray-800">
            <main className="container flex gap-8">
              <Sidebar className="w-1/4 bg-white border-r border-gray-300 p-6 shadow-sm h-screen"/>
              <div className="flex-1">
                <div className='flex-1 w-full'>
                  <div className="p-8 flex justify-between">
                    <h1 className="text-3xl font-bold text-black">Dashboard</h1>
                  </div>
                </div>
              </div>
            </main>
          </div>
          </>
  );
}
