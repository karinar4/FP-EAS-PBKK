import RegisterForm from '../../components/RegisterForm';
import { Navbar, NavbarBrand } from "@nextui-org/react";

export default function LoginPage() {
  return (
    <div className="min-h-screen flex flex-col">
      <Navbar position="static">
        <NavbarBrand>
          <p className="font-bold text-2xl">ACME</p>
        </NavbarBrand>
      </Navbar>

      <div className="flex-grow flex items-center justify-center px-4 py-6">
        <div className="w-full max-w-[400px]">
          <RegisterForm />
        </div>
      </div>
    </div>
  );
}
