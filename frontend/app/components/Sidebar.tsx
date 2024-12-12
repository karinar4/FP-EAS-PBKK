"use client";

import { Button } from '@nextui-org/react';
import React, { useState } from 'react';
import { useRouter } from 'next/navigation';

export default function Sidebar(props) {
  const router = useRouter();

  return (
    <aside className="w-1/4 bg-white border-r border-gray-300 p-6 shadow-sm" {...props}>
      <div>
        <div className="space-y-2">
          <Button onClick={() => router.push('/admin/product')} className="w-full bg-white border border-yellow-400 shadow-sm">
            Product
          </Button>
          <Button onClick={() => router.push('/admin/image')} className="w-full bg-white border border-yellow-400 shadow-sm">
            Image
          </Button>
          <Button onClick={() => router.push('/admin/brand')} className="w-full bg-white border border-yellow-400 shadow-sm">
            Brand
          </Button>
          <Button onClick={() => router.push('/admin/category')} className="w-full bg-white border border-yellow-400 shadow-sm">
            Category
          </Button>
          <Button onClick={() => router.push('/admin/transaction')} className="w-full bg-white border border-yellow-400 shadow-sm">
            Transaction
          </Button>
          <Button onClick={() => router.push('/admin/payment')} className="w-full bg-white border border-yellow-400 shadow-sm">
            Payment
          </Button>
        </div>
      </div>
    </aside>
  );
}
