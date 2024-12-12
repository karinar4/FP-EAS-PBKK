"use client";

import NavigationBar from "@/app/components/NavigationBar";
import React, { useState, useEffect } from "react";
import { Input, Form, Button, Card, CardBody, Select, SelectItem } from "@nextui-org/react";

export default function CreateProductForm() {
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

  const getTokenFromCookies = () => {
    const cookies = document.cookie.split('; ');
    const tokenCookie = cookies.find((cookie) => cookie.startsWith('auth-token='));
    return tokenCookie ? tokenCookie.split('=')[1] : null;
  };

  useEffect(() => {
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

    fetchCategories();
    fetchBrands();
  }, []);

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

      const response = await fetch("http://localhost:3000/api/v1/product/", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          'Authorization': `Bearer ${token}`,
        },
        body: JSON.stringify(formData),
      });

      if (!response.ok) {
        throw new Error("Failed to create product");
      }

      const data = await response.json();
      alert("Product created successfully!");
      console.log(data);

      // Reset form after successful submission
      setFormData({
        name: "",
        description: "",
        price: 0,
        stock: 0,
        category_id: "",
        brand_id: "",
      });
    } catch (error) {
      console.error("Error creating product:", error);
    }
  };

  return (
    <>
      <NavigationBar />
      <main className='flex flex-col p-6 bg-gray-50 h-screen items-center'>
        <h1 className='font-bold text-gray-800 text-2xl mb-6'>Create Product</h1>
        <Card shadow='sm' className='p-3 w-[500px] items-center'>
          <CardBody>
            <Form className='items-center' onSubmit={handleSubmit}>
              <Input
                type="text"
                name="name"
                label="Name"
                labelPlacement="outside"
                value={formData.name}
                onChange={handleChange}
                className='mb-4'
                isRequired
              />
              <Input
                type="text"
                name="description"
                label="Description"
                labelPlacement="outside"
                value={formData.description}
                onChange={handleChange}
                className='mb-4'
                isRequired
              />
              <Input
                type="number"
                name="price"
                label="Price"
                labelPlacement="outside"
                value={formData.price}
                onChange={handleChange}
                className='mb-4'
                isRequired
              />
              <Input
                type="number"
                name="stock"
                label="Stock"
                labelPlacement="outside"
                value={formData.stock}
                onChange={handleChange}
                className='mb-4'
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
              <Button type="submit" className="text-md font-medium w-40 mt-5 bg-yellow-500" radius="sm">
                Create
              </Button>
            </Form>
          </CardBody>
        </Card>
      </main>
    </>
  );
}
