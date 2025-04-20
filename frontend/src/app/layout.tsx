import './globals.css';
import AppProvider from '../provider/AppProvider';

export const metadata = {
  title: 'Task Timer',
  description: 'Time tracking app',
};

export default function RootLayout({ children }: { children: React.ReactNode }) {
  return (
    <html lang="en">
      <body>
        <AppProvider>
          {children}
        </AppProvider>
      </body>
    </html>
  );
}
