<?php

namespace App\Repositories\Eloquent\Financial;

use App\Enum\Financial\FinancialReleasesStatus;
use App\Models\Financial\FinancialRelease;
use App\Repositories\Eloquent\EloquentBaseRepository;
use App\Repositories\Interface\Financial\FinancialReleaseInterface;
use Carbon\Carbon;

class EloquentFinancialReleaseRepository extends EloquentBaseRepository implements FinancialReleaseInterface
{
    protected function modelClass(): string
    {
        return FinancialRelease::class;
    }

    public function create(array $data): FinancialRelease
    {
        return $this->query()->create($data);
    }

    public function findByUuid(string $uuid, int $userId, array $with = [], ?bool $lock = false): ?FinancialRelease
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

    public function payFinancialRelease(FinancialRelease $financialRelease, Carbon $paidAt): FinancialRelease
    {
        $financialRelease->update([
            'status' => FinancialReleasesStatus::PAID->value,
            'paid_at' => $paidAt
        ]);

        return $financialRelease->fresh();
    }

    public function cancelFinancialRelease(FinancialRelease $financialRelease, Carbon $cancelledAt): FinancialRelease
    {
        $financialRelease->update([
            'status' => FinancialReleasesStatus::CANCELLED->value,
            'paid_at' => $cancelledAt
        ]);

        return $financialRelease->fresh();
    }

    public function updateFinancialAccountBalance(string $uuid, string $financialReleaseUuid, int $userId, float $value)
    {

    }
}
