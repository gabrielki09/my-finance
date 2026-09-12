<?php

namespace App\Repositories\Interface\Financial;

use App\Models\Financial\FinancialAccount;
use App\Repositories\Interface\BaseInterface;

interface FinancialAccountInterface extends BaseInterface
{
    public function findByUuid(string $uuid, int $userId, array $with = [], ?bool $lock = false): ?FinancialAccount;
}
