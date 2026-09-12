<?php

namespace App\Service\Financial;

use App\Enum\Financial\FinancialReleasesStatus;
use App\Exceptions\BusinessException;
use App\Models\Financial\FinancialRelease;
use App\Repositories\Interface\Financial\FinancialAccountInterface;
use App\Repositories\Interface\Financial\FinancialReleaseInterface;
use Carbon\Carbon;
use Illuminate\Database\Eloquent\ModelNotFoundException;
use Illuminate\Support\Facades\DB;

class FinancialReleaseService
{
    public function __construct(
        protected FinancialReleaseInterface $eloquentFinancialReleaseRepository,
        protected FinancialAccountInterface $eloquentFinancialAccountRepository
    ){}

    public function all()
    {
        return $this->eloquentFinancialReleaseRepository->all();
    }

        public function find(string|int $id, int $userId): FinancialRelease
    {
        $financialRelease = $this->eloquentFinancialReleaseRepository->find($id, $userId);

        if ( ! $financialRelease ) throw new ModelNotFoundException('Lançamento financeiro não localizada.');

        return $financialRelease;
    }

    public function findByUuid(string $uuid, int $userId): FinancialRelease
    {
        $financialRelease = $this->eloquentFinancialReleaseRepository->findByUuid($uuid, $userId);

        if ( ! $financialRelease ) throw new ModelNotFoundException('Lançamento financeiro não localizada.');

        return $financialRelease;
    }

    public function create(array $data): FinancialRelease
    {
        return $this->eloquentFinancialReleaseRepository->create($data);
    }

    public function update(array $data, string $uuid)
    {
        return $this->eloquentFinancialReleaseRepository->update($this->findByUuid($uuid, $data['user_id']), $data);
    }

    public function delete(string $uuid, int $userId)
    {
        $this->eloquentFinancialReleaseRepository->delete($this->findByUuid($uuid, $userId));
    }

    public function active(string $uuid, int $userId)
    {
        $this->eloquentFinancialReleaseRepository->active($this->findByUuid($uuid, $userId));
    }

    public function payFinancialRelease(string $uuid, int $userId): FinancialRelease
    {
        return DB::transaction(function() use ($uuid, $userId) {
            $financialRelease = $this->findByUuid($uuid, $userId);

            if ($financialRelease->status == FinancialReleasesStatus::CANCELLED->value)
            {
                throw new BusinessException('Esse lançamento financeiro já está cancelado.');
            }

            if ($financialRelease->status == FinancialReleasesStatus::PAID->value)
            {
                throw new BusinessException('Esse lançamento financeiro já está pago.');
            }

            return $this->eloquentFinancialReleaseRepository->payFinancialRelease($financialRelease, Carbon::now());
        });
    }

    public function cancelFinancialRelease(string $uuid, int $userId): FinancialRelease
    {
        $financialRelease = $this->findByUuid($uuid, $userId);

        if ($financialRelease->status == FinancialReleasesStatus::CANCELLED->value)
        {
            throw new BusinessException('Esse lançamento financeiro já está cancelado.');
        }

        return $this->eloquentFinancialReleaseRepository->cancelFinancialRelease($financialRelease, Carbon::now());
    }

    public function updateFinancialAccountBalance(string $financialReleaseUuid, int $userId, float $value)
    {
        $financialRelease = $this->findByUuid($financialReleaseUuid, $userId);
        $financialAccount = $this->eloquentFinancialAccountRepository->findByUuid($financialRelease->account_id, $userId, [], true);

    }
}
