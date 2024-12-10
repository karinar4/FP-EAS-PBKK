'use client';

import React, { useState, useEffect } from 'react';
import { useRouter } from 'next/navigation';
import { Form, Input, Button, Card, CardHeader, CardBody, Link, Alert } from '@nextui-org/react';
import { jwtDecode, JwtPayload } from 'jwt-decode';

interface CustomJwtPayload extends JwtPayload {
  role: string; // Menambahkan properti role pada tipe payload
}

export default function LoginForm() {
  const [error, setError] = useState<string>('');
  const router = useRouter();

  const checkToken = () => {
    const token = document.cookie.split(';').find(c => c.trim().startsWith('auth-token='));
      
    if (token){
      const decodedToken = jwtDecode<CustomJwtPayload>(token.split('=')[1]);
      const role = decodedToken.role; // Mendapatkan role dari token
      // Redirect berdasarkan role
      if (role === 'admin') {
        router.push('/admin/dashboard');
      } else {
        router.push('/dashboard');
      }
    }
  }

  useEffect(() => {
    checkToken();
  }, []);

  const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();

    const data = Object.fromEntries(new FormData(e.currentTarget));

    try {
      const response = await fetch('http://localhost:3000/api/v1/auth/login', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          email: data.email,
          password: data.password,
        }),
      });

      const result = await response.json();

      if (response.ok) {
        document.cookie = `auth-token=${result.data.token}; path=/;`;

        checkToken();
      } else {
        setError(`Error: ${result.message}`);
        console.error(`Error: ${result.message}`);
      }
    } catch (error: any) {
      setError(`Error: ${error.message}`);
      console.error(`Error: ${error.message}`);
    }
  };

  return (
    <Card className="w-full max-w-[400px] p-3" shadow="sm" radius="sm">
      <CardHeader className="pb-4 pt-3 px-4 flex-col items-center">
        <h1 className="font-bold text-4xl">Login</h1>
      </CardHeader>
      <CardBody>
        <Form
          className="gap-3"
          validationBehavior="native"
          onSubmit={handleSubmit}
        >
          {error && <Alert color="danger">{error}</Alert>}

          {/* Email */}
          <Input
            classNames={{
              label: ['text-md'],
              inputWrapper: ['shadow-none'],
            }}
            variant="underlined"
            isRequired
            errorMessage="Please enter a valid email"
            label="Email"
            labelPlacement="inside"
            name="email"
            type="email"
          />

          {/* Password */}
          <Input
            classNames={{
              label: ['text-md'],
              inputWrapper: ['shadow-none'],
            }}
            variant="underlined"
            isRequired
            errorMessage="Please enter your password"
            label="Password"
            labelPlacement="inside"
            name="password"
            type="password"
          />

          <div className="flex pt-7 pb-1 items-center w-full">
            <Button color="primary" type="submit" className="text-md font-bold w-full" radius="sm">
              Login
            </Button>
          </div>

          <p className="text-sm">
            Don't have an account?{' '}
            <Link href="/register" className="text-primary font-semibold">
              Sign Up
            </Link>
          </p>
        </Form>
      </CardBody>
    </Card>
  );
}
