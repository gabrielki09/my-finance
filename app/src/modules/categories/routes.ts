import type { RouteRecordRaw } from "vue-router";

export const categoryRoutes: RouteRecordRaw[] = [
    {
        path: '/categories',
        children: [
            {
                path: '',
                component: () => import('./pages/CategoryHome.vue')
            }
        ],
    }
];
