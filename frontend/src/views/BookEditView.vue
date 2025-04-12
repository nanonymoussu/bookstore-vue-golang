<template>
  <div>
    <div
      class="pb-5 border-b border-gray-200 sm:flex sm:items-center sm:justify-between"
    >
      <h1 class="text-2xl font-bold text-gray-900">Edit Book</h1>
      <div class="mt-3 flex space-x-3 sm:mt-0">
        <router-link
          :to="{ name: 'book-details', params: { id } }"
          class="btn btn-secondary"
        >
          Cancel
        </router-link>
      </div>
    </div>

    <div class="mt-6">
      <div v-if="loading" class="text-center py-10">
        <div class="text-gray-500">Loading book data...</div>
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
            <h3 class="text-sm font-medium text-red-800">Error</h3>
            <div class="mt-2 text-sm text-red-700">
              <p>{{ error }}</p>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
