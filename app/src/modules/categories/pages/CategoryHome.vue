<template>
    <q-page padding>
        <div class="p-4 block">
            <div>
                <q-table
                    title="Categorias"
                    :columns="columns"
                    :rows="categories"
                    row-key="id"
                    hide-no-data
                >
                    <template v-slot:top-right>
                        <q-btn
                            color="primary"
                            label="Cadastrar"
                            @click="showFormDialog = true"
                            no-caps
                        />
                    </template>

                    <template v-slot:body="props">
                        <q-tr>


                        </q-tr>
                    </template>
                </q-table>

                <CategoryFormDialog
                    v-model="showFormDialog"
                    :loading="loading"
                    @save="createCategory"
                />
            </div>
        </div>
    </q-page>
</template>

<script setup lang="ts">
    import { onMounted, ref } from 'vue';
    import type { CategoryPayload } from '../types/category.types'
    import type { QTableColumn } from 'quasar';
    import { storeToRefs } from 'pinia';
    import { useCategoryStore } from '../stores/category.store';
    import CategoryFormDialog from '../components/CategoryFormDialog.vue';
import { handleApiError } from '@/utils/api-error-handler';

    const categoryStore = useCategoryStore();
    const { categories, loading } = storeToRefs(categoryStore);

    const showFormDialog = ref<boolean>(false);

    const columns: QTableColumn[] = [
        {
            label: 'ID',
            name: 'id',
            field: 'id',
            align: 'left'
        },
        {
            label: 'Nome',
            name: 'name',
            field: 'name',
            align: 'left'
        },
        {
            label: 'Ações',
            name: 'actions',
            field: 'actions',
            align: 'left'
        }
    ];

    onMounted(async () => {
        void categoryStore.fetchCategories();
    });

    async function createCategory(payload: CategoryPayload)
    {
        try {
            await categoryStore.createCategory(payload);
            showFormDialog.value = false;
        } catch (error) {
            handleApiError(
                error,
                'Não foi possível cadastrar a categoria.'
            );
        };
    };
</script>
