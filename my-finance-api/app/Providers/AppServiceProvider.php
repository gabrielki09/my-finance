<?php

namespace App\Providers;

use App\Repositories\Eloquent\Financial\EloquentFinancialAccountRepository;
use App\Repositories\Eloquent\Financial\EloquentFinancialCategoryRepository;
use App\Repositories\Eloquent\User\EloquentUserRepository;
use App\Repositories\Interface\Financial\FinancialAccountInterface;
use App\Repositories\Interface\Financial\FinancialCategoryInterface;
use App\Repositories\Interface\User\UserInterface;
use Illuminate\Support\ServiceProvider;

class AppServiceProvider extends ServiceProvider
{
    /**
     * Register any application services.
     */
    public function register(): void
    {
        $this->app->bind(
            UserInterface::class,
            EloquentUserRepository::class
        );

        $this->app->bind(
            FinancialCategoryInterface::class,
            EloquentFinancialCategoryRepository::class
        );

        $this->app->bind(
            FinancialAccountInterface::class,
            EloquentFinancialAccountRepository::class
        );
    }

    /**
     * Bootstrap any application services.
     */
    public function boot(): void
    {
        //
    }
}
