export type CategoryType = 'income' | 'expense' | 'both';
export const categoryTypesArray: Array<{
    label: string,
    value: CategoryType
}> = [
    {
        label: 'Entradas',
        value: 'income',
    },
    {
        label: 'Saídas',
        value: 'expense',
    },
    {
        label: 'Ambas',
        value: 'both',
    }
];

export interface CategoryPayload {
    parentId: number | null;
    name: string;
    type: CategoryType;
};

export interface Category {
    id: number;
    parentId?: number | null;
    name: string;
    type: CategoryType;
    createdAt: Date;
    updatedAt: Date;
    deletedAt?: Date;
};
