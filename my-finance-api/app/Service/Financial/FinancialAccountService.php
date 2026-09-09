<?php

namespace App\Service\Financial;

use App\Models\Financial\FinancialAccount;
use App\Repositories\Interface\Financial\FinancialAccountInterface;
use Illuminate\Database\Eloquent\ModelNotFoundException;

class FinancialAccountService
{
    public function __construct(
        protected FinancialAccountInterface $eloquentFinancialAccountRepository
    ){}

    public function all()
    {
        return $this->eloquentFinancialAccountRepository->all();
    }

        public function find(string|int $id): FinancialAccount
    {
        $financialAccount = $this->eloquentFinancialAccountRepository->find($id);

        if ( ! $financialAccount ) throw new ModelNotFoundException('Categoria financeira não localizada.');

        return $financialAccount;
    }

    public function findByUuid(string $uuid): FinancialAccount
    {
        $financialAccount = $this->eloquentFinancialAccountRepository->findByUuid($uuid);

        if ( ! $financialAccount ) throw new ModelNotFoundException('Conta financeira não localizada.');

        return $financialAccount;
    }

    public function create(array $data): FinancialAccount
    {
        return $this->eloquentFinancialAccountRepository->create($data);
    }

    public function update(array $data, string $uuid)
    {
        return $this->eloquentFinancialAccountRepository->update($this->findByUuid($uuid), $data);
    }

    public function delete(string $uuid)
    {
        $this->eloquentFinancialAccountRepository->delete($this->findByUuid($uuid));
    }

}
