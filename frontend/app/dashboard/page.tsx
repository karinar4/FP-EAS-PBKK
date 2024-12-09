'use client';

import React, { useEffect, useState } from 'react';
import LogoutButton from '../components/LogoutButton';
import NavigationBar from '../components/NavigationBar'
import {NavbarContent, Dropdown, DropdownTrigger, DropdownMenu, DropdownItem, Avatar, Link} from "@nextui-org/react";
import Image from 'next/image';

export default function Dashboard() {
  const [user, setUser] = useState<{ data: { email: string; name: string } } | null>(null);

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

  const handleLogout = LogoutButton();

  return (
    <div>
      <NavigationBar
        customButtons={
          <>
            <NavbarContent as="div" justify="end">
              <Link href="#">
                <Image
                  src="/shopping-cart.png"
                  alt="Cart"
                  width={35}
                  height={35}
                  className="mx-3"
                />
              </Link>
              <Dropdown placement="bottom-end">
                <DropdownTrigger>
                  <Avatar
                    showFallback 
                    isBordered
                    as="button"
                    className="transition-transform"
                    color="secondary"
                    name={user?.data.name}
                    size="sm"
                    src="https://images.unsplash.com/broken"
                  />
                </DropdownTrigger>
                <DropdownMenu aria-label="Profile Actions" variant="flat">
                  <DropdownItem key="profile" className="h-14 gap-2">
                    <p className="font-semibold">Signed in as</p>
                    <p className="font-semibold">{ user?.data.email }</p>
                  </DropdownItem>
                  <DropdownItem key="transactions">Transactions</DropdownItem>
                  <DropdownItem key="logout" color="danger" onClick={handleLogout}>
                    Log Out
                  </DropdownItem>
                </DropdownMenu>
              </Dropdown>
            </NavbarContent>
          </>
        }
      />

    
      <h1>Welcome to the Dashboard</h1>
      </div>
  );
};