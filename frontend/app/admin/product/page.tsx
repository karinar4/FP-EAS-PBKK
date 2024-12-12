"use client";

import React, { useEffect, useState } from 'react';
import LogoutButton from '@/app/components/LogoutButton';
import Sidebar from '@/app/components/Sidebar';
import { Navbar, NavbarContent, Dropdown, DropdownTrigger, DropdownMenu, DropdownItem, Avatar, Table, TableHeader, TableColumn, TableBody, TableRow, TableCell, Button } from "@nextui-org/react";
import NavigationBar from '@/app/components/NavigationBar';
import { useRouter } from 'next/navigation';

export const VerticalDotsIcon = ({ size = 24, width, height, ...props }) => {
  return (
    <svg
      aria-hidden="true"
      fill="none"
      focusable="false"
      height={size || height}
      role="presentation"
      viewBox="0 0 24 24"
      width={size || width}
      {...props}
    >
      <path
        d="M12 10c-1.1 0-2 .9-2 2s.9 2 2 2 2-.9 2-2-.9-2-2-2zm0-6c-1.1 0-2 .9-2 2s.9 2 2 2 2-.9 2-2-.9-2-2-2zm0 12c-1.1 0-2 .9-2 2s.9 2 2 2 2-.9 2-2-.9-2-2-2z"
        fill="currentColor"
      />
    </svg>
  );
};

export default function ProductAdmin() {
  const [user, setUser] = useState<{ data: { email: string; name: string } } | null>(null);
  const [products, setProducts] = useState<any[]>([]);
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

    // Fetch product data from API
    const fetchProductData = async () => {
      try {
        const token = getTokenFromCookies();

        if (!token) {
          throw new Error('No authentication token found in cookies.');
        }

        const response = await fetch('http://localhost:3000/api/v1/product/', {
          method: 'GET',
          headers: {
            'Content-Type': 'application/json',
            'Authorization': `Bearer ${token}`,
          },
        });
        if (!response.ok) {
          throw new Error('Failed to fetch product data');
        }
        const data = await response.json();
        setProducts(data.data.products);
      } catch (error) {
        console.error('Error fetching product data:', error);
      }
    };

    fetchUserData();
    fetchProductData();
  }, []);

  const handleLogout = () => {
    // Implement logout logic, such as clearing cookies or redirecting
    LogoutButton();
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
                color="secondary"
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
          <Sidebar />
          <div className="flex-1">
            <div className="bg-white w-full shadow-sm">
              
            </div>

          <div className='flex-1 w-full'>
            <div className="p-8 flex justify-between">
              <h1 className="text-3xl font-bold text-black">Product</h1>
              <Button className="bg-yellow-500 font-medium" onClick={() => router.push('/admin/product/create')}>Create New</Button>
            </div>

            {products.length > 0 ? (
              <Table className="mx-8 max-w-5xl">
                <TableHeader>
                  <TableColumn>Id</TableColumn>
                  <TableColumn>Name</TableColumn>
                  <TableColumn>Price</TableColumn>
                  <TableColumn>Stock</TableColumn>
                  <TableColumn>Category</TableColumn>
                  <TableColumn>Brand</TableColumn>
                  <TableColumn>Action</TableColumn>
                </TableHeader>
                <TableBody>
                  {products.map((product) => (
                    <TableRow key={product.id}>
                      <TableCell className="py-4">{product.id}</TableCell>
                      <TableCell className="py-4">{product.name}</TableCell>
                      <TableCell className="py-4">{product.price}</TableCell>
                      <TableCell className="py-4">{product.stock}</TableCell>
                      <TableCell className="py-4">{product.category.name}</TableCell>
                      <TableCell className="py-4">{product.brand.name}</TableCell>
                      <TableCell>
                        <div className="relative flex justify-end items-center gap-2">
                          <Dropdown className="bg-background border-1 border-default-200">
                            <DropdownTrigger>
                              <Button isIconOnly radius="full" size="sm" variant="light">
                                <VerticalDotsIcon className="text-default-400" width={20} height={20} />
                              </Button>
                            </DropdownTrigger>
                            <DropdownMenu>
                              <DropdownItem key="edit">Edit</DropdownItem>
                              <DropdownItem key="delete">Delete</DropdownItem>
                            </DropdownMenu>
                          </Dropdown>
                        </div>
                      </TableCell>
                    </TableRow>
                  ))}
                </TableBody>
              </Table>
            ) : (
              <p>No Products</p>
            )}
          </div>
          </div>
        </main>
      </div>
    </>
  );
}
