"use client";

import React, { useEffect, useState } from "react";
import LogoutButton from "@/app/components/LogoutButton";
import Sidebar from "@/app/components/Sidebar";
import {
  Navbar,
  NavbarContent,
  Dropdown,
  DropdownTrigger,
  DropdownMenu,
  DropdownItem,
  Avatar,
  Table,
  TableHeader,
  TableColumn,
  TableBody,
  TableRow,
  TableCell,
  Button,
} from "@nextui-org/react";
import { useRouter } from "next/navigation";

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

export default function TransactionAdmin() {
  const [user, setUser] = useState<{ data: { email: string; name: string } } | null>(null);
  const [transactions, setTransactions] = useState<any[]>([]);
  const router = useRouter();

  useEffect(() => {
    const getTokenFromCookies = () => {
      const cookies = document.cookie.split("; ");
      const tokenCookie = cookies.find((cookie) => cookie.startsWith("auth-token="));
      return tokenCookie ? tokenCookie.split("=")[1] : null;
    };

    // Fetch user data from API
    const fetchUserData = async () => {
      try {
        const token = getTokenFromCookies();

        if (!token) {
          throw new Error("No authentication token found in cookies.");
        }

        const response = await fetch("http://localhost:3000/api/v1/auth/me", {
          method: "GET",
          headers: {
            "Content-Type": "application/json",
            Authorization: `Bearer ${token}`,
          },
        });
        if (!response.ok) {
          throw new Error("Failed to fetch user data");
        }
        const data = await response.json();
        setUser(data);
      } catch (error) {
        console.error("Error fetching user data:", error);
      }
    };

    // Fetch transaction data from API
    const fetchTransactionData = async () => {
      try {
        const token = getTokenFromCookies();

        if (!token) {
          throw new Error("No authentication token found in cookies.");
        }

        const response = await fetch("http://localhost:3000/api/v1/transaction/", {
          method: "GET",
          headers: {
            "Content-Type": "application/json",
            Authorization: `Bearer ${token}`,
          },
        });
        if (!response.ok) {
          throw new Error("Failed to fetch transaction data");
        }
        const data = await response.json();
        setTransactions(data.data.transactions);
      } catch (error) {
        console.error("Error fetching transaction data:", error);
      }
    };

    fetchUserData();
    fetchTransactionData();
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
            <div className="p-8 flex justify-between">
              <h1 className="text-3xl font-bold text-black">Transactions</h1>
            </div>

            {transactions && transactions.length > 0 ? (
              <Table className="mx-8 max-w-5xl">
                <TableHeader>
                  <TableColumn>ID</TableColumn>
                  <TableColumn>Transaction Date</TableColumn>
                  <TableColumn>Total Quantity</TableColumn>
                  <TableColumn>Total Price</TableColumn>
                  <TableColumn>Status</TableColumn>
                  <TableColumn>User Name</TableColumn>
                  <TableColumn>User Email</TableColumn>
                  <TableColumn>Action</TableColumn>
                </TableHeader>
                <TableBody>
                  {transactions.map((transaction) => (
                    <TableRow key={transaction.id}>
                      <TableCell>{transaction.id}</TableCell>
                      <TableCell>
                        {new Date(transaction.transaction_date).toLocaleString()}
                      </TableCell>
                      <TableCell>{transaction.total_quantity}</TableCell>
                      <TableCell>{transaction.total_price}</TableCell>
                      <TableCell>{transaction.status}</TableCell>
                      <TableCell>{transaction.user.name}</TableCell>
                      <TableCell>{transaction.user.email}</TableCell>
                      <TableCell>
                        <div className="relative flex justify-end items-center gap-2">
                          <Dropdown className="bg-background border-1 border-default-200">
                            <DropdownTrigger>
                              <Button isIconOnly radius="full" size="sm" variant="light">
                                <VerticalDotsIcon className="text-default-400" width={20} height={20} />
                              </Button>
                            </DropdownTrigger>
                            <DropdownMenu>
                              <DropdownItem key="edit" onClick={() => router.push(`/admin/transaction/update/${transaction.id}`)}>View Details</DropdownItem>
                            </DropdownMenu>
                          </Dropdown>
                        </div>
                      </TableCell>
                    </TableRow>
                  ))}
                </TableBody>
              </Table>
            ) : (
              <p className="mx-8">No Transactions Found</p>
            )}
          </div>
        </main>
      </div>
    </>
  );
}
