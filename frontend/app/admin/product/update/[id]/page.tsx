"use client";

import React, { useState, useEffect } from "react";
import { Navbar, NavbarContent, Dropdown, DropdownTrigger, DropdownMenu, DropdownItem, Avatar, Input, Form, Button, Card, CardBody, Select, SelectItem, Alert } from "@nextui-org/react";
import LogoutButton from '@/app/components/LogoutButton';
import { useRouter, useParams } from 'next/navigation';

export default function UpdateProductForm() {
  const { id } = useParams(); // Get product ID from URL params
  const [formData, setFormData] = useState({
    name: "",
    description: "",
    price: 0,
    stock: 0,
    category_id: "",
    brand_id: "",
  });
  const [categories, setCategories] = useState([]);
  const [brands, setBrands] = useState([]);
  const [alert, setAlert] = useState({ show: false, message: '' });
  const [isVisible, setIsVisible] = useState(true);
  const [user, setUser] = useState<{ data: { email: string; name: string } } | null>(null);
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

    const fetchCategories = async () => {
      try {
        const token = getTokenFromCookies();
        if (!token) {
          throw new Error('No authentication token found in cookies.');
        }

        const response = await fetch('http://localhost:3000/api/v1/category/', {
          method: 'GET',
          headers: {
            'Content-Type': 'application/json',
            'Authorization': `Bearer ${token}`,
          },
        });
        if (!response.ok) {
          throw new Error('Failed to fetch categories');
        }
        const data = await response.json();
        setCategories(data.data.categories);
      } catch (error) {
        console.error('Error fetching categories:', error);
      }
    };

    const fetchBrands = async () => {
      try {
        const token = getTokenFromCookies();
        if (!token) {
          throw new Error('No authentication token found in cookies.');
        }

        const response = await fetch('http://localhost:3000/api/v1/brand/', {
          method: 'GET',
          headers: {
            'Content-Type': 'application/json',
            'Authorization': `Bearer ${token}`,
          },
        });
        if (!response.ok) {
          throw new Error('Failed to fetch brands');
        }
        const data = await response.json();
        setBrands(data.data.brands);
      } catch (error) {
        console.error('Error fetching brands:', error);
      }
    };

    const fetchProduct = async () => {
      try {
        const token = getTokenFromCookies();
        if (!token) {
          throw new Error('No authentication token found in cookies.');
        }

        const response = await fetch(`http://localhost:3000/api/v1/product/${id}`, {
          method: 'GET',
          headers: {
            'Content-Type': 'application/json',
            'Authorization': `Bearer ${token}`,
          },
        });
        if (!response.ok) {
          throw new Error('Failed to fetch product');
        }
        const data = await response.json();
        setFormData({
          name: data.data.name,
          description: data.data.description,
          price: data.data.price,
          stock: data.data.stock,
          category_id: data.data.category_id,
          brand_id: data.data.brand_id,
        });
      } catch (error) {
        console.error('Error fetching product:', error);
      }
    };

    fetchUserData();
    fetchCategories();
    fetchBrands();
    fetchProduct(); // Fetch product data to pre-fill form
  }, [id]);

  const handleChange = (e) => {
    const { name, value } = e.target;
    setFormData((prevData) => ({
      ...prevData,
      [name]: name === "price" || name === "stock" ? Number(value) : value,
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

      const response = await fetch(`http://localhost:3000/api/v1/product/${id}`, {
        method: "PUT", // Use PUT for updating
        headers: {
          "Content-Type": "application/json",
          'Authorization': `Bearer ${token}`,
        },
        body: JSON.stringify({
          ...formData,
          category_id: formData.category_id.currentKey,
          brand_id: formData.brand_id.currentKey,
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
      console.error("Error updating product:", error);
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
                name="name"
                label="Name"
                labelPlacement="outside"
                value={formData.name}
                onChange={handleChange}
                className="mb-4"
                isRequired
              />
              <Input
                type="text"
                name="description"
                label="Description"
                labelPlacement="outside"
                value={formData.description}
                onChange={handleChange}
                className="mb-4"
                isRequired
              />
              <Input
                type="number"
                name="price"
                label="Price"
                labelPlacement="outside"
                value={formData.price}
                onChange={handleChange}
                className="mb-4"
                isRequired
              />
              <Input
                type="number"
                name="stock"
                label="Stock"
                labelPlacement="outside"
                value={formData.stock}
                onChange={handleChange}
                className="mb-4"
                isRequired
              />
              <Select
                label="Category"
                labelPlacement="outside"
                selectedKeys={formData.category_id}
                onSelectionChange={(value) => handleSelectChange("category_id", value)}
                isRequired
              >
                {categories.map((category) => (
                  <SelectItem key={category.id} value={category.id}>
                    {category.name}
                  </SelectItem>
                ))}
              </Select>
              <Select
                label="Brand"
                labelPlacement="outside"
                selectedKeys={formData.brand_id}
                onSelectionChange={(value) => handleSelectChange("brand_id", value)}
                isRequired
              >
                {brands.map((brand) => (
                  <SelectItem key={brand.id} value={brand.id}>
                    {brand.name}
                  </SelectItem>
                ))}
              </Select>
              <div className="flex gap-5">
                <Button type="submit" className="text-md font-medium w-40 mt-5 bg-yellow-500" radius="sm">
                  Update
                </Button>
                <Button onClick={() => router.push('/admin/product')} className="text-md font-medium w-40 mt-5 bg-yellow-500">
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
