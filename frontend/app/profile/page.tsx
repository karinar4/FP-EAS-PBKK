'use client';

import React, { useEffect, useState } from 'react';
import NavigationBar from '@/app/components/NavigationBar'
import Image from 'next/image';
import { Input, Form, Button, Card, CardBody, Alert } from "@nextui-org/react";


export default function Profile() {
    const [user, setUser] = useState({
        id: '',
        name: '',
        email: '',
        telephone: '',
        address: '',
    });
    const [alert, setAlert] = useState({ show: false, message: '' });
    const [isVisible, setIsVisible] = useState(true);


    const getTokenFromCookies = () => {
        const cookies = document.cookie.split('; ');
        const tokenCookie = cookies.find((cookie) => cookie.startsWith('auth-token='));
        return tokenCookie ? tokenCookie.split('=')[1] : null;
      };

  useEffect(() => {
    

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
        setUser(data.data);
      } catch (error) {
        console.error('Error fetching user data:', error);
      }
    };

    fetchUserData();
  }, []);

  const handleInputChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const { name, value } = e.target;
    setUser((prevUser) => ({
      ...prevUser,
      [name]: value,
    }));
  };

  const handleUpdate = async (e: React.FormEvent) => {
    e.preventDefault();
    try {
      const token = getTokenFromCookies();

        if (!token) {
          throw new Error('No authentication token found in cookies.');
        }

        console.log(token);

      const response = await fetch('http://localhost:3000/api/v1/auth/' + user.id, {
        method: 'PUT',
        headers: {
          'Content-Type': 'application/json',
          'Authorization': `Bearer ${token}`,
        },
        body: JSON.stringify(user),
      });

      if (!response.ok) {
        console.log('Failed to update user data');
      } else {
        const updatedData = await response.json();
        setUser(updatedData.data);
        console.log('Profile updated successfully!');
        setAlert({
          show: true,
          message: "Profile updated successfully!",
        });
      }

    } catch (error) {
      console.error('Error updating user data:', error);
      console.log('Failed to update profile.');
    }
  };

  return (
    <div>
      <NavigationBar />
      {alert.show && (
        <Alert
          color="success"
          title="Profile updated successfully"
          variant="faded"
          isVisible={isVisible}
          onClose={() => setIsVisible(false)}
        >
          {alert.message}
        </Alert>
      )}
      <main className='flex flex-col p-6 bg-gray-50 h-screen items-center'>
          <h1 className='font-bold text-gray-800 text-2xl'>Personal Info</h1>
          <p className='text-gray-600 py-2'>You can update your profile photo and personal details here.</p>
      <Image
        src="/profile-picture.png"
        height={120}
        width={120}
        alt="Dummy Image"
        className="rounded-full aspect-square object-cover mt-4 mb-8"
      />
      <Card shadow='sm' className='p-3 w-[500px] items-center'>
        <CardBody>
            <Form className='items-center'  onSubmit={handleUpdate}>
                <Input
                    name="name"
                    label="Name"
                    labelPlacement="outside"
                    value={user.name}
                    onChange={handleInputChange}
                    type="text"
                    placeholder="Enter your name"
                    className='mb-4'
                    isRequired
                />
                <Input
                    name="email"
                    label="Email"
                    labelPlacement="outside"
                    value={user.email}
                    onChange={handleInputChange}
                    type="text"
                    placeholder="Enter your email"
                    className='mb-4'
                    isRequired
                />
                <Input
                    name="telephone"
                    label="Phone Number"
                    labelPlacement="outside"
                    value={user.telephone}
                    onChange={handleInputChange}
                    type="text"
                    placeholder="Enter your phone number"
                    className='mb-4'
                    isRequired
                />
                <Input
                    name="address"
                    label="Address"
                    labelPlacement="outside"
                    value={user.address}
                    onChange={handleInputChange}
                    type="text"
                    placeholder="Enter your address"
                    className='mb-4'
                    isRequired
                />
                <Button color="primary" type="submit" className="text-md font-medium w-40" radius="sm">
                    Update
                </Button>
            </Form>
        </CardBody>
      </Card>
      </main>
    </div>
  );
};