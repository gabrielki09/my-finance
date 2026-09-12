<?php

namespace App\Repositories\Interface\Financial;

use App\Models\Financial\FinancialAccount;
use App\Models\Financial\FinancialRelease;
use App\Repositories\Interface\BaseInterface;
use Carbon\Carbon;

interface FinancialReleaseInterface extends BaseInterface
{
    public function findByUuid(string $uuid, int $userId, array $with = [], ?bool $lock = false): ?FinancialRelease;
    public function payFinancialRelease(FinancialRelease $financialRelease, Carbon $paidAt): FinancialRelease;
    public function cancelFinancialRelease(FinancialRelease $financialRelease, Carbon $cancelledAt): FinancialRelease;
    public function updateFinancialAccountBalance(string $uuid, string $financialReleaseUuid, int $userId, float $value);
}
