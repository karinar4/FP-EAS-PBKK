"use client";

import { Card, Input, Button, Textarea, CardBody } from "@nextui-org/react";
import Navbar from "../components/NavigationBar";

export default function ContactPage() {
    return (
        <>
            <Navbar />
            <main className="flex flex-col items-center p-8 bg-gray-50 min-h-screen">
                
                <section>
                    <div className="container mx-auto text-center">
                        <h1 className="text-3xl font-bold text-gray-800">About Us</h1>
                        <p className="mt-4 text-xl text-gray-700">Leading the Future of Robotics and Drone Technology Rentals</p>
                    </div>
                </section>

                <section>
                    <div className="max-w-3xl mt-8">
                        <Card className="border border-yellow-400 p-5 rounded-lg" shadow="sm">
                            <CardBody className="items-center justify-center">
                                <h2 className="text-3xl font-semibold text-gray-800">Our Mission</h2>
                            <p className="mt-4 text-lg text-gray-600">
                                At RoboDrones, our mission is to provide cutting-edge robotics and drone technologies to enhance business operations, support education, and enable exciting personal experiences. We are committed to providing top-tier rental services with a focus on customer satisfaction, reliability, and innovation.
                            </p>
                            </CardBody>
                        </Card>
                        
                    </div>
                </section>

                <section className="py-10 bg-gray-50 max-w-3xl">
                    <div className="container mx-auto text-center">
                        <h2 className="text-3xl font-semibold text-gray-800">What We Offer</h2>
                        <p className="mt-4 text-lg text-gray-600">Explore our wide range of rental options designed for various industries and use cases.</p>

                        <div className="mt-12 grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-8">
                            <div className="border border-yellow-400 p-6 rounded-lg bg-white shadow-md">
                                <h3 className="text-xl font-semibold text-gray-800">Robots for Business</h3>
                                <p className="mt-4 text-gray-600">Optimize your workflow and reduce operational costs with our robotic solutions tailored for business needs.</p>
                            </div>
                            <div className="border border-yellow-400 p-6 rounded-lg bg-white shadow-md">
                                <h3 className="text-xl font-semibold text-gray-800">Drone Rentals</h3>
                                <p className="mt-4 text-gray-600">Capture stunning aerial footage or perform inspections with our high-quality drones available for rental.</p>
                            </div>
                            <div className="border border-yellow-400 p-6 rounded-lg bg-white shadow-md">
                                <h3 className="text-xl font-semibold text-gray-800">Training & Support</h3>
                                <p className="mt-4 text-gray-600">Receive expert guidance and training to make the most of our advanced equipment and technology.</p>
                            </div>
                        </div>
                    </div>
                </section>

                <section className="py-16 px-4 text-center bg-yellow-50 rounded-md max-w-3xl">
                    <div className="mx-auto">
                        <h2 className="text-3xl font-semibold text-gray-800">Why Choose Us?</h2>
                        <p className="mt-4 text-lg text-gray-600">Discover the reasons our customers trust us with their robotics and drone needs:</p>

                        <ul className="mt-8 space-y-4 text-left mx-auto max-w-3xl text-gray-600">
                            <li className="flex items-center">
                                <span className="mr-3">✔</span>
                                Top-of-the-line equipment from the best manufacturers.
                            </li>
                            <li className="flex items-center">
                                <span className=" mr-3">✔</span>
                                Flexible rental terms to accommodate long-term or short-term projects.
                            </li>
                            <li className="flex items-center">
                                <span className=" mr-3">✔</span>
                                Affordable pricing and access to high-tech equipment without the commitment of full ownership.
                            </li>
                            <li className="flex items-center">
                                <span className="mr-3">✔</span>
                                Expert technical support and training available to maximize your rental experience.
                            </li>
                        </ul>
                    </div>
                </section>

                <section className="py-16  text-center">
                    <div className="container mx-auto">
                        <h2 className="text-3xl font-semibold text-gray-800">Get in Touch</h2>
                        <p className="mt-4 text-lg text-gray-600">Ready to rent a robot or drone? Contact us today to learn more or to get started with your rental.</p>

                        <a href="/contact" className="mt-8 inline-block px-8 py-3 bg-yellow-500 text-white font-semibold rounded-lg hover:bg-yellow-400">Contact Us</a>
                    </div>
                </section>
            </main>
        </>
    );
}
