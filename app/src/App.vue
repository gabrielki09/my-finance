<template>
    <router-view />
</template>

<script setup lang="ts">
    import { onMounted } from 'vue';
    import { api } from './boot/axios';
    import { useRouter } from 'vue-router';

    const router = useRouter();

    onMounted(async () => {
        const { data } = await api.get('/health');

        if (!data.success)
        {
            router.push('/internal-server-error');
            return;
        };
    });
</script>
