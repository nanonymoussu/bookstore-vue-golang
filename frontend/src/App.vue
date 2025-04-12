<template>
  <div class="min-h-screen bg-gray-100">
    <nav class="bg-white shadow">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="flex justify-between h-16">
          <div class="flex">
            <div class="flex-shrink-0 flex items-center">
              <router-link to="/" class="text-xl font-bold text-blue-600">
                Bookstore
              </router-link>
            </div>
            <div class="ml-6 flex items-center space-x-4">
              <router-link
                to="/"
                class="px-3 py-2 rounded-md text-sm font-medium hover:bg-gray-100"
              >
                Home
              </router-link>
              <router-link
                to="/books/create"
                class="px-3 py-2 rounded-md text-sm font-medium hover:bg-gray-100"
              >
                Add Book
              </router-link>
            </div>
          </div>
        </div>
      </div>
    </nav>

    <main class="max-w-7xl mx-auto py-6 sm:px-6 lg:px-8">
      <router-view />
    </main>
  </div>
</template>

<script>
export default {
  name: "App",
};
</script>
