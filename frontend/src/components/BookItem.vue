<template>
  <div class="bg-white shadow overflow-hidden sm:rounded-lg">
    <div class="px-4 py-5 sm:px-6 flex justify-between items-center">
      <div>
        <h3 class="text-lg leading-6 font-medium text-gray-900">
          {{ book.title }}
        </h3>
        <p class="mt-1 max-w-2xl text-sm text-gray-500">by {{ book.author }}</p>
      </div>
      <div class="flex space-x-2">
        <router-link
          :to="{ name: 'book-details', params: { id: book.id } }"
          class="btn btn-secondary"
        >
          View
        </router-link>
        <router-link
          :to="{ name: 'book-edit', params: { id: book.id } }"
          class="btn btn-primary"
        >
          Edit
        </router-link>
        <button @click="$emit('delete', book.id)" class="btn btn-danger">
          Delete
        </button>
      </div>
    </div>
    <div class="border-t border-gray-200">
      <dl>
        <div
          class="bg-gray-50 px-4 py-3 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-6"
        >
          <dt class="text-sm font-medium text-gray-500">ISBN</dt>
          <dd class="mt-1 text-sm text-gray-900 sm:mt-0 sm:col-span-2">
            {{ book.isbn }}
          </dd>
        </div>
        <div class="bg-white px-4 py-3 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-6">
          <dt class="text-sm font-medium text-gray-500">Price</dt>
          <dd class="mt-1 text-sm text-gray-900 sm:mt-0 sm:col-span-2">
            ${{ book.price.toFixed(2) }}
          </dd>
        </div>
        <div
          class="bg-gray-50 px-4 py-3 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-6"
        >
          <dt class="text-sm font-medium text-gray-500">Stock</dt>
          <dd class="mt-1 text-sm text-gray-900 sm:mt-0 sm:col-span-2">
            {{ book.stock }} items
          </dd>
        </div>
      </dl>
    </div>
  </div>
</template>

<script>
export default {
  name: "BookItem",
  props: {
    book: {
      type: Object,
      required: true,
    },
  },
  emits: ["delete"],
};
</script>
