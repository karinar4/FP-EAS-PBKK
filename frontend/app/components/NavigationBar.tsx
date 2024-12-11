"use client";

import { Navbar, NavbarBrand, NavbarContent, NavbarItem, Link, Dropdown, DropdownTrigger, DropdownMenu, DropdownItem, Avatar } from "@nextui-org/react";
import Image from 'next/image';
import { useEffect, useState } from "react";
import { useRouter } from 'next/navigation';

export const AcmeLogo = () => {
  return (
    <svg fill="none" height="36" viewBox="0 0 32 32" width="36">
      <path
        clipRule="evenodd"
        d="M17.6482 10.1305L15.8785 7.02583L7.02979 22.5499H10.5278L17.6482 10.1305ZM19.8798 14.0457L18.11 17.1983L19.394 19.4511H16.8453L15.1056 22.5499H24.7272L19.8798 14.0457Z"
        fill="currentColor"
        fillRule="evenodd"
      />
    </svg>
  );
};

export default function NavigationBar() {
  // const [isLoggedIn, setIsLoggedIn] = useState(false);
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
          // setIsLoggedIn(false);
        } else {
          const data = await response.json();
            // setIsLoggedIn(true);
          setUser(data);
        }
      } catch (error) {
        console.log('Error fetching user data:', error);
      }
    };

    fetchUserData();
    // console.log(isLoggedIn);
  }, []);

  const handleLogout = () => {
    document.cookie = 'auth-token=; Max-Age=0';
    // setIsLoggedIn(false);
    setUser(null);
    router.push("/");
  };
  
  return (
    <Navbar position="static" maxWidth="full">
      <NavbarBrand className="w-1/3">
        <AcmeLogo />
        <p className="font-bold text-2xl">ACME</p>
      </NavbarBrand>
      <NavbarContent className="hidden sm:flex gap-4" justify="center">
        <NavbarItem>
          <Link color="foreground" href="#">
            About
          </Link>
        </NavbarItem>
        <NavbarItem>
          <Link color="foreground" href="contact">
            Contact
          </Link>
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
                color="secondary"
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
                color="secondary"
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
                <DropdownItem key="transactions">Transactions</DropdownItem>
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
