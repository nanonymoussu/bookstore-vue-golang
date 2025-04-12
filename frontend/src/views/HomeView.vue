<template>
  <div>
    <div
      class="pb-5 border-b border-gray-200 sm:flex sm:items-center sm:justify-between"
    >
      <h1 class="text-2xl font-bold text-gray-900">Books</h1>
      <div class="mt-3 sm:mt-0 sm:ml-4">
        <router-link to="/books/create" class="btn btn-primary">
          Add Book
        </router-link>
      </div>
    </div>

    <div class="mt-6">
      <book-list
        :books="books"
        :loading="loading"
        :current-page="currentPage"
        :total-count="totalCount"
        :page-size="pageSize"
        @delete="deleteBook"
        @page-change="handlePageChange"
      />
    </div>
  </div>
</template>

<script>
import BookList from "@/components/BookList.vue";
import bookService from "@/services/bookService";

export default {
  name: "HomeView",
  components: {
    BookList,
  },
  data() {
    return {
      books: [],
      loading: true,
      currentPage: 1,
      pageSize: 10,
      totalCount: 0,
    };
  },
  created() {
    this.fetchBooks();
  },
  methods: {
    async fetchBooks() {
      this.loading = true;
      try {
        const response = await bookService.getBooks(
          this.currentPage,
          this.pageSize
        );
        this.books = response.books;
        this.totalCount = response.totalCount;
      } catch (error) {
        console.error("Failed to fetch books:", error);
        // Handle error (show notification)
      } finally {
        this.loading = false;
      }
    },
    async deleteBook(id) {
      if (!confirm("Are you sure you want to delete this book?")) {
        return;
      }

      try {
        await bookService.deleteBook(id);
        // Remove from local list
        this.books = this.books.filter((book) => book.id !== id);
        // Refetch if it was the last item on a page
        if (this.books.length === 0 && this.currentPage > 1) {
          this.currentPage -= 1;
          this.fetchBooks();
        } else {
          // Update total count
          this.totalCount -= 1;
        }
      } catch (error) {
        console.error("Failed to delete book:", error);
        // Handle error (show notification)
      }
    },
    handlePageChange(page) {
      this.currentPage = page;
      this.fetchBooks();
    },
  },
};
</script>
