import '@/styles/globals.css';
import AppProvider from '@/provider/AppProvider';
import { ReactNode } from 'react';

export const metadata = {
  title: 'Task Timer',
  description: 'Time tracking app',
};

export default function RootLayout({ children }: { children: ReactNode }) {
  return (
    <html lang="ja">
      <body>
        <AppProvider>
          {children}
        </AppProvider>
      </body>
    </html>
  );
}
