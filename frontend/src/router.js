import { createRouter, createWebHistory } from "vue-router";
import HomeView from "./views/HomeView.vue";
import BookDetailsView from "./views/BookDetailsView.vue";
import BookCreateView from "./views/BookCreateView.vue";
import BookEditView from "./views/BookEditView.vue";

const routes = [
  {
    path: "/",
    name: "home",
    component: HomeView,
  },
  {
    path: "/books/:id",
    name: "book-details",
    component: BookDetailsView,
    props: true,
  },
  {
    path: "/books/create",
    name: "book-create",
    component: BookCreateView,
  },
  {
    path: "/books/:id/edit",
    name: "book-edit",
    component: BookEditView,
    props: true,
  },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
