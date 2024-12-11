'use client';

import React, { useEffect, useState } from 'react';
import NavigationBar from '@/app/components/NavigationBar';
import { useRouter } from 'next/navigation';
import { Card, Table, TableHeader, TableBody, TableColumn, TableRow, TableCell, Tooltip, Textarea, Button, DateInput, CardBody } from "@nextui-org/react";
import {parseAbsolute, today, getLocalTimeZone} from "@internationalized/date";
import Image from 'next/image';

export const DeleteIcon = (props: any) => {
  return (
    <svg
      aria-hidden="true"
      fill="none"
      focusable="false"
      height="1em"
      role="presentation"
      viewBox="0 0 20 20"
      width="1em"
      {...props}
    >
      <path
        d="M17.5 4.98332C14.725 4.70832 11.9333 4.56665 9.15 4.56665C7.5 4.56665 5.85 4.64998 4.2 4.81665L2.5 4.98332"
        stroke="currentColor"
        strokeLinecap="round"
        strokeLinejoin="round"
        strokeWidth={1.5}
      />
      <path
        d="M7.08331 4.14169L7.26665 3.05002C7.39998 2.25835 7.49998 1.66669 8.90831 1.66669H11.0916C12.5 1.66669 12.6083 2.29169 12.7333 3.05835L12.9166 4.14169"
        stroke="currentColor"
        strokeLinecap="round"
        strokeLinejoin="round"
        strokeWidth={1.5}
      />
      <path
        d="M15.7084 7.61664L15.1667 16.0083C15.075 17.3166 15 18.3333 12.675 18.3333H7.32502C5.00002 18.3333 4.92502 17.3166 4.83335 16.0083L4.29169 7.61664"
        stroke="currentColor"
        strokeLinecap="round"
        strokeLinejoin="round"
        strokeWidth={1.5}
      />
      <path
        d="M8.60834 13.75H11.3833"
        stroke="currentColor"
        strokeLinecap="round"
        strokeLinejoin="round"
        strokeWidth={1.5}
      />
      <path
        d="M7.91669 10.4167H12.0834"
        stroke="currentColor"
        strokeLinecap="round"
        strokeLinejoin="round"
        strokeWidth={1.5}
      />
    </svg>
  );
};

export default function Cart() {
  const [user, setUser] = useState<{ data: { id: string; email: string; name: string } } | null>(null);
  // const [cartCreated, setCartCreated] = useState(false);
  const [cart, setCart] = useState<{ data: { id: string; total_price: number; total_quantity: number, user_id: string } } | null>(null);
  const [cartProducts, setCartProducts] = useState<any[]>([]);
  const [value, setValue] = React.useState(null);
  const router = useRouter();

  const getTokenFromCookies = () => {
    const cookies = document.cookie.split('; ');
    const tokenCookie = cookies.find((cookie) => cookie.startsWith('auth-token='));
    return tokenCookie ? tokenCookie.split('=')[1] : null;
  };

  useEffect(() => {
    const fetchUserDataAndCart = async () => {
      try {
        const token = getTokenFromCookies();

        if (!token) {
          console.log('No authentication token found in cookies.');
          router.push('/login');
          return;
        }

        // Fetch user data
        const userResponse = await fetch('http://localhost:3000/api/v1/auth/me', {
          method: 'GET',
          headers: {
            'Content-Type': 'application/json',
            'Authorization': `Bearer ${token}`,
          },
        });

        if (!userResponse.ok) {
          console.log('Failed to fetch user data');
          router.push('/login');
          return;
        }

        const userData = await userResponse.json();
        setUser(userData);

        const cartResponse = await fetch(`http://localhost:3000/api/v1/cart/${userData.data.id}`, {
          method: 'GET',
          headers: {
            'Content-Type': 'application/json',
            'Authorization': `Bearer ${token}`,
          },
        });

        if (cartResponse.ok) {
          const cartData = await cartResponse.json();
          if (cartData.data) {
            console.log('Cart exists:', cartData);
            // setCartCreated(true);
            setCart(cartData);
            console.log(cartData);
            await fetchCartProducts(cartData.data.id, token);
          } else {
            console.log('No cart found, creating cart...');
            await createCart(userData.data.id, token);
          }
        } else {
          console.log('Error fetching cart data:', cartResponse);
        }
      } catch (error) {
        console.log('Error fetching user or cart data:', error);
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
          setCart(data);
          // setCartCreated(true);
        } else {
          console.log('Error creating cart:', response);
        }
      } catch (error) {
        console.log('Error creating cart:', error);
      }
    };

    fetchUserDataAndCart();
  }, []);

  const fetchCartProducts = async (cartId: string, token: string) => {
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
        console.log('Products in cart:', data);
        setCartProducts(data.data.cart_products || []);
      } else {
        console.log('Error fetching cart products:', response);
      }
    } catch (error) {
      console.log('Error fetching cart products:', error);
    }
  };

  const updateCartProduct = async (cartId: string, ProductId: string, updatedData: any) => {
    const token = getTokenFromCookies();
    if (!token) return;

    try {
      const response = await fetch(`http://localhost:3000/api/v1/cart_product/${cartId}/${ProductId}`, {
        method: 'PUT',
        headers: {
          'Content-Type': 'application/json',
          'Authorization': `Bearer ${token}`,
        },
        body: JSON.stringify(updatedData),
      });

      if (response.ok) {
        console.log('Product updated:', updatedData);
        if (user){
          fetchCart(user.data.id, token);
        }
      } else {
        console.log('Error updating cart product:', response);
      }
    } catch (error) {
      console.log('Error updating cart product:', error);
    }
  };

  const fetchCart = async (userId: string, token: string) => {
    try {
      const response = await fetch(`http://localhost:3000/api/v1/cart/` + userId, {
        method: 'GET',
        headers: {
          'Content-Type': 'application/json',
          Authorization: `Bearer ${token}`,
        },
      });

      if (response.ok) {
        const data = await response.json();
        setCart(data);
      } else {
        console.log('Error fetching cart:', response);
      }
    } catch (error) {
      console.log('Error fetching cart:', error);
    }
  };

  const deleteCartProduct = async (cartId: string, productId:string) => {
    try {
      const token = getTokenFromCookies();
      if (!token) return;

      const response = await fetch(`http://localhost:3000/api/v1/cart_product/${cartId}/${productId}` , {
        method: 'DELETE',
        headers: {
          'Content-Type': 'application/json',
          Authorization: `Bearer ${token}`,
        },
      });

      if (response.ok) {
        const data = await response.json();
        await fetchCartProducts(cartId, token);
      } else {
        console.log('Error fetching cart:', response);
      }
    } catch (error) {
      console.log('Error fetching cart:', error);
    }
  };

  const handleQuantityChange = (cartId: string, productId: string, change: number) => {
    const updatedCartProducts = cartProducts.map((cartProduct) => {
      if (cartProduct.cart_id === cartId && cartProduct.product.id === productId) {
        const updatedQuantity = Math.max(Math.min(cartProduct.quantity + change, cartProduct.product.stock), 1);
        const updatedPrice = updatedQuantity * cartProduct.product.price;
  
        if (updatedQuantity !== cartProduct.quantity || updatedPrice !== cartProduct.price) {
          return { ...cartProduct, quantity: updatedQuantity, price: updatedPrice };
        }
  
        return cartProduct;
      }
      return cartProduct;
    });
  
    const hasChanges = updatedCartProducts.some(
      (cartProduct, index) => cartProduct !== cartProducts[index]
    );
  
    if (hasChanges) {
      setCartProducts(updatedCartProducts);
      const updatedData = {
        quantity: updatedCartProducts.find(cartProduct => cartProduct.cart_id === cartId && cartProduct.product.id === productId)?.quantity,
        price: updatedCartProducts.find(cartProduct => cartProduct.cart_id === cartId && cartProduct.product.id === productId)?.price,
      };
      console.log(updatedData);
      updateCartProduct(cartId, productId, updatedData);
    }
  };

  const handleDateChange = (cartId: any, productId: any, startDate: any, endDate: any) => {
    let start, end
    if (typeof startDate == "string"){
      start = parseAbsolute(startDate, "Asia/Bangkok");
    } else if (typeof startDate == "object") {
      start = startDate
    }

    if (typeof endDate == "string"){
      end = parseAbsolute(endDate, "Asia/Bangkok");
    } else if (typeof endDate == "object") {
      end = endDate
    }
      
    const duration = Math.ceil((end.toDate().getTime() - start.toDate().getTime()) / (1000 * 60 * 60 * 24)); // Duration in days
    console.log("Duration in days:", duration);

    const displayDuration = isNaN(duration) ? 'Invalid duration' : duration;
      
    const updatedCartProducts = cartProducts.map((cartProduct) => {
      if (cartProduct.product.id === productId) {
        const updatedPrice = (duration > 0 && !isNaN(duration))? duration * cartProduct.quantity * cartProduct.product.price : cartProduct.price;
        console.log("Duration:", duration);
        console.log("Updated Price:", updatedPrice);
        return {
          ...cartProduct,
          rent_start_date: startDate.toString().split('[')[0],
          rent_end_date: endDate.toString().split('[')[0], 
          price: updatedPrice
        };
      }
      return cartProduct;
    });

    const hasChanges = updatedCartProducts.some(
      (cartProduct, index) => cartProduct !== cartProducts[index]
    );

    if(hasChanges){
      setCartProducts(updatedCartProducts);
      const updatedData = {
        rent_start_date: updatedCartProducts.find(cartProduct => cartProduct.cart_id === cartId && cartProduct.product.id === productId)?.rent_start_date,
        rent_end_date: updatedCartProducts.find(cartProduct => cartProduct.cart_id === cartId && cartProduct.product.id === productId)?.rent_end_date,
        price: updatedCartProducts.find(cartProduct => cartProduct.cart_id === cartId && cartProduct.product.id === productId)?.price,
      };
      updateCartProduct(cartId, productId, updatedData);
    }
  };
  
  return (
    <div>
      <NavigationBar />
      <main className="flex flex-col items-center p-8 bg-gray-50 min-h-screen">
        {cartProducts.length > 0 ? (
        <><h1 className="text-3xl font-bold mb-8 text-gray-800">Cart</h1>
        <Table aria-label="Cart table"
            classNames={{
              th: ["bg-transparent", "text-default-500", "border-b", "border-divider"],
              tbody: ["border-b"]
            }}
            bottomContent={<div className="py-2 px-2 flex flex-col justify-end items-end">
              <div className='pb-4'>
                <p className='font-medium text-lg'>Total Quantity: {cart?.data.total_quantity}</p>
                <p className='font-medium text-lg'>Total Price: ${cart?.data.total_price}</p>
              </div>
              <Button radius='sm' className="bg-yellow-500 hover:bg-yellow-600 text-white text-md font-bold" onClick={() => router.push('/checkout')}>Checkout</Button>
            </div>}>
            <TableHeader>
              <TableColumn>NAME</TableColumn>
              <TableColumn>PRICE</TableColumn>
              <TableColumn>RENT START DATE</TableColumn>
              <TableColumn>RENT END DATE</TableColumn>
              <TableColumn>QUANTITY</TableColumn>
              <TableColumn>TOTAL PRICE</TableColumn>
              <TableColumn> </TableColumn>
            </TableHeader>
            <TableBody>
              {cartProducts.map((cartProduct) => (
                <TableRow key={cartProduct.product.id}>
                  <TableCell className='py-4'>{cartProduct.product.name}</TableCell>
                  <TableCell className='py-4'>{cartProduct.product.price}</TableCell>
                  <TableCell className='py-4'>
                    <DateInput
                      aria-label='rent start date'
                      hideTimeZone
                      defaultValue={parseAbsolute(cartProduct.rent_start_date, "Asia/Bangkok")}
                      classNames={{
                        inputWrapper: "bg-transparent shadow-none"
                      }}
                      minValue={today(getLocalTimeZone()).add({ days: 1 })}
                      value={parseAbsolute(cartProduct.rent_start_date, 'Asia/Bangkok')}
                      onChange={(date) => {
                        console.log("Selected date:", date);
                        handleDateChange(
                          cartProduct.cart_id, cartProduct.product.id, date, cartProduct.rent_end_date);
                      } } />
                  </TableCell>
                  <TableCell className='py-4'>
                    <DateInput
                      aria-label='rent end date'
                      hideTimeZone
                      defaultValue={parseAbsolute(cartProduct.rent_end_date, "Asia/Bangkok")}
                      classNames={{
                        inputWrapper: "bg-transparent shadow-none"
                      }}
                      minValue={parseAbsolute(cartProduct.rent_start_date, "Asia/Bangkok")}
                      value={parseAbsolute(cartProduct.rent_end_date, 'Asia/Bangkok')}
                      onChange={(date) => {
                        console.log("Selected date:", date);
                        handleDateChange(
                          cartProduct.cart_id, cartProduct.product.id, cartProduct.rent_start_date, date);
                      } } />
                  </TableCell>
                  <TableCell className='py-4'>
                    <div className='flex flex-row items-center justify-center rounded-full border border-gray-300 w-24 h-8'>
                      <button className='text-xl pl-4' onClick={() => handleQuantityChange(cartProduct.cart_id, cartProduct.product.id, -1)}>
                        <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24">
                          <path d="M19 13H5C4.44772 13 4 13.4477 4 14C4 14.5523 4.44772 15 5 15H19C19.5523 15 20 14.5523 20 14C20 13.4477 19.5523 13 19 13Z" />
                        </svg>
                      </button>
                      <p className='px-4'>{cartProduct.quantity}</p>
                      <button className='text-xl pr-4' onClick={() => handleQuantityChange(cartProduct.cart_id, cartProduct.product.id, 1)}>
                        <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24">
                          <path d="M19 11H13V5C13 4.44772 12.5523 4 12 4C11.4477 4 11 4.44772 11 5V11H5C4.44772 11 4 11.4477 4 12C4 12.5523 4.44772 13 5 13H11V19C11 19.5523 11.4477 20 12 20C12.5523 20 13 19.5523 13 19V13H19C19.5523 13 20 12.5523 20 12C20 11.4477 19.5523 11 19 11Z" />
                        </svg>
                      </button>
                    </div>
                  </TableCell>
                  <TableCell className='py-4'>{cartProduct.price}</TableCell>
                  <TableCell className='py-4'>
                    <Tooltip color="danger" content="Delete">
                      <span className="text-lg text-danger cursor-pointer active:opacity-50" onClick={() => deleteCartProduct(cartProduct.cart_id, cartProduct.product.id)}>
                        <DeleteIcon />
                      </span>
                    </Tooltip>
                  </TableCell>
                </TableRow>
              ))}
            </TableBody>
          </Table></>
          ) : (
            <Card fullWidth>
              <CardBody>
                <div className='flex  flex-col h-full items-center justify-center p-12'>
                  <Image
                    src="/empty_cart.png"
                    alt="Cart"
                    width={450}
                    height={450}
                    className="mx-3"
                  />
                  <h2 className='font-bold text-3xl'>Your cart is empty!</h2>
                  <p className='font-normal text-xl'>Browse products and add them to your cart.</p>
                  <Button
                    className=" bg-yellow-500 hover:bg-yellow-600 text-white font-bold text-xl mt-5">
                    Book Now!
                  </Button>
                </div>
              </CardBody>
            </Card>
          )}
      </main>
    </div>
  );
}