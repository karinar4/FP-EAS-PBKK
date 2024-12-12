"use client";

import React, { useState, useEffect } from "react";
import { Navbar, NavbarContent, Dropdown, DropdownTrigger, DropdownMenu, DropdownItem, Avatar, Input, Form, Button, Card, CardBody, Select, SelectItem, Alert } from "@nextui-org/react";
import LogoutButton from '@/app/components/LogoutButton';
import { useRouter, useParams } from 'next/navigation';

export default function UpdateImageForm() {
    const { id } = useParams();
    const [formData, setFormData] = useState({
        product_id: "",
        url: "",
    });
    const [products, setProducts] = useState([]);
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

        const fetchImage = async () => {
            try {
                const token = getTokenFromCookies();
                if (!token) {
                    throw new Error('No authentication token found in cookies.');
                }

                const response = await fetch(`http://localhost:3000/api/v1/image/${id}`, {
                    method: 'GET',
                    headers: {
                        'Content-Type': 'application/json',
                        'Authorization': `Bearer ${token}`,
                    },
                });
                if (!response.ok) {
                    throw new Error('Failed to fetch brand');
                }
                const data = await response.json();
                setFormData({
                    product_id: data.data.name,
                    url: data.data.origin,
                });
            } catch (error) {
                console.error('Error fetching brand:', error);
            }
        };

        fetchUserData();
        fetchImage(); // Fetch product data to pre-fill form
    }, [id]);

    const handleChange = (e) => {
        const { name, value } = e.target;
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

            const response = await fetch(`http://localhost:3000/api/v1/image/${id}`, {
                method: "PUT", // Use PUT for updating
                headers: {
                    "Content-Type": "application/json",
                    'Authorization': `Bearer ${token}`,
                },
                body: JSON.stringify({
                    ...formData,
                }),
            });

            if (!response.ok) {
                throw new Error("Failed to update brand");
            } else {
                const data = await response.json();
                setAlert({
                    show: true,
                    message: "Brand updated successfully!",
                });
            }

            setFormData({
                product_id: "",
                url: "",
            });
        } catch (error) {
            console.error("Error updating brand:", error);
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
                <h1 className="font-bold text-gray-800 text-2xl mb-6">Update Image</h1>
                <Card shadow="sm" className="p-3 w-[500px] items-center">
                    <CardBody>
                        <Form className="items-center" onSubmit={handleSubmit}>
                            <Input
                                type="text"
                                name="product_id"
                                label="ProductID"
                                labelPlacement="outside"
                                value={formData.product_id}
                                onChange={handleChange}
                                className="mb-4"
                                isRequired
                            />
                            <Input
                                type="text"
                                name="url"
                                label="URL"
                                labelPlacement="outside"
                                value={formData.url}
                                onChange={handleChange}
                                className="mb-4"
                                isRequired
                            />
                            <div className="flex gap-5">
                                <Button type="submit" className="text-md font-medium w-40 mt-5 bg-yellow-500" radius="sm">
                                    Update
                                </Button>
                                <Button onClick={() => router.push('/admin/image')} className="text-md font-medium w-40 mt-5 bg-yellow-500">
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