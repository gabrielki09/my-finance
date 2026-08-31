<?php

use Symfony\Component\HttpFoundation\JsonResponse;

function apiSuccess(
    string $message = 'Sucesso!',
    mixed $data = [],
    int $status = 200,
    ?bool $success = true,
): JsonResponse {
    return response()->json([
        'success' => $success,
        'message' => $message,
        'data' => $data,
        'status' => $status

    ], $status);
};

function apiError(
    string $message = 'Erro ao processar a operação',
    mixed $data = [],
    bool $success = false,
    int $status = 400
): JsonResponse {
    return response()->json([
        'success' => $success,
        'message' => $message,
        'data' => $data,
        'status' => $status

    ], $status);
};
