<?php

namespace App\Repositories\Eloquent\Financial;

use App\Models\Financial\FinancialAccount;
use App\Repositories\Eloquent\EloquentBaseRepository;
use App\Repositories\Interface\Financial\FinancialCategoryInterface;

class EloquentFinancialAccountRepository extends EloquentBaseRepository implements FinancialCategoryInterface
{
    protected function modelClass(): string
    {
        return FinancialAccount::class;
    }
}
