import type { RouteRecordRaw } from 'vue-router';
import { categoryRoutes } from '@/modules/categories/routes';

const routes: RouteRecordRaw[] = [
    {
        path: '/',
        component: () => import('@/layouts/MainLayout.vue'),
        children: [
            ...categoryRoutes
        ],
    },

    // Always leave this as last one,
    // but you can also remove it
    {
        path: '/:catchAll(.*)*',
        component: () => import('@/pages/ErrorPages/ErrorNotFound.vue'),
    },
    {
        path: '/internal-server-error',
        component: () => import('@/pages/ErrorPages/InternalError.vue')
    }
];

export default routes;
