<?php

use App\Http\Controllers\UserController;
use Illuminate\Support\Facades\Route;

require_once __DIR__."/Auth/AuthRoute.php";
Route::apiResource('users', UserController::class);

Route::middleware('auth:sanctum')->group(function() {
    require_once __DIR__."/Financial/FinancialCategory/FinancialCategoryRoute.php";
    require_once __DIR__."/Financial/FinancialAccount/FinancialAccountRoute.php";
});
