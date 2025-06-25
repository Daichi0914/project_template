'use client';

import React from 'react';
import { Provider } from 'jotai';

export default function AppProvider({ children }: { children: React.ReactNode }) {
  return (
    <Provider>
      {children}
    </Provider>
  );
}
