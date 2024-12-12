"use client";

import NavigationBar from "@/app/components/NavigationBar";
import { useState, useEffect } from "react";
import { useRouter } from "next/router";

interface Product {
  id: string;
  name: string;
  description: string;
  price: number;
  stock: number;
  category: { id: string; name: string };
  brand: { id: string; name: string };
}

interface Image {
  id: string;
  product_id: string;
  url: string;
}

type QuantityChangeType = "increase" | "decrease";

export default function ProductPage(): JSX.Element {
  // const router = useRouter();
  // const { id: productId } = router.query;

  const [quantity, setQuantity] = useState<number>(1);
  const [currentImageIndex, setCurrentImageIndex] = useState<number>(0);
  const [product, setProduct] = useState<Product | null>(null);
  const [images, setImages] = useState<Image[]>([]);
  
  const productId = "027b9afd-8eb5-48dc-a797-52205a844efe"; // Example product ID

  useEffect(() => {
    if (!productId) return;

    fetch(`http://localhost:3000/api/v1/product/${productId}`)
      .then((response) => response.json())
      .then((data) => {
        if (data.status) {
          setProduct(data.data);
        } else {
          console.error(data.error);
        }
      })
      .catch((err) => console.error("Failed to fetch product details:", err));

    fetch(`http://localhost:3000/api/v1/image/product/${productId}`)
      .then((response) => response.json())
      .then((data) => {
        if (data.status) {
          setImages(data.data.images);
        } else {
          console.error(data.error);
        }
      })
      .catch((err) => console.error("Failed to fetch product images:", err));
  }, [productId]);

  const handleQuantityChange = (type: QuantityChangeType): void => {
    setQuantity((prev) => (type === "increase" ? prev + 1 : Math.max(prev - 1, 1)));
  };

  const handleImageChange = (direction: "prev" | "next"): void => {
    setCurrentImageIndex((prev) => {
      if (direction === "prev") {
        return prev === 0 ? images.length - 1 : prev - 1;
      }
      return prev === images.length - 1 ? 0 : prev + 1;
    });
  };

  if (!product) {
    return <div>Loading...</div>;
  }

  return (
    <div className="min-h-screen bg-gray-50">
      <NavigationBar />
      <main className="container mx-auto p-6">
        <div className="flex flex-col md:flex-row gap-6">
          <div className="flex-1">
            <div className="relative">
              {images.length > 0 ? (
                <img
                  src={images[currentImageIndex].url}
                  alt={product.name}
                  className="w-full rounded-lg shadow-md"
                />
              ) : (
                <div className="w-full h-64 bg-gray-200 flex items-center justify-center rounded-lg">
                  No Image Available
                </div>
              )}
              <button
                className="absolute left-2 top-1/2 transform -translate-y-1/2 bg-gray-200 px-2 py-1 rounded"
                onClick={() => handleImageChange("prev")}
                disabled={images.length === 0}
              >
                ◀
              </button>
              <button
                className="absolute right-2 top-1/2 transform -translate-y-1/2 bg-gray-200 px-2 py-1 rounded"
                onClick={() => handleImageChange("next")}
                disabled={images.length === 0}
              >
                ▶
              </button>
            </div>
          </div>

          <div className="flex-1 space-y-4">
            <h1 className="text-2xl font-bold">{product.name}</h1>
            <p className="text-xl text-gray-600">${product.price.toFixed(2)}</p>
            <div className="flex items-center space-x-4">
              <span>Quantity:</span>
              <div className="flex items-center space-x-2">
                <button
                  className="px-2 py-1 bg-gray-200 rounded"
                  onClick={() => handleQuantityChange("decrease")}
                >
                  -
                </button>
                <span>{quantity}</span>
                <button
                  className="px-2 py-1 bg-gray-200 rounded"
                  onClick={() => handleQuantityChange("increase")}
                >
                  +
                </button>
              </div>
            </div>
            <p className="text-gray-500">Available: {product.stock}</p>
            <div className="border p-4 rounded-lg space-y-4">
              <h2 className="text-lg font-bold">Rental Period</h2>
              <div className="space-y-2">
                <label>Select rental start date:</label>
                <input
                  type="date"
                  className="border px-3 py-2 rounded w-full"
                />
              </div>

              <div className="space-y-2">
                <label>Select rental end date:</label>
                <input
                  type="date"
                  className="border px-3 py-2 rounded w-full"
                />
              </div>
            </div>

            <button className="w-full bg-blue-500 text-white py-3 rounded-lg shadow hover:bg-blue-600">
              ADD TO CART
            </button>
          </div>
        </div>

        <section className="mt-10">
          <h2 className="text-xl font-bold">Product Description</h2>
          <p className="mt-2 text-gray-600">{product.description}</p>
        </section>
      </main>
    </div>
  );
}