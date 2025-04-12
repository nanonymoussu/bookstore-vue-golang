<template>
  <div>
    <div
      class="pb-5 border-b border-gray-200 sm:flex sm:items-center sm:justify-between"
    >
      <h1 class="text-2xl font-bold text-gray-900">Add New Book</h1>
      <div class="mt-3 sm:mt-0">
        <router-link :to="{ name: 'home' }" class="btn btn-secondary">
          Cancel
        </router-link>
      </div>
    </div>

    <div class="mt-6">
      <div class="bg-white shadow px-4 py-5 sm:rounded-lg sm:p-6">
        <book-form
          :initial-data="bookData"
          submit-button-text="Create Book"
          @submit="createBook"
          @cancel="$router.push({ name: 'home' })"
        />
      </div>
    </div>
  </div>
</template>

<script>
import BookForm from "@/components/BookForm.vue";
import bookService from "@/services/bookService";

export default {
  name: "BookCreateView",
  components: {
    BookForm,
  },
  data() {
    return {
      bookData: {
        title: "",
        author: "",
        isbn: "",
        price: 0,
        stock: 0,
      },
    };
  },
  methods: {
    async createBook(book) {
      try {
        const createdBook = await bookService.createBook(book);
        this.$router.push({
          name: "book-details",
          params: { id: createdBook.id },
        });
      } catch (error) {
        console.error("Failed to create book:", error);
        // Handle error (show notification)
        alert("Failed to create book. Please try again.");
      }
    },
  },
};
</script>
