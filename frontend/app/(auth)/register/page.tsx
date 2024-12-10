import NavigationBar from '@/app/components/NavigationBar';
import RegisterForm from '@/app/components/RegisterForm';

export default function LoginPage() {
  return (
    <div className="min-h-screen flex flex-col">
      <NavigationBar />

      <div className="flex-grow flex items-center justify-center px-4 py-6">
        <div className="w-full max-w-[400px]">
          <RegisterForm />
        </div>
      </div>
    </div>
  );
}
