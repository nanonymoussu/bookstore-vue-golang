<template>
  <form @submit.prevent="handleSubmit" class="space-y-6">
    <div>
      <label for="title" class="form-label">Title</label>
      <input
        id="title"
        v-model="formData.title"
        type="text"
        required
        class="form-input"
      />
    </div>

    <div>
      <label for="author" class="form-label">Author</label>
      <input
        id="author"
        v-model="formData.author"
        type="text"
        required
        class="form-input"
      />
    </div>

    <div>
      <label for="isbn" class="form-label">ISBN</label>
      <input
        id="isbn"
        v-model="formData.isbn"
        type="text"
        required
        class="form-input"
      />
    </div>

    <div>
      <label for="price" class="form-label">Price</label>
      <input
        id="price"
        v-model.number="formData.price"
        type="number"
        step="0.01"
        min="0"
        required
        class="form-input"
      />
    </div>

    <div>
      <label for="stock" class="form-label">Stock</label>
      <input
        id="stock"
        v-model.number="formData.stock"
        type="number"
        min="0"
        required
        class="form-input"
      />
    </div>

    <div class="flex justify-end">
      <button
        type="button"
        @click="$emit('cancel')"
        class="btn btn-secondary mr-3"
      >
        Cancel
      </button>
      <button type="submit" class="btn btn-primary">
        {{ submitButtonText }}
      </button>
    </div>
  </form>
</template>

<script>
export default {
  name: "BookForm",
  props: {
    initialData: {
      type: Object,
      default: () => ({
        title: "",
        author: "",
        isbn: "",
        price: 0,
        stock: 0,
      }),
    },
    submitButtonText: {
      type: String,
      default: "Save",
    },
  },
  emits: ["submit", "cancel"],
  data() {
    return {
      formData: { ...this.initialData },
    };
  },
  watch: {
    initialData: {
      handler(newVal) {
        this.formData = { ...newVal };
      },
      deep: true,
    },
  },
  methods: {
    handleSubmit() {
      this.$emit("submit", { ...this.formData });
    },
  },
};
</script>
