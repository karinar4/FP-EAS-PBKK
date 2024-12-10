"use client";

import { Card, Input, Button, Textarea } from "@nextui-org/react";
import Navbar from "../components/NavigationBar";

export default function ContactPage() {
    return (
        <>
            <Navbar />
            <main className="flex flex-col items-center p-8 bg-gray-50 min-h-screen">
                <h1 className="text-3xl font-bold mb-8 text-gray-800">Contact Us</h1>
                <div className="flex flex-col lg:flex-row justify-between w-full max-w-7xl gap-8">
                    <div className="flex-1 bg-white shadow-md rounded-lg p-6">
                        <h2 className="text-2xl font-semibold mb-4 text-gray-700">
                            Contact Details
                        </h2>
                        <p className="mb-6 text-gray-600">
                            We're here to help and can be reached in three ways!
                        </p>
                        <div className="space-y-6">
                            <Card className="border border-yellow-400 p-4 rounded-lg">
                                <h3 className="font-semibold text-gray-800">Email</h3>
                                <p className="text-gray-600">
                                    Use the form below for order and technical support.
                                </p>
                            </Card>
                            <Card className="border border-yellow-400 p-4 rounded-lg">
                                <h3 className="font-semibold text-gray-800">Phone</h3>
                                <p className="text-gray-600">
                                    +1 (941) 444-0021 <br />
                                    Monday - Friday: 11:00 AM - 6:00 PM EST
                                </p>
                            </Card>
                            <Card className="border border-yellow-400 p-4 rounded-lg">
                                <h3 className="font-semibold text-gray-800">Live Chat</h3>
                                <p className="text-gray-600">
                                    Visit any page and start a chat for shopping assistance.
                                </p>
                            </Card>
                        </div>

                        <div className="mt-8 space-y-4 text-sm text-gray-600">
                            <p>
                                We can cancel or remove items from orders until they're packed. Please
                                contact us as soon as possible!
                            </p>
                            <p>
                                Need pictures to help us resolve an issue? Email us directly at{" "}
                                <a
                                    href="mailto:Support@RoboRent.com"
                                    className="text-blue-600 underline"
                                >
                                    Support@RoboRent.com
                                </a>{" "}
                                with images and relevant details.
                            </p>
                        </div>
                    </div>

                    <div className="flex-1 bg-white shadow-md rounded-lg p-6">
                        <h2 className="text-2xl font-semibold mb-4 text-gray-700">
                            Write Us
                        </h2>
                        <form className="space-y-6">
                            <div className="grid grid-cols-1 sm:grid-cols-2 gap-4">
                                <Input
                                    label="Name"
                                    placeholder="Name"
                                    isRequired
                                    className="bg-gray-100"
                                />
                                <Input
                                    label="Email"
                                    placeholder="Email"
                                    isRequired
                                    type="email"
                                    className="bg-gray-100"
                                />
                            </div>
                            <div className="grid grid-cols-1 sm:grid-cols-2 gap-4">
                                <Input
                                    label="Telephone"
                                    placeholder="Phone Number"
                                    className="bg-gray-100"
                                />
                                <Input
                                    label="Product"
                                    placeholder="Product"
                                    className="bg-gray-100"
                                />
                            </div>
                            <Input
                                label="Order No"
                                placeholder="Order Number"
                                className="bg-gray-100"
                            />
                            <Textarea
                                label="What's on your mind?"
                                placeholder="Write your message here"
                                rows={5}
                                isRequired
                                className="bg-gray-100"
                            />
                            <Button
                                type="submit"
                                color="primary"
                                className="w-full bg-yellow-500 hover:bg-yellow-600 text-white"
                            >
                                Submit
                            </Button>
                        </form>
                    </div>
                </div>
            </main>
        </>
    );
}
