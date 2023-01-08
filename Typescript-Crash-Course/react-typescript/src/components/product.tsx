import { IProduct } from "../App";

interface ProductProps {
  product: IProduct;
  handleAddToCart(id: number): void;
}

const Product = ({ product, handleAddToCart }: ProductProps) => {
  return (
    <div>
      <h2>{product.title}</h2>
      <p>{product.price}</p>
      <button onClick={() => handleAddToCart(product.id)}>Add To Cart</button>
    </div>
  );
};

export default Product;
