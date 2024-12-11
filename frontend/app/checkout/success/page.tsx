'use client';

import React from 'react';
import { Button, Card, CardBody } from "@nextui-org/react";
import { useRouter } from 'next/navigation';
import Image from 'next/image';
import NavigationBar from '@/app/components/NavigationBar';

export default function TransactionSuccess() {
    const router = useRouter();

    return (
        <div>
            <NavigationBar />
            <main className="flex flex-col items-center p-8 bg-gray-50 min-h-screen">
            <Card fullWidth>
              <CardBody>
                <div className='flex  flex-col h-full items-center justify-center px-12 py-28'>
                  <Image
                    src="/success.png"
                    alt="Cart"
                    width={300}
                    height={300}
                    className="mx-3"
                  />
                  <h2 className='font-bold text-3xl'>Transaction Successful!</h2>
                  <p className='font-normal text-xl'>Your rental transaction has been successfully completed. Thank you for choosing our service!</p>
                  <Button
                    className=" bg-yellow-500 hover:bg-yellow-600 text-white font-bold text-xl mt-5" onClick={() => router.push('/transaction')}>
                    All Transactions
                  </Button>
                </div>
              </CardBody>
            </Card>
            </main>
        </div>
    );
}


