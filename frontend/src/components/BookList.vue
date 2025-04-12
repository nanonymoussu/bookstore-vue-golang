<template>
  <div>
    <div v-if="loading" class="text-center py-10">
      <div class="text-gray-500">Loading books...</div>
    </div>

    <div v-else-if="!books.length" class="text-center py-10">
      <div class="text-gray-500">
        No books found. Add some books to get started.
      </div>
      <router-link to="/books/create" class="btn btn-primary mt-4"
        >Add Book</router-link
      >
    </div>

    <div v-else class="space-y-6">
      <book-item
        v-for="book in books"
        :key="book.id"
        :book="book"
        @delete="$emit('delete', book.id)"
      />

      <div v-if="totalPages > 1" class="flex justify-center mt-6">
        <nav
          class="relative z-0 inline-flex shadow-sm -space-x-px"
          aria-label="Pagination"
        >
          <button
            @click="$emit('page-change', currentPage - 1)"
            :disabled="currentPage === 1"
            class="relative inline-flex items-center px-2 py-2 rounded-l-md border border-gray-300 bg-white text-sm font-medium text-gray-500 hover:bg-gray-50"
            :class="{ 'opacity-50 cursor-not-allowed': currentPage === 1 }"
          >
            Previous
          </button>

          <template v-for="page in totalPages" :key="page">
            <button
              @click="$emit('page-change', page)"
              class="relative inline-flex items-center px-4 py-2 border border-gray-300 bg-white text-sm font-medium"
              :class="
                page === currentPage
                  ? 'bg-blue-50 text-blue-600 z-10'
                  : 'text-gray-500 hover:bg-gray-50'
              "
            >
              {{ page }}
            </button>
          </template>

          <button
            @click="$emit('page-change', currentPage + 1)"
            :disabled="currentPage === totalPages"
            class="relative inline-flex items-center px-2 py-2 rounded-r-md border border-gray-300 bg-white text-sm font-medium text-gray-500 hover:bg-gray-50"
            :class="{
              'opacity-50 cursor-not-allowed': currentPage === totalPages,
            }"
          >
            Next
          </button>
        </nav>
      </div>
    </div>
  </div>
</template>

<script>
import BookItem from "@/components/BookItem.vue";

export default {
  name: "BookList",
  components: {
    BookItem,
  },
  props: {
    books: {
      type: Array,
      required: true,
    },
    loading: {
      type: Boolean,
      default: false,
    },
    currentPage: {
      type: Number,
      default: 1,
    },
    totalCount: {
      type: Number,
      default: 0,
    },
    pageSize: {
      type: Number,
      default: 10,
    },
  },
  emits: ["delete", "page-change"],
  computed: {
    totalPages() {
      return Math.ceil(this.totalCount / this.pageSize);
    },
  },
};
</script>
