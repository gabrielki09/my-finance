<?php

namespace App\Repositories\Eloquent\Financial;

use App\Models\Financial\FinancialAccount;
use App\Repositories\Eloquent\EloquentBaseRepository;
use App\Repositories\Interface\Financial\FinancialAccountInterface;

class EloquentFinancialAccountRepository extends EloquentBaseRepository implements FinancialAccountInterface
{
    protected function modelClass(): string
    {
        return FinancialAccount::class;
    }

    public function findByUuid(string $uuid, int $userId, array $with = [], ?bool $lock = false): ?FinancialAccount
    {
        if ($lock)
        {
            return $this->query()
                    ->with($with)
                    ->where('user_id', $userId)
                    ->where('uuid', $uuid)
                    ->lockForUpdate()
                    ->first();
        }

        return $this->query()
                    ->with($with)
                    ->where('user_id', $userId)
                    ->where('uuid', $uuid)
                    ->first();
    }
}
