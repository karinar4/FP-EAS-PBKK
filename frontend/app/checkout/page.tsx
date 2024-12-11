'use client';

import React, { useEffect, useState } from 'react';
import NavigationBar from '@/app/components/NavigationBar';
import { Card, CardBody, Alert, Button, Divider, Form, Input, Link } from "@nextui-org/react";
import { useRouter } from 'next/navigation';

export default function Checkout() {
    const [user, setUser] = useState<{ data: { id: string; email: string; name: string; telephone: string; address: string } } | null>(null);
    const [alert, setAlert] = useState({ show: false, message: '' });
    const [cart, setCart] = useState<{ data: { id: string; total_price: number; total_quantity: number, user_id: string } } | null>(null);
    const [cartProducts, setCartProducts] = useState<any[]>([]);
    const [accountNumber, setAccountNumber] = useState("");
    const router = useRouter();

    const getTokenFromCookies = () => {
        const cookies = document.cookie.split('; ');
        const tokenCookie = cookies.find((cookie) => cookie.startsWith('auth-token='));
        return tokenCookie ? tokenCookie.split('=')[1] : null;
    };
    const token = getTokenFromCookies();

    if (!token) {
        throw new Error('No authentication token found in cookies.');
    }

    const fetchUserData = async () => {
        try {
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
            setUser(prevUser => {
                if (prevUser?.data.id !== data.data.id) {  // Update only if data has changed
                    return data;
                }
                return prevUser;
            });
            if (data.data.telephone === "" || data.data.address === "") {
                const missingFields = [];
                if (data.data.telephone === "") missingFields.push('phone number');
                if (data.data.address === "") missingFields.push('address');
                setAlert({
                    show: true,
                    message: `Please fill in the following fields: ${missingFields.join(', ')}`,
                });
            }
            return data;
        } catch (error) {
            console.error('Error fetching user data:', error);
        }
    };

    const fetchCart = async (userId: string) => {
        try {
            const response = await fetch(`http://localhost:3000/api/v1/cart/` + userId, {
                method: 'GET',
                headers: {
                    'Content-Type': 'application/json',
                    Authorization: `Bearer ${token}`,
                },
            });

            const data = await response.json();
            if (response.ok && data.data.id !== cart?.data.id) {
                setCart(data); // Update cart only if the cart has changed
            }
            return data;
        } catch (error) {
            console.log('Error fetching cart:', error);
        }
    };

    const fetchCartProducts = async (cartId: string) => {
        try {
            const response = await fetch(`http://localhost:3000/api/v1/cart_product/` + cartId, {
                method: 'GET',
                headers: {
                    'Content-Type': 'application/json',
                    Authorization: `Bearer ${token}`,
                },
            });

            if (response.ok) {
                const data = await response.json();
                if (data.data.cart_products !== cartProducts) {
                    setCartProducts(data.data.cart_products || []); // Only update if products changed
                }
            } else {
                console.log('Error fetching cart products:', response);
            }
        } catch (error) {
            console.log('Error fetching cart products:', error);
        }
    };

    const createCart = async (userId: string, token: string) => {
        try {
          const response = await fetch('http://localhost:3000/api/v1/cart/', {
            method: 'POST',
            headers: {
              'Content-Type': 'application/json',
              'Authorization': `Bearer ${token}`,
            },
            body: JSON.stringify({ user_id: userId }),
          });
  
          if (response.ok) {
            const data = await response.json();
            console.log('Cart created:', data);
            // setCart(data);
            // setCartCreated(true);
          } else {
            console.log('Error creating cart:', response);
          }
        } catch (error) {
          console.log('Error creating cart:', error);
        }
      };

    useEffect(() => {
        const fetchData = async () => {
            const user_data = await fetchUserData();
            if (user_data?.data.id) {
                console.log(user_data);
                const cart_data = await fetchCart(user_data.data.id);
                if (cart_data?.data.id) {
                    console.log(cart_data);
                    await fetchCartProducts(cart_data.data.id);
                }
            }
        };

        fetchData();
    }, []);  // Run only once on mount
    
    const handleTransaction = async (e: React.FormEvent) => {
        e.preventDefault();
        if (!user || !cart || cartProducts.length === 0) return;

        let transactionId: string | undefined;
        
        console.log(cartProducts[0]);

        try {
            const response = await fetch('http://localhost:3000/api/v1/transaction/', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                    'Authorization': `Bearer ${token}`,
                },
                body: JSON.stringify({
                    total_quantity: cart.data.total_quantity,
                    total_price: cart?.data.total_price,
                    user_id: user.data.id,
                    product_transactions: cartProducts.map(cartProduct => ({
                        cart_id: cartProduct.cart_id,
                        product_id: cartProduct.product.id, // Only include the product ID
                        rent_start_date: cartProduct.rent_start_date,
                        rent_end_date: cartProduct.rent_end_date,
                        quantity: cartProduct.quantity,
                        price: cartProduct.price,
                    })),
                }),
            });

            if (response.ok) {
                const data = await response.json();
                transactionId = data.data.id;
                console.log('Transaction created');
            } else {
                console.log('Failed to create transaction');
            }
        } catch (error) {
            console.error('Error creating transaction data:', error);
        }

        if (transactionId) {
            try {
                const response = await fetch('http://localhost:3000/api/v1/payment/', {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json',
                        'Authorization': `Bearer ${token}`,
                    },
                    body: JSON.stringify({
                        account_number: accountNumber,
                        transaction_id: transactionId,
                    }),
                });

                if (response.ok) {
                    const updatedData = await response.json();
                    setUser(updatedData.data);
                    console.log('Profile updated successfully!');
                } else {
                    console.log('Failed to process payment');
                }
            } catch (error) {
                console.error('Error processing payment data:', error);
            }
        }

        if (transactionId) {
            // Update the stock for each product in the cart
            for (const cartProduct of cartProducts) {
                const updatedStockQuantity = cartProduct.product.stock - cartProduct.quantity;
    
                try {
                    const response = await fetch(`http://localhost:3000/api/v1/product/${cartProduct.product.id}`, {
                        method: 'PUT',
                        headers: {
                            'Content-Type': 'application/json',
                            'Authorization': `Bearer ${token}`,
                        },
                        body: JSON.stringify({
                            stock: updatedStockQuantity,
                        }),
                    });
    
                    if (response.ok) {
                        const updatedProduct = await response.json();
                        console.log(`Stock updated for product ${cartProduct.product.id}`);
                    } else {
                        console.log('Failed to update product stock');
                    }
                } catch (error) {
                    console.error('Error updating product stock:', error);
                }
            }
        }

        try {
            const token = getTokenFromCookies();
            if (!token) return;
      
            const response = await fetch(`http://localhost:3000/api/v1/cart/${cart.data.id}` , {
              method: 'DELETE',
              headers: {
                'Content-Type': 'application/json',
                Authorization: `Bearer ${token}`,
              },
            });
      
            if (response.ok) {
              const data = await response.json();
            } else {
              console.log('Error fetching cart:', response);
            }
          } catch (error) {
            console.log('Error fetching cart:', error);
          }

          createCart(user?.data.id, token)

          router.push('/checkout/success');
    };

    return (
        <div>
            <NavigationBar />
            {alert.show && (
                <Alert
                    color="warning"
                    title="Missing Information"
                    endContent={
                        <Button color="warning" size="md" variant="flat" onClick={() => router.push('/profile')}>
                            Profile
                        </Button>
                    }
                    variant="faded"
                >
                    {alert.message}
                </Alert>
            )}
            <main className="flex flex-col items-center p-8 bg-gray-50 min-h-screen">
                <h1 className="text-3xl font-bold mb-8 text-gray-800">Checkout</h1>
                <div className='flex flex-col justify-between w-full max-w-7xl gap-6 lg:flex-row '>
                    <div className='lg:w-2/3'>
                        <Card className='p-3'>
                            <CardBody className='items-center'>
                                <h2 className='font-bold text-xl'>Order Summary</h2>
                                <p>{cart?.data.total_quantity} items | ${cart?.data.total_price}</p>
                                <Link href='/cart'>Edit Cart</Link>
                                <Divider className="my-4" />
                                {cartProducts.map((cartProduct) => (
                                    <div className='flex flex-row w-full justify-between mb-4' key={cartProduct.product.id}>
                                        <p>{cartProduct.quantity} x {cartProduct.product.name}</p>
                                        <p>{cartProduct.price}</p>
                                    </div>
                                ))}
                                <Divider className="mb-4" />
                                <div className='flex flex-row w-full justify-between mb-1'>
                                    <p>Subtotal</p>
                                    <p>${cart?.data.total_price}</p>
                                </div>
                                <div className='flex flex-row w-full justify-between'>
                                    <p>Shipping</p>
                                    <p>Free</p>
                                </div>
                                <Divider className="my-4" />
                                <div className='flex flex-row w-full justify-between'>
                                    <p className='text-lg'>Total Price</p>
                                    <p className='text-xl font-bold'>${cart?.data.total_price}</p>
                                </div>
                            </CardBody>
                        </Card>
                    </div>
                    <div className='lg:w-1/3'>
                        <Card className='p-3'>
                            <CardBody>
                                <h2 className='font-bold text-xl mb-4'>Customer</h2>
                                <h3 className='font-semibold text-md mb-1'>Name</h3>
                                <p className='text-md'>{user?.data?.name || '-'}</p>
                                <h3 className='font-semibold text-md mt-3 mb-1'>Email</h3>
                                <p className='text-md'>{user?.data?.email || '-'}</p>
                                <h3 className='font-semibold text-md mt-3 mb-1'>Address</h3>
                                <p className='text-md'>{user?.data?.address || '-'}</p>
                                <h3 className='font-semibold text-md mt-3 mb-1'>Phone Number</h3>
                                <p className='text-md'>{user?.data?.telephone || '-'}</p>
                                <Divider className="my-4" />
                                <h2 className='font-bold text-xl mb-4'>Payment</h2>
                                <Form onSubmit={handleTransaction}>
                                    <Input
                                        isRequired
                                        label="Account Number"
                                        type="text"
                                        variant="bordered"
                                        name="account_number"
                                        value={accountNumber}
                                        onChange={(e) => setAccountNumber(e.target.value)}
                                    />
                                    <Button color="primary" type="submit" className="text-md font-bold w-full mt-5" radius="sm" isDisabled={alert.show}>Pay</Button>
                                </Form>
                            </CardBody>
                        </Card>
                    </div>
                </div>
            </main>
        </div>
    );
}
