import { api } from "@/boot/axios";

import type {
    Category,
    CategoryPayload
} from '../types/category.types'
import { ApiResponse } from "@/types/api/api.types";

export const categoryService = {
    async getAll(): Promise<Category[]> {
        const { data } = await api.get<ApiResponse<Category[]>>('/categories');

        return data.data;
    },

    async create(payload: CategoryPayload): Promise<Category> {
        const { data } = await api.post<ApiResponse<Category>>(
            '/categories',
            payload
        );

        return data.data;
    }
};
