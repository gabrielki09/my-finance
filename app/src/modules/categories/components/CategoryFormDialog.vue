<script setup lang="ts">
    import { computed, reactive, watch } from 'vue';
    import {
        categoryTypesArray,
        type CategoryPayload
    } from '../types/category.types';
    import { useCategoryStore } from '../stores/category.store';
    import { storeToRefs } from 'pinia';

    const modelValue = defineModel<boolean>({ required: true });

    defineProps<{
        loading?: boolean
    }>();

    const emit = defineEmits<{
        save: [payload: CategoryPayload]
    }>();

    const categoryStore = useCategoryStore();
    const { categories } = storeToRefs(categoryStore);

    const parentOptions = computed(() => {
        return categories.value.filter(c => c.parentId === null);
    });

    const form = reactive<CategoryPayload>({
        parentId: null,
        name: '',
        type: 'both',
    });

    watch(modelValue, (opened) => {
        if (opened) {
            form.name = '';
            form.type = 'both';
            form.parentId = null;
        };
    });

    const submit = () => {
        emit('save', {
            name: form.name.trim(),
            parentId: form.parentId,
            type: form.type
        });
    };
</script>

<template>
    <q-dialog
        v-model="modelValue"
        persistent
    >
        <q-card class="w-full max-w-lg">
            <q-form
                @submit.prevent="submit"
                class="q-gutter-md"
            >
                <q-card-section>
                    <h6>Nova categoria</h6>
                </q-card-section>

                <q-card-section class="m-4">
                    <q-select
                        v-model="form.parentId"
                        :options="parentOptions"
                        option-value="id"
                        option-label="name"
                        emit-value
                        map-options
                        clearable
                        outlined
                        label="Categoria principal"
                    />

                    <q-input
                        v-model="form.name"
                        label="Nome"
                        outlined
                        autofocus
                        :rules="[
                            value => Boolean(value?.trim()) || 'Informe o nome'
                        ]"
                    />

                    <q-select
                        v-model="form.type"
                        :options="categoryTypesArray"
                        option-value="value"
                        option-label="label"
                        emit-value
                        map-options
                        label="Tipo da categoria"
                    />
                </q-card-section>

                <q-card-actions align="right">
                    <q-btn
                        flat
                        label="Cancelar"
                        :disable="loading"
                        @click="modelValue = false"
                    />

                    <q-btn
                        type="submit"
                        color="primary"
                        label="Salvar"
                        :loading="loading"
                    />
                </q-card-actions>

            </q-form>
        </q-card>
    </q-dialog>
</template>
