"use client";

import React from "react";
import { Form, Input, Button, Card, CardHeader, CardBody, Link, Alert } from "@nextui-org/react";
import { useRouter } from 'next/navigation';

export default function RegisterForm() {
  const [error, setError] = React.useState<string>('');
  const router = useRouter();

  const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();

    const data = Object.fromEntries(new FormData(e.currentTarget));

    if (data.password !== data.confirm_password) {
      setError("Password and Confirm Password do not match");
      return;
    }

    try {
      const response = await fetch("http://localhost:3000/api/v1/auth/register", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          name: data.name,
          email: data.email,
          password: data.password,
          confirm_password: data.confirm_password,
        }),
      });

      if (response.ok) {
        const result = await response.json();
        console.log(`Registration successful: ${JSON.stringify(result)}`);

        router.push('/login');
      } else {
        const error = await response.json();
        setError(`Error: ${error.message}`);
        console.error(`Error: ${error.message}`);
      }
    } catch (error: any) {
      setError(`Error: ${error.message}`);
      console.error(`Error: ${error.message}`);
    }
  };

  return (
      <Card className="w-full max-w-[400px] p-3" shadow="sm" radius="sm">
        <CardHeader className="pb-4 pt-3 px-4 flex-col items-center">
            <h1 className="font-bold text-4xl">Sign Up</h1>
        </CardHeader>
        <CardBody>
          <Form
            className="gap-3"
            validationBehavior="native"
            onSubmit={handleSubmit}
          >

            {error && <Alert color="danger">
              {error}
            </Alert>}

            {/* Full Name */}
            <Input
              classNames={{ 
                label: ["text-md"],
                inputWrapper: ["shadow-none"]
              }}
              variant="underlined"
              isRequired
              errorMessage="Please enter your full name"
              label="Full Name"
              labelPlacement="inside"
              name="name"
              type="text"
            />

            {/* Email */}
            <Input
              classNames={{ 
                label: ["text-md"],
                inputWrapper: ["shadow-none"]
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
                label: ["text-md"],
                inputWrapper: ["shadow-none"]
              }}
              variant="underlined"
              isRequired
              errorMessage="Please enter a password"
              label="Password"
              labelPlacement="inside"
              name="password"
              type="password"
            />

            {/* Confirm Password */}
            <Input
              classNames={{ 
                label: ["text-md"],
                inputWrapper: ["shadow-none"]
              }}
              variant="underlined"
              isRequired
              errorMessage="Please confirm your password"
              label="Confirm Password"
              labelPlacement="inside"
              name="confirm_password"
              type="password"
            />

            <div className="flex pt-7 pb-1 items-center w-full">
              <Button color="primary" type="submit" className="text-md font-bold w-full" radius="sm">
                Sign Up
              </Button>
            </div>

            <p className="text-sm">Already have an account? <Link href="/login" className="text-primary font-semibold">Login</Link></p>
          </Form>
        </CardBody>
      </Card>
  );
}
