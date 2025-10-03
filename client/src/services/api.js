// api.js
import axios from 'axios';

const API_URL = 'http://localhost:8080';

const api = axios.create({
  baseURL: API_URL,
  headers: {
    'Content-Type': 'application/json',
    'Alif':"hello"
  },
});

// CRUD operations

export const getAllProducts = async () => {
  try {
    const response = await api.get('/products');
    return response.data;
  } catch (error) {
    throw error;
  }
};

export const createProduct = async (productData) => {
  if(productData?.price) {
    productData.price=Number(productData.price);
  }
  const payload = {
    title: productData.title,
    description: productData.description,
    price: productData.price,
    image_url: productData.image_url || ' ', 
  };

  try {
    const response = await api.post('/products', payload);
    return response.data;
  } catch (error) {
    console.error('Error creating product:', error);
    throw error;
  }
};

export const updateProduct = async (productId, productData) => {
  try {
    const response = await api.put(`/products/${productId}`, productData);
    return response.data;
  } catch (error) {
    throw error;
  }
};

export const deleteProduct = async (productId) => {
  try {
    const response = await api.delete(`/products/${productId}`);
    return response.data;
  } catch (error) {
    throw error;
  }
};

export default api;
