<template>
    <q-layout view="lHh Lpr lFf">
        <q-header elevated>
            <q-toolbar>
                <q-btn flat dense round icon="menu" aria-label="Menu" @click="toggleLeftDrawer" />

                <q-toolbar-title>
                    Meu financeiro
                </q-toolbar-title>
            </q-toolbar>
        </q-header>

        <q-drawer v-model="leftDrawerOpen" show-if-above bordered>
            <q-list>
                <q-item-label header>
                    Finance App
                </q-item-label>

                <EssentialLink v-for="link in linksList" :key="link.label" v-bind="link" />
            </q-list>
        </q-drawer>

        <q-page-container>
            <router-view />
        </q-page-container>
    </q-layout>
</template>

<script setup lang="ts">
    import { ref } from 'vue';
    import EssentialLink, { type EssentialLinkProps } from '@/components/EssentialLink.vue';

    const linksList: EssentialLinkProps[] = [
        {
            label: 'Home',
            caption: 'dashboard',
            icon: 'dashboard',
            link: '/'
        },
        {
            label: 'Categorias',
            caption: 'categorias financeiras',
            icon: 'category',
            link: '/categories'
        },
    ];

    const leftDrawerOpen = ref<boolean>(false);

    const toggleLeftDrawer = () => {
        leftDrawerOpen.value = !leftDrawerOpen.value;
    };
</script>
