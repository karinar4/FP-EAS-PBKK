'use client';

import React from 'react';
import { useRouter } from 'next/navigation';

const LogoutButton = () => {
  const router = useRouter();

  const handleLogout = (e: React.MouseEvent) => {
    console.log("Tes");
    e.preventDefault();
    e.stopPropagation(); 
    document.cookie = 'auth-token=; path=/; max-age=0';

    router.push('/');
  };

  return handleLogout;
};

export default LogoutButton;
