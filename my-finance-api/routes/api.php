<?php

use App\Http\Controllers\Financial\FinancialCategoryController;
use App\Http\Controllers\UserController;
use Illuminate\Support\Facades\Route;

require_once __DIR__."/Auth/AuthRoute.php";

Route::apiResource('users', UserController::class);

Route::prefix('categories')->controller(FinancialCategoryController::class)->group(function() {
    Route::get('', 'all');
    Route::post('', 'store');
    Route::get('/{uuid}', 'show');
    Route::put('/{uuid}', 'update');
    Route::delete('/{uuid}', 'delete');
});
