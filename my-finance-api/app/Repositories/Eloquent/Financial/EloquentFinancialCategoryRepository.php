<?php

namespace App\Repositories\Eloquent\Financial;

use App\Models\Financial\FinancialCategory;
use App\Repositories\Eloquent\EloquentBaseRepository;
use App\Repositories\Interface\Financial\FinancialCategoryInterface;

class EloquentFinancialCategoryRepository extends EloquentBaseRepository implements FinancialCategoryInterface
{
    protected function modelClass(): string
    {
        return FinancialCategory::class;
    }
}
