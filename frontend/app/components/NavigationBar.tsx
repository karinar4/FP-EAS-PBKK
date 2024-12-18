"use client";

import { Navbar, NavbarBrand, NavbarContent, NavbarItem, Link, Dropdown, DropdownTrigger, DropdownMenu, DropdownItem, Avatar } from "@nextui-org/react";
import Image from 'next/image';
import { useEffect, useState } from "react";
import { useRouter } from 'next/navigation';

export default function NavigationBar() {
  const [user, setUser] = useState<{ data: { email: string; name: string } } | null>(null);
  const router = useRouter();

  useEffect(() => {
    const getTokenFromCookies = () => {
      const cookies = document.cookie.split('; ');
      const tokenCookie = cookies.find((cookie) => cookie.startsWith('auth-token='));
      return tokenCookie ? tokenCookie.split('=')[1] : null;
    };

    const fetchUserData = async () => {
      try {
        const token = getTokenFromCookies();

        if (!token) {
          console.log('No authentication token found in cookies.');
        }

        const response = await fetch('http://localhost:3000/api/v1/auth/me', {
          method: 'GET',
          headers: {
            'Content-Type': 'application/json',
            'Authorization': `Bearer ${token}`,
          },
        });

        if (!response.ok) {
          console.log('Failed to fetch user data');
          setUser(null);
        } else {
          const data = await response.json();
          setUser(data);
        }
      } catch (error) {
        console.log('Error fetching user data:', error);
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
    <Navbar position="static" maxWidth="full">
      <NavbarBrand className="w-1/3" onClick={() => router.push('/')}>
        <p className="font-bold text-2xl cursor-pointer">RoboRent</p>
      </NavbarBrand>
      <NavbarContent className="hidden sm:flex gap-4" justify="center">
        <NavbarItem>
          <Link color="foreground" href="/">Home</Link>
        </NavbarItem>
        <NavbarItem>
          <Link color="foreground" href="/catalog">Catalog</Link>
        </NavbarItem>
        <NavbarItem>
          <Link color="foreground" href="/about">About</Link>
        </NavbarItem>
        <NavbarItem>
          <Link color="foreground" href="/contact">Contact</Link>
        </NavbarItem>
      </NavbarContent>
      <NavbarContent as="div" justify="end">
        <Link href="/cart">
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
            {user ? (
              <Avatar
                showFallback 
                isBordered
                as="button"
                className="transition-transform"
                color="warning"
                name={user?.data.name}
                size="sm"
                src=""
              />
            ) : (
              <Avatar
                showFallback
                isBordered
                as="button"
                className="transition-transform"
                color="warning"
                size="sm"
                src=""
              />
            )}
          </DropdownTrigger>
          <DropdownMenu aria-label="Profile Actions" variant="flat">
            {user ? (
              <>
                <DropdownItem key="profile" className="h-14 gap-2" onClick={() => (window.location.href = '/profile')}>
                  <p className="font-semibold">Signed in as</p>
                  <p className="font-semibold">{ user?.data.email }</p>
                </DropdownItem>
                <DropdownItem key="transactions" onClick={() => router.push('/transaction')}>Transactions</DropdownItem>
                <DropdownItem key="logout" color="danger" onClick={handleLogout}>
                  Log Out
                </DropdownItem>
              </>
            ) : (
              <>
                <DropdownItem key="login" onClick={() => (window.location.href = '/login')}>
                  Login
                </DropdownItem>
                <DropdownItem key="signup" onClick={() => (window.location.href = '/register')}>
                  Sign Up
                </DropdownItem>
              </>
            )}
          </DropdownMenu>
        </Dropdown>
      </NavbarContent>
    </Navbar>
  );
}
