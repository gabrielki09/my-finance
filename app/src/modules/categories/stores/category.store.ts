import { defineStore } from "pinia";
import type { Category, CategoryPayload } from '../types/category.types'
import { ref } from "vue";
import { categoryService } from '../services/category.service';

export const useCategoryStore = defineStore('category', () => {
    const categories = ref<Category[]>([]);
    const loading = ref<boolean>(false);

    async function fetchCategories() {
        loading.value = true;

        try {
            categories.value = await categoryService.getAll();
        } finally {
            loading.value = false
        };
    };

    async function createCategory(payload: CategoryPayload) {
        loading.value = true;

        try {
            const category = await categoryService.create(payload);
            categories.value.push(category);
        } finally {
            loading.value = false;
        };
    };

    return {
        categories,
        loading,
        fetchCategories,
        createCategory,
    };
});
