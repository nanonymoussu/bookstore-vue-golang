import axios from "axios";

const API_URL = "/api";

export default {
  async getBooks(page = 1, pageSize = 10) {
    try {
      const response = await axios.get(`${API_URL}/books`, {
        params: { page, pageSize },
      });
      return response.data;
    } catch (error) {
      console.error("Error fetching books:", error);
      throw error;
    }
  },

  async getBook(id) {
    try {
      const response = await axios.get(`${API_URL}/books/${id}`);
      return response.data;
    } catch (error) {
      console.error(`Error fetching book ${id}:`, error);
      throw error;
    }
  },

  async createBook(book) {
    try {
      const response = await axios.post(`${API_URL}/books`, book);
      return response.data;
    } catch (error) {
      console.error("Error creating book:", error);
      throw error;
    }
  },

  async updateBook(id, book) {
    try {
      const response = await axios.put(`${API_URL}/books/${id}`, book);
      return response.data;
    } catch (error) {
      console.error(`Error updating book ${id}:`, error);
      throw error;
    }
  },

  async deleteBook(id) {
    try {
      await axios.delete(`${API_URL}/books/${id}`);
      return true;
    } catch (error) {
      console.error(`Error deleting book ${id}:`, error);
      throw error;
    }
  },
};
