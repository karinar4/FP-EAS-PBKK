'use client';

import React, { useEffect, useState } from 'react';
import NavigationBar from '@/app/components/NavigationBar';
import { Card, CardBody, Chip, Divider } from '@nextui-org/react';
import { useRouter } from 'next/navigation';

export default function Transaction() {
  const [user, setUser] = useState<{ id: string; email: string; name: string; telephone: string; address: string } | null>(null);
  const [transactions, setTransactions] = useState<any[]>([]);
  const [transactionProducts, setTransactionProducts] = useState<Map<string, any[]>>(new Map());
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
      if (!response.ok) throw new Error('Failed to fetch user data');
      const data = await response.json();
      setUser(data);
      return data;
    } catch (error) {
      console.error('Error fetching user data:', error);
    }
  };

  const fetchTransactions = async (userId: string) => {
    try {
      const response = await fetch(`http://localhost:3000/api/v1/transaction/user/${userId}`, {
        method: 'GET',
        headers: {
          'Content-Type': 'application/json',
          Authorization: `Bearer ${token}`,
        },
      });
      if (!response.ok) throw new Error('Failed to fetch transaction data');
      const data = await response.json();
      setTransactions(data.data.transactions);
      return data.data.transactions;
    } catch (error) {
      console.error('Error fetching transactions:', error);
    }
  };

  const fetchTransactionProducts = async (transactionId: string) => {
    try {
      const response = await fetch(`http://localhost:3000/api/v1/product_transaction/${transactionId}`, {
        method: 'GET',
        headers: {
          'Content-Type': 'application/json',
          Authorization: `Bearer ${token}`,
        },
      });
      if (!response.ok) throw new Error('Failed to fetch transaction products');
      const data = await response.json();
      return data.data.product_transactions || [];
    } catch (error) {
      console.error('Error fetching transaction products:', error);
      return [];
    }
  };

  useEffect(() => {
    const fetchData = async () => {
      const user_data = await fetchUserData();
      if (user_data?.data?.id) {
        const transaction_data = await fetchTransactions(user_data.data.id);
        const productsMap = new Map();
        for (const transaction of transaction_data) {
          const products = await fetchTransactionProducts(transaction.id);
          productsMap.set(transaction.id, products);
        }
        setTransactionProducts(productsMap);
      }
    };

    fetchData();
  }, []);

  return (
    <div>
      <NavigationBar />
      <main className="flex flex-col items-center p-8 bg-gray-50 min-h-screen gap-5">
        <h1 className="text-3xl font-bold mb-8 text-gray-800">Transactions</h1>
        {transactions.length > 0 ? (
          transactions.map((transaction) => {
            const products = transactionProducts.get(transaction.id) || [];
            return (
              <Card className="w-full max-w-5xl p-3 shadow-md" key={transaction.id}>
                <CardBody>
                  <div className="flex flex-row justify-between items-center text-yellow-500 font-semibold">
                    {
                        (() => {
                            const transactionDate = new Date(transaction.transaction_date);
                            const day = String(transactionDate.getDate()).padStart(2, '0');
                            const month = String(transactionDate.getMonth() + 1).padStart(2, '0');
                            const year = transactionDate.getFullYear();
                            const hours = String(transactionDate.getHours()).padStart(2, '0');
                            const minutes = String(transactionDate.getMinutes()).padStart(2, '0');
                            const seconds = String(transactionDate.getSeconds()).padStart(2, '0');
                            const formattedDate = `${day}-${month}-${year} ${hours}:${minutes}:${seconds}`;
                    
                            return formattedDate;
                        })()
                    }
                    <Chip color="warning" variant="solid" className="text-md p-4">
                      {transaction.status === "completed" ? "Completed" : transaction.status === "pending" ? "Pending" : transaction.status === "in_progress" ? "In Progress" : "Failed"}
                    </Chip>
                  </div>
                  <Divider className="my-4" />
                  {products.length > 0 ? (
                    products.map((product, index) => (
                      <div className="flex flex-row w-full justify-between mb-4 items-center" key={`${product.product.id}-${transaction.id}-${index}`}>
                        <div className="flex-col">
                          <p>
                            {product.quantity} x {product?.product?.name || '-'}
                          </p>
                          <p className="text-gray-500">
                            {
                                (() => {
                                const rentStartDate = new Date(product.rent_start_date);
                                const dayStart = String(rentStartDate.getDate()).padStart(2, '0');
                                const monthStart = String(rentStartDate.getMonth() + 1).padStart(2, '0'); 
                                const yearStart = rentStartDate.getFullYear();
                                const hoursStart = String(rentStartDate.getHours()).padStart(2, '0');
                                const minutesStart = String(rentStartDate.getMinutes()).padStart(2, '0');
                                const secondsStart = String(rentStartDate.getSeconds()).padStart(2, '0');
                                const formattedStartDate = `${dayStart}-${monthStart}-${yearStart} ${hoursStart}:${minutesStart}:${secondsStart}`;

                                const rentEndDate = new Date(product.rent_end_date);
                                const dayEnd = String(rentEndDate.getDate()).padStart(2, '0');
                                const monthEnd = String(rentEndDate.getMonth() + 1).padStart(2, '0');
                                const yearEnd = rentEndDate.getFullYear();
                                const hoursEnd = String(rentEndDate.getHours()).padStart(2, '0');
                                const minutesEnd = String(rentEndDate.getMinutes()).padStart(2, '0');
                                const secondsEnd = String(rentEndDate.getSeconds()).padStart(2, '0');
                                const formattedEndDate = `${dayEnd}-${monthEnd}-${yearEnd} ${hoursEnd}:${minutesEnd}:${secondsEnd}`;

                                return `${formattedStartDate} - ${formattedEndDate}`;
                                })()
                            }
                            </p>
                        </div>
                        <p className="text-yellow-500">{product.price}</p>
                      </div>
                    ))
                  ) : (
                    <p>No products found.</p>
                  )}
                  <Divider className="mb-4" />
                  <div className="flex flex-row w-full justify-end items-center">
                    <p className="w-20 text-lg">Total:</p>
                    <p className="text-xl font-bold text-yellow-500">{transaction.total_price}</p>
                  </div>
                </CardBody>
              </Card>
            );
          })
        ) : (
          <p>No transactions found.</p>
        )}
      </main>
    </div>
  );
}
