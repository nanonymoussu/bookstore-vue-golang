<template>
  <div>
    <div
      class="pb-5 border-b border-gray-200 sm:flex sm:items-center sm:justify-between"
    >
      <h1 class="text-2xl font-bold text-gray-900">Book Details</h1>
      <div class="mt-3 flex space-x-3 sm:mt-0">
        <router-link :to="{ name: 'home' }" class="btn btn-secondary">
          Back to Books
        </router-link>
        <router-link
          :to="{ name: 'book-edit', params: { id } }"
          class="btn btn-primary"
        >
          Edit
        </router-link>
      </div>
    </div>

    <div class="mt-6">
      <div v-if="loading" class="text-center py-10">
        <div class="text-gray-500">Loading book details...</div>
      </div>

      <div v-else-if="error" class="bg-red-50 p-4 rounded-md">
        <div class="flex">
          <div class="flex-shrink-0">
            <!-- Error icon -->
            <svg
              class="h-5 w-5 text-red-400"
              xmlns="http://www.w3.org/2000/svg"
              viewBox="0 0 20 20"
              fill="currentColor"
              aria-hidden="true"
            >
              <path
                fill-rule="evenodd"
                d="M10 18a8 8 0 100-16 8 8 0 000 16zM8.707 7.293a1 1 0 00-1.414 1.414L8.586 10l-1.293 1.293a1 1 0 101.414 1.414L10 11.414l1.293 1.293a1 1 0 001.414-1.414L11.414 10l1.293-1.293a1 1 0 00-1.414-1.414L10 8.586 8.707 7.293z"
                clip-rule="evenodd"
              />
            </svg>
          </div>
          <div class="ml-3">
            <h3 class="text-sm font-medium text-red-800">
              Error loading book details
            </h3>
            <div class="mt-2 text-sm text-red-700">
              <p>{{ error }}</p>
            </div>
          </div>
        </div>
      </div>

      <div v-else class="bg-white shadow overflow-hidden sm:rounded-lg">
        <div class="px-4 py-5 sm:px-6">
          <h3 class="text-lg leading-6 font-medium text-gray-900">
            {{ book.title }}
          </h3>
          <p class="mt-1 max-w-2xl text-sm text-gray-500">
            by {{ book.author }}
          </p>
        </div>
        <div class="border-t border-gray-200">
          <dl>
            <div
              class="bg-gray-50 px-4 py-5 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-6"
            >
              <dt class="text-sm font-medium text-gray-500">ISBN</dt>
              <dd class="mt-1 text-sm text-gray-900 sm:mt-0 sm:col-span-2">
                {{ book.isbn }}
              </dd>
            </div>
            <div
              class="bg-white px-4 py-5 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-6"
            >
              <dt class="text-sm font-medium text-gray-500">Price</dt>
              <dd class="mt-1 text-sm text-gray-900 sm:mt-0 sm:col-span-2">
                ${{ book.price ? book.price.toFixed(2) : "0.00" }}
              </dd>
            </div>
            <div
              class="bg-gray-50 px-4 py-5 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-6"
            >
              <dt class="text-sm font-medium text-gray-500">Stock</dt>
              <dd class="mt-1 text-sm text-gray-900 sm:mt-0 sm:col-span-2">
                {{ book.stock }} items
              </dd>
            </div>
            <div
              class="bg-white px-4 py-5 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-6"
            >
              <dt class="text-sm font-medium text-gray-500">Added on</dt>
              <dd class="mt-1 text-sm text-gray-900 sm:mt-0 sm:col-span-2">
                {{ formatDate(book.createdAt) }}
              </dd>
            </div>
            <div
              class="bg-gray-50 px-4 py-5 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-6"
            >
              <dt class="text-sm font-medium text-gray-500">Last updated</dt>
              <dd class="mt-1 text-sm text-gray-900 sm:mt-0 sm:col-span-2">
                {{ formatDate(book.updatedAt) }}
              </dd>
            </div>
          </dl>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import bookService from "@/services/bookService";

export default {
  name: "BookDetailsView",
  props: {
    id: {
      type: [String, Number],
      required: true,
    },
  },
  data() {
    return {
      book: {},
      loading: true,
      error: null,
    };
  },
  created() {
    this.fetchBook();
  },
  methods: {
    async fetchBook() {
      this.loading = true;
      this.error = null;

      try {
        this.book = await bookService.getBook(this.id);
      } catch (error) {
        console.error("Failed to fetch book details:", error);
        this.error =
          "Could not load book details. The book might have been deleted or you may not have permission to view it.";
      } finally {
        this.loading = false;
      }
    },
    formatDate(dateString) {
      if (!dateString) return "N/A";

      const date = new Date(dateString);
      return new Intl.DateTimeFormat("en-US", {
        year: "numeric",
        month: "long",
        day: "numeric",
        hour: "2-digit",
        minute: "2-digit",
      }).format(date);
    },
  },
};
</script>
