import { ReactNode } from 'react';
import {NextUIProvider} from "@nextui-org/react";

interface ProviderProps {
    children: ReactNode;
}

export default function Provider({ children }: ProviderProps) {
  return (
    <NextUIProvider>
      {children}
    </NextUIProvider>
  );
}