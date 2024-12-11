"use client";

import { useState, useEffect } from "react";
import {
  Input,
  Card,
  CardBody,
  CardHeader,
  Button,
  Pagination,
  Select,
  SelectItem,
} from "@nextui-org/react";
import NavigationBar from "../components/NavigationBar";

const CatalogPage = () => {
  const [products, setProducts] = useState([]);
  const [categories, setCategories] = useState([]);
  const [currentPage, setCurrentPage] = useState(1);
  const [totalPages, setTotalPages] = useState(1);
  const [searchQuery, setSearchQuery] = useState("");
  const [sortOption, setSortOption] = useState("latest");
  const [selectedCategory, setSelectedCategory] = useState(null);

  const itemsPerPage = 16;

  useEffect(() => {
    const fetchProducts = async () => {
      try {
        const response = await fetch(`http://localhost:3000/api/v1/product/`);
        const data = await response.json();

        const productsWithImages = await Promise.all(
          data.data.products.map(async (product) => {
            const imageResponse = await fetch(
              `http://localhost:3000/api/v1/image/product/${product.id}`
            );
            const imageData = await imageResponse.json();

            return {
              ...product,
              thumbnail:
                imageData?.data?.images?.[0]?.url ||
                "https://via.placeholder.com/150",
            };
          })
        );

        setProducts(productsWithImages);
      } catch (error) {
        console.error("Failed to fetch products:", error);
      }
    };

    const fetchCategories = async () => {
      try {
        const response = await fetch(`http://localhost:3000/api/v1/category/`);
        const data = await response.json();
        setCategories(data.data.categories || []);
      } catch (error) {
        console.error("Failed to fetch categories:", error);
      }
    };

    fetchProducts();
    fetchCategories();
  }, []);

  const handlePageChange = (page) => setCurrentPage(page);

  const handleCategoryClick = (categoryId) => {
    setSelectedCategory(categoryId);
    setCurrentPage(1);
  };

  const sortedProducts = products.sort((a, b) => {
    if (sortOption === "priceLow") return a.price - b.price;
    if (sortOption === "priceHigh") return b.price - a.price;
    return 0;
  });

  const filteredProducts = sortedProducts.filter((product) => {
    const matchesCategory =
      !selectedCategory || String(product.category.id) === String(selectedCategory);
    const matchesSearch = product.name
      .toLowerCase()
      .includes(searchQuery.toLowerCase());
    return matchesCategory && matchesSearch;
  });

  useEffect(() => {
    setTotalPages(Math.ceil(filteredProducts.length / itemsPerPage));
    if (currentPage > Math.ceil(filteredProducts.length / itemsPerPage)) {
      setCurrentPage(1);
    }
  }, [filteredProducts]);

  const paginatedProducts = filteredProducts.slice(
    (currentPage - 1) * itemsPerPage,
    currentPage * itemsPerPage
  );

  return (
    <div className="min-h-screen bg-gray-200 text-gray-800">
      <NavigationBar />

      <main className="container mx-auto px-6 py-8 flex gap-8">
        <aside className="w-1/4 bg-gray-100 border-r border-gray-300 p-6 rounded-lg shadow-sm">
          <div className="mb-6">
            <h2 className="text-xl font-semibold mb-3 text-gray-700">Search</h2>
            <Input
              placeholder="Search products..."
              value={searchQuery}
              onChange={(e) => setSearchQuery(e.target.value)}
              radius="sm"
              size="lg"
              className="bg-white"
              variant="bordered"
              color="primary"
            />
          </div>

          <div>
            <h2 className="text-xl font-semibold mb-4 text-gray-700">Browse Categories</h2>
            <div className="space-y-2">
              {categories.map((category, index) => (
                <div key={category.id}>
                  <Button
                    variant={selectedCategory === category.id ? "solid" : "light"}
                    color={selectedCategory === category.id ? "primary" : "default"}
                    onClick={() => handleCategoryClick(category.id)}
                    className={`w-full text-left font-medium ${selectedCategory === category.id ? "text-white" : "text-gray-700"
                      } hover:bg-primary-50`}
                  >
                    {category.name}
                  </Button>
                  {index !== categories.length - 1 && (
                    <hr className="border-t border-gray-300 my-2" />
                  )}
                </div>
              ))}
            </div>
          </div>

          <div className="mt-8">
            <Button
              onClick={() => handleCategoryClick(null)}
              variant="ghost"
              color="danger"
              className="w-full font-semibold"
            >
              Reset Filters
            </Button>
          </div>
        </aside>

        <section className="flex-1">
          <div className="flex justify-between items-center mb-6">
            <p>
              Showing {(currentPage - 1) * itemsPerPage + 1}–
              {Math.min(currentPage * itemsPerPage, filteredProducts.length)} of{" "}
              {filteredProducts.length} results
            </p>
            <Select
              label="Sort by"
              placeholder="Select sort option"
              selectedKeys={[sortOption]}
              onChange={(e) => setSortOption(e.target.value)}
              className="w-48"
            >
              <SelectItem key="latest" value="latest">
                Latest
              </SelectItem>
              <SelectItem key="priceLow" value="priceLow">
                Price: Low to High
              </SelectItem>
              <SelectItem key="priceHigh" value="priceHigh">
                Price: High to Low
              </SelectItem>
            </Select>
          </div>

          <div className="grid grid-cols-2 md:grid-cols-4 gap-6">
            {paginatedProducts.map((product) => (
              <Card key={product.id} shadow="sm" isHoverable>
                <CardHeader className="p-0">
                  <img
                    src={product.thumbnail}
                    alt={product.name}
                    className="w-full h-40 object-cover rounded-t-lg"
                  />
                </CardHeader>
                <CardBody>
                  <h3 className="text-md font-semibold">{product.name}</h3>
                  <p className="text-primary font-bold">${product.price}</p>
                </CardBody>
              </Card>
            ))}
          </div>

          <div className="flex justify-center mt-8">
            <Pagination
              total={totalPages}
              initialPage={1}
              page={currentPage}
              onChange={handlePageChange}
              color="primary"
            />
          </div>
        </section>
      </main>
    </div>
  );
};

export default CatalogPage;
