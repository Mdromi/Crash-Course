import { useState } from 'react'
import reactLogo from './assets/react.svg'
import './App.css'
import Product from './components/product'

export interface IProduct {
  id: number;
  title: string;
  price: number;
}

function App() {
  const [products, setProducts] = useState<IProduct[]>([
    {
      id: 1,
      title: 'Iphone',
      price: 300
    },
    {
      id: 2,
      title: 'MAC',
      price: 3300
    },
  ])

  function handleAddToCart(id: number) {
    console.log('Clicked - ', id);
    
  }

  return (
    <div className="App">
      {products.map(product => <Product product={product} key={product.id} handleAddToCart={handleAddToCart}/>)}
    </div>
  )
}

export default App
