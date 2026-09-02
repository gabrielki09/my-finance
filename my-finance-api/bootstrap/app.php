<?php

use App\Exceptions\BusinessException;
use Illuminate\Auth\Access\AuthorizationException;
use Illuminate\Foundation\Application;
use Illuminate\Foundation\Configuration\Exceptions;
use Illuminate\Foundation\Configuration\Middleware;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\Log;

return Application::configure(basePath: dirname(__DIR__))
    ->withRouting(
        api: __DIR__.'/../routes/api.php',
        web: __DIR__.'/../routes/web.php',
        commands: __DIR__.'/../routes/console.php',
        health: '/up',
    )
    ->withMiddleware(function (Middleware $middleware): void {
        //
    })
    ->withExceptions(function (Exceptions $exceptions): void {

        $exceptions->render(function(AuthorizationException $e, Request $r) {
            Log::error('Usuário não autenticado: ', [
                'e' => $e
            ]);

            return apiError(
                message: 'Usuário não autenticado',
                status: 401
            );
        });

        $exceptions->render(function(BusinessException $e, Request $r) {
            Log::error('Erro de regra de negócio:', [
                'e' => $e
            ]);

            return apiError(
                message: $e->getMessage()
            );
        });

        $exceptions->shouldRenderJsonWhen(
            fn (Request $request) => $request->is('api/*') || $request->expectsJson(),
        );
    })->create();
