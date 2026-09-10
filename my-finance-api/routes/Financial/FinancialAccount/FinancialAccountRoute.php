<?php

use App\Http\Controllers\Financial\FinancialCategoryController;
use Illuminate\Support\Facades\Route;

Route::prefix('accounts')->controller(FinancialCategoryController::class)->group(function() {
    Route::get('', 'all');
    Route::post('', 'store');
    Route::get('/{uuid}', 'show');
    Route::put('/{uuid}', 'update');
    Route::patch('/{uuid}/activate', 'active');
    Route::patch('/{uuid}/deactivate', 'delete');
});
