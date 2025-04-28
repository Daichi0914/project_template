import { render, screen } from '@testing-library/react';
import Home from '../src/app/page';

describe('Home page', () => {
  it('renders the template text', () => {
    render(<Home />);
    expect(screen.getByText('テンプレート')).toBeInTheDocument();
  });
});
