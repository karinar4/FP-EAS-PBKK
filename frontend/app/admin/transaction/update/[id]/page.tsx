"use client";

import React, { useState, useEffect } from "react";
import { Navbar, NavbarContent, Dropdown, DropdownTrigger, DropdownMenu, DropdownItem, Avatar, Input, Form, Button, Card, CardBody, Select, SelectItem, Alert } from "@nextui-org/react";
import LogoutButton from '@/app/components/LogoutButton';
import { useRouter, useParams } from 'next/navigation';
import Transaction from "@/app/transaction/page";



export default function UpdateTransactionForm() {
  const Status = {
    pending: "pending",
    in_progress: "in_progress",
    completed: "completed",
    cancelled: "cancelled"
  };

  const { id } = useParams(); // Get product ID from URL params
  const [formData, setFormData] = useState({
    transaction_date: "",
    total_quantity: 0,
    total_price: 0.0,
    status: Status.pending,
  });
  const [categories, setCategories] = useState([]);
  const [transactions, setTransactions] = useState([]);
  const [alert, setAlert] = useState({ show: false, message: '' });
  const [isVisible, setIsVisible] = useState(true);
  const [user, setUser] = useState<{ data: { id: string; email: string; name: string } } | null>(null);
  const router = useRouter();

  const getTokenFromCookies = () => {
    const cookies = document.cookie.split('; ');
    const tokenCookie = cookies.find((cookie) => cookie.startsWith('auth-token='));
    return tokenCookie ? tokenCookie.split('=')[1] : null;
  };

  useEffect(() => {
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

    const fetchTransactionData = async () => {
      try {
        const token = getTokenFromCookies();
        if (!token) {
          throw new Error('No authentication token found in cookies.');
        }

        const response = await fetch(`http://localhost:3000/api/v1/transaction/${id}`, {
          method: 'GET',
          headers: {
            'Content-Type': 'application/json',
            'Authorization': `Bearer ${token}`,
          },
        });
        if (!response.ok) {
          throw new Error('Failed to fetch transaction');
        }
        const data = await response.json();
        console.log(data);
        const transactionDate = new Date(data.data.transaction_date);
        setFormData({
          // name: data.data.name,
          // description: data.data.description,
          // price: data.data.price,
          // stock: data.data.stock,
          // category_id: data.data.category_id,
          // brand_id: data.data.brand_id,
          transaction_date: transactionDate.toLocaleString(),
          total_quantity: data.data.total_quantity,
          total_price: data.data.total_price,
          status: data.data.status,
        });
      } catch (error) {
        console.error('Error fetching transaction:', error);
      }
    };

    fetchUserData();
    fetchTransactionData(); // Fetch product data to pre-fill form
  }, [id]);

  const handleChange = (e) => {
    const { name, value } = e.target;
    setFormData((prevData) => ({
      ...prevData,
      [name]: name === "total_quantity" || name === "total_price" ? Number(value) : value,
    }));
  };

  const handleSelectChange = (name, value) => {
    setFormData((prevData) => ({
      ...prevData,
      [name]: value,
    }));
  };

  const handleSubmit = async (e) => {
    e.preventDefault();

    try {
      console.log(formData);
      const token = getTokenFromCookies();

      if (!token) {
        throw new Error('No authentication token found in cookies.');
      }

      console.log(formData);
      const response = await fetch(`http://localhost:3000/api/v1/transaction/${id}`, {
        method: "PUT", // Use PUT for updating
        headers: {
          "Content-Type": "application/json",
          'Authorization': `Bearer ${token}`,
        },
        body: JSON.stringify({
          status: formData.status.currentKey,
        }),
      });

      if (!response.ok) {
        throw new Error("Failed to update product");
      } else {
        const data = await response.json();
        setAlert({
          show: true,
          message: "Product updated successfully!",
        });
      }
    } catch (error) {
      console.error("Error updating transaction:", error);
    }
  };

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

      {alert.show && (
        <Alert
          color="success"
          title="Product updated successfully"
          variant="faded"
          isVisible={isVisible}
          onClose={() => setIsVisible(false)}
        >
          {alert.message}
        </Alert>
      )}

      <main className="flex flex-col p-6 bg-gray-50 h-screen items-center">
        <h1 className="font-bold text-gray-800 text-2xl mb-6">Update Product</h1>
        <Card shadow="sm" className="p-3 w-[500px] items-center">
          <CardBody>
            <Form className="items-center" onSubmit={handleSubmit}>
              <Input
                type="text"
                name="transaction_date"
                label="Transaction Date"
                labelPlacement="outside"
                value={formData.transaction_date}
                // onChange={handleChange}
                className="mb-4"
                // isRequired
              />
              <Input
                type="number"
                name="total_quantity"
                label="Total Quantity"
                labelPlacement="outside"
                value={formData.total_quantity}
                // onChange={handleChange}
                className="mb-4"
                // isRequired
              />
              <Input
                type="number"
                name="total_price"
                label="Total Price"
                labelPlacement="outside"
                value={formData.total_price}
                // onChange={handleChange}
                className="mb-4"
                // isRequired
              />
              <Select
                label="Status"
                labelPlacement="outside"
                selectedKeys={formData.status}
                onSelectionChange={(value) => handleSelectChange("status", value)}
                isRequired
              >
                <SelectItem key={Status.pending} value={Status.pending}>
                  Pending
                </SelectItem>
                <SelectItem key={Status.in_progress} value={Status.in_progress}>
                  In Progress
                </SelectItem>
                <SelectItem key={Status.completed} value={Status.completed}>
                  Completed
                </SelectItem>
                <SelectItem key={Status.cancelled} value={Status.cancelled}>
                  Cancelled
                </SelectItem>
              </Select>
              <div className="flex gap-5">
                <Button type="submit" className="text-md font-medium w-40 mt-5 bg-yellow-500" radius="sm">
                  Update
                </Button>
                <Button onClick={() => router.push('/admin/transaction/')} className="text-md font-medium w-40 mt-5 bg-yellow-500">
                  Back
                </Button>
              </div>
            </Form>
          </CardBody>
        </Card>
      </main>
    </>
  );
}
