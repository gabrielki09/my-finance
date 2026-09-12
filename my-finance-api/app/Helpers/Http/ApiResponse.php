<?php

use Symfony\Component\HttpFoundation\JsonResponse;

function apiSuccess(
    string $message = 'Sucesso!',
    mixed $data = [],
    int $status = 200,
): JsonResponse {
    return response()->json([
        'success' => true,
        'message' => $message,
        'data' => $data,
        'status' => $status

    ], $status);
};

function apiError(
    string $message = 'Erro ao processar a operação',
    int $status = 400,
    mixed $data = [],
): JsonResponse {
    return response()->json([
        'success' => false,
        'message' => $message,
        'data' => $data,
        'status' => $status

    ], $status);
};
