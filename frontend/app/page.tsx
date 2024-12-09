import {NavbarContent, NavbarItem, Link, Button} from "@nextui-org/react";
import NavigationBar from "./components/NavigationBar"
import RegisterForm from "./components/RegisterForm";

export default function Home() {
  return (
    <div>
      <NavigationBar
        customButtons={
          <>
            <NavbarContent justify="end">
              <NavbarItem className="hidden lg:flex">
                <Link href="/login">Login</Link>
              </NavbarItem>
              <NavbarItem>
                <Button as={Link} color="primary" href="/register" variant="flat">
                  Sign Up
                </Button>
              </NavbarItem>
            </NavbarContent>
          </>
        }
      />
    </div>
  );
}
