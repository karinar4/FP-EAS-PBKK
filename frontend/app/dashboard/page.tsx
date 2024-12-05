'use client';

import { useEffect } from 'react';
import { useRouter } from 'next/navigation';

const DashboardPage = () => {
  const router = useRouter();

  useEffect(() => {
    // Check if the auth-token exists in cookies
    const token = document.cookie.split('; ').find(row => row.startsWith('auth-token='));

    if (!token) {
      router.push('/login');
    }
  }, [router]);

  return (
    <div>
      <h1>Welcome to your Dashboard!</h1>
    </div>
  );
};

export default DashboardPage;
