"use client";

import NavigationBar from "@/app/components/NavigationBar";
import { useState, useEffect } from "react";
import { useRouter, useParams } from "next/navigation";
import { Alert } from "@nextui-org/react";

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
  const router = useRouter();
  const { id } = useParams();

  const [quantity, setQuantity] = useState<number>(1);
  const [currentImageIndex, setCurrentImageIndex] = useState<number>(0);
  const [product, setProduct] = useState<Product | null>(null);
  const [images, setImages] = useState<Image[]>([]);
  const [user, setUser] = useState<{ data: { id: string; email: string; name: string } } | null>(null);
  const [rentStartDate, setRentStartDate] = useState<string>("");
  const [rentEndDate, setRentEndDate] = useState<string>("");
  const [cart, setCart] = useState<{ data: { id: string; total_price: number; total_quantity: number, user_id: string } } | null>(null);
  const [alert, setAlert] = useState({ show: false, message: '' });
  const [isVisible, setIsVisible] = useState(true);
  
  useEffect(() => {
    if (!id) {
      console.log(id);
      return;
    }

    const fetchData = async () => {
      try {
        // Fetch product details
        const productResponse = await fetch(
          `http://localhost:3000/api/v1/product/${id}`
        );
        if (!productResponse.ok) {
          throw new Error("Failed to fetch product details");
        }
        const productData = await productResponse.json();
        if (productData.status) {
          setProduct(productData.data);
          console.log("Set Product");
        } else {
          console.error(productData.error);
        }

        // Fetch product images
        const imagesResponse = await fetch(
          `http://localhost:3000/api/v1/image/product/${id}`
        );
        if (!imagesResponse.ok) {
          throw new Error("Failed to fetch product images");
        }
        const imagesData = await imagesResponse.json();
        if (imagesData.status) {
          setImages(imagesData.data.images);
        } else {
          console.error(imagesData.error);
        }
      } catch (error) {
        console.error("Error fetching data:", error);
      }
    };

    fetchData();
    fetchUser();
  }, [id]);

  useEffect(() => {
    if (user) {
      fetchCart(user.data.id);
      console.log("fetch cart");
    }
  }, [user]);

  const handleQuantityChange = (type: QuantityChangeType): void => {
    setQuantity((prev) => (type === "increase" ? prev + 1 : Math.max(prev - 1, 1)));
  };

  const getTokenFromCookies = () => {
    const cookies = document.cookie.split('; ');
    const tokenCookie = cookies.find((cookie) => cookie.startsWith('auth-token='));
    return tokenCookie ? tokenCookie.split('=')[1] : null;
  };

  const fetchCart = async (userId: string) => {
    try {
      const token = getTokenFromCookies();

        if (!token) {
          console.log('No authentication token found in cookies.');
          router.push('/login');
          return;
        }

      const response = await fetch(`http://localhost:3000/api/v1/cart/` + userId, {
        method: 'GET',
        headers: {
          'Content-Type': 'application/json',
          Authorization: `Bearer ${token}`,
        },
      });

      if (response.ok) {
        const data = await response.json();
        setCart(data);
        console.log("Cart data", data);
        return data;
      } else {
        console.log('Error fetching cart:', response);
      }
    } catch (error) {
      console.log('Error fetching cart:', error);
    }
  };

  const fetchUser = async () => {
    try {
      const token = getTokenFromCookies();

      if (!token) {
        console.log('No authentication token found in cookies.');
        router.push('/login');
        return;
      }

      // Fetch user data
      const userResponse = await fetch('http://localhost:3000/api/v1/auth/me', {
        method: 'GET',
        headers: {
          'Content-Type': 'application/json',
          'Authorization': `Bearer ${token}`,
        },
      });

      if (!userResponse.ok) {
        console.log('Failed to fetch user data');
        router.push('/login');
        return;
      }

      const userData = await userResponse.json();
      setUser(userData);
    } catch (error) {
      console.log('Error fetching user or cart data:', error);
    } 
  };

  const handleImageChange = (direction: "prev" | "next"): void => {
    setCurrentImageIndex((prev) => {
      if (direction === "prev") {
        return prev === 0 ? images.length - 1 : prev - 1;
      }
      return prev === images.length - 1 ? 0 : prev + 1;
    });
  };

  const handleAddToCart = async (): Promise<void> => {
    if (!product) return;

    const formattedRentStartDate = rentStartDate
    ? new Date(rentStartDate).toISOString()
    : null;
  const formattedRentEndDate = rentEndDate
    ? new Date(rentEndDate).toISOString()
    : null;

    if (!formattedRentStartDate || !formattedRentEndDate) {
      console.error("Both rent start and end dates must be selected.");
      return;
    }
      const startDate = new Date(formattedRentStartDate);
    const endDate = new Date(formattedRentEndDate);
    const durationInDays = Math.ceil(
      (endDate.getTime() - startDate.getTime()) / (1000 * 60 * 60 * 24)
    );

    if (durationInDays <= 0) {
      console.error("Invalid rental period.");
      return;
    }

    // Hitung total harga
    const totalPrice = durationInDays * product.price * quantity;

    const cartProductData = {
      product_id: product.id,
      cart_id: cart?.data.id,
      quantity: quantity,
      price: totalPrice,
      rent_start_date: formattedRentStartDate,
      rent_end_date: formattedRentEndDate,
    };

    try {
      const token = getTokenFromCookies();

      if (!token) {
        console.log('No authentication token found in cookies.');
        router.push('/login');
        return;
      }

      console.log(cartProductData);
      const response = await fetch("http://localhost:3000/api/v1/cart_product/", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          'Authorization': `Bearer ${token}`,
        },
        body: JSON.stringify(cartProductData),
      });

      if (response.ok) {
        const data = await response.json();
        if (data.status) {
          console.log("Product added to cart:", data);
          // Optionally, redirect or update UI to reflect cart update
          setAlert({
            show: true,
            message: "Product added to cart",
        });
        } else {
          console.error("Failed to add product to cart:", data.error);
        }
      } else {
        console.error("Failed to create cart product");
      }
    } catch (error) {
      console.error("Error adding to cart:", error);
    }
  };

  if (!product) {
    return <div>Loading...</div>;
  }

  return (
    <>
      <NavigationBar />
    <div className="min-h-screen bg-gray-50">
      {alert.show && (
        <Alert
          color="success"
          title="Product added to cart"
          variant="faded"
          isVisible={isVisible}
          onClose={() => setIsVisible(false)}
        >
          {alert.message}
        </Alert>
      )}
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
                  type="datetime-local"
                  className="border px-3 py-2 rounded w-full"
                  value={rentStartDate}
                  onChange={(e) => setRentStartDate(e.target.value)}
                />
              </div>

              <div className="space-y-2">
                <label>Select rental end date:</label>
                <input
                  type="datetime-local"
                  className="border px-3 py-2 rounded w-full"
                  value={rentEndDate}
                  onChange={(e) => setRentEndDate(e.target.value)}
                />
              </div>
            </div>

            <button
              className="w-full bg-blue-500 text-white py-3 rounded-lg shadow hover:bg-blue-600"
              onClick={handleAddToCart}
            >
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
    </>
  );
}
