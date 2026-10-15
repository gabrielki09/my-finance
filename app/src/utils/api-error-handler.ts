import axios from "axios";
import { Notify } from "quasar";

interface ApiErrorDetail {
    message?: string;
    fields?: Record<string, string | string[]>;
};

interface ApiErrorResponse {
    message?: string;
    success?: boolean;
    error?: ApiErrorDetail
};

const getFieldError = (
    fields?: Record<string, string | string[]>
): string | null => {
    if (!fields) return null;

    for (const messages of Object.values(fields)) {
        console.log(messages);

        if (messages.length > 0) return messages[0] ?? null;
    };

    return null;
};

export const getApiErrorMessage = (
    error: unknown,
    fallback = 'Não possível concluir a operação.'
): string => {
    if (!axios.isAxiosError<ApiErrorResponse>(error)) return error instanceof Error ? error.message : fallback;

    if (!error.response) return 'Não foi possível conectar ao servidor.';

    const response = error.response.data;

    console.log(response.error?.fields);

    const fieldError = getFieldError(response?.error?.fields);

    if (fieldError) return fieldError;

    if (response?.error?.message) return response.error?.message;

    if (response.message) return response.message;

    return fallback;
};

export const handleApiError = (
    error: unknown,
    fallback?: string
): void => {
    console.error(error);

    Notify.create({
        type: 'negative',
        position: 'top',
        message: getApiErrorMessage(error, fallback),
        timeout: 5000,
        actions: [
            {
                icon: 'close',
                color: 'white'
            }
        ]
    });
};
